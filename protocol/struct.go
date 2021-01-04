package protocol

import (
	"go.dedis.ch/onet/v3"
	"go.dedis.ch/onet/v3/network"
)

const DefaultProtocolName = "Paxos"

func init() {
	network.RegisterMessages(
		Prepare{}, Promise{}, Accept{},
	)
}

type Prepare struct {
	suggestN int64
	Sender   string
}

type Promise struct {
	suggestN int64
	Sender   string
}

type Accept struct {
	suggestN int64
	value    []byte
	Sender   string
}

type Accepted struct {
	suggestN int64
	// value    []byte
	Sender string
}

type StructPrepare struct {
	*onet.TreeNode
	Prepare
}

type StructPromise struct {
	*onet.TreeNode
	Promise
}

type StructAccept struct {
	*onet.TreeNode
	Accept
}

type StructAccepted struct {
	*onet.TreeNode
	Accepted
}
