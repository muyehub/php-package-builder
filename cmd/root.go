package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

// Execute 命令执行函数
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(builderCmd)
}
