package protocol

import (
	"go.dedis.ch/onet/v3/network"
)

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
	value    []byte
	Sender   string
}

type StructPrepare struct {
}

type StructPromise struct {
}

type StructAccept struct {
}

type StructAccepted struct {
}
