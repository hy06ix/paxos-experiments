package main

import (
	// Service needs to be imported here to be instantiated.
	_ "github.com/hy06ix/paxos-experiments/simulation"
	"go.dedis.ch/onet/v3/simul"
)

func main() {
	simul.Start()
}
