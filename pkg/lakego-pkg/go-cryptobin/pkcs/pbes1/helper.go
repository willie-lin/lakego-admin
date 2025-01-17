package pbes1

// PEMCipher 列表
var PEMCipherMap = map[string]Cipher{
    // pkcs12 模式
    "SHA1And3DES":    SHA1And3DES,
    "SHA1And2DES":    SHA1And2DES,
    "SHA1AndRC2_128": SHA1AndRC2_128,
    "SHA1AndRC2_40":  SHA1AndRC2_40,
    "SHA1AndRC4_128": SHA1AndRC4_128,
    "SHA1AndRC4_40":  SHA1AndRC4_40,

    // pkcs5-v1.5 模式
    "MD2AndDES":      MD2AndDES,
    "MD2AndRC2_64":   MD2AndRC2_64,
    "MD5AndDES":      MD5AndDES,
    "MD5AndRC2_64":   MD5AndRC2_64,
    "SHA1AndDES":     SHA1AndDES,
    "SHA1AndRC2_64":  SHA1AndRC2_64,
}

// 获取 Cipher 类型
func GetCipherFromName(name string) Cipher {
    if data, ok := PEMCipherMap[name]; ok {
        return data
    }

    return PEMCipherMap["MD5AndDES"]
}

// 检测 Cipher 类型
func CheckCipherFromName(name string) bool {
    if _, ok := PEMCipherMap[name]; ok {
        return true
    }

    return false
}
