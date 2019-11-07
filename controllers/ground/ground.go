package ground

import (
	"math/rand"
	"time"
)

type Edge struct{
	Weight	int
}

func Create(size int)([][]Edge){
	rand.Seed(time.Now().UTC().UnixNano())
	edges := make([][]Edge, size+1)
	Weight:= 0
	for x := 0; x < size+1; x++ {
		edges[x] = make([]Edge, size+1)
		for y := 0; y < size+1; y++ {
			Weight = rand.Intn(10) + 1
			edges[x][y] = Edge{Weight}
		}
    }
	return edges
}