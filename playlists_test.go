package utube_test

import (
	"github.com/smartwalle/utube"
	"testing"
)

func TestYoutube_Playlists(t *testing.T) {
	t.Log("========== Playlists ==========")
	var c = GetYoutube()
	var p = utube.GetPlaylistsParam{}
	p.ChannelId = "UCgZd5ygXFoQry9KDGlddSBg"
	var ps, err = c.GetPlaylists(p)
	if err != nil {
		t.Fatal(err)
		return
	}

	for _, p := range ps.Items {
		t.Log(p.Id, p.Snippet.Title)
	}
}

func TestYoutube_GetPlaylistItems(t *testing.T) {
	t.Log("========== PlaylistItems ==========")
	var c = GetYoutube()
	var p = utube.GetPlaylistItemsParam{}
	p.PlaylistId = "PLKD8bdRWqx74T40SxvhxFsp-oFzOAT0dj"
	var ps, err = c.GetPlaylistItems(p)
	if err != nil {
		t.Fatal(err)
		return
	}

	for _, p := range ps.Items {
		t.Log(p.Id, p.Snippet.ResourceId.GetId())
	}

}
