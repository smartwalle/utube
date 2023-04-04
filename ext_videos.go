package utube

//
//import (
//	"context"
//	"github.com/smartwalle/form"
//	"github.com/smartwalle/ngx"
//	"net/url"
//)
//
//const (
//	kGetVideoInfoURL = "https://www.youtube.com/get_video_info"
//)
//
//func GetVideoInfoWithVideoId(videoId string) (videoInfo *VideoInfo, err error) {
//	var req = ngx.NewRequest("GET", kGetVideoInfoURL)
//	req.SetParam("video_id", videoId)
//
//	//result, err := request.Request("GET", kGetVideoInfoURL, url.Values{"video_id": []string{videoId}})
//	//if err != nil {
//	//	return nil, err
//	//}
//	var rep = req.Exec(context.Background())
//
//	var queryStr = rep.MustString()
//	videoInfo, err = VideoInfoWithQueryString(queryStr)
//	return videoInfo, err
//}
//
//func VideoInfoWithQueryString(query string) (videoInfo *VideoInfo, err error) {
//	param, err := url.ParseQuery(query)
//	if err != nil {
//		return nil, err
//	}
//
//	err = form.Bind(param, &videoInfo)
//	if err != nil {
//		return nil, err
//	}
//	return videoInfo, nil
//}
