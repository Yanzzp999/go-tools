package cmd

import (
	"fmt"
	"os"

	"github.com/Yanzzp999/go-tools/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	targetPath  string
	oldString   string
	newString   string
	recursive   bool
	previewMode bool
)

// fileManagerCmd represents the file manager command
var fileManagerCmd = &cobra.Command{
	Use:   "filemgr",
	Short: "文件管理工具",
	Long:  `提供文件和目录的批量重命名功能，支持递归处理子目录`,
}

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "批量重命名文件和目录",
	Long: `批量重命名指定路径下的文件和目录，支持字符串替换和递归处理。
	
示例:
  # 在当前目录下将所有包含 "old" 的文件名替换为 "new"
  go-tools filemgr rename --path . --old "old" --new "new"
  
  # 删除文件名中的指定字符串（替换为空字符串）
  go-tools filemgr rename --path . --old "删除" --new ""
  
  # 递归处理所有子目录
  go-tools filemgr rename --path . --old "old" --new "new" --recursive
  
  # 预览模式，只显示将要进行的操作，不实际执行
  go-tools filemgr rename --path . --old "old" --new "new" --preview`,
	Run: func(cmd *cobra.Command, args []string) {
		// 验证必要参数
		if targetPath == "" {
			fmt.Println("错误: 必须指定目标路径")
			fmt.Println("使用 --path 参数指定路径")
			return
		}

		if oldString == "" {
			fmt.Println("错误: 必须指定需要替换的字符串")
			fmt.Println("使用 --old 参数指定需要替换的字符串")
			return
		}

		// 注意：newString 可以为空字符串，表示删除指定的字符串		// 检查路径是否存在
		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			fmt.Printf("错误: 路径不存在: %s\n", targetPath)
			return
		}

		// 预览模式
		if previewMode {
			if err := utils.PreviewRename(targetPath, oldString, newString, recursive); err != nil {
				fmt.Printf("预览失败: %v\n", err)
				return
			}
			fmt.Println("\n提示: 这只是预览，没有实际执行重命名操作")
			fmt.Println("要执行实际操作，请移除 --preview 参数")
			return
		}

		// 执行重命名操作
		fmt.Printf("开始批量重命名操作...\n")
		fmt.Printf("路径: %s\n", targetPath)
		if newString == "" {
			fmt.Printf("删除字符串: '%s'\n", oldString)
		} else {
			fmt.Printf("替换: '%s' -> '%s'\n", oldString, newString)
		}
		fmt.Printf("递归: %t\n", recursive)
		fmt.Println("==================")

		if err := utils.RenameFilesAndDirs(targetPath, oldString, newString, recursive); err != nil {
			fmt.Printf("重命名操作失败: %v\n", err)
			return
		}

		fmt.Println("批量重命名操作完成!")
	},
}

func init() {
	rootCmd.AddCommand(fileManagerCmd)
	fileManagerCmd.AddCommand(renameCmd)

	// 重命名命令的参数
	renameCmd.Flags().StringVarP(&targetPath, "path", "p", "", "目标路径 (必需)")
	renameCmd.Flags().StringVarP(&oldString, "old", "o", "", "需要替换的字符串 (必需)")
	renameCmd.Flags().StringVarP(&newString, "new", "n", "", "替换后的字符串 (可为空字符串，表示删除)")
	renameCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "是否递归处理子目录")
	renameCmd.Flags().BoolVar(&previewMode, "preview", false, "预览模式，只显示将要进行的操作，不实际执行")

	// 标记必需参数
	renameCmd.MarkFlagRequired("path")
	renameCmd.MarkFlagRequired("old")
}
