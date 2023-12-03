package string2hash

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func GetResult(parameters *Parameters) {
	if parameters.error == false {
		switch parameters.hash {
		case "md5":
			hash := md5.Sum([]byte(parameters.text))
			hashSlice := make([]byte, 0)
			hashSlice = append(hashSlice, hash[:]...)
			parameters.base64 = base64.StdEncoding.EncodeToString(hashSlice)
			parameters.result = hex.EncodeToString(hashSlice)
		case "sha1":
			hash := sha1.Sum([]byte(parameters.text))
			hashSlice := make([]byte, 0)
			hashSlice = append(hashSlice, hash[:]...)
			parameters.base64 = base64.StdEncoding.EncodeToString(hashSlice)
			parameters.result = hex.EncodeToString(hashSlice)
		}
		fmt.Println(parameters.result)
	}
}
