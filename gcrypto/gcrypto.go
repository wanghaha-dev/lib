package gcrypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

// MD5 MD5加密
func MD5(data string) string {
	bytes := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", bytes)
}

// MD5Salt MD5加盐
func MD5Salt(salt string, data string) string {
	hash := md5.New()
	hash.Write([]byte(salt))
	hash.Write([]byte(data))
	sum := hash.Sum(nil)
	return hex.EncodeToString(sum)
}

// SHA1 sha1加密
func SHA1(data string) string {
	hash := sha1.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// SHA1Salt sha1加密-加盐
func SHA1Salt(salt string, data string) string {
	hash := sha1.New()
	hash.Write([]byte(salt))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// SHA256 sha256加密
func SHA256(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// SHA256Salt sha256加密-加盐
func SHA256Salt(salt string, data string) string {
	hash := sha256.New()
	hash.Write([]byte(salt))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// SHA512 sha512加密
func SHA512(data string) string {
	hash := sha512.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// SHA512Salt sha512加密-加盐
func SHA512Salt(salt string, data string) string {
	hash := sha512.New()
	hash.Write([]byte(salt))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
