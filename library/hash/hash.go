package hash

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

func Hash(req interface{}) (string, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	hash := md5.New()
	hash.Write(jsonData)

	return hex.EncodeToString(hash.Sum(nil)), nil
}
