package govk

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

var authURLTpl = "https://oauth.vk.com/access_token?client_id=%s&client_secret=%s&v=%s&grant_type=client_credentials"
var apiEndpointTpl = "https://api.vk.com/method/%s?%s"

// Client to vk api
type Client struct {
	clientID     string
	clientSecret string
	apiVersion   string
	AccessToken  string
	callFunc     func(uri string, s interface{}) error
}

// NewClient yiled new Client structure
func NewClient(id, secret, v string) *Client {
	return &Client{
		clientID:     id,
		clientSecret: secret,
		apiVersion:   v,
		callFunc:     call,
	}
}

// Auth is call server authorization method and get access_token
func (c *Client) Auth() error {
	uri := fmt.Sprintf(authURLTpl, c.clientID, c.clientSecret, c.apiVersion)
	res := SuccessAuthResponse{}
	err := c.callFunc(uri, &res)
	if err != nil {
		return err
	}
	c.AccessToken = res.AccessToken
	return nil
}

// OrdersGet call order.get vk api method
// size - max 1000, default 0
// offset - default 0
// test - 1 or 0, enable or disable test mode
func (c *Client) OrdersGet(size, offset int, test int) ([]OrderResponse, error) {
	v := url.Values{}
	v.Add("count", strconv.Itoa(size))
	v.Add("offset", strconv.Itoa(offset))
	v.Add("test_mode", strconv.Itoa(test))
	v.Add("version", c.apiVersion)
	v.Add("access_token", c.AccessToken)

	uri := buildURLForMethod("orders.get", v)
	res := []OrderResponse{}
	err := c.send(uri, &res)
	return res, err
}

func buildURLForMethod(method string, p url.Values) string {
	return fmt.Sprintf(apiEndpointTpl, method, p.Encode())
}

func (c *Client) send(uri string, r interface{}) error {
	res := VkResponse{}
	err := c.callFunc(uri, &res)
	if err != nil {
		return err
	}
	if res.Error != nil {
		e := &VkErrorResponse{}
		err = json.Unmarshal(res.Error, e)
		if err != nil {
			return fmt.Errorf("error: \"%s\" on parsing error_message, response: \"%s\"", err.Error(), res.Error)
		}
		return fmt.Errorf("error response: %s", e.Message)
	}
	err = json.Unmarshal(res.Response, r)
	if err != nil {
		return fmt.Errorf("error: \"%s\" parse response, response:  \"%s\"", err.Error(), res.Response)
	}
	return nil
}

func call(uri string, s interface{}) error {
	r, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return makeHTTPErrorResponse(r.Body)
	}
	return json.NewDecoder(r.Body).Decode(s)
}

func makeHTTPErrorResponse(r io.Reader) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return fmt.Errorf("error response, msg: %s", body)
}
