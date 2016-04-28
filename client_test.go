package govk

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	test "github.com/smartystreets/goconvey/convey"
	"os"
)

func TestOrdersGet(t *testing.T) {
	cli := NewClient("123123", "fakeSecret", "5.23")

	test.Convey("orders.get", t, func() {

		test.Convey("missing access_token", func() {
			msg := `error response: User authorization failed: no access_token passed.`
			_, err := cli.OrdersGet(10, 0, 0)
			test.So(err, test.ShouldNotBeNil)
			test.So(err.Error(), test.ShouldEqual, msg)
		})

		test.Convey("success result", func() {
			expected := []OrderResponse{}
			expected = append(expected, OrderResponse{
				ID:         "123",
				AppOrderID: "123",
				Status:     "status",
				UserID:     "123",
				ReceiverID: "123",
				Item:       "item",
				Amount:     "3",
				Date:       "13213",
			})
			msg := `[{"id": "123", "app_order_id": "123", "status": "status", "user_id": "123", "receiver_id": "123", "item": "item", "amount": "3", "date": "13213"}]`
			mockCallFunc := func(uri string, r interface{}) error {
				a := r.(*VkResponse)
				a.Response = json.RawMessage([]byte(msg))
				return nil
			}
			cli.callFunc = mockCallFunc
			res, err := cli.OrdersGet(10, 0, 0)
			test.So(err, test.ShouldBeNil)
			test.So(res, test.ShouldResemble, expected)
		})
	})

}

func TestAuth(t *testing.T) {
	cli := NewClient("123123", "fakeSecret", "5.23")

	test.Convey("check auth errors", t, func() {

		test.Convey("blocked client id", func() {
			msg := `error response, msg: {"error":"invalid_client","error_description":"client_id is blocked"}`
			err := cli.Auth()
			test.So(err, test.ShouldNotBeNil)
			test.So(err.Error(), test.ShouldEqual, msg)
		})

		test.Convey("invalid client id", func() {
			msg := `error response, msg: {"error":"invalid_client","error_description":"client_id is invalid"}`
			cli.clientID = "invalidID"
			err := cli.Auth()
			test.So(err, test.ShouldNotBeNil)
			test.So(err.Error(), test.ShouldEqual, msg)
		})
	})

	test.Convey("check auth success", t, func() {
		mockCallFunc := func(uri string, r interface{}) error {
			a := r.(*SuccessAuthResponse)
			a.AccessToken = "123"
			return nil
		}
		cli.callFunc = mockCallFunc
		err := cli.Auth()
		test.So(err, test.ShouldBeNil)
		test.So(cli.AccessToken, test.ShouldEqual, "123")
	})
}

func TestBuildURLForMethod(t *testing.T) {

	test.Convey("build without params", t, func() {
		expected := "https://api.vk.com/method/order.get?"
		method := "order.get"
		uri := buildURLForMethod(method, url.Values{})
		test.So(expected, test.ShouldEqual, uri)
	})

	test.Convey("build with params", t, func() {
		expected := "https://api.vk.com/method/order.get?doo=yaa&same=13"
		method := "order.get"
		v := url.Values{}
		v.Add("same", "13")
		v.Add("doo", "yaa")
		uri := buildURLForMethod(method, v)
		test.So(expected, test.ShouldEqual, uri)
	})
}

func TestClient(t *testing.T) {

	test.Convey("create new Client instance", t, func() {
		id := "123"
		secret := "secret"
		version := "5.23"
		cli := NewClient(id, secret, version)

		test.So(id, test.ShouldEqual, cli.clientID)
		test.So(secret, test.ShouldEqual, cli.clientSecret)
		test.So(version, test.ShouldEqual, cli.apiVersion)
		test.So("", test.ShouldEqual, cli.AccessToken)
	})
}

func TestSendFunction(t *testing.T) {
	cli := NewClient("123123", "fakeSecret", "5.23")

	test.Convey("vk api call error gotten", t, func() {
		msg := "error"
		mockCallFunc := func(uri string, r interface{}) error {
			return fmt.Errorf(msg)
		}
		cli.callFunc = mockCallFunc
		err := cli.send("uri", VkResponse{})
		test.So(err, test.ShouldNotBeNil)
		test.So(err.Error(), test.ShouldEqual, msg)
	})

	test.Convey("invalid error response", t, func() {
		msg := `error: "invalid character 'i' looking for beginning of object key string" on parsing error_message, response: "{invalid_json"`
		mockCallFunc := func(uri string, r interface{}) error {
			s := r.(*VkResponse)
			s.Error = json.RawMessage([]byte("{invalid_json"))
			return nil
		}
		cli.callFunc = mockCallFunc
		err := cli.send("uri", []OrderResponse{})
		test.So(err.Error(), test.ShouldEqual, msg)
	})

	test.Convey("invalid success response", t, func() {
		msg := `error: "invalid character 'i' looking for beginning of object key string" parse response, response:  "{invalid_json"`
		mockCallFunc := func(uri string, r interface{}) error {
			s := r.(*VkResponse)
			s.Response = json.RawMessage([]byte("{invalid_json"))
			return nil
		}
		cli.callFunc = mockCallFunc
		err := cli.send("uri", []OrderResponse{})
		test.So(err.Error(), test.ShouldEqual, msg)
	})

	test.Convey("success response", t, func() {
		expected := SuccessAuthResponse{
			AccessToken: "123",
			Expire:      0,
		}
		mockCallFunc := func(uri string, r interface{}) error {
			s := r.(*VkResponse)
			s.Response = json.RawMessage([]byte(`{"access_token": "123", "expires_in": 0}`))
			return nil
		}
		cli.callFunc = mockCallFunc
		res := SuccessAuthResponse{}
		err := cli.send("uri", &res)
		test.So(err, test.ShouldBeNil)
		test.So(res, test.ShouldResemble, expected)
	})
}

func TestCallFunction(t *testing.T) {

	test.Convey("call bad endpoint", t, func() {
		msg := `Get hui://http: unsupported protocol scheme "hui"`
		uri := "hui://http"
		err := call(uri, VkResponse{})
		test.So(err, test.ShouldNotBeNil)
		test.So(err.Error(), test.ShouldEqual, msg)
	})
}

type FakeReader struct{}

func (f FakeReader) Read(b []byte) (n int, err error) {
	return 0, fmt.Errorf("error")
}

func TestHandleHTTPError(t *testing.T) {

	test.Convey("bad reader", t, func() {
		r := FakeReader{}
		f := bufio.NewReader(r)
		err := makeHTTPErrorResponse(f)
		test.So(err, test.ShouldNotBeNil)
		test.So(err.Error(), test.ShouldEqual, "error")
	})

	test.Convey("expected error body gotten", t, func() {
		f := bufio.NewReader(os.Stdin)
		err := makeHTTPErrorResponse(f)
		test.So(err, test.ShouldNotBeNil)
		test.So(err.Error(), test.ShouldEqual, "error response, msg: ")
	})
}
