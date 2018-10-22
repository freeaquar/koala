package outbound

type Revealer interface {
	// inspect protocol
	Inspect(request []byte) (ok bool)
	Parse(request []byte) (revealData RevealData, err error)
	PreMatch(revealData1, revealData2 RevealData) (ok bool)
}

type RevealData struct {
	Handler Revealer
	Must    []byte
}

var Revealers = []Revealer{
	HTTPRevealer{},
	ThriftRevealer{},
	UnknownRevealer{},
}

type model map[interface{}]interface{}
