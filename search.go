package utube

import "net/http"

const (
	kSearchAPI = "/v3/search"
)

// Search https://developers.google.cn/youtube/v3/docs/search/list
func (this *Client) Search(param SearchParam) (results *SearchResults, err error) {
	var api = this.BuildAPI(kSearchAPI)
	err = this.doRequest(http.MethodGet, api, param, &results)
	return results, err
}

func (this *Client) SearchVideo(param SearchParam) (results *SearchResults, err error) {
	param.Type = "video"

	var api = this.BuildAPI(kSearchAPI)
	err = this.doRequest(http.MethodGet, api, param, &results)
	return results, err
}

func (this *Client) SearchChannel(param SearchParam) (results *SearchResults, err error) {
	param.Type = "channel"

	var api = this.BuildAPI(kSearchAPI)
	err = this.doRequest(http.MethodGet, api, param, &results)
	return results, err
}

func (this *Client) SearchPlaylist(param SearchParam) (results *SearchResults, err error) {
	param.Type = "playlist"

	var api = this.BuildAPI(kSearchAPI)
	err = this.doRequest(http.MethodGet, api, param, &results)
	return results, err
}
