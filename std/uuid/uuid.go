package uuid

import (
	"encoding/hex"
	"math/rand"
	"sync"
)

// IDGenerator is a type that generates unique IDs
type IDGenerator struct {
	mu  sync.Mutex
	rnd *rand.Rand
}

// New creates a new IDGenerator, using the given seed
func New(seed int64) *IDGenerator {
	return &IDGenerator{
		rnd: rand.New(rand.NewSource(seed)),
	}
}

// NewID generates a new unique ID
func (g *IDGenerator) NewID(id []byte) (string, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	_, err := g.rnd.Read(id[:])

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(id[:]), nil
}
