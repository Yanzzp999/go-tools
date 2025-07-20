package cmd

import (
	"fmt"
	"os"

	"github.com/Yanzzp999/go-tools/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	hashType   string
	hashFile   string
	hashString string
)

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "计算文件或字符串的哈希值",
	Long:  `支持 MD5、SHA1、SHA256、BLAKE3、xxhash 哈希算法计算文件或字符串的哈希值`,
	Run: func(cmd *cobra.Command, args []string) {
		var result string
		var err error

		// 验证哈希类型
		var ht utils.HashType
		switch hashType {
		case "md5":
			ht = utils.MD5
		case "sha1":
			ht = utils.SHA1
		case "sha256":
			ht = utils.SHA256
		case "blake3":
			ht = utils.BLAKE3
		case "xxhash":
			ht = utils.XXHASH
		default:
			fmt.Printf("不支持的哈希类型: %s\n", hashType)
			fmt.Println("支持的类型: md5, sha1, sha256, blake3, xxhash")
			return
		}

		if hashFile != "" {
			// 检查文件是否存在
			if _, err := os.Stat(hashFile); os.IsNotExist(err) {
				fmt.Printf("文件不存在: %s\n", hashFile)
				return
			}
			result, err = utils.HashFile(hashFile, ht)
			if err != nil {
				fmt.Printf("计算文件哈希失败: %v\n", err)
				return
			}
			fmt.Printf("%s (%s): %s\n", hashFile, hashType, result)
		} else if hashString != "" {
			result, err = utils.HashString(hashString, ht)
			if err != nil {
				fmt.Printf("计算字符串哈希失败: %v\n", err)
				return
			}
			fmt.Printf("字符串 '%s' (%s): %s\n", hashString, hashType, result)
		} else {
			fmt.Println("请指定要计算哈希的文件或字符串")
			fmt.Println("使用 --file 指定文件路径")
			fmt.Println("使用 --string 指定字符串")
		}
	},
}

func init() {
	rootCmd.AddCommand(hashCmd)

	hashCmd.Flags().StringVarP(&hashType, "type", "t", "md5", "哈希类型 (md5, sha1, sha256, blake3, xxhash)")
	hashCmd.Flags().StringVarP(&hashFile, "file", "f", "", "要计算哈希的文件路径")
	hashCmd.Flags().StringVarP(&hashString, "string", "s", "", "要计算哈希的字符串")
}
