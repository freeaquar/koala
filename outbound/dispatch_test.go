package outbound

import (
	"testing"
	"github.com/stretchr/testify/require"
	"fmt"
	"encoding/json"
)

var request = `POST /gulfstream/driver/v2/driver/dFinishOrder HTTP/1.0
host: api.udache.com
x-forwarded-for: 112.17.236.196, 112.17.236.196
connection: close
content-length: 1794
x-real-ip: 112.17.236.196
didi-header-rid: 0635535a5bc6e89a000060bedbc6db30
x-forwarded-proto: http
user-agent: Android/6.0.1 didihttp OneNet/2.1.0.68 com.sdu.didi.gsui/5.1.24
accept-encoding: gzip
cip: 112.17.236.196
cityid: 86
content-type: application/x-www-form-urlencoded
didi-header-omgid: evdwVhzLQk-OS2Q1EAUNoQ
didi-header-spanid: 6144280284046801549
flowtag: 1
inner-source: trans
wsgsig: dd01-UtMTLB/cm0Pthvg3GZ7rVUNDYSqTEuAc1e50FeblFEOLUQOwwmDKRUuDuNIFPy4OeO4nZkfnZ+HJswdK/yKN2XtsYhrxXani+zAP1fQ6WB6ezywx5KekQl3SouSnIfIEdJLbqUfoYol9TonTbR4lGiNVuoT3S7DkbvDPvkJTYNbGnNgTa3R/ZEWpu7T0Qo8Oc70OukbqFoVwYhLKbR0PZUiT
didi-header-hint-content: {"Cityid":86}
x-routing-trace: gs_api_nginx_hna

a3_token=P9yQI7nN9WGXu1SsYxXYoR%2FXnjhOiUndnIMoySDpzU%2Bn9CsHpQ3yfvBcv7n6gVNtVVM%2FbQz0oSA25MVoU3VTEQlredhJ7rIsoM60sCCiz%2FsZYsstua1D4O1GVfy2cNZbFhuqbKZKIm3QB98Msbnk8OquWy1LEYQisheTysiHr2z4%2BSnY6WcDSffGdC9RtQTD&datatype=2&dest_lat=29.345328354702122&new_price_conf=1&ticket=OJqjvzCXglMUS-Fzy_HgW1RbTSPMUdneGrIRN3IMndAkzDmuwzAMANG7TE0YJGVKFtvf_ztkcZZGARKkMnL3wEj9BrMxlKRMOinCMNKE4aSrqgqjkNaie1hrtcx155ncKUj-_hEOJAhHMhaN2b0szWqrKpxJN2ElN16P9_O0kvoRLqRFacW1lS5cSSw8lh5ReyDcfs_7nn8DAAD__w%3D%3D&maptype=soso&appversioncode=321&highway_fee=%E8%AF%B7%E8%BE%93%E5%85%A5%E9%87%91%E9%A2%9D&imei=863670033562410&appversion=5.1.24&lang=zh-CN&park_fee=%E8%AF%B7%E8%BE%93%E5%85%A5%E9%87%91%E9%A2%9D&platform_type=2&bridge_fee=%E8%AF%B7%E8%BE%93%E5%85%A5%E9%87%91%E9%A2%9D&channel=2010000001&is_offline_paid=0&dest_lng=120.10726630826768&product_id=260&utc_offset=480&dest_timestamp=1539762330&terminal_id=1&deviceid=7852c53bb1047946641cdbdc23646f0f&lng=120.10726630826768&os=6.0.1&oid=TWpnNE5UZzBOekEyT0RBd09ETTFOek14&plutus_data=pbv3H4sIAAAAAAAA%2F%2ByUPWgUURDHd8yqkWuMNjHVuZUfm2PmvZ2d3RVxAzZaaRMionE%2F4xUmeqSIBOGUFLE8UvhVKHZJE7ELdoqIYPADhFSCohZ2goWlXMhbkmgR%2B8wW%2B2P%2B8%2BY%2FPIbXe%2Bf9g08791oDt3b0rnyfX2y3Z2BfjVhxEDL7IfffnXv06vNRp0baZz8IhDQdstAOrNg6PSQo6KWiSuSCk%2B4hhamXljrLM8w8ivIkFUpREs48VqUq0yLjouQsJJ9SSc8Mj1iXvs7svgxXe6asNsBtsNaiU9F92K%2BQgkHCQZI6caRVpHm%2B0p9AuQTLb7pxLH4O1utK%2BQDWCthfwEaW8AfUuj9Xa5RA%2FayKfsPc0rcX79q207F717Dv17P8RM%2Fiy%2FjAakl9Nj5oModNhkyD0MCQgVMGzho4ZyAx0DRwzcB1Azer2f51HfcqegyzMO2kxVhzfHSyeaVwImIdik8eousUU0k2ObqqVgL7ZIRiPDdppZW4TqvIJlp50XKi88wYNFCxFiJRniiXOVQNZM0ivs%2BkxMX134UbCzD9n6Mwb2UUz%2BMGrg%2Ff1X%2B7P4XhLbrjBlPc4IUubm58ZHsntndiU2M1sCcYO9mcGDke9z1cXnj7cdfFuL%2FTfUjr8CcAAP%2F%2FOetnUVEFAAA%3D&biz_type=2&lat=29.345328354702122&location_cityid=86&model=OPPO+R9s
`

var httpObj = HTTP{}

func Test_Inspect(t *testing.T) {
	should := require.New(t)
	should.True(httpObj.Inspect(request))
}

func Test_Parse(t *testing.T) {
	should := require.New(t)
	_, err := httpObj.Parse(request)
	should.Nil(err)
}

func Test_reveal(t *testing.T) {
	should := require.New(t)
	request = "POST /gulfstream/driver/v2/driver/dFinishOrder?a=1&b=9#x0 HTTP/1.0\r\nhost: api.udache.com\r\n" +
		"x-forwarded-for: 112.17.236.196, 112.17.236.196\r\n\r\n" +
		"a3_token=P9yQI7nN9WGXu1SsYxXYoR%2FXnjhOiUndnIMoySDpzU%2Bn9CsHpQ3yfvBcv7n6gVNtVVM%2FbQz0oSA25MVoU3VTEQlredhJ7rIsoM60sCCiz%2FsZYsstua1D4O1GVfy2cNZbFhuqbKZKIm3QB98Msbnk8OquWy1LEYQisheTysiHr2z4%2BSnY6WcDSffGdC9RtQTD&datatype=2&dest_lat=29.345328354702122&new_price_conf=1&ticket=OJqjvzCXglMUS-Fzy_HgW1RbTSPMUdneGrIRN3IMndAkzDmuwzAMANG7TE0YJGVKFtvf_ztkcZZGARKkMnL3wEj9BrMxlKRMOinCMNKE4aSrqgqjkNaie1hrtcx155ncKUj-_hEOJAhHMhaN2b0szWqrKpxJN2ElN16P9_O0kvoRLqRFacW1lS5cSSw8lh5ReyDcfs_7nn8DAAD__w%3D%3D&maptype=soso&appversioncode=321&highway_fee=%E8%AF%B7%E8%BE%93%E5%85%A5%E9%87%91%E9%A2%9D&imei=863670033562410&appversion=5.1.24&lang=zh-CN&park_fee=%E8%AF%B7%E8%BE%93%E5%85%A5%E9%87%91%E9%A2%9D&platform_type=2&bridge_fee=%E8%AF%B7%E8%BE%93%E5%85%A5%E9%87%91%E9%A2%9D&channel=2010000001&is_offline_paid=0&dest_lng=120.10726630826768&product_id=260&utc_offset=480&dest_timestamp=1539762330&terminal_id=1&deviceid=7852c53bb1047946641cdbdc23646f0f&lng=120.10726630826768&os=6.0.1&oid=TWpnNE5UZzBOekEyT0RBd09ETTFOek14&plutus_data=pbv3H4sIAAAAAAAA%2F%2ByUPWgUURDHd8yqkWuMNjHVuZUfm2PmvZ2d3RVxAzZaaRMionE%2F4xUmeqSIBOGUFLE8UvhVKHZJE7ELdoqIYPADhFSCohZ2goWlXMhbkmgR%2B8wW%2B2P%2B8%2BY%2FPIbXe%2Bf9g08791oDt3b0rnyfX2y3Z2BfjVhxEDL7IfffnXv06vNRp0baZz8IhDQdstAOrNg6PSQo6KWiSuSCk%2B4hhamXljrLM8w8ivIkFUpREs48VqUq0yLjouQsJJ9SSc8Mj1iXvs7svgxXe6asNsBtsNaiU9F92K%2BQgkHCQZI6caRVpHm%2B0p9AuQTLb7pxLH4O1utK%2BQDWCthfwEaW8AfUuj9Xa5RA%2FayKfsPc0rcX79q207F717Dv17P8RM%2Fiy%2FjAakl9Nj5oModNhkyD0MCQgVMGzho4ZyAx0DRwzcB1Azer2f51HfcqegyzMO2kxVhzfHSyeaVwImIdik8eousUU0k2ObqqVgL7ZIRiPDdppZW4TqvIJlp50XKi88wYNFCxFiJRniiXOVQNZM0ivs%2BkxMX134UbCzD9n6Mwb2UUz%2BMGrg%2Ff1X%2B7P4XhLbrjBlPc4IUubm58ZHsntndiU2M1sCcYO9mcGDke9z1cXnj7cdfFuL%2FTfUjr8CcAAP%2F%2FOetnUVEFAAA%3D&biz_type=2&lat=29.345328354702122&location_cityid=86&model=OPPO+R9s"

	method, uri, version, args, headers, body := httpObj.revealRequest(request)

	fmt.Println("=============================")
	fmt.Println(method)
	fmt.Println(uri)
	fmt.Println(parse(version))
	fmt.Println(parse(args))
	fmt.Println(parse(headers))
	fmt.Println(body)

	should.True(true)
}

func Test_match(t *testing.T) {
	should :=require.New(t)
	request1 := "POST /gulfstream/driver/v2/driver/dFinishOrder?a=1&b=9#x0 HTTP/1.0\r\nhost: api.udache.com\r\n" +
		"x-forwarded-for: 112.17.236.196, 112.17.236.196\r\n\r\n" +
		"a3_token=P9yQI7nN9WGXu1SsYxXYoR%2FXnjhOiUndnIMoySDpzU%2Bn9CsHpQ3yfvBcv7n6gVNtVVM%2FbQz0oSA25MVoU3VTEQlredhJ7rIsoM60sCCiz%2FsZYsstua1D4O1GVfy2cNZbFhuqbKZKIm3QB98Msbnk8OquWy1LEYQisheTysiHr2z4%2BSnY6WcDSffGdC9RtQTD&datatype=2&dest_lat=29.345328354702122&new_price_conf=1&ticket=OJqjvzCXglMUS-Fzy_HgW1RbTSPMUdneGrIRN3IMndAkzDmuwzAMANG7TE0YJGVKFtvf_ztkcZZGARKkMnL3wEj9BrMxlKRMOinCMNKE4aSrqgqjkNaie1hrtcx155ncKUj-_hEOJAhHMhaN2b0szWqrKpxJN2ElN16P9_O0kvoRLqRFacW1lS5cSSw8lh5ReyDcfs_7nn8DAAD__w%3D%3D&maptype=soso&appversioncode=321&highway_fee=%E8%AF%B7%E8%BE%93%E5%85%A5%E9%87%91%E9%A2%9D&imei=863670033562410&appversion=5.1.24&lang=zh-CN&park_fee=%E8%AF%B7%E8%BE%93%E5%85%A5%E9%87%91%E9%A2%9D&platform_type=2&bridge_fee=%E8%AF%B7%E8%BE%93%E5%85%A5%E9%87%91%E9%A2%9D&channel=2010000001&is_offline_paid=0&dest_lng=120.10726630826768&product_id=260&utc_offset=480&dest_timestamp=1539762330&terminal_id=1&deviceid=7852c53bb1047946641cdbdc23646f0f&lng=120.10726630826768&os=6.0.1&oid=TWpnNE5UZzBOekEyT0RBd09ETTFOek14&plutus_data=pbv3H4sIAAAAAAAA%2F%2ByUPWgUURDHd8yqkWuMNjHVuZUfm2PmvZ2d3RVxAzZaaRMionE%2F4xUmeqSIBOGUFLE8UvhVKHZJE7ELdoqIYPADhFSCohZ2goWlXMhbkmgR%2B8wW%2B2P%2B8%2BY%2FPIbXe%2Bf9g08791oDt3b0rnyfX2y3Z2BfjVhxEDL7IfffnXv06vNRp0baZz8IhDQdstAOrNg6PSQo6KWiSuSCk%2B4hhamXljrLM8w8ivIkFUpREs48VqUq0yLjouQsJJ9SSc8Mj1iXvs7svgxXe6asNsBtsNaiU9F92K%2BQgkHCQZI6caRVpHm%2B0p9AuQTLb7pxLH4O1utK%2BQDWCthfwEaW8AfUuj9Xa5RA%2FayKfsPc0rcX79q207F717Dv17P8RM%2Fiy%2FjAakl9Nj5oModNhkyD0MCQgVMGzho4ZyAx0DRwzcB1Azer2f51HfcqegyzMO2kxVhzfHSyeaVwImIdik8eousUU0k2ObqqVgL7ZIRiPDdppZW4TqvIJlp50XKi88wYNFCxFiJRniiXOVQNZM0ivs%2BkxMX134UbCzD9n6Mwb2UUz%2BMGrg%2Ff1X%2B7P4XhLbrjBlPc4IUubm58ZHsntndiU2M1sCcYO9mcGDke9z1cXnj7cdfFuL%2FTfUjr8CcAAP%2F%2FOetnUVEFAAA%3D&biz_type=2&lat=29.345328354702122&location_cityid=86&model=OPPO+R9s"
	request2 := "POST /gulfstream/driver/v2/driver/dFinishOrder?a=1&b=9#x0 HTTP/1.0\r\nhost: api.udache.com\r\n" +
		"x-forwarded-for: 112.17.236.196, 112.17.236.196\r\n\r\n" +
		"a3_token=P9yQI7nN9WGXu1SsYxXYoR%2FXnjhOiUndnIMoySDpzU%2Bn9CsHpQ3yfvBcv7n6gVNtVVM%2FbQz0oSA25MVoU3VTEQlredhJ7rIsoM60sCCiz%2FsZYsstua1D4O1GVfy2cNZbFhuqbKZKIm3QB98Msbnk8OquWy1LEYQisheTysiHr2z4%2BSnY6WcDSffGdC9RtQTD&datatype=2&dest_lat=29.345328354702122&new_price_conf=1&ticket=OJqjvzCXglMUS-Fzy_HgW1RbTSPMUdneGrIRN3IMndAkzDmuwzAMANG7TE0YJGVKFtvf_ztkcZZGARKkMnL3wEj9BrMxlKRMOinCMNKE4aSrqgqjkNaie1hrtcx155ncKUj-_hEOJAhHMhaN2b0szWqrKpxJN2ElN16P9_O0kvoRLqRFacW1lS5cSSw8lh5ReyDcfs_7nn8DAAD__w%3D%3D&maptype=soso&appversioncode=321&highway_fee=%E8%AF%B7%E8%BE%93%E5%85%A5%E9%87%91%E9%A2%9D&imei=863670033562410&appversion=5.1.24&lang=zh-CN&park_fee=%E8%AF%B7%E8%BE%93%E5%85%A5%E9%87%91%E9%A2%9D&platform_type=2&bridge_fee=%E8%AF%B7%E8%BE%93%E5%85%A5%E9%87%91%E9%A2%9D&channel=2010000001&is_offline_paid=0&dest_lng=120.10726630826768&product_id=260&utc_offset=480&dest_timestamp=1539762330&terminal_id=1&deviceid=7852c53bb1047946641cdbdc23646f0f&lng=120.10726630826768&os=6.0.1&oid=TWpnNE5UZzBOekEyT0RBd09ETTFOek14&plutus_data=pbv3H4sIAAAAAAAA%2F%2ByUPWgUURDHd8yqkWuMNjHVuZUfm2PmvZ2d3RVxAzZaaRMionE%2F4xUmeqSIBOGUFLE8UvhVKHZJE7ELdoqIYPADhFSCohZ2goWlXMhbkmgR%2B8wW%2B2P%2B8%2BY%2FPIbXe%2Bf9g08791oDt3b0rnyfX2y3Z2BfjVhxEDL7IfffnXv06vNRp0baZz8IhDQdstAOrNg6PSQo6KWiSuSCk%2B4hhamXljrLM8w8ivIkFUpREs48VqUq0yLjouQsJJ9SSc8Mj1iXvs7svgxXe6asNsBtsNaiU9F92K%2BQgkHCQZI6caRVpHm%2B0p9AuQTLb7pxLH4O1utK%2BQDWCthfwEaW8AfUuj9Xa5RA%2FayKfsPc0rcX79q207F717Dv17P8RM%2Fiy%2FjAakl9Nj5oModNhkyD0MCQgVMGzho4ZyAx0DRwzcB1Azer2f51HfcqegyzMO2kxVhzfHSyeaVwImIdik8eousUU0k2ObqqVgL7ZIRiPDdppZW4TqvIJlp50XKi88wYNFCxFiJRniiXOVQNZM0ivs%2BkxMX134UbCzD9n6Mwb2UUz%2BMGrg%2Ff1X%2B7P4XhLbrjBlPc4IUubm58ZHsntndiU2M1sCcYO9mcGDke9z1cXnj7cdfFuL%2FTfUjr8CcAAP%2F%2FOetnUVEFAAA%3D&biz_type=2&lat=29.345328354702122&location_cityid=86&model=OPPO+R9s"
	r1, _ := httpObj.Parse(request1)
	r2, _ := httpObj.Parse(request2)
	fmt.Println(httpObj.Match(r1, r2))
	should.True(true)
}

func parse(hello interface{}) (result string){
	tmp, _ := json.Marshal(hello)
	return string(tmp)
}
