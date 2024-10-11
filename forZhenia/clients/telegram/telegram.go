package telegram

import (
	"net/http"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host, token string) Client {
	return Client{
		host:     host,
		basePath: NewBasePath(token),
		client:   http.Client{},
	}

}
func NewBasePath(token string) string {
	return "bot" + token
}
func (c *Client) Updates() {

}
func (c *Client) SendMessage(message string) {

}
