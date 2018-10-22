package protocol

type Revealer interface {
	// inspect protocol
	Inspect(request []byte) (ok bool)
	// parse key signature
	Parse(request []byte) (revealData RevealData, err error)
	// match key signature
	PreMatch(revealData1, revealData2 RevealData) (ok bool)
}

type RevealData struct {
	Handler Revealer
	Must    []byte
}

var RevealHandler = []Revealer{
	HTTPRevealer{},
	//ThriftRevealer{},
	UnknownRevealer{},
}

type model map[interface{}]interface{}
