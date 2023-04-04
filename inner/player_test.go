package inner_test

import (
	"github.com/smartwalle/utube/inner"
	"testing"
)

func TestGetPlayer(t *testing.T) {
	t.Log("========== GetPlayer ==========")
	var _, err = inner.GetPlayer(inner.ClientTypeWeb, "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8", "OtmWv_cgZqk")
	if err != nil {
		t.Fatal(err)
		return
	}

	//t.Log(info.StreamingData.Formats[0].URL)
}
