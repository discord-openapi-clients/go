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

// checks if the SetGuildMfaLevelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetGuildMfaLevelRequest{}

// SetGuildMfaLevelRequest struct for SetGuildMfaLevelRequest
type SetGuildMfaLevelRequest struct {
	Level GuildMFALevel `json:"level"`
}

type _SetGuildMfaLevelRequest SetGuildMfaLevelRequest

// NewSetGuildMfaLevelRequest instantiates a new SetGuildMfaLevelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetGuildMfaLevelRequest(level GuildMFALevel) *SetGuildMfaLevelRequest {
	this := SetGuildMfaLevelRequest{}
	this.Level = level
	return &this
}

// NewSetGuildMfaLevelRequestWithDefaults instantiates a new SetGuildMfaLevelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetGuildMfaLevelRequestWithDefaults() *SetGuildMfaLevelRequest {
	this := SetGuildMfaLevelRequest{}
	return &this
}

// GetLevel returns the Level field value
func (o *SetGuildMfaLevelRequest) GetLevel() GuildMFALevel {
	if o == nil {
		var ret GuildMFALevel
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *SetGuildMfaLevelRequest) GetLevelOk() (*GuildMFALevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *SetGuildMfaLevelRequest) SetLevel(v GuildMFALevel) {
	o.Level = v
}

func (o SetGuildMfaLevelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetGuildMfaLevelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["level"] = o.Level
	return toSerialize, nil
}

func (o *SetGuildMfaLevelRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"level",
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

	varSetGuildMfaLevelRequest := _SetGuildMfaLevelRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetGuildMfaLevelRequest)

	if err != nil {
		return err
	}

	*o = SetGuildMfaLevelRequest(varSetGuildMfaLevelRequest)

	return err
}

type NullableSetGuildMfaLevelRequest struct {
	value *SetGuildMfaLevelRequest
	isSet bool
}

func (v NullableSetGuildMfaLevelRequest) Get() *SetGuildMfaLevelRequest {
	return v.value
}

func (v *NullableSetGuildMfaLevelRequest) Set(val *SetGuildMfaLevelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetGuildMfaLevelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetGuildMfaLevelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetGuildMfaLevelRequest(val *SetGuildMfaLevelRequest) *NullableSetGuildMfaLevelRequest {
	return &NullableSetGuildMfaLevelRequest{value: val, isSet: true}
}

func (v NullableSetGuildMfaLevelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetGuildMfaLevelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

