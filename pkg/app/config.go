package app

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

const configFlagName = "config"

var cfgFile string

//配置文件相关flag
func init() {
	pflag.StringVarP(&cfgFile, "config", "c", "C:\\Users\\jimmy\\GolandProjects\\api\\config.json", "Read configuration from specified `FILE`, \"+\n\t\t\"support JSON, TOML, YAML, HCL, or Java properties formats.")
}

func addConfigFlag(basename string, fs *pflag.FlagSet) {
	//传进来的flagset是一个名字叫global的flagset,是空的
	logrus.Info("addConfigFlag")
	fs.AddFlag(pflag.Lookup(configFlagName))

	cobra.OnInitialize(func() {

		if cfgFile != "" {
			logrus.Infof("print cfg file:%v ", cfgFile)
			viper.SetConfigFile(cfgFile)
		} else {
			viper.AddConfigPath(".")

			viper.SetConfigName(basename)
		}

		if err := viper.ReadInConfig(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: failed to read configuration file(%s): %v\n", cfgFile, err)
			os.Exit(1)
		}
		logrus.Info("add config success ")
	})

}
