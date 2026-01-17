package internal

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// 使用Argon2算法对密码进行哈希处理
type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

// 创建盐值
func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// 对输入字符串进行哈希处理并返回哈希值
func hashString(password string) string {
	p := &params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
	// 生成对应的盐
	salt, _ := generateRandomBytes(p.saltLength)
	hash := argon2.IDKey([]byte(password),
		salt,
		p.iterations,
		p.memory,
		p.parallelism,
		p.keyLength)
	// 储存盐和哈希值
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	// 返回整合字符串
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)
	return encodedHash
}

// 分割哈希字符串
func decodeHash(encodedHash string) (p *params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	// 判断字段总数
	if len(vals) != 6 {
		return nil, nil, nil, errors.New("invalid hash format")
	}
	// 先验证版本是否匹配
	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errors.New("incompatible argon2 version")
	}
	// 解析参数
	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}
	// 解码盐值
	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))
	// 解码哈希值
	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))
	return p, salt, hash, nil
}

// 验证密码是否匹配
func decodeHashString(password, encodedHash string) (match bool, err error) {
	// 解析整合字符串
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}
	// 使用相同的参数对输入密码进行哈希处理
	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)
	// 使用恒定时间比较防止时序攻击
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		// 密码匹配
		return true, nil
	}
	return false, nil
}
