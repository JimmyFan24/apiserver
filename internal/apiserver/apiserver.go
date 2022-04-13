package apiserver

import (
	"apiserver/internal/apiserver/config"
	"apiserver/internal/apiserver/options"
	"apiserver/pkg/app"
	"github.com/sirupsen/logrus"
)

func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp(
		"API Server",
		basename,
		app.WithOptions(opts),
		app.WithRunFunc(run(opts)),
		app.WithDefaultValidArgs(),
		app.WithDescription("This is an apiserver"))
	return application
}
func run(opt *options.Options) app.RunFunc {
	return func(basename string) error {
		logrus.Info("apiserver run func...")
		cfg, err := config.CreateConfigFromOptions(opt)
		if err != nil {
			return err
		}
		return Run(cfg)
	}
}
