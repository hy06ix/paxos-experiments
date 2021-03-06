module github.com/hy06ix/paxos-experiments

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/dedis/cothority_template v0.0.0-20200612094046-dae253bd6f36
	github.com/hy06ix/paxos-experiments/protocol v0.0.0
	github.com/stretchr/testify v1.5.1
	go.dedis.ch/cothority/v3 v3.4.5
	go.dedis.ch/kyber/v3 v3.0.13
	go.dedis.ch/onet/v3 v3.2.6
	go.dedis.ch/protobuf v1.0.11
	golang.org/x/sys v0.0.0-20200523222454-059865788121
	gopkg.in/urfave/cli.v1 v1.20.0
)

replace github.com/hy06ix/paxos-experiments/protocol v0.0.0 => ./protocol

go 1.15
