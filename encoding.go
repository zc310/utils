package utils

import (
	"bytes"
	"compress/flate"
	"compress/zlib"
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"html"
	"io"
	"io/ioutil"
	"net/url"
	"strings"
)

func Base64Enc(a []byte) string {
	return base64.StdEncoding.EncodeToString(a)
}
func Base64Dec(a string) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(a)
	return b, err
}

func HexEnc(a []byte) string {
	return hex.EncodeToString(a)
}
func HexDec(a string) ([]byte, error) {
	b, err := hex.DecodeString(a)
	return b, err
}
func URLEnc(v string) string {
	return url.QueryEscape(v)
}
func URLDec(v string) (string, error) {
	return url.QueryUnescape(v)
}

func Base64ZlibEnc(a []byte) string {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(a)
	w.Close()
	return base64.StdEncoding.EncodeToString(b.Bytes())
}
func Base64ZlibDec(a string) ([]byte, error) {
	bv, err := base64.StdEncoding.DecodeString(a)
	if err != nil {
		return []byte{}, err
	}
	r, err := zlib.NewReader(bytes.NewReader(bv))
	if err != nil {
		return []byte{}, err
	}
	var ba bytes.Buffer
	ba.ReadFrom(r)
	r.Close()
	return ba.Bytes(), nil

}
func Base64FlateDec(a string) ([]byte, error) {
	bv, err := base64.StdEncoding.DecodeString(a)
	if err != nil {
		return []byte{}, err
	}
	r := flate.NewReader(bytes.NewReader(bv))
	var ba bytes.Buffer
	ba.ReadFrom(r)
	r.Close()
	return ba.Bytes(), nil
}

func HTMLDec(v string) (string, error) {
	return html.UnescapeString(v), nil
}
func HTMLEnc(v string) string {
	return html.EscapeString(v)
}

func Base64TwoZlibEnc(a []byte) string {
	var b, b2 bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(a)
	w.Close()

	w = zlib.NewWriter(&b2)
	w.Write(b.Bytes())
	w.Close()
	return base64.StdEncoding.EncodeToString(b2.Bytes())
}
func Base64TwoZlibDec(v string) ([]byte, error) {
	bv, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return []byte{}, err
	}
	r, err := zlib.NewReader(bytes.NewReader(bv))
	if err != nil {
		return []byte{}, err
	}
	var ba bytes.Buffer
	ba.ReadFrom(r)
	r.Close()
	r, err = zlib.NewReader(bytes.NewReader(ba.Bytes()))
	if err != nil {
		return []byte{}, err
	}
	ba.Reset()
	ba.ReadFrom(r)
	r.Close()
	return ba.Bytes(), nil
}

func UTF16EncStr(v string) string {
	b, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(v), unicode.UTF16(unicode.LittleEndian, unicode.ExpectBOM).NewEncoder()))
	if err != nil {
		return ""
	}
	return string(b)
}
func UTF16DecStr(v string) (string, error) {
	b, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(v), unicode.UTF16(unicode.LittleEndian, unicode.ExpectBOM).NewDecoder()))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func GB2312EncStr(v string) string {
	b, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(v), simplifiedchinese.HZGB2312.NewEncoder()))
	if err != nil {
		return ""
	}
	return string(b)
}
func GB2312DecStr(v string) (string, error) {
	b, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(v), simplifiedchinese.HZGB2312.NewDecoder()))
	if err != nil {
		return "", err
	}
	return string(b), nil
}
func GB18030EncStr(v string) string {
	b, err := ToGB18030Str(v)
	if err != nil {
		return ""
	}
	return string(b)
}
func GB18030DecStr(v string) (string, error) {
	tr := GB18030Dec(strings.NewReader(v))
	b, err := ioutil.ReadAll(tr)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func GBKEncStr(v string) string {
	b, err := ToGBKStr(v)
	if err != nil {
		return ""
	}
	return string(b)
}
func GBKDecStr(v string) (string, error) {
	tr := GbkDec(strings.NewReader(v))
	b, err := ioutil.ReadAll(tr)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func GbkEnc(r io.Reader) *transform.Reader {
	return transform.NewReader(r, simplifiedchinese.GBK.NewEncoder())
}
func GbkDec(r io.Reader) *transform.Reader {
	return transform.NewReader(r, simplifiedchinese.GBK.NewDecoder())
}
func GB18030Dec(r io.Reader) *transform.Reader {
	return transform.NewReader(r, simplifiedchinese.GB18030.NewDecoder())
}

func ToGBKStr(s string) ([]byte, error) {
	tr := GbkEnc(strings.NewReader(s))
	return ioutil.ReadAll(tr)
}
func GB18030Enc(r io.Reader) *transform.Reader {
	return transform.NewReader(r, simplifiedchinese.GB18030.NewEncoder())
}
func ToGB18030Str(s string) ([]byte, error) {
	tr := GB18030Enc(strings.NewReader(s))
	return ioutil.ReadAll(tr)
}
