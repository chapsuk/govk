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
	language     int
	callFunc     func(uri string, s interface{}) error
}

// NewClient yiled new Client structure
func NewClient(id, secret, v string, lng int) *Client {
	return &Client{
		clientID:     id,
		clientSecret: secret,
		language:     lng,
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

// UserIsAppUser call user.isAppUser vk api method
// return true if user install current app (by clientID)
func (c *Client) UserIsAppUser(id int, token string) (bool, error) {
	v := url.Values{}
	v.Add("user_id", strconv.Itoa(id))
	v.Add("access_token", token)

	uri := c.buildURLForMethod("users.isAppUser", v)
	var yep string
	err := c.send(uri, &yep)
	if yep == "1" {
		return true, err
	}
	return false, err
}

// OrdersGet call order.get vk api method
// size - max 1000, default 0
// offset - default 0
// test - 1 or 0, enable or disable test mode
func (c *Client) OrdersGet(count, offset int, test int) ([]OrderResponse, error) {
	v := url.Values{}
	v.Add("count", strconv.Itoa(count))
	v.Add("offset", strconv.Itoa(offset))
	v.Add("test_mode", strconv.Itoa(test))
	v.Add("access_token", c.AccessToken)

	uri := c.buildURLForMethod("orders.get", v)
	res := []OrderResponse{}
	err := c.send(uri, &res)
	return res, err
}

// DatabaseGetCountries call database.getCountries vk api method
// count - max 1000, default 100
// offset - default 0
// code - "RU,UA,BY" for exampl
// all - if true return all countries
func (c *Client) DatabaseGetCountries(count, offset int, all bool, code string) ([]CountryResponse, error) {
	v := url.Values{}
	if count > 0 {
		v.Add("count", strconv.Itoa(count))
	}
	if offset > 0 {
		v.Add("offset", strconv.Itoa(offset))
	}
	if all {
		v.Add("need_all", "1")
	} else {
		v.Add("need_all", "0")
	}
	if code != "" {
		v.Add("code", code)
	}

	uri := c.buildURLForMethod("database.getCountries", v)
	res := GetContriesResponse{}
	err := c.send(uri, &res)
	return res.Items, err
}

// DatabaseGetCities call database.getCities vk api method
// count - max 1000, default 100
// offset - default 0
// all - if false return only important
// countryID - required param
// regionID - optional
// query - part of city name
func (c *Client) DatabaseGetCities(count, offset int, all bool, countryID, regionID int, query string) ([]CityResponse, error) {
	v := url.Values{}
	if count > 0 {
		v.Add("count", strconv.Itoa(count))
	}
	if offset > 0 {
		v.Add("offset", strconv.Itoa(offset))
	}
	if all {
		v.Add("need_all", "1")
	} else {
		v.Add("need_all", "0")
	}
	if countryID != 0 {
		v.Add("country_id", strconv.Itoa(countryID))
	}
	if regionID != 0 {
		v.Add("region_id", strconv.Itoa(regionID))
	}
	if query != "" {
		v.Add("q", query)
	}

	uri := c.buildURLForMethod("database.getCities", v)
	res := GetCitiesResponse{}
	err := c.send(uri, &res)
	return res.Items, err
}

// DatabaseGetCitiesByID call database.getCitiesById vk api method
func (c *Client) DatabaseGetCitiesByID(ids string) ([]CityByIDResponse, error) {
	v := url.Values{}
	v.Add("city_ids", ids)

	uri := c.buildURLForMethod("database.getCitiesById", v)
	res := []CityByIDResponse{}
	err := c.send(uri, &res)
	return res, err
}

func (c Client) buildURLForMethod(method string, p url.Values) string {
	p.Add("v", c.apiVersion)
	p.Add("lang", strconv.Itoa(c.language))
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
