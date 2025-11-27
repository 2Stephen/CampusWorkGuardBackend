package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Argon2id 参数（可按你服务器性能调整）
var (
	argonTime    uint32 = 1         // 迭代次数
	argonMemory  uint32 = 64 * 1024 // 64 MB
	argonThreads uint8  = 4
	argonKeyLen  uint32 = 32
	saltLen             = 16
)

// 可选 pepper：应用层全局密钥（提高安全性），建议保存在环境变量或 KMS。
// var pepper = os.Getenv("PASSWORD_PEPPER")

// GenerateSalt 生成安全随机 salt
func GenerateSalt(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// HashPassword 使用 Argon2id 对给定密码进行哈希，返回自描述字符串
// 格式:
// $argon2id$v=19$m=65536,t=1,p=4$<base64(salt)>$<base64(hash)>
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password empty")
	}

	// 可选：password = password + pepper

	salt, err := GenerateSalt(saltLen)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, argonTime, argonMemory, argonThreads, argonKeyLen)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encoded := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		19, argonMemory, argonTime, argonThreads, b64Salt, b64Hash)

	return encoded, nil
}

// VerifyPassword 比较明文密码与存储的 encodedHash 是否匹配
func VerifyPassword(password, encodedHash string) (bool, error) {
	if password == "" || encodedHash == "" {
		return false, errors.New("empty password or hash")
	}

	parts := strings.Split(encodedHash, "$")
	// parts: ["", "argon2id", "v=19", "m=...,t=...,p=...", "<salt>", "<hash>"]
	if len(parts) != 6 {
		return false, errors.New("invalid hash format")
	}

	// 解析参数
	params := parts[3]
	var memory uint32
	var time uint32
	var threads uint8

	for _, kv := range strings.Split(params, ",") {
		if strings.HasPrefix(kv, "m=") {
			v := strings.TrimPrefix(kv, "m=")
			m64, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return false, err
			}
			memory = uint32(m64)
		} else if strings.HasPrefix(kv, "t=") {
			v := strings.TrimPrefix(kv, "t=")
			t64, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return false, err
			}
			time = uint32(t64)
		} else if strings.HasPrefix(kv, "p=") {
			v := strings.TrimPrefix(kv, "p=")
			p64, err := strconv.ParseUint(v, 10, 8)
			if err != nil {
				return false, err
			}
			threads = uint8(p64)
		}
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}
	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	// 可选：password = password + pepper

	computed := argon2.IDKey([]byte(password), salt, time, memory, threads, uint32(len(expectedHash)))

	// 常量时间比较，防止时序攻击
	if subtle.ConstantTimeCompare(computed, expectedHash) == 1 {
		return true, nil
	}
	return false, nil
}
