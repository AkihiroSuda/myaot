package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	if err := newRootCommand().Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "myaot",
		Short:         "An experimental AOT(-ish) compiler",
		Args:          cobra.NoArgs,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	flags := cmd.PersistentFlags()
	flags.Bool("debug", false, "debug mode")

	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if debug, _ := cmd.Flags().GetBool("debug"); debug {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	}

	cmd.AddCommand(
		newCompileCommand(),
	)
	return cmd
}
