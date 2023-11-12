package cmd

import (
	"fmt"
	"os"

	"github.com/Astemirdum/url-loader/pkg/cmd/subcmd"
	"github.com/spf13/cobra"
)

const version = "1.0"

var rootCmd = &cobra.Command{
	Use:     "url-loaderutil",
	Version: version,
	Short:   "A Simple URL-Loader Utility",
	Long: `URL-Loader Utility is a Cobra framework-based CLI application 
that provides simple load url content functions.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is a Cobra based URL-Loader Utility.")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(subcmd.NewCmdSubstr())
}
