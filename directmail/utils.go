package directmail

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//control over the request lifecycle
var httpClient = &http.Client{
	Timeout: 2000 * time.Millisecond,
}

func SetClientTimeout(duration time.Duration) {
	httpClient.Timeout = duration
}

func httpReqWithParams(method, url, body string) ([]byte, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("ReadAll err : %v\n", err)
		}

		var result EmailSendSingleResp
		err = json.Unmarshal(data, &result)
		if err != nil {
			fmt.Printf("Unmarshal err: %v\n", err)
		}

		fmt.Printf("httpReqWithParams err: %v, result: %v\n", resp, result)
		return nil, fmt.Errorf(resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}
