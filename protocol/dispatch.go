package protocol

//
//import (
//	"net/http"
//	"net/url"
//	"strings"
//)
//
//var HTTP_METHODS = map[string]bool{
//	http.MethodGet:     true,
//	http.MethodPost:    true,
//	http.MethodPut:     true,
//	http.MethodDelete:  true,
//	http.MethodHead:    true,
//	http.MethodOptions: true,
//	http.MethodTrace:   true,
//	http.MethodConnect: true,
//}
//
//type HTTPChunks struct {
//	ProtocolChunks
//
//	Method  Chunk
//	Uri     Chunk
//	Version Chunk
//	Args    Chunk
//	Headers Chunk
//	Body    Chunk
//}
//
//type HTTPWeight struct {
//	Method  int
//	Uri     int
//	Version int
//	Args    int
//	Headers int
//	Body    int
//}
//
//var weight = struct {
//	Method  int
//	Uri     int
//	Version int
//	Args    int
//	Headers int
//	Body    int
//}{
//	Method:  PRIORITY_LOW,
//	Uri:     PRIORITY_MUST,
//	Version: PRIORITY_IGNORE,
//	Args:    PRIORITY_MID,
//	Headers: PRIORITY_LOW,
//	Body:    PRIORITY_MID,
//}
//
//type HTTP struct {
//	Revealer
//}
//
//func (this *HTTP) Inspect(request string) (ok bool) {
//	for method, _ := range HTTP_METHODS {
//		if !strings.HasPrefix(request, method) {
//			continue
//		}
//		if !strings.Contains(request, "HTTP/") {
//			continue
//		}
//		return true
//	}
//	return false
//}
//
//func (this *HTTP) Parse(request string) (chunks HTTPChunks, err error) {
//	httpChunks := HTTPChunks{}
//	method, uri, version, args, headers, body := this.revealRequest(request)
//
//	httpChunks.Uri = Chunk{value: uri, weight: PRIORITY_MUST}
//	httpChunks.Args = Chunk{value: args, weight: PRIORITY_HIGH}
//	httpChunks.Method = Chunk{value: method, weight: PRIORITY_MID}
//	httpChunks.Headers = Chunk{value: headers, weight: PRIORITY_LOW}
//	httpChunks.Version = Chunk{value: version, weight: PRIORITY_LOW}
//	httpChunks.Body = Chunk{value: body, weight: PRIORITY_LOW}
//	httpChunks.StartByUnknown = false
//
//	return httpChunks, nil
//}
//
//func (this *HTTP) Match(chunks1, chunks2 HTTPChunks) (score int) {
//	score = 0
//	if weight.Method > 0 && chunks1.Method == chunks2.Method {
//		score += weight.Method
//	}
//	if weight.Uri > 0 && chunks1.Uri == chunks2.Uri {
//		score += weight.Uri
//	}
//	if weight.Version > 0 && chunks1.Version == chunks2.Version {
//		score += weight.Version
//	}
//	//if weight.Args > 0 && chunks1.Args == chunks2.Args {
//	//	score += weight.Args
//	//}
//	//if weight.Headers > 0 && chunks1.Headers == chunks2.Headers {
//	//	score += weight.Headers
//	//}
//	if weight.Body > 0 && chunks1.Body == chunks2.Body {
//		score += weight.Body
//	}
//
//	return score
//}
//
//func (this *HTTP) revealRequest(request string) (
//	method, uri, version string, args map[string][]string,
//	headers map[string]string, body string) {
//
//	data := strings.Split(request, "\r\n\r\n")
//	headerLines, bodyLines := data[0], data[1]
//	data = strings.SplitN(headerLines, "\r\n", 2)
//	firstLine, headerLines := data[0], data[1]
//
//	method, uri, version, args = this.revealFirstLine(firstLine)
//	headers = this.revealHeaderLines(headerLines)
//	body = bodyLines
//
//	return method, uri, version, args, headers, body
//}
//
//func (this *HTTP) revealFirstLine(firstLine string) (
//	method, uri, version string, args map[string][]string) {
//	data := strings.Fields(firstLine)
//	if len(data) != 3 {
//		return
//	}
//	method, uri, version = data[0], data[1], data[2]
//	if _, ok := HTTP_METHODS[method]; !ok {
//		return
//	}
//
//	requestUrl, err := url.Parse(uri)
//	if err != nil {
//		return
//	}
//	uri = requestUrl.EscapedPath()
//	args = requestUrl.Query()
//
//	return method, uri, version, args
//}
//
//func (this *HTTP) revealHeaderLines(headerLines string) (headers map[string]string) {
//	data := strings.Split(headerLines, "\r\n")
//	headers = make(map[string]string)
//	for _, line := range data {
//		values := strings.SplitN(line, ":", 2)
//		headers[values[0]] = strings.TrimSpace(values[1])
//	}
//
//	return headers
//}
