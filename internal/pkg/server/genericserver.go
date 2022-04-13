package server

import "github.com/sirupsen/logrus"

type GenericAPIServer struct {
}

// Run spawns the http server. It only returns when the port cannot be listened on initially.
func (s *GenericAPIServer) Run() error {
	logrus.Info("genericserver run func...")
	return nil
}
