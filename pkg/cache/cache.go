package cache

import (
	"encoding/hex"
	"github.com/minio/highwayhash"
)

type Cache map[string]string

func (t Cache) C(text string) string {
	hf, err := highwayhash.New64(make([]byte, 32))
	if err != nil {
		panic(err)
	}
	id := hex.EncodeToString(hf.Sum([]byte(text)))
	if _, exist := t[id]; !exist {
		t[id] = text
	}
	return t[id]
}
