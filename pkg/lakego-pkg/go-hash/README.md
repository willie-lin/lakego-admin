## go-hash


### 项目介绍

*  常用的摘要算法
*  算法包括: (MD2/MD4/MD5/Hmac/Ripemd160/SHA1/SHA3/SHA256/SM3)


### 下载安装

~~~go
go get -u github.com/deatil/go-hash
~~~


### 使用

~~~go
package main

import (
    "fmt"
    "github.com/deatil/go-hash/hash"
)

func main() {
    // MD5 获取摘要
    md5Data := hash.
        FromString("useData"). // 输入数据, 需要获取摘要的数据
        MD5().                 // 摘要方式
        ToHexString()          // 输出数据, 返回数据类型

    fmt.Println("MD5 结果：", md5Data)

    // NewMD5 获取摘要
    md5Data2 := hash.
        Hashing().
        NewMD5().                 // 摘要方式
        Write([]byte("useData")). // 需要获取摘要的数据
        Sum(nil).                 // 计算
        ToHexString()             // 输出数据, 返回数据类型

    fmt.Println("MD5 结果2：", md5Data2)

    // =============

    // HmacSHA1 获取摘要
    hmacMD5Data := hash.
        FromString("useData").                    // 输入数据, 需要获取摘要的数据
        Hmac(hash.HmacSHA1.New, []byte("pass")).  // 摘要方式
        ToHexString()                             // 输出数据, 返回数据类型

    fmt.Println("HmacSHA1 结果：", hmacMD5Data)

    // HmacSHA1 获取摘要
    hmacMD5Data2 := hash.
        Hashing().                                   // 输入数据, 为空
        NewHmac(hash.HmacSHA1.New, []byte("pass")).  // 摘要方式
        Write([]byte("useData")).                    // 需要获取摘要的数据
        Sum(nil).                                    // 计算
        ToHexString()                                // 输出数据, 返回数据类型

    fmt.Println("HmacSHA1 结果2：", hmacMD5Data2)
}

~~~


### 输入输出数据

*  输入数据:
`FromBytes(data []byte)`, `FromString(data string)`, `FromBase64String(data string)`, `FromHexString(data string)`, `FromReader(reader io.Reader)`
*  输出数据:
`String() string`, `ToBytes() []byte`, `ToString() string`, `ToBase64String() string`, `ToHexString() string`, `ToReader() io.Reader`


### 常用算法

*  直接使用:
`Adler32()`,
`Blake2b_256()`, `Blake2b_384()`, `Blake2b_512()`, `Blake2s_256()`, `Blake2s_128()`,
`CRC16_X25()`, `CRC16_Modbus()`,
`CRC32_IEEE()`, `CRC32_Castagnoli()`, `CRC32_Koopman()`,
`CRC64_ISO()`, `CRC64_ECMA()`,
`Fnv32()`, `Fnv32a()`, `Fnv64()`, `Fnv64a()`, `Fnv128()`, `Fnv128a()`,
`Hmac(h func() hash.Hash, secret []byte)`, 
`Keccak256()`, `Keccak512()`,
`Maphash()`,
`MD2()`, `MD4()`, `MD5()`, `MD5SHA1()`,
`Murmur32()`, `Murmur32WithSeed(seed uint32)`,
`Murmur64()`, `Murmur64WithSeed(seed uint32)`,
`Murmur128()`, `Murmur128WithSeed(seed uint32)`,
`Ripemd160()`,
`SHA1()`, `SHA224()`, `SHA256()`, `SHA384()`, `SHA512()`, `SHA512_224()`, `SHA512_256()`,
`SHA3_224()`, `SHA3_256()`, `SHA3_384()`, `SHA3_512()`,
`Shake128()`, `Shake256()`,
`SM3()`

*  New 使用:
`NewAdler32()`,
`NewBlake2b_256(key []byte)`, `NewBlake2b_384(key []byte)`, `NewBlake2b_512(key []byte)`, `NewBlake2s_256(key []byte)`, `NewBlake2s_128(key []byte)`,
`NewCRC32_IEEE()`, `NewCRC32_Castagnoli()`, `NewCRC32_Koopman()`,
`NewCRC64_ISO()`, `NewCRC64_ECMA()`,
`NewFnv32()`, `NewFnv32a()`, `NewFnv64()`, `NewFnv64a()`, `NewFnv128()`, `NewFnv128a()`,
`NewHmac(h func() hash.Hash, secret []byte)`,
`NewKeccak256()`, `NewKeccak512()`,
`NewMaphash()`,
`NewMD2()`, `NewMD4()`, `NewMD5()`,
`NewMurmur32()`, `NewMurmur32WithSeed(seed uint32)`,
`NewMurmur64()`, `NewMurmur64WithSeed(seed uint32)`,
`NewMurmur128()`, `NewMurmur128WithSeed(seed uint32)`,
`NewRipemd160()`,
`NewSHA1()`, `NewSHA224()`, `NewSHA256()`, `NewSHA384()`, `NewSHA512()`, `NewSHA512_224()`, `NewSHA512_256()`,
`NewSHA3_224()`, `NewSHA3_256()`, `NewSHA3_384()`, `NewSHA3_512()`,
`NewSM3()`

*  Hmac 可用 hash:
`HmacADLER32`, `HmacMD2`, `HmacMD4`, `HmacMD5`, `HmacSHA1`, `HmacSHA224`, `HmacSHA256`, `HmacSHA384`, `HmacSHA512`, `HmacSHA512_224`, `HmacSHA512_256`, `HmacRIPEMD160`, `HmacSHA3_224`, `HmacSHA3_256`, `HmacSHA3_384`, `HmacSHA3_512`


### 开源协议

*  本软件包遵循 `Apache2` 开源协议发布，在保留本软件包版权的情况下提供个人及商业免费使用。


### 版权

*  本软件包所属版权归 deatil(https://github.com/deatil) 所有。
