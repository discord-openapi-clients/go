/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discord

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GuildProductPurchaseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildProductPurchaseResponse{}

// GuildProductPurchaseResponse struct for GuildProductPurchaseResponse
type GuildProductPurchaseResponse struct {
	ListingId string `json:"listing_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	ProductName string `json:"product_name"`
}

type _GuildProductPurchaseResponse GuildProductPurchaseResponse

// NewGuildProductPurchaseResponse instantiates a new GuildProductPurchaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildProductPurchaseResponse(listingId string, productName string) *GuildProductPurchaseResponse {
	this := GuildProductPurchaseResponse{}
	this.ListingId = listingId
	this.ProductName = productName
	return &this
}

// NewGuildProductPurchaseResponseWithDefaults instantiates a new GuildProductPurchaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildProductPurchaseResponseWithDefaults() *GuildProductPurchaseResponse {
	this := GuildProductPurchaseResponse{}
	return &this
}

// GetListingId returns the ListingId field value
func (o *GuildProductPurchaseResponse) GetListingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListingId
}

// GetListingIdOk returns a tuple with the ListingId field value
// and a boolean to check if the value has been set.
func (o *GuildProductPurchaseResponse) GetListingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListingId, true
}

// SetListingId sets field value
func (o *GuildProductPurchaseResponse) SetListingId(v string) {
	o.ListingId = v
}

// GetProductName returns the ProductName field value
func (o *GuildProductPurchaseResponse) GetProductName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value
// and a boolean to check if the value has been set.
func (o *GuildProductPurchaseResponse) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductName, true
}

// SetProductName sets field value
func (o *GuildProductPurchaseResponse) SetProductName(v string) {
	o.ProductName = v
}

func (o GuildProductPurchaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildProductPurchaseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["listing_id"] = o.ListingId
	toSerialize["product_name"] = o.ProductName
	return toSerialize, nil
}

func (o *GuildProductPurchaseResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"listing_id",
		"product_name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGuildProductPurchaseResponse := _GuildProductPurchaseResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildProductPurchaseResponse)

	if err != nil {
		return err
	}

	*o = GuildProductPurchaseResponse(varGuildProductPurchaseResponse)

	return err
}

type NullableGuildProductPurchaseResponse struct {
	value *GuildProductPurchaseResponse
	isSet bool
}

func (v NullableGuildProductPurchaseResponse) Get() *GuildProductPurchaseResponse {
	return v.value
}

func (v *NullableGuildProductPurchaseResponse) Set(val *GuildProductPurchaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildProductPurchaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildProductPurchaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildProductPurchaseResponse(val *GuildProductPurchaseResponse) *NullableGuildProductPurchaseResponse {
	return &NullableGuildProductPurchaseResponse{value: val, isSet: true}
}

func (v NullableGuildProductPurchaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildProductPurchaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


