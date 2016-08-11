package govk

import (
	"encoding/json"
)

// SuccessAuthResponse is success vk api response structure
type SuccessAuthResponse struct {
	AccessToken string `json:"access_token"`
	Expire      int    `json:"expires_in"`
}

// ErrorAuthResponse structure of vk api failed authorization response
type ErrorAuthResponse struct {
	Message     string `json:"error"`
	Description string `json:"error_description"`
}

// VkResponse root response structure
type VkResponse struct {
	Response json.RawMessage `json:"response"`
	Error    json.RawMessage `json:"error"`
}

// VkErrorResponse error response structure
type VkErrorResponse struct {
	Code    int     `json:"error_code"`
	Message string  `json:"error_msg"`
	Params  []Param `json:"request_params"`
}

// Param of query
type Param struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// OrderResponse item of order.get response result list
type OrderResponse struct {
	ID         string `json:"id" db:"order_id"`
	AppOrderID string `json:"app_order_id" db:"app_order_id"`
	Status     string `json:"status" db:"status"`
	UserID     string `json:"user_id" db:"user_id"`
	ReceiverID string `json:"receiver_id" db:"receiver_id"`
	Item       string `json:"item" db:"item"`
	Amount     string `json:"amount" db:"amount"`
	Date       string `json:"date" db:"date"`
}

// GetContriesResponse database.getCountries response
type GetContriesResponse struct {
	Count int               `json:"count"`
	Items []CountryResponse `json:"items"`
}

// CountryResponse is country item
type CountryResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// GetRegionsResponse database.getRegions response
type GetRegionsResponse struct {
	Count int               `json:"count"`
	Items []CountryResponse `json:"items"`
}

// RegionResponse region item
type RegionResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// GetCitiesResponse database.getCities response
type GetCitiesResponse struct {
	Count int            `json:"count"`
	Items []CityResponse `json:"items"`
}

// CityResponse is city item
type CityResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Area      string `json:"area"`
	Region    string `json:"region"`
	Important int    `json:"important"`
}

// CityByIDResponse city item
type CityByIDResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
