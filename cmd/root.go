package cmd

import (
	"fmt"
	"gin-template/version"
	"github.com/spf13/cobra"
	"os"
)

var vers bool

var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
			return nil
		}
		return fmt.Errorf("no flags find")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
func init() {
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "the version")
}
