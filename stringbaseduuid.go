package randomstring

import (
	"encoding/base32"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

type StringBasedUUID struct {
	rand *rand.Rand
}

func NewStringBasedUUID() *StringBasedUUID {
	p := new(StringBasedUUID)
	p.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	return p
}

func (sbu *StringBasedUUID) Generate() string {
	u := uuid.New()
	data := make([]byte, 0)
	for _, b := range u {
		data = append(data, b)
	}
	dst := make([]byte, base32.StdEncoding.EncodedLen(len(data)))
	base32.StdEncoding.Encode(dst, data)
	s := make([]string, 0)
	for _, c := range strings.Split(string(dst), "") {
		if c != "=" {
			s = append(s, sbu.down(c))
		}
	}
	return strings.Join(s, "")
}

func (sbu *StringBasedUUID) down(c string) string {
	if sbu.rand.Intn(2) > 0 {
		return strings.ToLower(c)
	} else {
		return c
	}
}
