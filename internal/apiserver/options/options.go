package options

import (
	genericoptions "apiserver/internal/pkg/options"
	"apiserver/pkg/app"
)

type Options struct {
	RedisOptions            *genericoptions.RedisOptions           `json:"redis" mapstructure:"redis"`
	GenericServerRunOptions *genericoptions.ServerRunOptions       `json:"server" mapstructure:"server"`
	GRPCOptions             *genericoptions.GRPCOptions            `json:"grpc" mapstructure:"grpc"`
	InsecureServing         *genericoptions.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	JwtOptions              *genericoptions.JwtOptions             `json:"jwt" mapstructure:"jwt"`
	FeatureOptions          *genericoptions.FeatureOptions         `json:"feature"  mapstructure:"feature"`
	SecureServing           *genericoptions.SecureServingOptions   `json:"secure"   mapstructure:"secure"`
}

func (o *Options) Flags() (fss app.NameFlagSets) {
	o.FeatureOptions.AddFlags(fss.FlagSet("feature"))
	o.RedisOptions.AddFlags(fss.FlagSet("redis"))
	o.GenericServerRunOptions.AddFlags(fss.FlagSet("server"))
	o.JwtOptions.AddFlags(fss.FlagSet("jwt"))
	o.GRPCOptions.AddFlags(fss.FlagSet("grpc"))
	o.SecureServing.AddFlags(fss.FlagSet("secure"))
	o.InsecureServing.AddFlags(fss.FlagSet("insecure"))
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
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		InsecureServing:         genericoptions.NewInsecureServingOptions(),
		SecureServing:           genericoptions.NewSecureServingOptions(),
		JwtOptions:              genericoptions.NewJwtOptions(),
		FeatureOptions:          genericoptions.NewFeatureOptions(),
		GRPCOptions:             genericoptions.NewGRPCOptions(),
	}

	return &o
}

var _ app.CliOptions = &Options{}
