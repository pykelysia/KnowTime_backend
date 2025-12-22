package internal

import (
	"crypto/sha256"
	"fmt"
)

// HashString 使用SHA256算法对字符串进行哈希处理
func hashString(input string) string {
	hash := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", hash)
}
