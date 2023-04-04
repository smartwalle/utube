# utube

一个 Youtube 相关的工具类库

##### GetVideoMeta

本方法在 inner 子包中，不需要授权认证, 即可根据 Youtube 视频的 ID 获取该视频的详细信息, 包含下载地址。

```
func GetVideoMeta(key, videoId string) (videoMeta *VideoMeta, err error)
```