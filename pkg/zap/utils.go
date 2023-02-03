package zap

import (
	"fmt"

	mh "github.com/multiformats/go-multihash"
)

func ShortenedCid(cid mh.Multihash) string {
	return fmt.Sprintf("%X…%X", []byte(cid[2:6]), []byte(cid[len(cid)-4:]))
}
