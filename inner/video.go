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

// GetVideoMeta 获取视频信息
// 可以获取到视频的下载地址，下载地址在 VideoMeta.StreamingData.Formats 中
//
// 本方法需要的 key 无法申请，目前已经以下 key 可用：
// AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8
// AIzaSyCtkvNIR1HCEwzsqK6JuE6KqpyjusIRI30
// AIzaSyA8eiZmM1FaDVjRy-df2KTyQ_vz_yYM39w
// AIzaSyC8UYZpvA2eknNex0Pjid0_eTLJoDu6los
// AIzaSyCjc_pVEDi4qsv5MtC2dMXzpIaDoRFLsxw
// AIzaSyDHQ9ipnphqTzDqZsbtd8_Ru4_kiKVQe2k
func GetVideoMeta(key, videoId string) (videoMeta *VideoMeta, err error) {
	var req = ngx.NewRequest(http.MethodPost, kGetVideoURL)
	req.AddQuery("key", key)

	req.SetContentType(ngx.ContentTypeJSON)

	var nBody = strings.NewReader(fmt.Sprintf(`{"context": {"client": {"clientName": "WEB","clientVersion": "2.20200720.00.02"}},"videoId": "%s"}`, videoId))
	req.SetBody(nBody)

	var rsp = req.Exec(context.Background())

	if err = rsp.UnmarshalJSON(&videoMeta); err != nil {
		return nil, err
	}
	return videoMeta, nil
}
