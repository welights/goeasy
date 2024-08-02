package checksumutil

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

type Checksum struct {
	Md5    string `json:"md5,omitempty"`
	Sha1   string `json:"sha1,omitempty"`
	Sha256 string `json:"sha256,omitempty"`
}

func Md5(data []byte) string {
	h := md5.Sum(data)
	return hex.EncodeToString(h[:])
}

func Md5Str(data string) string {
	return Md5([]byte(data))
}

func Sha1(data []byte) string {
	h := sha1.Sum(data)
	return hex.EncodeToString(h[:])
}

func Sha1Str(data string) string {
	return Sha1([]byte(data))
}

func Sha256(data []byte) string {
	h := sha256.Sum256(data)
	return hex.EncodeToString(h[:])
}

func Sha256Str(data string) string {
	return Sha256([]byte(data))
}

func Checksums(data []byte) Checksum {
	md5Hasher := md5.New()
	sha1Hasher := sha1.New()
	sha256Hasher := sha256.New()
	hasher := io.MultiWriter(md5Hasher, sha1Hasher, sha256Hasher)
	_, _ = hasher.Write(data)
	return Checksum{
		Md5:    fmt.Sprintf("%x", md5Hasher.Sum(nil)),
		Sha1:   fmt.Sprintf("%x", sha1Hasher.Sum(nil)),
		Sha256: fmt.Sprintf("%x", sha256Hasher.Sum(nil)),
	}
}

func ChecksumsWithReader(r io.Reader) Checksum {
	md5Hasher := md5.New()
	sha1Hasher := sha1.New()
	sha256Hasher := sha256.New()
	hasher := io.MultiWriter(md5Hasher, sha1Hasher, sha256Hasher)
	_, _ = io.Copy(hasher, r)
	return Checksum{
		Md5:    fmt.Sprintf("%x", md5Hasher.Sum(nil)),
		Sha1:   fmt.Sprintf("%x", sha1Hasher.Sum(nil)),
		Sha256: fmt.Sprintf("%x", sha256Hasher.Sum(nil)),
	}
}
