package cmd

import (
	"github.com/muyehub/php-package-builder/internal/builder"
	"github.com/spf13/cobra"
)

var (
	name  string
	email string
)

var builderCmd = &cobra.Command{
	Use:   "builder",
	Short: "php 包创建",
	Long:  "支持 php 自定义包的创建",
	Run: func(cmd *cobra.Command, args []string) {
		builder.Build(name, email)
	},
}

func init() {
	builderCmd.Flags().StringVarP(&name, "name", "n", "", "请输入包名")
	builderCmd.Flags().StringVarP(&email, "email", "e", "", "请输入邮箱")
}
