package outbound

type ThriftRevealer struct {
	Revealer
	Methods map[string]bool
}

/*
 * contain `HTTP`
 * start by HTTP method
 */
func (this *ThriftRevealer) Inspect(request []byte) (ok bool) {
	return
}

/*
 * store thrift name as must match segment
 */
func (this *ThriftRevealer) Parse(request []byte) (revealData RevealData, err error) {
	return
}

func (this *ThriftRevealer) PreMatch(revealData1, revealData2 RevealData) (ok bool) {
	return
}
