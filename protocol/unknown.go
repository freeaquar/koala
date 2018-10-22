package protocol

type UnknownRevealer struct {
	Revealer
}

func (this UnknownRevealer) Inspect(request []byte) (ok bool) {
	return true
}

func (this UnknownRevealer) Parse(request []byte) (revealData RevealData, err error) {
	return revealData, nil
}

func (this UnknownRevealer) PreMatch(revealData1, revealData2 RevealData) (ok bool) {
	return true
}
