package crypto

import (
    "strconv"
)

// 类型
var TypeMultiple = NewTypeSet[Multiple, string](maxMultiple)

// 模式
var TypeMode = NewTypeSet[Mode, string](maxMode)

// 补码
var TypePadding = NewTypeSet[Padding, string](maxPadding)

// 构造函数
func NewTypeSet[N TypeName, D any](max N) *TypeSet[N, D] {
    return &TypeSet[N, D]{
        max:   max,
        names: NewDataSet[N, D](),
    }
}

// 名称类型
type TypeName interface {
    ~uint | ~int
}

// 类型数据
type TypeSet[N TypeName, D any] struct {
    // 最大值
    max N

    // 数据
    names *DataSet[N, D]
}

// 生成新序列
func (this *TypeSet[N, D]) Generate() N {
    old := this.max
    this.max++

    return old
}

// 设置
func (this *TypeSet[N, D]) Names() *DataSet[N, D] {
    return this.names
}

// ===================

// 加密类型
type Multiple uint

func (this Multiple) String() string {
    switch this {
        case Aes:
            return "Aes"
        case Des:
            return "Des"
        case TwoDes:
            return "TwoDes"
        case TripleDes:
            return "TripleDes"
        case Twofish:
            return "Twofish"
        case Blowfish:
            return "Blowfish"
        case Tea:
            return "Tea"
        case Xtea:
            return "Xtea"
        case Cast5:
            return "Cast5"
        case RC2:
            return "RC2"
        case RC4:
            return "RC4"
        case RC4MD5:
            return "RC4MD5"
        case RC5:
            return "RC5"
        case Idea:
            return "Idea"
        case SM4:
            return "SM4"
        case Chacha20:
            return "Chacha20"
        case Chacha20poly1305:
            return "Chacha20poly1305"
        case Chacha20poly1305X:
            return "Chacha20poly1305X"
        case Xts:
            return "Xts"
        case Salsa20:
            return "Salsa20"
        case Seed:
            return "Seed"
        case Aria:
            return "Aria"
        case Camellia:
            return "Camellia"
        default:
            if TypeMultiple.Names().Has(this) {
                return (TypeMultiple.Names().Get(this))()
            }

            return "unknown multiple value " + strconv.Itoa(int(this))
    }
}

const (
    Aes Multiple = 1 + iota
    Des
    TwoDes
    TripleDes
    Twofish
    Blowfish
    Tea
    Xtea
    Cast5
    RC2
    RC4
    RC4MD5
    RC5
    Idea
    SM4
    Chacha20
    Chacha20poly1305
    Chacha20poly1305X
    Xts
    Salsa20
    Seed
    Aria
    Camellia
    maxMultiple
)

// ===================

// 加密模式
type Mode uint

func (this Mode) String() string {
    switch this {
        case ECB:
            return "ECB"
        case CBC:
            return "CBC"
        case PCBC:
            return "PCBC"
        case CFB:
            return "CFB"
        case CFB1:
            return "CFB1"
        case CFB8:
            return "CFB8"
        case CFB16:
            return "CFB16"
        case CFB32:
            return "CFB32"
        case CFB64:
            return "CFB64"
        case CFB128:
            return "CFB128"
        case OFB:
            return "OFB"
        case OFB8:
            return "OFB8"
        case CTR:
            return "CTR"
        case GCM:
            return "GCM"
        case CCM:
            return "CCM"
        default:
            if TypeMode.Names().Has(this) {
                return (TypeMode.Names().Get(this))()
            }

            return "unknown mode value " + strconv.Itoa(int(this))
    }
}

const (
    ECB  Mode = 1 + iota
    CBC
    PCBC
    CFB
    CFB1
    CFB8
    CFB16
    CFB32
    CFB64
    CFB128
    OFB
    OFB8
    CTR
    GCM
    CCM
    maxMode
)

// ===================

// 补码类型
type Padding uint

func (this Padding) String() string {
    switch this {
        case NoPadding:
            return "NoPadding"
        case ZeroPadding:
            return "ZeroPadding"
        case PKCS5Padding:
            return "PKCS5Padding"
        case PKCS7Padding:
            return "PKCS7Padding"
        case X923Padding:
            return "X923Padding"
        case ISO10126Padding:
            return "ISO10126Padding"
        case ISO7816_4Padding:
            return "ISO7816_4Padding"
        case ISO97971Padding:
            return "ISO97971Padding"
        case TBCPadding:
            return "TBCPadding"
        case PKCS1Padding:
            return "PKCS1Padding"
        default:
            if TypePadding.Names().Has(this) {
                return (TypePadding.Names().Get(this))()
            }

            return "unknown padding value " + strconv.Itoa(int(this))
    }
}

const (
    NoPadding Padding = 1 + iota
    ZeroPadding
    PKCS5Padding
    PKCS7Padding
    X923Padding
    ISO10126Padding
    ISO7816_4Padding
    ISO97971Padding
    TBCPadding
    PKCS1Padding
    maxPadding
)
