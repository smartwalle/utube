package utube

const (
	kCommentThreadsAPI = "/v3/commentThreads"
	kCommentsAPI       = "/v3/comments"
)

func (this *Client) GetCommentThreads(param GetCommentThreadsParam) (results *CommentThreads, err error) {
	var api = this.BuildAPI(kCommentThreadsAPI)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}

func (this *Client) GetComments(param GetCommentsParams) (results *Comments, err error) {
	var api = this.BuildAPI(kCommentsAPI)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}
