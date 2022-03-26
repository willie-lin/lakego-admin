package hash

import (
    "crypto"
    "crypto/hmac"
    "encoding/hex"
)

// HmacMd5 签名
func HmacMd5(message string, secret string) string {
    return HmacSign(crypto.MD5, message, secret)
}

// HmacSHA1 签名
func HmacSHA1(message string, secret string) string {
    return HmacSign(crypto.SHA1, message, secret)
}

// HmacSha224 签名
func HmacSha224(message string, secret string) string {
    return HmacSign(crypto.SHA224, message, secret)
}

// HmacSha256 签名
func HmacSha256(message string, secret string) string {
    return HmacSign(crypto.SHA256, message, secret)
}

// HmacSha384 签名
func HmacSha384(message string, secret string) string {
    return HmacSign(crypto.SHA384, message, secret)
}

// HmacSha512 签名
func HmacSha512(message string, secret string) string {
    return HmacSign(crypto.SHA512, message, secret)
}

// 签名
func HmacSign(hash crypto.Hash, message string, secret string) string {
    hasher := hmac.New(hash.New, []byte(secret))
    hasher.Write([]byte(message))
    return hex.EncodeToString(hasher.Sum(nil))
}

