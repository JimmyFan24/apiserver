package apiserver

import (
	"apiserver/internal/apiserver/config"
	genericoptions "apiserver/internal/pkg/options"
	"fmt"
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
	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}
	extraConfig, err := buildExtraConfig(cfg)
	if err != nil {
		return nil, err
	}
	genericServer, err := genericConfig.Complete().New()
	logrus.Infof(cfg.RedisOptions.Host)
	return &apiServer{}, nil
}

type ExtraConfig struct {
	Addr       string
	MaxMsgSize int
	ServerCert genericoptions.GeneratableKeyCert
}

func buildExtraConfig(cfg *config.Config) (*ExtraConfig, error) {
	return &ExtraConfig{
		Addr:       fmt.Sprintf("%s:%d", cfg.GRPCOptions.BindAddress, cfg.GRPCOptions.BindPort),
		MaxMsgSize: cfg.GRPCOptions.MaxMsgSize,
		ServerCert: cfg.SecureServing.ServerCert,
	}, nil
}
