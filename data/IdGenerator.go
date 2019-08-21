package data

import (
	"encoding/hex"
	"math/rand"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

type TIdGenerator struct{}

var IdGenerator *TIdGenerator = &TIdGenerator{}

func (c *TIdGenerator) NextShort() string {
	value := 100000000 + rand.Int63n(899999999)
	return strconv.FormatInt(value, 10)
}

func (c *TIdGenerator) NextLong() string {
	value, ok := uuid.NewV4()
	if ok != nil {
		return IdGenerator.NextShort()
	}
	return hex.EncodeToString(([]byte)(value[:]))
}
