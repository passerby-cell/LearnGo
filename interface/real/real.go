package real

import (
	"net/http"
	"net/http/httputil"
)

type Receiver struct {
}

func (receiver Receiver) GetRes(s string) string {
	response, err := http.Get(s)
	if err != nil {
		panic(err)
		return ""
	}
	dumpResponse, err := httputil.DumpResponse(response, true)
	response.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(dumpResponse)

}
