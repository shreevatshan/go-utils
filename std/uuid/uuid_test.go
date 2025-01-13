package uuid

import (
	"math/rand"
	"testing"
	"time"
)

func TestIDGenerator_NewID(t *testing.T) {
	g := &IDGenerator{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	var base [8]byte
	id, _ := g.NewID(base[:])

	if id == "" {
		t.Errorf("id is empty")
	}
}
