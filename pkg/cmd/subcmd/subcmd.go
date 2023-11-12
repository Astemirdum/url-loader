package subcmd

import (
	"fmt"
	"os"

	"github.com/Astemirdum/url-loader/pkg/loader"
	"github.com/spf13/cobra"
)

func NewCmdSubstr() *cobra.Command {
	var glimit int
	const glimitDefult = 5

	cmd := &cobra.Command{
		Use:     "file",
		Short:   "Obtains a file-path",
		Aliases: []string{"file"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filePath := args[0]
			f, err := os.Open(filePath)
			if err != nil {
				fmt.Printf("%v: file=%s", err, filePath)
				return
			}
			defer f.Close()

			processor := loader.NewProcessor(glimit)
			fmt.Printf("processor run filePath=%s, glimit=%d\n", filePath, glimit)
			if err := processor.Run(f); err != nil {
				fmt.Fprintf(os.Stderr, "processor.Run err: %v\n", err)
			}
		},
	}

	cmd.Flags().IntVarP(&glimit, "glimit", "g", glimitDefult, fmt.Sprintf("go number default %d", glimitDefult))
	return cmd
}
