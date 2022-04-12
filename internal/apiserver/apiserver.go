package apiserver

import (
	"apiserver/internal/options"
	"apiserver/pkg/app"
	"fmt"
)

func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp(
		"API Server",
		basename, app.WithOptions(opts),
		app.WithRunFunc(run(opts)),
		app.WithDefaultValidArgs(),
		app.WithDescription("This is an apiserver"))
	return application
}
func run(opt *options.Options) app.RunFunc {
	return func(basename string) error {
		fmt.Println("running...")
		fmt.Println()
		return nil
	}
}
