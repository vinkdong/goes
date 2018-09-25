package socket

import (
	"github.com/vinkdong/gox/log"
	"net/url"
)

type Client struct {
	ServerHost string
	Scheme     string
	Path       string
}

func (c *Client) Run() {
	u := url.URL{
		Scheme: "ws",
		Host:   c.ServerHost,
		Path:   c.Path,
	}
	log.Info(u)
}
