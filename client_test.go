package govk_test

import (
	"encoding/json"
	"testing"

	"github.com/chapsuk/govk"
	test "github.com/smartystreets/goconvey/convey"
)

func TestOrdersGet(t *testing.T) {
	cli := govk.NewClient("123123", "fakeSecret", "5.53", 0)

	test.Convey("orders.get", t, func() {

		test.Convey("missing access_token", func() {
			msg := `error response: User authorization failed: no access_token passed.`
			_, err := cli.OrdersGet(10, 0, 0)
			test.So(err, test.ShouldNotBeNil)
			test.So(err.Error(), test.ShouldEqual, msg)
		})

		test.Convey("success result", func() {
			expected := []govk.OrderResponse{}
			expected = append(expected, govk.OrderResponse{
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
				a := r.(*govk.VkResponse)
				a.Response = json.RawMessage([]byte(msg))
				return nil
			}
			cli.CallFunc = mockCallFunc
			res, err := cli.OrdersGet(10, 0, 0)
			test.So(err, test.ShouldBeNil)
			test.So(res, test.ShouldResemble, expected)
		})
	})

}

func TestAuth(t *testing.T) {
	cli := govk.NewClient("123123", "fakeSecret", "5.53", 0)

	test.Convey("check auth errors", t, func() {

		test.Convey("blocked client id", func() {
			msg := `error response, msg: {"error":"invalid_client","error_description":"client_id is blocked"}`
			err := cli.Auth()
			test.So(err, test.ShouldNotBeNil)
			test.So(err.Error(), test.ShouldEqual, msg)
		})

		test.Convey("invalid client id", func() {
			msg := `error response, msg: {"error":"invalid_client","error_description":"client_id is invalid"}`
			cli.ClientID = "invalidID"
			err := cli.Auth()
			test.So(err, test.ShouldNotBeNil)
			test.So(err.Error(), test.ShouldEqual, msg)
		})
	})

	test.Convey("check auth success", t, func() {
		mockCallFunc := func(uri string, r interface{}) error {
			a := r.(*govk.SuccessAuthResponse)
			a.AccessToken = "123"
			return nil
		}
		cli.CallFunc = mockCallFunc
		err := cli.Auth()
		test.So(err, test.ShouldBeNil)
		test.So(cli.AccessToken, test.ShouldEqual, "123")
	})
}

func TestClient(t *testing.T) {

	test.Convey("create new Client instance", t, func() {
		id := "123"
		secret := "secret"
		version := "5.53"
		cli := govk.NewClient(id, secret, version, 0)

		test.So(id, test.ShouldEqual, cli.ClientID)
		test.So(secret, test.ShouldEqual, cli.ClientSecret)
		test.So(version, test.ShouldEqual, cli.APIVersion)
		test.So("", test.ShouldEqual, cli.AccessToken)
	})
}
