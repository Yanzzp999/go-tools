package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	prettyPrint bool
	outputFile  string
)

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "JSON 格式化和验证工具",
	Long:  `用于格式化、验证和美化 JSON 数据的工具`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var input []byte
		var err error

		// 从参数或标准输入读取
		if len(args) > 0 {
			input = []byte(args[0])
		} else {
			fmt.Println("请输入 JSON 数据 (按 Ctrl+D 结束):")
			input, err = os.ReadFile("/dev/stdin")
			if err != nil {
				fmt.Printf("读取输入失败: %v\n", err)
				return
			}
		}

		// 验证 JSON
		var jsonData interface{}
		if err := json.Unmarshal(input, &jsonData); err != nil {
			fmt.Printf("无效的 JSON: %v\n", err)
			return
		}

		// 格式化输出
		var output []byte
		if prettyPrint {
			output, err = json.MarshalIndent(jsonData, "", "  ")
		} else {
			output, err = json.Marshal(jsonData)
		}

		if err != nil {
			fmt.Printf("格式化失败: %v\n", err)
			return
		}

		// 输出结果
		if outputFile != "" {
			err = os.WriteFile(outputFile, output, 0644)
			if err != nil {
				fmt.Printf("写入文件失败: %v\n", err)
				return
			}
			fmt.Printf("结果已保存到: %s\n", outputFile)
		} else {
			fmt.Println(string(output))
		}
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)

	jsonCmd.Flags().BoolVarP(&prettyPrint, "pretty", "p", true, "美化输出格式")
	jsonCmd.Flags().StringVarP(&outputFile, "output", "o", "", "输出到文件")
}
