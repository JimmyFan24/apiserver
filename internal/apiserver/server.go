package apiserver

import (
	"apiserver/internal/apiserver/config"
	genericoption "apiserver/internal/pkg/options"
	genericapiserver "apiserver/internal/pkg/server"
	"github.com/sirupsen/logrus"
)

type apiServer struct {
	redisOptions     *genericoption.RedisOptions
	gRPCAPIServer    *grpcAPIServer
	genericApiServer *genericapiserver.GenericAPIServer
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
func buildGenericConfig(cfg *config.Config) (genericConfig *genericapiserver.Config, lastErr error) {
	genericConfig = genericapiserver.NewConfig()

	if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}
	if lastErr = cfg.FeatureOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}
	if lastErr = cfg.JwtOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}
	if lastErr = cfg.InsecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}
	if lastErr = cfg.SecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	return
}
