package utube_test

import (
	"github.com/smartwalle/utube"
	"testing"
)

// 搜索指定 Channel 下的视频
func TestYoutube_SearchVideoWithChannel(t *testing.T) {
	t.Log("========== Search ==========")
	var c = GetYoutube()
	var p = utube.SearchParam{}
	p.ChannelId = "UCBR8-60-B28hp2BmDPdntcQ "
	var rs, err = c.SearchVideo(p)
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, s := range rs.Items {
		t.Log(s.GetId(), s.Id.Kind, s.Snippet.Title)
	}
}
