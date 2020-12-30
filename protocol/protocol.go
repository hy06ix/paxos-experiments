package protocol

/*
The `NewProtocol` method is used to define the protocol and to register
the handlers that will be called if a certain type of message is received.
The handlers will be treated according to their signature.

The protocol-file defines the actions that the protocol needs to do in each
step. The root-node will call the `Start`-method of the protocol. Each
node will only use the `Handle`-methods, and not call `Start` again.
*/

import (
	"sync"
	"time"

	"go.dedis.ch/onet/v3"
	"go.dedis.ch/onet/v3/log"
	"go.dedis.ch/onet/v3/network"
)

const Name = "Paxos"

func init() {
	network.RegisterMessages(Prepare{}, Promise{}, Accept{}, Accepted{})
	_, err := onet.GlobalProtocolRegister(Name, NewProtocol)
	if err != nil {
		panic(err)
	}
}

type VerificationFn func(msg []byte, data []byte) bool

// TemplateProtocol holds the state of a given protocol.
//
// For this example, it defines a channel that will receive the number
// of children. Only the root-node will write to the channel.
type PaxosProtocol struct {
	*onet.TreeNodeInstance

	stoppedOnce sync.Once

	ChannelPrepare  chan StructPrepare
	ChannelPromise  chan StructPromise
	ChannelAccept   chan StructAccept
	ChannelAccepted chan StructAccepted
}

// Check that *TemplateProtocol implements onet.ProtocolInstance
var _ onet.ProtocolInstance = (*PaxosProtocol)(nil)

// NewProtocol initialises the structure for use in one round
func NewProtocol(n *onet.TreeNodeInstance) (onet.ProtocolInstance, error) {
	t := &PaxosProtocol{
		TreeNodeInstance: n,
	}
	if err := n.RegisterChannels(&t.ChannelPrepare, &t.ChannelPromise, &t.ChannelAccept); err != nil {
		return nil, err
	}
	return t, nil
}

// Start sends the Announce-message to all children
func (paxos *PaxosProtocol) Start() error {
	log.Lvl1(paxos.ServerIdentity(), "Starting TemplateProtocol")

	if paxos.IsRoot() {
		go func() {
			if err := paxos.SendToChildrenInParallel(&Prepare{}); len(err) > 0 {
				log.Lvl2(paxos.ServerIdentity(), "failed to send announce to all children")
			}
		}()
	}

	return nil
}

// Dispatch implements the main logic of the protocol. The function is only
// called once. The protocol is considered finished when Dispatch returns and
// Done is called.
func (paxos *PaxosProtocol) Dispatch() error {
	log.Lvl3(paxos.ServerIdentity(), "Started node")
	log.Lvl3("Sleeping dispatch for keys")
	time.Sleep(time.Duration(4) * time.Second)

	// set threshold

	if !paxos.IsRoot() {
		// verifyChan := make(chan bool, 1)

		log.Lvl2(paxos.ServerIdentity(), "Waiting for prepare")
		prepare, channelOpen := <-paxos.ChannelPrepare
		if !channelOpen {
			return nil
		}
		prepare = prepare // need delete

		log.Lvl2(paxos.ServerIdentity(), "Received prepare. Verifying...")
		go func() {
			// verifyChan <- paxos.verificationFn()
		}()

		// prepare -> promise

		if err := paxos.SendToParent(&Promise{}); err != nil {
			log.Lvl3(paxos.ServerIdentity(), "error while broadcasting promise message")
		}
	} else {
		// Root
	}

	// Set timeout

	if paxos.IsRoot() {

	}

	return nil
}

func (paxos *PaxosProtocol) Shutdown() error {
	paxos.stoppedOnce.Do(func() {
		close(paxos.ChannelPrepare)
		close(paxos.ChannelPromise)
		close(paxos.ChannelAccept)
		close(paxos.ChannelAccepted)
	})

	return nil
}
