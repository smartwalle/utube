package utube_test

import (
	"github.com/smartwalle/utube"
	"testing"
)

func TestYoutube_GetChannels(t *testing.T) {
	t.Log("========== GetChannels ==========")
	var c = GetYoutube()
	var p = utube.GetChannelsParam{}
	p.ForUsername = "RiPTrippers"
	var rs, err = c.GetChannels(p)
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, r := range rs.Items {
		t.Log(r.Id, r.Snippet.Title)
	}
}
