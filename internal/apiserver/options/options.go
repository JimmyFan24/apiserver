package options

import (
	genericoptions "apiserver/internal/pkg/options"
	"apiserver/pkg/app"
)

type Options struct {
	RedisOptions            *genericoptions.RedisOptions
	GenericServerRunOptions *genericoptions.ServerRunOptions
}

func (o *Options) Flags() (fss app.NameFlagSets) {
	o.RedisOptions.AddFlags(fss.FlagSet("redis"))
	o.GenericServerRunOptions.AddFlags(fss.FlagSet("server"))
	return fss
}

func (o *Options) Validate() []error {
	var errs []error
	errs = append(errs, o.RedisOptions.Validate()...)
	return errs
}

func NewOptions() *Options {

	o := Options{
		RedisOptions:            genericoptions.NewRedisOptions(),
		GenericServerRunOptions: genericoptions.NewServerRunOptions()}

	return &o
}

var _ app.CliOptions = &Options{}
