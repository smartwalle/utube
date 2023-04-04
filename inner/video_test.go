package inner_test

import (
	"github.com/smartwalle/utube/inner"
	"testing"
)

func TestGetVideoMeta(t *testing.T) {
	t.Log("========== GetVideoMeta ==========")
	var info, err = inner.GetVideoMeta("AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8", "OtmWv_cgZqk")
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(info.StreamingData.Formats[0].URL)
}
