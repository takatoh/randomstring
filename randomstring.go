package randomstring

import (
	"math/rand"
	"time"
)

type RandomString struct {
	pool string
	r    *rand.Rand
	l    int
}

func New(pool string) *RandomString {
	p := new(RandomString)
	p.pool = pool
	p.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	p.l = len(pool)
	return p
}

func (rs *RandomString) Rand(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		x := rs.r.Intn(rs.l)
		s = s + rs.pool[x:x+1]
	}
	return s
}
