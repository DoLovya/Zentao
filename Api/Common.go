package Api

import (
	"encoding/json"
	"io"
	"net/http"
)

const URL = "https://chan.aqrose.com/"

type cache struct {
	Token string
}

var Cache cache

func UnmarshalResponse(resp *http.Response, data any) error {
	respByteStream, _ := io.ReadAll(resp.Body)
	err := json.Unmarshal(respByteStream, data)
	if err != nil {
		return err
	}
	return nil
}

func Get(url string, v any) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	// 设置请求头
	req.Header.Set("Token", Cache.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err1 := UnmarshalResponse(resp, v)
	if err1 != nil {
		return err1
	}
	return nil
}
