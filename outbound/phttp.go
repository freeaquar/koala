package outbound

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
)

type HTTPRevealer struct {
	Revealer
	Methods map[string]bool
}

func NewHTTPRevealer() *HTTPRevealer {
	return &HTTPRevealer{
		Methods: map[string]bool{
			http.MethodGet:     true,
			http.MethodPost:    true,
			http.MethodPut:     true,
			http.MethodDelete:  true,
			http.MethodHead:    true,
			http.MethodOptions: true,
			http.MethodTrace:   true,
			http.MethodConnect: true,
		},
	}
}

/*
 * contain `HTTP`
 * start by HTTP method
 */
func (this *HTTPRevealer) Inspect(request []byte) (ok bool) {

	if !bytes.Contains(request, []byte("HTTP/")) {
		return false
	}

	index := bytes.Index(request, []byte(" "))
	if _, ok := this.Methods[string(request[:index])]; ok {
		return true
	}

	return false
}

/*
 * store uri as must match segment
 */
func (this *HTTPRevealer) Parse(request []byte) (revealData RevealData, err error) {
	revealData.Handler = this
	revealData.Must = this.revealUri(request)
	if len(revealData.Must) <= 0 {
		err = errors.New("Can not reveal uri")
	}
	return revealData, err
}

func (this *HTTPRevealer) PreMatch(revealData1, revealData2 RevealData) (ok bool) {
	if 0 == bytes.Compare(revealData1.Must, revealData2.Must) {
		return true
	}
	return false
}

func (this *HTTPRevealer) revealUri(request []byte) (uri []byte) {
	data := bytes.Split(request, []byte("\r\n\r\n"))
	headerLines := data[0]
	data = bytes.SplitN(headerLines, []byte("\r\n"), 2)
	firstLine := data[0]

	_, uri, _, _ = this.revealFirstLine(firstLine)
	return uri
}

func (this *HTTPRevealer) revealFirstLine(firstLine []byte) (
	method, uri, version []byte, args map[string][]string) {

	data := bytes.Fields(firstLine)
	if len(data) != 3 {
		return
	}
	method, uri, version = data[0], data[1], data[2]
	if _, ok := this.Methods[string(method)]; !ok {
		return
	}

	requestUrl, err := url.Parse(string(uri))
	if err != nil {
		return
	}
	uri = []byte(requestUrl.EscapedPath())
	args = requestUrl.Query()

	return method, uri, version, args
}
