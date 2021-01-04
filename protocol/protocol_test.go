package protocol

/*
The test-file should at the very least run the protocol for a varying number
of nodes. It is even better practice to test the different methods of the
protocol, as in Test Driven Development.
*/

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.dedis.ch/kyber/v3/suites"
	"go.dedis.ch/onet/v3"
	"go.dedis.ch/onet/v3/log"
)

var tSuite = suites.MustFind("Ed25519")

func TestMain(m *testing.M) {
	log.MainTest(m)
}

// Tests a 2, 5 and 13-node system. It is good practice to test different
// sizes of trees to make sure your protocol is stable.
func TestNode(t *testing.T) {
	log.SetDebugVisible(2)

	nodes := []int{5}
	for _, nbrNodes := range nodes {
		local := onet.NewLocalTest(tSuite)
		_, _, tree := local.GenBigTree(nbrNodes, nbrNodes, nbrNodes-1, true)
		log.Lvl3(tree.Dump())

		pi, err := local.CreateProtocol(DefaultProtocolName, tree)
		require.Nil(t, err)

		protocol := pi.(*PaxosProtocol)
		// require.NoError(t, protocol.Start())

		// timeout := network.WaitRetry * time.Duration(network.MaxRetryConnect*nbrNodes*2) * time.Millisecond

		err = protocol.Start()
		if err != nil {
			local.CloseAll()
			t.Fatal(err)
		}

		select {
		case <-protocol.ChannelFinish:
			log.Lvl1("End of Round")
		case <-time.After(time.Second * time.Duration(10)):
			t.Fatal("Didn't finish in time")
		}
		local.CloseAll()
	}
}
