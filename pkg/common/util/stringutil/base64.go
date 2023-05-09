package stringutil

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
)

func DecodeBase64(i string) ([]byte, error) {
	return ioutil.ReadAll(base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(i)))
}
