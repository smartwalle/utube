package utube

import "net/http"

const (
	kVideoAPI        = "/v3/videos"
	kVideoCategories = "/v3/videoCategories"
)

// GetVideos https://developers.google.cn/youtube/v3/docs/videos/list
func (this *Client) GetVideos(param GetVideosParam) (results *Videos, err error) {
	var api = this.BuildAPI(kVideoAPI)
	err = this.doRequest(http.MethodGet, api, param, &results)
	return results, err
}

// GetVideoCategories https://developers.google.cn/youtube/v3/docs/videoCategories/list
func (this *Client) GetVideoCategories(param GetVideoCategoriesParam) (results *VideoCategories, err error) {
	var api = this.BuildAPI(kVideoCategories)
	err = this.doRequest(http.MethodGet, api, param, &results)
	return results, err
}
