module simul

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/dedis/cothority_template v0.0.0-20200612094046-dae253bd6f36 // indirect
	go.dedis.ch/onet/v3 v3.2.6
	hy06ix/protocol v0.0.0
)

replace hy06ix/protocol v0.0.0 => ../protocol
