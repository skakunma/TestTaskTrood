package spacy

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/skakunma/TestTaskTrood/internal/config"
)

type Response struct {
	Entities string `json:"entities"`
}

var ErrBadRequest = errors.New("status code from service is not OK")

func CreateRequest(ctx context.Context, cfg *config.Config, message string) (string, error) {
	data := map[string]string{"text": message}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", cfg.SpacyHost+"/process", bytes.NewBuffer(jsonData))

	if err != nil {
		return "", err
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", ErrBadRequest
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	return response.Entities, nil
}
