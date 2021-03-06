package randomstring

import (
	"math/rand"
	"time"
)

type RandomString struct {
	pool []rune
	r    *rand.Rand
	l    int
}

func New(pool string) *RandomString {
	p := new(RandomString)
	p.pool = uniq([]rune(pool))
	p.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	p.l = len(p.pool)
	return p
}

func (rs *RandomString) Rand(n int) string {
	b := make([]rune, n)
	for i := 0; i < n; i++ {
		x := rs.r.Intn(rs.l)
		b[i] = rs.pool[x]
	}
	return string(b)
}

func uniq(s []rune) []rune {
	m := make(map[rune]struct{})
	for _, e := range s {
		m[e] = struct{}{}
	}
	u := []rune{}
	for i := range m {
		u = append(u, i)
	}
	return u
}
