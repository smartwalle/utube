package utube

import "net/http"

const (
	kChannelsAPI = "/v3/channels"
)

// GetChannels https://developers.google.com/youtube/v3/docs/channels/list
func (this *Client) GetChannels(param GetChannelsParam) (results *Channels, err error) {
	var api = this.BuildAPI(kChannelsAPI)
	err = this.doRequest(http.MethodGet, api, param, &results)
	return results, err
}
