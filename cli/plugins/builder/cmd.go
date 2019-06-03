package builder

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/project-flogo/cli/common"
	"github.com/spf13/cobra"
)

func init() {
	enterpriseBuilderCmd.AddCommand(fePluginCommand)
	common.RegisterPlugin(enterpriseBuilderCmd)
}

var outFile string

var enterpriseBuilderCmd = &cobra.Command{
	Use:   "febuild",
	Short: "Flogo Enterprise builder support",
	Long:  "Using this command to build Flogo Enterprise apps",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

var fePluginCommand = &cobra.Command{
	Use:   "upgrade [flogo.json]",
	Short: "upgrade flogo.json",
	Long:  "Upgrades the flogo.json file",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 && args[0] != "" {

			binFolder, excName := getFEBuilder()
			err := runBuilder(binFolder, excName, args)
			if err != nil {
				fmt.Println("Build app eror:", err.Error())
				os.Exit(1)
			}

		}
	},
}

func runBuilder(builderBin string, builderName string, args []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

	buildArgs := make([]string, len(args)+1)
	buildArgs[0] = "build"

	for i, v := range args {
		buildArgs[i+1] = v
	}

	builderCmd := exec.CommandContext(ctx, builderName, buildArgs...)
	builderCmd.Dir = builderBin

	builderCmd.Stdout = os.Stdout
	builderCmd.Stderr = os.Stderr
	defer func() {
		cancel()
	}()

	err := builderCmd.Run()
	if err != nil {
		if err.Error() == "signal: killed" {
			return fmt.Errorf("build app timeout")
		} else {
			return fmt.Errorf("build app failed, %s", err.Error())
		}
	}
	return nil
}

func getFEBuilder() (string, string) {
	flogoHome, ok := os.LookupEnv("FLOGO_HOME")
	if !ok {
		fmt.Fprintf(os.Stderr, "Flogo Enterprise Home(FLOGO_HOME) env var must set")
		os.Exit(1)
	}

	builderBin := filepath.Join(flogoHome, "bin")
	var builderName string
	if runtime.GOOS == "linux" {
		builderName = "builder-linux_amd64"
	} else if runtime.GOOS == "darwin" {
		builderName = "builder-darwin_amd64"
	} else if runtime.GOOS == "windows" {
		builderName = "builder-windows_amd64.exe"
	} else {
		fmt.Fprintf(os.Stderr, "Unsupport os [%s] defect,  supported os [linux,windows,darwin]", runtime.GOOS)
		os.Exit(1)
	}

	if _, err := os.Stat(filepath.Join(builderBin, builderName)); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "builder executable [%s] not exist", filepath.Join(builderBin, builderName))
		os.Exit(1)
	}

	return builderBin, builderName
}
