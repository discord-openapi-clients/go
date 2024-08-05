/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GatewayResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayResponse{}

// GatewayResponse struct for GatewayResponse
type GatewayResponse struct {
	Url string `json:"url"`
}

type _GatewayResponse GatewayResponse

// NewGatewayResponse instantiates a new GatewayResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayResponse(url string) *GatewayResponse {
	this := GatewayResponse{}
	this.Url = url
	return &this
}

// NewGatewayResponseWithDefaults instantiates a new GatewayResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayResponseWithDefaults() *GatewayResponse {
	this := GatewayResponse{}
	return &this
}

// GetUrl returns the Url field value
func (o *GatewayResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *GatewayResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *GatewayResponse) SetUrl(v string) {
	o.Url = v
}

func (o GatewayResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *GatewayResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varGatewayResponse := _GatewayResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGatewayResponse)

	if err != nil {
		return err
	}

	*o = GatewayResponse(varGatewayResponse)

	return err
}

type NullableGatewayResponse struct {
	value *GatewayResponse
	isSet bool
}

func (v NullableGatewayResponse) Get() *GatewayResponse {
	return v.value
}

func (v *NullableGatewayResponse) Set(val *GatewayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayResponse(val *GatewayResponse) *NullableGatewayResponse {
	return &NullableGatewayResponse{value: val, isSet: true}
}

func (v NullableGatewayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

