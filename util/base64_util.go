package util

import (
	"encoding/base64"
	"fmt"
)

func Base64Encoding(data []byte) string {

	sEnc := base64.StdEncoding.EncodeToString(data)
	fmt.Println(sEnc)
	return sEnc
}
