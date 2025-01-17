package rsa

import (
    "errors"
    "crypto/rand"
    "crypto/x509"
    "encoding/pem"

    cryptobin_tool "github.com/deatil/go-cryptobin/tool"
    cryptobin_rsa "github.com/deatil/go-cryptobin/rsa"
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

// 生成私钥 pem 数据, PKCS1 别名
// 使用:
// obj := New().GenerateKey(2048)
// priKey := obj.CreatePrivateKey().ToKeyString()
func (this Rsa) CreatePrivateKey() Rsa {
    return this.CreatePKCS1PrivateKey()
}

// 生成私钥带密码 pem 数据, PKCS1 别名
func (this Rsa) CreatePrivateKeyWithPassword(password string, opts ...string) Rsa {
    return this.CreatePKCS1PrivateKeyWithPassword(password, opts...)
}

// 生成公钥 pem 数据
func (this Rsa) CreatePublicKey() Rsa {
    return this.CreatePKCS1PublicKey()
}

// ====================

// 生成 PKCS1 私钥
func (this Rsa) CreatePKCS1PrivateKey() Rsa {
    if this.privateKey == nil {
        err := errors.New("Rsa: privateKey error.")
        return this.AppendError(err)
    }

    x509PrivateKey := x509.MarshalPKCS1PrivateKey(this.privateKey)

    privateBlock := &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: x509PrivateKey,
    }

    this.keyData = pem.EncodeToMemory(privateBlock)

    return this
}

// 生成 PKCS1 私钥带密码 pem 数据
// CreatePKCS1PrivateKeyWithPassword("123", "AES256CBC")
// PEMCipher: DESCBC | DESEDE3CBC | AES128CBC | AES192CBC | AES256CBC
func (this Rsa) CreatePKCS1PrivateKeyWithPassword(password string, opts ...string) Rsa {
    if this.privateKey == nil {
        err := errors.New("Rsa: privateKey error.")
        return this.AppendError(err)
    }

    opt := "AES256CBC"
    if len(opts) > 0 {
        opt = opts[0]
    }

    // 加密方式
    cipher, err := cryptobin_tool.GetPEMCipher(opt)
    if err != nil {
        err := errors.New("Rsa: PEMCipher not exists.")
        return this.AppendError(err)
    }

    // 生成私钥
    x509PrivateKey := x509.MarshalPKCS1PrivateKey(this.privateKey)

    // 生成加密数据
    privateBlock, err := x509.EncryptPEMBlock(
        rand.Reader,
        "RSA PRIVATE KEY",
        x509PrivateKey,
        []byte(password),
        cipher,
    )
    if err != nil {
        return this.AppendError(err)
    }

    this.keyData = pem.EncodeToMemory(privateBlock)

    return this
}

// 生成 pcks1 公钥 pem 数据
func (this Rsa) CreatePKCS1PublicKey() Rsa {
    if this.publicKey == nil {
        err := errors.New("Rsa: publicKey error.")
        return this.AppendError(err)
    }

    x509PublicKey := x509.MarshalPKCS1PublicKey(this.publicKey)

    publicBlock := &pem.Block{
        Type:  "RSA PUBLIC KEY",
        Bytes: x509PublicKey,
    }

    this.keyData = pem.EncodeToMemory(publicBlock)

    return this
}

// ====================

// 生成 PKCS8 私钥 pem 数据
func (this Rsa) CreatePKCS8PrivateKey() Rsa {
    if this.privateKey == nil {
        err := errors.New("Rsa: privateKey error.")
        return this.AppendError(err)
    }

    x509PrivateKey, err := x509.MarshalPKCS8PrivateKey(this.privateKey)
    if err != nil {
        return this.AppendError(err)
    }

    privateBlock := &pem.Block{
        Type:  "PRIVATE KEY",
        Bytes: x509PrivateKey,
    }

    this.keyData = pem.EncodeToMemory(privateBlock)

    return this
}

// 生成 PKCS8 私钥带密码 pem 数据
// CreatePKCS8PrivateKeyWithPassword("123", "AES256CBC", "SHA256")
func (this Rsa) CreatePKCS8PrivateKeyWithPassword(password string, opts ...any) Rsa {
    if this.privateKey == nil {
        err := errors.New("Rsa: privateKey error.")
        return this.AppendError(err)
    }

    opt, err := cryptobin_pkcs8s.ParseOpts(opts...)
    if err != nil {
        return this.AppendError(err)
    }

    // 生成私钥
    x509PrivateKey, err := x509.MarshalPKCS8PrivateKey(this.privateKey)
    if err != nil {
        return this.AppendError(err)
    }

    // 生成加密数据
    privateBlock, err := cryptobin_pkcs8s.EncryptPEMBlock(
        rand.Reader,
        "ENCRYPTED PRIVATE KEY",
        x509PrivateKey,
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
func (this Rsa) CreatePKCS8PublicKey() Rsa {
    if this.publicKey == nil {
        err := errors.New("Rsa: publicKey error.")
        return this.AppendError(err)
    }

    x509PublicKey, err := x509.MarshalPKIXPublicKey(this.publicKey)
    if err != nil {
        return this.AppendError(err)
    }

    publicBlock := &pem.Block{
        Type:  "PUBLIC KEY",
        Bytes: x509PublicKey,
    }

    this.keyData = pem.EncodeToMemory(publicBlock)

    return this
}

// ====================

// 生成私钥 xml 数据
func (this Rsa) CreateXMLPrivateKey() Rsa {
    if this.privateKey == nil {
        err := errors.New("Rsa: privateKey error.")
        return this.AppendError(err)
    }

    xmlPrivateKey, err := cryptobin_rsa.MarshalXMLPrivateKey(this.privateKey)
    if err != nil {
        return this.AppendError(err)
    }

    this.keyData = xmlPrivateKey

    return this
}

// 生成公钥 xml 数据
func (this Rsa) CreateXMLPublicKey() Rsa {
    if this.publicKey == nil {
        err := errors.New("Rsa: publicKey error.")
        return this.AppendError(err)
    }

    xmlPublicKey, err := cryptobin_rsa.MarshalXMLPublicKey(this.publicKey)
    if err != nil {
        return this.AppendError(err)
    }

    this.keyData = xmlPublicKey

    return this
}
