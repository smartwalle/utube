package utube_test

import (
	"github.com/smartwalle/utube"
	"testing"
)

func TestYoutube_GetVideos(t *testing.T) {
	t.Log("========== Videos ==========")
	var c = GetYoutube()
	var p = utube.GetVideosParam{}
	p.Chart = "mostPopular"
	var vs, err = c.GetVideos(p)
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, v := range vs.Items {
		t.Log(v.Id, v.Snippet.Title, v.Snippet.Thumbnails.GetThumbnailURL())
	}
}

func TestYoutube_GetVideoCategories(t *testing.T) {
	t.Log("========== Video Categories ==========")
	var c = GetYoutube()
	var p = utube.GetVideoCategoriesParam{}
	p.RegionCode = "US"
	p.HL = "zh_CN"
	var vs, err = c.GetVideoCategories(p)
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, v := range vs.Items {
		t.Log(v.Snippet.ChannelId, v.Snippet.Title)
	}
}
