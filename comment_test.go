package utube_test

import (
	"github.com/smartwalle/utube"
	"testing"
)

func TestYoutube_GetCommentThreads(t *testing.T) {
	t.Log("========== GetCommentThreads ==========")
	var c = GetYoutube()
	var p = utube.GetCommentThreadsParam{}
	p.Part.ShowSnippet()
	p.Part.ShowReplies()
	p.VideoId = "OtmWv_cgZqk"
	var cts, err = c.GetCommentThreads(p)
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, c := range cts.Items {
		t.Log(c.Id, c.Kind)
	}
}

func TestYoutube_GetComments(t *testing.T) {
	t.Log("========== GetComments ==========")
	var c = GetYoutube()
	var p = utube.GetCommentsParams{}
	p.Part.ShowSnippet()
	p.AppendId("UgzkzMAiQ9qi_-syaBN4AaABAg")
	p.AppendId("Ugxh6h4yTi99QJF_wxd4AaABAg")

	var cs, err = c.GetComments(p)
	if err != nil {
		t.Fatal(err)
		return
	}

	for _, c := range cs.Items {
		t.Log(c.Id, c.Snippet.TextDisplay)
	}
}
