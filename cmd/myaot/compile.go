package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AkihiroSuda/myaot/pkg/ccutil"
	"github.com/AkihiroSuda/myaot/pkg/compile"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newCompileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "compile FILE",
		Aliases: []string{"c"},
		Short:   "Compile an riscv32 ELF binary to a native binary",
		Args:    cobra.MinimumNArgs(1),
		RunE:    compileAction,

		DisableFlagsInUseLine: true,
	}

	flags := cmd.Flags()
	flags.StringP("output", "o", "a.out", "Output file, like \"a.out\", or \"a.c\"")

	return cmd
}

func compileAction(cmd *cobra.Command, args []string) error {

	flags := cmd.Flags()
	outFilePath, err := flags.GetString("output")
	if err != nil {
		return err
	}

	inFilePath := args[0]
	inFile, err := os.Open(inFilePath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outCPath := outFilePath + ".c"
	outIsC := filepath.Ext(outFilePath) == ".c"
	if outIsC {
		outCPath = outFilePath
	}
	outC, err := os.Create(outCPath)
	if err != nil {
		return err
	}
	defer outC.Close()

	logrus.Infof("Compiling %s --> %s", inFilePath, outCPath)
	if err = compile.Compile(outC, inFile); err != nil {
		return err
	}
	if err = outC.Close(); err != nil {
		return err
	}

	if !outIsC {
		logrus.Infof("Compiling %s --> %s", outCPath, outFilePath)
		cc, err := ccutil.CC()
		if err != nil {
			return err
		}
		// Use -O0 by default to shorten the compilation time
		ccCmd := exec.Command(cc, "-O0", "-o", outFilePath, outCPath)
		ccCmd.Stdout = os.Stdout
		ccCmd.Stderr = os.Stderr
		logrus.Debugf("Running %v", ccCmd.Args)
		if err = ccCmd.Run(); err != nil {
			return fmt.Errorf("failed to run %v: %w", ccCmd.Args, err)
		}

		logrus.Infof("Removing %s", outCPath)
		if err = os.RemoveAll(outCPath); err != nil {
			return err
		}
	}
	logrus.Infof("Done: %s", outFilePath)
	return nil
}
