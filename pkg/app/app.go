package app

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/marmotedu/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

var (
	progressMessage = color.GreenString("==>")
	//nolint: deadcode,unused,varcheck
	usageTemplate = fmt.Sprintf(`%s{{if .Runnable}}
  %s{{end}}{{if .HasAvailableSubCommands}}
  %s{{end}}{{if gt (len .Aliases) 0}}

%s
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

%s
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

%s{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  %s {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

%s
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

%s
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

%s{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "%s --help" for more information about a command.{{end}}
`,
		color.CyanString("Usage:"),
		color.GreenString("{{.UseLine}}"),
		color.GreenString("{{.CommandPath}} [command]"),
		color.CyanString("Aliases:"),
		color.CyanString("Examples:"),
		color.CyanString("Available Commands:"),
		color.GreenString("{{rpad .Name .NamePadding }}"),
		color.CyanString("Flags:"),
		color.CyanString("Global Flags:"),
		color.CyanString("Additional help topics:"),
		color.GreenString("{{.CommandPath}} [command]"),
	)
)

type App struct {
	name        string
	basename    string
	description string
	options     CliOptions
	runFunc     RunFunc
	commands    []*Command
	args        cobra.PositionalArgs
	cmd         *cobra.Command

	silence   bool
	noVersion bool
	noConfig  bool
}

type RunFunc func(basename string) error
type OptApp func(a *App)

func NewApp(name string, basename string, opts ...OptApp) *App {
	a := &App{
		name:     name,
		basename: basename,
	}
	for _, o := range opts {
		o(a)
	}
	a.buildCammand()
	return a
}
func (a *App) buildCammand() {
	cmd := cobra.Command{
		Use:           a.basename,
		Short:         a.name,
		Long:          a.description,
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          a.args,
	}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = true
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	if len(a.commands) > 0 {
		for _, command := range a.commands {
			cmd.AddCommand(command.cobraCommand())
		}
		cmd.SetHelpCommand(helpCommand(a.name))
	}

	if a.runFunc != nil {
		cmd.RunE = a.runCommand
	}
	var nameFlagSets NameFlagSets
	if a.options != nil {

		nameFlagSets = a.options.Flags()

		//为root command生成flagset,并且遍历赋值
		fs := cmd.Flags()
		for _, f := range nameFlagSets.FlagSets {

			fs.AddFlagSet(f)
		}

		usageFmt := "Usage:\n  %s\n"

		cols, _, _ := TerminalSize(cmd.OutOrStdout())
		cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
			PrintSections(cmd.OutOrStdout(), nameFlagSets, cols)
		})
		cmd.SetUsageFunc(func(cmd *cobra.Command) error {
			fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
			PrintSections(cmd.OutOrStderr(), nameFlagSets, cols)
			return nil
		})
	}
	if !a.noConfig {

		addConfigFlag(a.basename, nameFlagSets.FlagSet("global"))
	}

	a.cmd = &cmd
}

// Run is used to launch the application.
func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}
func (a *App) applyOptionRules() error {
	if completeableOptions, ok := a.options.(CompleteableOptions); ok {
		if err := completeableOptions.Complete(); err != nil {
			return err
		}
	}

	if errs := a.options.Validate(); len(errs) != 0 {
		return errors.NewAggregate(errs)
	}
	if printableOptions, ok := a.options.(PrintableOptions); ok && !a.silence {
		logrus.Infof("%v Config: `%s`", progressMessage, printableOptions.String())
	}

	if printableOptions, ok := a.options.(PrintableOptions); ok && !a.silence {
		logrus.Infof("%v Config: `%s`", progressMessage, printableOptions.String())
	}

	return nil
}
func (a *App) runCommand(cmd *cobra.Command, args []string) error {
	//PrintFlags(cmd.Flags())
	if a.options != nil {

		if err := a.applyOptionRules(); err != nil {
			return err
		}
	}
	// run application
	if a.runFunc != nil {
		return a.runFunc(a.basename)
	}
	if !a.noConfig {

		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}

		if err := viper.Unmarshal(a.options); err != nil {
			return err
		}
	}

	return nil
}

// PrintFlags logs the flags in the flagset.
func PrintFlags(flags *pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		logrus.Infof("FLAG: --%s=%q", flag.Name, flag.Value)
	})
}

func WithOptions(opt CliOptions) OptApp {
	return func(a *App) {
		a.options = opt
	}
}
func WithRunFunc(runfunc RunFunc) OptApp {
	return func(a *App) {
		a.runFunc = runfunc
	}
}
func WithNoConfig() OptApp {
	return func(a *App) {
		a.noConfig = true
	}
}
func WithDescription(desc string) OptApp {

	return func(a *App) {
		a.description = desc
	}
}

func WithDefaultValidArgs() OptApp {
	return func(a *App) {
		a.args = func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		}
	}
}
