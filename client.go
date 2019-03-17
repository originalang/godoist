package togoist

type Client struct {
	Token     string
	SyncToken string
}

func NewClient(token string) *Client {
	return &Client {
		Token: token
		SyncToken: "*"
	}
}
