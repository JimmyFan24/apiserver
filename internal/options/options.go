package options

import (
	"apiserver/pkg/app"
	genericoptions "apiserver/pkg/options"
)

type Options struct {
	RedisOptions *genericoptions.RedisOptions
}

func (o Options) Flags() (fss app.NameFlagSets) {
	o.RedisOptions.AddFlags(fss.FlagSet("redis"))
	return fss
}

func (o *Options) Validate() []error {
	var errs []error
	errs = append(errs, o.RedisOptions.Validate()...)
	return errs
}

func NewOptions() *Options {
	o := Options{RedisOptions: genericoptions.NewRedisOptions()}
	return &o
}

var _ app.CliOptions = &Options{}
