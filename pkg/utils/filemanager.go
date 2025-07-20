package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// RenameFilesAndDirs 批量重命名文件和目录
// path: 目标路径
// oldStr: 需要替换的字符串
// newStr: 替换后的字符串
// recursive: 是否递归处理子目录
func RenameFilesAndDirs(path, oldStr, newStr string, recursive bool) error {
	// 检查路径是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("路径不存在: %s", path)
	}

	// 如果是文件，直接处理
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("无法获取路径信息: %v", err)
	}

	if !info.IsDir() {
		return renameFile(path, oldStr, newStr)
	}

	// 如果是目录，处理目录内容
	return renameInDirectory(path, oldStr, newStr, recursive)
}

// renameInDirectory 处理目录内的文件和子目录重命名
func renameInDirectory(dirPath, oldStr, newStr string, recursive bool) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("无法读取目录: %v", err)
	}

	// 先收集所有需要重命名的项目，避免在遍历过程中修改目录结构
	var renameItems []struct {
		oldPath string
		newPath string
		isDir   bool
	}

	for _, entry := range entries {
		oldPath := filepath.Join(dirPath, entry.Name())

		// 检查文件/目录名是否包含需要替换的字符串
		if strings.Contains(entry.Name(), oldStr) {
			newName := strings.ReplaceAll(entry.Name(), oldStr, newStr)
			newPath := filepath.Join(dirPath, newName)

			renameItems = append(renameItems, struct {
				oldPath string
				newPath string
				isDir   bool
			}{
				oldPath: oldPath,
				newPath: newPath,
				isDir:   entry.IsDir(),
			})
		}
	}

	// 执行重命名操作
	for _, item := range renameItems {
		if newStr == "" {
			fmt.Printf("重命名: %s -> %s (删除 '%s')\n", item.oldPath, item.newPath, oldStr)
		} else {
			fmt.Printf("重命名: %s -> %s\n", item.oldPath, item.newPath)
		}

		if err := os.Rename(item.oldPath, item.newPath); err != nil {
			fmt.Printf("重命名失败: %v\n", err)
			continue
		}

		fmt.Printf("重命名成功: %s\n", item.newPath)
	}

	// 如果需要递归处理，处理子目录
	if recursive {
		// 重新读取目录，因为可能有目录被重命名了
		entries, err := os.ReadDir(dirPath)
		if err != nil {
			return fmt.Errorf("重新读取目录失败: %v", err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				subDirPath := filepath.Join(dirPath, entry.Name())
				if err := renameInDirectory(subDirPath, oldStr, newStr, recursive); err != nil {
					fmt.Printf("处理子目录失败 %s: %v\n", subDirPath, err)
				}
			}
		}
	}

	return nil
}

// renameFile 重命名单个文件
func renameFile(filePath, oldStr, newStr string) error {
	dir := filepath.Dir(filePath)
	fileName := filepath.Base(filePath)

	if !strings.Contains(fileName, oldStr) {
		fmt.Printf("文件名不包含需要替换的字符串: %s\n", fileName)
		return nil
	}

	newFileName := strings.ReplaceAll(fileName, oldStr, newStr)
	newFilePath := filepath.Join(dir, newFileName)

	if newStr == "" {
		fmt.Printf("重命名: %s -> %s (删除 '%s')\n", filePath, newFilePath, oldStr)
	} else {
		fmt.Printf("重命名: %s -> %s\n", filePath, newFilePath)
	}

	if err := os.Rename(filePath, newFilePath); err != nil {
		return fmt.Errorf("重命名失败: %v", err)
	}

	fmt.Printf("重命名成功: %s\n", newFilePath)
	return nil
}

// PreviewRename 预览重命名操作，不实际执行
func PreviewRename(path, oldStr, newStr string, recursive bool) error {
	// 检查路径是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("路径不存在: %s", path)
	}

	fmt.Println("=== 重命名预览 ===")
	fmt.Printf("路径: %s\n", path)
	if newStr == "" {
		fmt.Printf("删除字符串: '%s'\n", oldStr)
	} else {
		fmt.Printf("替换: '%s' -> '%s'\n", oldStr, newStr)
	}
	fmt.Printf("递归: %t\n", recursive)
	fmt.Println("==================")

	// 如果是文件，直接处理
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("无法获取路径信息: %v", err)
	}

	if !info.IsDir() {
		return previewFile(path, oldStr, newStr)
	}

	// 如果是目录，预览目录内容
	return previewDirectory(path, oldStr, newStr, recursive)
}

// previewDirectory 预览目录内的重命名操作
func previewDirectory(dirPath, oldStr, newStr string, recursive bool) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("无法读取目录: %v", err)
	}

	for _, entry := range entries {
		oldPath := filepath.Join(dirPath, entry.Name())

		// 检查文件/目录名是否包含需要替换的字符串
		if strings.Contains(entry.Name(), oldStr) {
			newName := strings.ReplaceAll(entry.Name(), oldStr, newStr)
			newPath := filepath.Join(dirPath, newName)

			fileType := "文件"
			if entry.IsDir() {
				fileType = "目录"
			}

			if newStr == "" {
				fmt.Printf("[预览] %s: %s -> %s (删除 '%s')\n", fileType, oldPath, newPath, oldStr)
			} else {
				fmt.Printf("[预览] %s: %s -> %s\n", fileType, oldPath, newPath)
			}
		}

		// 如果需要递归处理且是目录
		if recursive && entry.IsDir() {
			subDirPath := filepath.Join(dirPath, entry.Name())
			if err := previewDirectory(subDirPath, oldStr, newStr, recursive); err != nil {
				fmt.Printf("预览子目录失败 %s: %v\n", subDirPath, err)
			}
		}
	}

	return nil
}

// previewFile 预览单个文件的重命名
func previewFile(filePath, oldStr, newStr string) error {
	dir := filepath.Dir(filePath)
	fileName := filepath.Base(filePath)

	if !strings.Contains(fileName, oldStr) {
		fmt.Printf("[预览] 文件名不包含需要替换的字符串: %s\n", fileName)
		return nil
	}

	newFileName := strings.ReplaceAll(fileName, oldStr, newStr)
	newFilePath := filepath.Join(dir, newFileName)

	if newStr == "" {
		fmt.Printf("[预览] 文件: %s -> %s (删除 '%s')\n", filePath, newFilePath, oldStr)
	} else {
		fmt.Printf("[预览] 文件: %s -> %s\n", filePath, newFilePath)
	}
	return nil
}
