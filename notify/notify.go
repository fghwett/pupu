package notify

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/fghwett/pupu/util"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Send(secretKey string, title string, content string) error {
	reqUrl := fmt.Sprintf("https://sctapi.ftqq.com/%s.send?title=%s&desp=%s", secretKey, title, url.QueryEscape(content))

	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return err
	}

	response := &Response{}
	resp, err := http.DefaultClient.Do(req)

	err = util.GetHTTPResponse(resp, reqUrl, err, response)
	if err != nil {
		return err
	}

	if response.Code != 0 {
		return fmt.Errorf(response.Message)
	}

	return nil
}
