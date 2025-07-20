package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"

	"github.com/cespare/xxhash/v2"
	"github.com/zeebo/blake3"
)

// HashType 定义支持的哈希类型
type HashType string

const (
	MD5    HashType = "md5"
	SHA1   HashType = "sha1"
	SHA256 HashType = "sha256"
	BLAKE3 HashType = "blake3"
	XXHASH HashType = "xxhash"
)

// HashFile 计算文件的哈希值
func HashFile(filePath string, hashType HashType) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var hasher hash.Hash
	switch hashType {
	case MD5:
		hasher = md5.New()
	case SHA1:
		hasher = sha1.New()
	case SHA256:
		hasher = sha256.New()
	case BLAKE3:
		hasher = blake3.New()
	case XXHASH:
		hasher = xxhash.New()
	default:
		return "", fmt.Errorf("不支持的哈希类型: %s", hashType)
	}

	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// HashString 计算字符串的哈希值
func HashString(input string, hashType HashType) (string, error) {
	var hash []byte
	switch hashType {
	case MD5:
		h := md5.Sum([]byte(input))
		hash = h[:]
	case SHA1:
		h := sha1.Sum([]byte(input))
		hash = h[:]
	case SHA256:
		h := sha256.Sum256([]byte(input))
		hash = h[:]
	case BLAKE3:
		h := blake3.Sum256([]byte(input))
		hash = h[:]
	case XXHASH:
		h := xxhash.Sum64([]byte(input))
		hashBytes := make([]byte, 8)
		for i := 0; i < 8; i++ {
			hashBytes[i] = byte(h >> (8 * (7 - i)))
		}
		hash = hashBytes
	default:
		return "", fmt.Errorf("不支持的哈希类型: %s", hashType)
	}

	return hex.EncodeToString(hash), nil
}
