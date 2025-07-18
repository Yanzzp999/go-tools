package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// 全局变量
	cfgFile string
	verbose bool
)

// rootCmd 代表基础命令
var rootCmd = &cobra.Command{
	Use:   "go-tools",
	Short: "一个实用的Go工具集合",
	Long: `go-tools 是一个包含多种实用功能的Go命令行工具。
这个工具提供了各种常用的功能，帮助提高开发效率。`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("欢迎使用 go-tools! 使用 --help 查看可用命令。")
	},
}

// Execute 添加所有子命令到根命令并设置适当的标志
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// 这里可以定义标志和配置设置
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "配置文件 (默认是 $HOME/.go-tools.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "详细输出")

	// Cobra 也支持本地标志，只有在直接调用此操作时才会运行
	rootCmd.Flags().BoolP("toggle", "t", false, "帮助信息切换")
}

// initConfig 读取配置文件和环境变量
func initConfig() {
	if cfgFile != "" {
		// 使用配置文件从标志
		viper.SetConfigFile(cfgFile)
	} else {
		// 查找主目录
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// 在主目录中搜索名为".go-tools"的配置（不带扩展名）
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".go-tools")
	}

	viper.AutomaticEnv() // 读取匹配的环境变量

	// 如果找到配置文件，读取它
	if err := viper.ReadInConfig(); err == nil {
		if verbose {
			logrus.Infof("使用配置文件: %s", viper.ConfigFileUsed())
		}
	}

	// 设置日志级别
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}
