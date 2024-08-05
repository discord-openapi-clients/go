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

// checks if the ChannelPermissionOverwriteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelPermissionOverwriteResponse{}

// ChannelPermissionOverwriteResponse struct for ChannelPermissionOverwriteResponse
type ChannelPermissionOverwriteResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Type ChannelPermissionOverwrites `json:"type"`
	Allow string `json:"allow"`
	Deny string `json:"deny"`
}

type _ChannelPermissionOverwriteResponse ChannelPermissionOverwriteResponse

// NewChannelPermissionOverwriteResponse instantiates a new ChannelPermissionOverwriteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPermissionOverwriteResponse(id string, type_ ChannelPermissionOverwrites, allow string, deny string) *ChannelPermissionOverwriteResponse {
	this := ChannelPermissionOverwriteResponse{}
	this.Id = id
	this.Type = type_
	this.Allow = allow
	this.Deny = deny
	return &this
}

// NewChannelPermissionOverwriteResponseWithDefaults instantiates a new ChannelPermissionOverwriteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPermissionOverwriteResponseWithDefaults() *ChannelPermissionOverwriteResponse {
	this := ChannelPermissionOverwriteResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ChannelPermissionOverwriteResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChannelPermissionOverwriteResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChannelPermissionOverwriteResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ChannelPermissionOverwriteResponse) GetType() ChannelPermissionOverwrites {
	if o == nil {
		var ret ChannelPermissionOverwrites
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChannelPermissionOverwriteResponse) GetTypeOk() (*ChannelPermissionOverwrites, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ChannelPermissionOverwriteResponse) SetType(v ChannelPermissionOverwrites) {
	o.Type = v
}

// GetAllow returns the Allow field value
func (o *ChannelPermissionOverwriteResponse) GetAllow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value
// and a boolean to check if the value has been set.
func (o *ChannelPermissionOverwriteResponse) GetAllowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allow, true
}

// SetAllow sets field value
func (o *ChannelPermissionOverwriteResponse) SetAllow(v string) {
	o.Allow = v
}

// GetDeny returns the Deny field value
func (o *ChannelPermissionOverwriteResponse) GetDeny() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Deny
}

// GetDenyOk returns a tuple with the Deny field value
// and a boolean to check if the value has been set.
func (o *ChannelPermissionOverwriteResponse) GetDenyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deny, true
}

// SetDeny sets field value
func (o *ChannelPermissionOverwriteResponse) SetDeny(v string) {
	o.Deny = v
}

func (o ChannelPermissionOverwriteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelPermissionOverwriteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["allow"] = o.Allow
	toSerialize["deny"] = o.Deny
	return toSerialize, nil
}

func (o *ChannelPermissionOverwriteResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"allow",
		"deny",
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

	varChannelPermissionOverwriteResponse := _ChannelPermissionOverwriteResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChannelPermissionOverwriteResponse)

	if err != nil {
		return err
	}

	*o = ChannelPermissionOverwriteResponse(varChannelPermissionOverwriteResponse)

	return err
}

type NullableChannelPermissionOverwriteResponse struct {
	value *ChannelPermissionOverwriteResponse
	isSet bool
}

func (v NullableChannelPermissionOverwriteResponse) Get() *ChannelPermissionOverwriteResponse {
	return v.value
}

func (v *NullableChannelPermissionOverwriteResponse) Set(val *ChannelPermissionOverwriteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPermissionOverwriteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPermissionOverwriteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPermissionOverwriteResponse(val *ChannelPermissionOverwriteResponse) *NullableChannelPermissionOverwriteResponse {
	return &NullableChannelPermissionOverwriteResponse{value: val, isSet: true}
}

func (v NullableChannelPermissionOverwriteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPermissionOverwriteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

