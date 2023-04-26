package utube

import (
	"context"
	"github.com/smartwalle/ngx"
	"net/http"
	"strings"
)

const (
	kYoutube = "https://www.googleapis.com/youtube"
)

type Client struct {
	Client      *http.Client
	key         string
	accessToken string
	host        string
}

func New(key, accessToken string) *Client {
	var nClient = &Client{}
	nClient.key = key
	nClient.accessToken = accessToken
	nClient.host = kYoutube
	return nClient
}

func (this *Client) BuildAPI(paths ...string) string {
	var path = this.host
	for _, p := range paths {
		p = strings.TrimSpace(p)
		if len(p) > 0 {
			if strings.HasSuffix(path, "/") {
				path = path + p
			} else {
				if strings.HasPrefix(p, "/") {
					path = path + p
				} else {
					path = path + "/" + p
				}
			}
		}
	}
	return path
}

func (this *Client) doRequest(method, url string, param Param, result interface{}) (err error) {
	var values = param.Values()
	if this.key != "" {
		values.Add("key", this.key)
	}

	var req = ngx.NewRequest(method, url, ngx.WithClient(this.Client))
	req.SetParams(values)
	if this.accessToken != "" {
		req.SetHeader("Authorization", "Bearer "+this.accessToken)
	}

	rsp := req.Exec(context.Background())

	switch rsp.StatusCode() {
	case http.StatusOK:
		if err = rsp.UnmarshalJSON(result); err != nil {
			return err
		}
	default:
		var rErr *ResponseError
		if err = rsp.UnmarshalJSON(&rErr); err != nil {
			return err
		}

		rErr.Method = method
		rErr.URL = rsp.Request().URL.String()
		return rErr
	}
	return nil
}
