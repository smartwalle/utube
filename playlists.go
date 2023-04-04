package utube

import "net/http"

const (
	kPlaylistsAPIAPI   = "/v3/playlists"
	kPlaylistsItemsAPI = "/v3/playlistItems"
)

// GetPlaylists https://developers.google.cn/youtube/v3/docs/playlists/list
func (this *Client) GetPlaylists(param GetPlaylistsParam) (results *Playlists, err error) {
	var api = this.BuildAPI(kPlaylistsAPIAPI)
	err = this.doRequest(http.MethodGet, api, param, &results)
	return results, err
}

// GetPlaylistItems https://developers.google.cn/youtube/v3/docs/playlistItems/list
func (this *Client) GetPlaylistItems(param GetPlaylistItemsParam) (results *PlaylistItems, err error) {
	var api = this.BuildAPI(kPlaylistsItemsAPI)
	err = this.doRequest(http.MethodGet, api, param, &results)
	return results, err
}
