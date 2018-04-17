package oschina

type OSChinaConfig struct {
	ApiKey      string
	ApiSecret   string
	RedirectURI string
}

//client_id=%s&client_secret=%s&grant_type=refresh_token&redirect_uri=%s&refresh_token=%s&dataType=json
type TokenRequest struct {
	ClientID     string `org:"client_id"`
	ClientSecret string `org:"client_secret"`
	GrantType    string `org:"grant_type"`
	RedirectURI  string `org:"redirect_uri"`
	RefreshToken string `org:"refresh_token"`
	DataType     string `org:"dataType"`
}

//access_token=%s&authorname=%s&page=1&pageSize=5&dataType=json
type BlogListRequest struct {
	AccessToken string `org:"access_token"`
	Authorname  string `org:"authorname"`
	Authoruid   string `org:"authoruid"`
	Page        int    `org:"page"`
	PageSize    int    `org:"pageSize"`
	DataType    string `org:"dataType"`
}

type BlogCommentRequest struct {
	AccessToken string `org:"access_token"`
	Blog        int    `org:"blog"`
	Content     string `org:"content"`
	DataType    string `org:"dataType"`
}

type TokenResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expires      int64  `json:"expires_in"`
}

type BlogListResponse struct {
	Count  int `json:"count"`
	Notice struct {
		ReplyCount int `json:"replyCount"`
		MsgCount   int `json:"msgCount"`
		FansCount  int `json:"fansCount"`
		ReferCount int `json:"referCount"`
	} `json:"notice"`
	Bloglist []struct {
		ID           int    `json:"id"`
		PubDate      string `json:"pubDate"`
		Author       string `json:"author"`
		Title        string `json:"title"`
		Authorid     int    `json:"authorid"`
		Type         int    `json:"type"`
		CommentCount int    `json:"commentCount"`
	} `json:"projectlist"`
}

type CommentResponse struct {
	Code     string `json:"error"`
	CodeDesc string `json:"error_description"`
}
