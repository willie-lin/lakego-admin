package dh

import (
    "errors"
    "crypto/rand"
    "encoding/pem"

    "github.com/deatil/go-cryptobin/dh/dh"
    cryptobin_pkcs8 "github.com/deatil/go-cryptobin/pkcs8"
    cryptobin_pkcs8s "github.com/deatil/go-cryptobin/pkcs8s"
)

type (
    // 配置
    Opts       = cryptobin_pkcs8.Opts
    // PBKDF2 配置
    PBKDF2Opts = cryptobin_pkcs8.PBKDF2Opts
    // Scrypt 配置
    ScryptOpts = cryptobin_pkcs8.ScryptOpts
)

var (
    // 获取 Cipher 类型
    GetCipherFromName = cryptobin_pkcs8.GetCipherFromName
    // 获取 hash 类型
    GetHashFromName   = cryptobin_pkcs8.GetHashFromName
)

// 生成私钥 pem 数据
// 使用:
// obj := New().GenerateKey("P2048")
// priKey := obj.CreatePrivateKey().ToKeyString()
func (this Dh) CreatePrivateKey() Dh {
    if this.privateKey == nil {
        err := errors.New("Dh: privateKey error.")
        return this.AppendError(err)
    }

    privateKey, err := dh.MarshalPrivateKey(this.privateKey)
    if err != nil {
        return this.AppendError(err)
    }

    privateBlock := &pem.Block{
        Type:  "PRIVATE KEY",
        Bytes: privateKey,
    }

    this.keyData = pem.EncodeToMemory(privateBlock)

    return this
}

// 生成 PKCS8 私钥带密码 pem 数据
// CreatePrivateKeyWithPassword("123", "AES256CBC", "SHA256")
func (this Dh) CreatePrivateKeyWithPassword(password string, opts ...any) Dh {
    if this.privateKey == nil {
        err := errors.New("Dh: privateKey error.")
        return this.AppendError(err)
    }

    opt, err := cryptobin_pkcs8s.ParseOpts(opts...)
    if err != nil {
        return this.AppendError(err)
    }

    // 生成私钥
    privateKey, err := dh.MarshalPrivateKey(this.privateKey)
    if err != nil {
        return this.AppendError(err)
    }

    // 生成加密数据
    privateBlock, err := cryptobin_pkcs8s.EncryptPEMBlock(
        rand.Reader,
        "ENCRYPTED PRIVATE KEY",
        privateKey,
        []byte(password),
        opt,
    )
    if err != nil {
        return this.AppendError(err)
    }

    this.keyData = pem.EncodeToMemory(privateBlock)

    return this
}

// 生成公钥 pem 数据
func (this Dh) CreatePublicKey() Dh {
    if this.publicKey == nil {
        err := errors.New("Dh: publicKey error.")
        return this.AppendError(err)
    }

    publicKeyBytes, err := dh.MarshalPublicKey(this.publicKey)
    if err != nil {
        return this.AppendError(err)
    }

    publicBlock := &pem.Block{
        Type:  "PUBLIC KEY",
        Bytes: publicKeyBytes,
    }

    this.keyData = pem.EncodeToMemory(publicBlock)

    return this
}

// 根据公钥和私钥生成密钥
func (this Dh) CreateSecretKey() Dh {
    if this.privateKey == nil {
        err := errors.New("Dh: privateKey error.")
        return this.AppendError(err)
    }

    if this.publicKey == nil {
        err := errors.New("Dh: publicKey error.")
        return this.AppendError(err)
    }

    this.secretData = dh.ComputeSecret(this.privateKey, this.publicKey)

    return this
}
