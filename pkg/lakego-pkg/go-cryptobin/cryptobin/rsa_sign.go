package cryptobin

import (
    "crypto/rsa"
    "crypto/rand"
)

// 私钥签名
func (this Rsa) Sign() Rsa {
    newHash := NewHash()

    hasher := newHash.GetCryptoHash(this.signHash)
    hashData := newHash.DataCryptoHash(this.signHash, this.data)

    this.paredData, this.Error = rsa.SignPKCS1v15(rand.Reader, this.privateKey, hasher, hashData)

    return this
}

// 公钥验证
// 使用原始数据[data]对比签名后数据
func (this Rsa) Very(data []byte) Rsa {
    newHash := NewHash()

    hasher := newHash.GetCryptoHash(this.signHash)
    hashData := newHash.DataCryptoHash(this.signHash, data)

    err := rsa.VerifyPKCS1v15(this.publicKey, hasher, hashData, this.data)
    if err != nil {
        this.veryed = false
        this.Error = err

        return this
    }

    this.veryed = true

    return this
}
