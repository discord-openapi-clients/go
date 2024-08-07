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

// checks if the UserSelectDefaultValueResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSelectDefaultValueResponse{}

// UserSelectDefaultValueResponse struct for UserSelectDefaultValueResponse
type UserSelectDefaultValueResponse struct {
	Type SnowflakeSelectDefaultValueTypes `json:"type"`
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _UserSelectDefaultValueResponse UserSelectDefaultValueResponse

// NewUserSelectDefaultValueResponse instantiates a new UserSelectDefaultValueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSelectDefaultValueResponse(type_ SnowflakeSelectDefaultValueTypes, id string) *UserSelectDefaultValueResponse {
	this := UserSelectDefaultValueResponse{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewUserSelectDefaultValueResponseWithDefaults instantiates a new UserSelectDefaultValueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSelectDefaultValueResponseWithDefaults() *UserSelectDefaultValueResponse {
	this := UserSelectDefaultValueResponse{}
	return &this
}

// GetType returns the Type field value
func (o *UserSelectDefaultValueResponse) GetType() SnowflakeSelectDefaultValueTypes {
	if o == nil {
		var ret SnowflakeSelectDefaultValueTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserSelectDefaultValueResponse) GetTypeOk() (*SnowflakeSelectDefaultValueTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserSelectDefaultValueResponse) SetType(v SnowflakeSelectDefaultValueTypes) {
	o.Type = v
}

// GetId returns the Id field value
func (o *UserSelectDefaultValueResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserSelectDefaultValueResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserSelectDefaultValueResponse) SetId(v string) {
	o.Id = v
}

func (o UserSelectDefaultValueResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSelectDefaultValueResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *UserSelectDefaultValueResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varUserSelectDefaultValueResponse := _UserSelectDefaultValueResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSelectDefaultValueResponse)

	if err != nil {
		return err
	}

	*o = UserSelectDefaultValueResponse(varUserSelectDefaultValueResponse)

	return err
}

type NullableUserSelectDefaultValueResponse struct {
	value *UserSelectDefaultValueResponse
	isSet bool
}

func (v NullableUserSelectDefaultValueResponse) Get() *UserSelectDefaultValueResponse {
	return v.value
}

func (v *NullableUserSelectDefaultValueResponse) Set(val *UserSelectDefaultValueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSelectDefaultValueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSelectDefaultValueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSelectDefaultValueResponse(val *UserSelectDefaultValueResponse) *NullableUserSelectDefaultValueResponse {
	return &NullableUserSelectDefaultValueResponse{value: val, isSet: true}
}

func (v NullableUserSelectDefaultValueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSelectDefaultValueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


