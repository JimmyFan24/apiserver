package apiserver

import (
	"apiserver/internal/apiserver/config"
	genericoption "apiserver/internal/pkg/options"
	genericserver "apiserver/internal/pkg/server"
	"github.com/sirupsen/logrus"
)

type apiServer struct {
	redisOptions     *genericoption.RedisOptions
	gRPCAPIServer    *grpcAPIServer
	genericApiServer *genericserver.GenericAPIServer
}

type preparedAPIServer struct {
	*apiServer
}

func (s *apiServer) PrepareRun() preparedAPIServer {
	logrus.Info("preparerun func...")
	return preparedAPIServer{}
}
func (s preparedAPIServer) Run() error {
	logrus.Info("preparerun run func...")
	return s.genericApiServer.Run()
}
func buildGenericConfig(cfg *config.Config) (genericConfig *genericserver.Config, lastErr error) {
	genericConfig = genericserver.NewConfig()
	if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}
	return
}
