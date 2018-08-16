package godisco

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Requester interface allowing testing further down the line
type Requester interface {
	do(*http.Request) ([]byte, int, error)
	Get(string) ([]byte, int, error)
	Post(string, []byte) ([]byte, int, error)
	Put(string, []byte) ([]byte, int, error)
}

// Client struct to keep track of new client
type Client struct {
	client *http.Client
	domain string
	key    string
	user   string
}

// NewClient Declare new HTTP client
func NewClient(ClientEndpoint string, ClientKey string, ClientUser string) (*Client, error) {
	// Send the request via a client
	client := &http.Client{}
	var domain string
	// Check if domain has proper protocol
	if strings.HasPrefix(ClientEndpoint, "http") {
		domain = ClientEndpoint
	} else {
		domain = "https://" + ClientEndpoint
	}
	return &Client{
		client: client,
		domain: domain,
		key:    ClientKey,
		user:   ClientUser,
	}, nil
}

// Get resource string
func (c *Client) Get(resource string) ([]byte, int, error) {
	url := fmt.Sprintf("%s%s", c.domain, resource)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}
	return c.do(req)
}

func (c *Client) do(req *http.Request) ([]byte, int, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 404 {
		err = fmt.Errorf("Received unexpected status %d while trying to retrieve the server data with \"%s\"", resp.StatusCode, string(body))
		return nil, resp.StatusCode, err
	}
	return body, resp.StatusCode, nil
}

// Post to resource string the data provided
func (c *Client) Post(resource string, data []byte) ([]byte, int, error) {
	apiAuth := url.Values{}
	apiAuth.Set("api_key", c.key)
	apiAuth.Add("api_username", c.user)
	url := fmt.Sprintf("%s%s?%s", c.domain, resource, apiAuth.Encode())
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	return c.do(req)
}

// Put to resource string the data provided
func (c *Client) Put(resource string, data []byte) ([]byte, int, error) {
	log.Printf("PUT Data: %v", string(data))
	apiAuth := url.Values{}
	apiAuth.Set("api_key", c.key)
	apiAuth.Add("api_username", c.user)
	url := fmt.Sprintf("%s%s?%s", c.domain, resource, apiAuth.Encode())
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	return c.do(req)
}
