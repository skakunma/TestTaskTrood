package manager

import "github.com/skakunma/TestTaskTrood/internal/config"

func SendToManager(cfg *config.Config, message string) error {
	//I did it this way. Too little time to implement sending messages

	cfg.Sugar.Infof("Сообщение менеджеру: %s", message)
	return nil

}
