package uuid

import (
	"encoding/hex"
	"math/rand"
	"sync"
)

type IDGenerator struct {
	mu  sync.Mutex
	rnd *rand.Rand
}

func New(seed int64) *IDGenerator {
	return &IDGenerator{
		rnd: rand.New(rand.NewSource(seed)),
	}
}

func (g *IDGenerator) NewID(id []byte) (string, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	_, err := g.rnd.Read(id[:])

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(id[:]), nil
}
