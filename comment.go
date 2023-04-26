package utube

import "net/http"

const (
	kCommentThreads = "/v3/commentThreads"
	kComments       = "/v3/comments"
)

func (this *Client) GetCommentThreads(param GetCommentThreadsParam) (results *CommentThreads, err error) {
	var api = this.BuildAPI(kCommentThreads)
	err = this.doRequest(http.MethodGet, api, param, &results)
	return results, err
}

func (this *Client) GetComments(param GetCommentsParams) (results *Comments, err error) {
	var api = this.BuildAPI(kComments)
	err = this.doRequest(http.MethodGet, api, param, &results)
	return results, err
}
