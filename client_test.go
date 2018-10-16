package godisco

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type testClient struct {
	client *http.Client
	domain string
	key    string
	user   string
}

func newTestClient() *testClient {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`OK`))
	}))
	client := &testClient{
		client: server.Client(),
		domain: server.URL,
		key:    "key",
		user:   "user",
	}
	return client
}

// Get resource string
func (c *testClient) Get(resource string) ([]byte, int, error) {
	return []byte{}, 0, nil
}

func (c *testClient) do(req *http.Request) ([]byte, int, error) {
	return []byte{}, 0, nil
}

// Post to resource string the data provided
func (c *testClient) Post(resource string, data []byte) ([]byte, int, error) {
	return []byte{}, 0, nil
}

// Put to resource string the data provided
func (c *testClient) Put(resource string, data []byte) ([]byte, int, error) {
	return []byte{}, 0, nil
}

func TestPut(t *testing.T) {
	client := newTestClient()
	_, _, err := client.Put("/resource", []byte{})

	if err != nil {
		t.Errorf("expected: %v actual: %v", nil, err)
	}

}
