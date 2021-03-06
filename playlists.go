package utube

const (
	k_PLAYLISTS_API      = "/v3/playlists"
	k_PLAYLIST_ITEMS_API = "/v3/playlistItems"
)

// GetPlaylists https://developers.google.cn/youtube/v3/docs/playlists/list
func (this *Client) GetPlaylists(param GetPlaylistsParam) (results *Playlists, err error) {
	var api = this.BuildAPI(k_PLAYLISTS_API)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}

// GetPlaylistItems https://developers.google.cn/youtube/v3/docs/playlistItems/list
func (this *Client) GetPlaylistItems(param GetPlaylistItemsParam) (results *PlaylistItems, err error) {
	var api = this.BuildAPI(k_PLAYLIST_ITEMS_API)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}
