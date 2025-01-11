package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// MD5 : 128位的哈希算法, 表示32位的十六进制字符串
// 从哈希值难以逆推出原始输入, 防止文件原始名称暴露, 用于数据完整性校验, 标识数据唯一性, 但是碰撞容易被发现


func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}


// SHA256 : 256位的哈希算法, 64字符的十六进制字符串

func EncodeSHA256(value string) string {
    hash := sha256.Sum256([]byte(value))
    return hex.EncodeToString(hash[:])
}