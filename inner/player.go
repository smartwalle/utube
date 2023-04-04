package inner

import (
	"context"
	"fmt"
	"github.com/smartwalle/ngx"
	"net/http"
	"strings"
)

const (
	kGetVideoURL = "https://www.youtube.com/youtubei/v1/player"
)

func GetPlayer(clientType ClientType, key, videoId string) (playerInfo *PlayerInfo, err error) {
	var req = ngx.NewRequest(http.MethodPost, kGetVideoURL)
	req.AddQuery("key", key)

	req.SetContentType(ngx.ContentTypeJSON)

	var nBody = strings.NewReader(fmt.Sprintf(clients[clientType], videoId))
	req.SetBody(nBody)

	var rsp = req.Exec(context.Background())

	if err = rsp.UnmarshalJSON(&playerInfo); err != nil {
		return nil, err
	}

	//fmt.Println(rsp.Error())
	fmt.Println(rsp.MustString())
	return playerInfo, nil
}
