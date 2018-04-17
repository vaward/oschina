package oschina

const BLOGLISTURL string = "https://www.oschina.net/action/openapi/user_blog_list"
const REFRESHURL string = "https://www.oschina.net/action/openapi/token"
const COMMENTURL string = "https://www.oschina.net/action/openapi/blog_comment_pub"

var bloglistreq = BlogListRequest{
	AccessToken: "",
	Authorname:  "",
	Authoruid:   "", //Authorname 二选一
	Page:        1,
	PageSize:    5,
}

var refreshtokenreq = TokenRequest{
	ClientID:     "",
	ClientSecret: "",
	RedirectURI:  "",
	GrantType:    "refresh_token",
	RefreshToken: "",
	DataType:     "json",
}

var commentreq = BlogCommentRequest{
	AccessToken: "",
	Blog:        22,
	Content:     "",
	DataType:    "json",
}

// GetBlog ..
/***
* 获取Blog列表
***/
func GetBlog(authorname, authoruid, accessToken string, bloglistres *BlogListResponse) error {
	bloglistreq.AccessToken = accessToken
	bloglistreq.Authorname = authorname
	bloglistreq.Authoruid = authoruid
	return HTTPPost(BLOGLISTURL, bloglistreq, bloglistres)
}

//GetRefreshToken 刷新token
func GetRefreshToken(o OSChinaConfig, refreshToken string, tokenResp *TokenResp) error {
	refreshtokenreq.ClientID = o.ApiKey
	refreshtokenreq.ClientSecret = o.ApiSecret
	refreshtokenreq.RedirectURI = o.RedirectURI
	refreshtokenreq.RefreshToken = refreshToken
	return HTTPPost(REFRESHURL, refreshtokenreq, tokenResp)
}

//PubComments 发表评论
func PubComments(accessToken, content string, blogid int, commentResponse *CommentResponse) error {
	commentreq.AccessToken = accessToken
	commentreq.Content = content
	commentreq.Blog = blogid
	return HTTPPost(COMMENTURL, commentreq, commentResponse)
}
