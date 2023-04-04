package utube_test

import (
	"github.com/smartwalle/utube"
	"os"
	"testing"
)

var client *utube.Client

func TestMain(m *testing.M) {
	client = utube.New("AIzaSyC8tmV8M2dgVtrvOhkeYxqhWEkXyl5XhW8", "")
	var exitCode = m.Run()
	os.Exit(exitCode)
}

func GetYoutube() *utube.Client {
	return client
}
