package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/skakunma/TestTaskTrood/internal/config"
	"github.com/skakunma/TestTaskTrood/internal/manager"
	"github.com/skakunma/TestTaskTrood/internal/spacy"
)

type Request struct {
	Message string `json: "messgae"`
}

func CreateMessage(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.Header.Get("Content-Type"), "application/json") {
			c.JSON(http.StatusBadRequest, "Content-Type must be application/json")
			return
		}

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Failed to read request body")
			cfg.Sugar.Error("Failed to read request body:", err)
			return
		}

		var req Request
		err = json.Unmarshal(body, &req)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid JSON format")
			cfg.Sugar.Error("Invalid JSON format:", err)
			return
		}
		v := validator.New()
		if err := v.Struct(req); err != nil {
			c.JSON(http.StatusBadRequest, `JSON is not correctly must me {"message": "your message"}`)
			cfg.Sugar.Error("JSON is not correctly")
			return
		}

		ctx := c.Request.Context()
		if oldAnswer, err := cfg.Store.GetAnswer(ctx, req.Message); oldAnswer != "" && err == nil {
			c.JSON(http.StatusOK, gin.H{"Answer": oldAnswer})
			return
		}

		answer, err := spacy.CreateRequest(ctx, cfg, req.Message)

		if err != nil && !errors.Is(err, spacy.ErrBadRequest) {
			cfg.Sugar.Errorf("Problem in request to spacy service %s", err.Error())
			c.JSON(http.StatusInternalServerError, "service error")
			return
		}

		if answer == "" {
			if err := manager.SendToManager(cfg, req.Message); err != nil {
				cfg.Sugar.Error("Problem in sending message to manager")
				c.JSON(http.StatusInternalServerError, "service error")
				return
			}
		}

		newAnswer, err := cfg.Store.GetAnswer(ctx, answer)

		if err != nil {
			c.JSON(http.StatusInternalServerError, "service error")
			return
		}

		if newAnswer == "" {
			if err := manager.SendToManager(cfg, req.Message); err != nil {
				cfg.Sugar.Error("Problem in sending message to manager")
				c.JSON(http.StatusInternalServerError, "service error")
				return
			}
			c.JSON(http.StatusOK, "Your request has been forwarded to the manager")
			return
		}

		c.JSON(http.StatusOK, gin.H{"Answer": newAnswer})
		return
	}
}
