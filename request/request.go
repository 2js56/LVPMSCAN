package request

import (
	"LVPMSCAN/Core"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/requests"
)

func Reqdata(url string) (*Core.FetchResult, error) {
	req := requests.Requests()
	req.SetTimeout(time.Duration(10))
	resp, err := req.Get(url)

	if err != nil {
		fmt.Println(err)
	}
	var headerString string
	req_data := Core.FetchResult{
		Url:           url,
		Content:       resp.Content(),
		Headers:       resp.R.Header,
		HeadersString: headerString,
		Certs:         GetCerts(resp.R),
	}
	return &req_data, nil
}

func GetCerts(resp *http.Response) []byte {
	var certs []byte
	if resp.TLS != nil {
		cert := resp.TLS.PeerCertificates[0]
		var str string
		if js, err := json.Marshal(cert); err == nil {
			certs = js
		}
		str = string(certs) + cert.Subject.String()
		certs = []byte(str)
	}
	return certs
}
