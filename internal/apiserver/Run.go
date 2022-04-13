package apiserver

import (
	"apiserver/internal/apiserver/config"
	"github.com/sirupsen/logrus"
)

func Run(cfg *config.Config) error {
	logrus.Info("Run  run func...")
	server, err := createApiServer(cfg)
	if err != nil {
		return err
	}

	return server.PrepareRun().Run()
}
func createApiServer(cfg *config.Config) (*apiServer, error) {
	_, _ = buildGenericConfig(cfg)
	logrus.Infof(cfg.RedisOptions.Host)
	return &apiServer{}, nil
}
