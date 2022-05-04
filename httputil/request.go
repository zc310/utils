package httputil

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"errors"
	"fmt"
	"strings"
)

var DefaultClient = http.DefaultClient

func GetRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return DefaultClient.Do(req)
}
func GetJSON(url string, v interface{}) error {
	resp, err := GetRequest(url)
	if err != nil {
		return err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	return json.NewDecoder(resp.Body).Decode(v)
}

func GetByte(url string) ([]byte, error) {
	resp, err := GetRequest(url)
	if err != nil {
		return []byte{}, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	if len(b) == 0 {
		err = io.EOF
	}
	return b, nil
}
func GetString(url string) (string, error) {
	b, err := GetByte(url)
	return string(b), err
}

func StatusCode(url string) (int, error) {
	resp, err := GetRequest(url)
	if err != nil {
		return 0, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	return resp.StatusCode, nil
}

func GetJs(url string) (string, error) {
	resp, err := GetRequest(url)
	if err != nil {
		return "", err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	if !strings.HasPrefix(resp.Header.Get("Content-Type"), "application/javascript") {
		return "", errors.New("error Content-Type:" + resp.Header.Get("Content-Type"))
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
