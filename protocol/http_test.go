package protocol

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Inspect(t *testing.T) {
	should := require.New(t)
	httpObj := HTTPRevealer{}
	requests := []byte("POST /gulfstream/passenger/v2/core/pNewOrder?_t=1540212743&appversion=5.2.22&channel=102&datatype=101&imei=5a4d1aaaf5f54276655f592c743e2927&lang=zh-CN&maptype=soso&model=iPhone&networkType=4G&os=12.0&sig=d27ccddcc552903513f2de294040be4625ac4f17&terminal_id=1&token=BoGETCmUS4qHbp4KQ6lCsjbg9tendHBPJDmLXzOAWOAkzDsOwjAQANG7TL2Kdtd2bG9Lzx34hE9jJBBVxN1RlOo1o1kZSpAmnRRhGGHCcMJUVYWRiI28UwgOR4QTAcKZ8N51rjZ3t-ZFhSuRhIVY-by-78tC6E-4EVayunqvWbgTWEs5ZW-1VYTHvnxu-T8AAP__&wsgenv=eV60AEgoiGE0s5VKAAAAA4ABAABe66x27PUF/I79zg5dsZtmBlK4qyKkXKCGkQWVgMu88e4VmSvO0EO5%2BpBTYj6mTX8IpOn6Dofg5a38PMjC0urw1QbQp5vtlKJBdk4dXwC25k4/bzKAkuA7EIYPIL/97PCzTyu/NFo91E8pu411CdWX2lPEG7mbz7oXxjoRe6q3hWYd1BRNFXv5DrWWb5CAHnqLGC2MBAq15oEMrAMMvZUrGuRi3b9E4YoACzubSP9fcnxEFlytiIaic69J5bHqtAng2QNA3%2B7%2BwLCg%2B5%2BGnGi1byxgBMNjPKMK8RxBKZMbbrENxDKlKcIBFiveI1EDogkIVxxMJ4c3DCyCWo%2BGI8U6S7QDC88vJiRqnPoJJKYr1SL3/sqTOX1HzkOZPYDMNbokYvzQP4ElCbXnab8gkPoqdAKLUDBWfH%2BoUfFSRIKScWevQAdQZEGcnDaJyrW8UiwzFcBbhq/PxDIuEprcPa40uA0wv6WX4UM28cxfZNKWgqMJtiF0IF6d4Xq4xJaoKKWIxCMD98KhvbAvrKRZrFHzN0KMmjxIwElEwa9Y6BCGYcJliM6zU6/XChgeYrL3QAA HTTP/1.0")
	should.True(httpObj.Inspect(requests))

	requests = []byte("\\x00\\x00\\x00\\x97\\x80\\x01\\x00\\x01\\x00\\x00\\x00\\x14MultiAreaInfoByCoord\\x00\\x00\\x00\\x00\\x0c\\x00\\x01\n\\x00\\x01\\x00\\x00\\x00>Jz\\xf3@\\x0b\\x00\\x02\\x00\\x00\\x00 4C351760F4A9486E060053B9167FD7FC\\x0f\\x00\\x03\\x0c\\x00\\x00\\x00\\x01\\x04\\x00\\x01@]\\x17\\x03\\xaf\\xb7\\xe9\\x10\\x04\\x00\\x02@C\\xec\\xbf\t\\x95\\xaa\\xf8\\x00\\x08\\x00\\x04\\x00\\x00\\x00\\x02\\x0c\\x00\\x05\\x08\\x00\\x01\\x00\\x00\\x00\\x02\\x00\\x0c\\x00\\x06\n\\x00\\x01\\x00\\x00\\x00\\x00\\x00\\x00\\x00\\x80\\x00\\x00\\x00")
	should.False(httpObj.Inspect(requests))
}
