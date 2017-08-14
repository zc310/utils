package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"hash/adler32"
	"hash/crc32"

	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

var list = map[string]func(b []byte) string{
	"md4":       MD4,
	"md5":       MD5,
	"sha1":      SHA1,
	"sha224":    SHA224,
	"sha256":    SHA256,
	"sha512":    SHA512,
	"sha384":    SHA384,
	"sha3_224":  Sha3_224,
	"sha3_256":  Sha3_256,
	"sha3_384":  Sha3_384,
	"sha3_512":  Sha3_512,
	"ripemd160": Ripemd160,
	"adler32":   Adler32,
	"crc32":     Crc32,
}

// Get get hash func
func Get(name string) (f func(b []byte) string, ok bool) {
	f, ok = list[name]
	return
}

// MD4 md4
func MD4(a []byte) string {
	h := md4.New()
	h.Write(a)
	return hex.EncodeToString(h.Sum(nil))
}

// MD5 md5
func MD5(a []byte) string {
	h := md5.New()
	h.Write(a)
	return hex.EncodeToString(h.Sum(nil))
}

// Md5Object md5(json(interface))
func Md5Object(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return MD5(b), nil
}

// SHA1 sha1
func SHA1(a []byte) string {
	h := sha1.New()
	h.Write(a)
	return hex.EncodeToString(h.Sum(nil))
}

// SHA224 sha224
func SHA224(a []byte) string {
	h := sha256.New224()
	h.Write(a)
	return hex.EncodeToString(h.Sum(nil))
}

// SHA256 sha256
func SHA256(a []byte) string {
	h := sha256.New()
	h.Write(a)
	return hex.EncodeToString(h.Sum(nil))
}

// SHA384 sha384
func SHA384(a []byte) string {
	h := sha512.New384()
	h.Write(a)
	return hex.EncodeToString(h.Sum(nil))
}

// SHA512 sha512
func SHA512(a []byte) string {
	h := sha512.New()
	h.Write(a)
	return hex.EncodeToString(h.Sum(nil))
}

// Sha3_224 sha3_224
func Sha3_224(a []byte) string {
	h := sha3.New224()
	h.Write(a)
	return hex.EncodeToString(h.Sum(nil))
}

// Sha3_256 sha3_256
func Sha3_256(a []byte) string {
	h := sha3.New256()
	h.Write(a)
	return hex.EncodeToString(h.Sum(nil))
}

// Sha3_384 sha3_384
func Sha3_384(a []byte) string {
	h := sha3.New384()
	h.Write(a)
	return hex.EncodeToString(h.Sum(nil))
}

// Sha3_512 sha3_512
func Sha3_512(a []byte) string {
	h := sha3.New512()
	h.Write(a)
	return hex.EncodeToString(h.Sum(nil))
}

// Ripemd160 Ripemd160
func Ripemd160(a []byte) string {
	h := ripemd160.New()
	h.Write(a)
	return hex.EncodeToString(h.Sum(nil))
}

// Adler32 Adler32
func Adler32(a []byte) string {
	h := adler32.New()
	h.Write(a)
	return hex.EncodeToString(h.Sum(nil))
}

// Crc32 Crc32
func Crc32(a []byte) string {
	h := crc32.NewIEEE()
	h.Write(a)
	return hex.EncodeToString(h.Sum(nil))
}
