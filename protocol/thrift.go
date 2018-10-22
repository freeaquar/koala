package protocol

type ThriftRevealer struct {
	Revealer
	Methods map[string]bool
}

func (this *ThriftRevealer) Inspect(request []byte) (ok bool) {
	return
}

func (this *ThriftRevealer) Parse(request []byte) (revealData RevealData, err error) {
	return
}

func (this *ThriftRevealer) PreMatch(revealData1, revealData2 RevealData) (ok bool) {
	return
}
