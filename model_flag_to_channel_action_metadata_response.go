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

// checks if the FlagToChannelActionMetadataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlagToChannelActionMetadataResponse{}

// FlagToChannelActionMetadataResponse struct for FlagToChannelActionMetadataResponse
type FlagToChannelActionMetadataResponse struct {
	ChannelId string `json:"channel_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _FlagToChannelActionMetadataResponse FlagToChannelActionMetadataResponse

// NewFlagToChannelActionMetadataResponse instantiates a new FlagToChannelActionMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlagToChannelActionMetadataResponse(channelId string) *FlagToChannelActionMetadataResponse {
	this := FlagToChannelActionMetadataResponse{}
	this.ChannelId = channelId
	return &this
}

// NewFlagToChannelActionMetadataResponseWithDefaults instantiates a new FlagToChannelActionMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlagToChannelActionMetadataResponseWithDefaults() *FlagToChannelActionMetadataResponse {
	this := FlagToChannelActionMetadataResponse{}
	return &this
}

// GetChannelId returns the ChannelId field value
func (o *FlagToChannelActionMetadataResponse) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *FlagToChannelActionMetadataResponse) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *FlagToChannelActionMetadataResponse) SetChannelId(v string) {
	o.ChannelId = v
}

func (o FlagToChannelActionMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlagToChannelActionMetadataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel_id"] = o.ChannelId
	return toSerialize, nil
}

func (o *FlagToChannelActionMetadataResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"channel_id",
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

	varFlagToChannelActionMetadataResponse := _FlagToChannelActionMetadataResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFlagToChannelActionMetadataResponse)

	if err != nil {
		return err
	}

	*o = FlagToChannelActionMetadataResponse(varFlagToChannelActionMetadataResponse)

	return err
}

type NullableFlagToChannelActionMetadataResponse struct {
	value *FlagToChannelActionMetadataResponse
	isSet bool
}

func (v NullableFlagToChannelActionMetadataResponse) Get() *FlagToChannelActionMetadataResponse {
	return v.value
}

func (v *NullableFlagToChannelActionMetadataResponse) Set(val *FlagToChannelActionMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFlagToChannelActionMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFlagToChannelActionMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlagToChannelActionMetadataResponse(val *FlagToChannelActionMetadataResponse) *NullableFlagToChannelActionMetadataResponse {
	return &NullableFlagToChannelActionMetadataResponse{value: val, isSet: true}
}

func (v NullableFlagToChannelActionMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlagToChannelActionMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

