package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示版本信息",
	Long:  `显示 go-tools 的版本信息`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-tools v1.0.0")
		fmt.Println("构建时间: 2025-07-18")
		fmt.Println("Go版本: go1.24.5")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
