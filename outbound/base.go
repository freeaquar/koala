package outbound

// save chunk value and its weight
//type Chunk map[interface{}]interface{}
type Chunk struct {
	value  interface{}
	weight int
}

// reveal request by protocol
// different Chunk part with different weight and strategy
type ProtocolChunks struct {
	StartByUnknown bool
	Unknown        Chunk
}

type ProtocolRevealer interface {
	Inspect(string) (ok bool)
	Parse(string) (chunks ProtocolChunks, err error)
	Match(ProtocolChunks, ProtocolChunks) (score int)
}

const (
	PRIORITY_MUST   = 100000
	PRIORITY_HIGH   = 100
	PRIORITY_MID    = 10
	PRIORITY_LOW    = 1
	PRIORITY_IGNORE = 0
)

type model map[string]string
