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

// checks if the ConnectedAccountGuildResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectedAccountGuildResponse{}

// ConnectedAccountGuildResponse struct for ConnectedAccountGuildResponse
type ConnectedAccountGuildResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Icon NullableString `json:"icon,omitempty"`
}

type _ConnectedAccountGuildResponse ConnectedAccountGuildResponse

// NewConnectedAccountGuildResponse instantiates a new ConnectedAccountGuildResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectedAccountGuildResponse(id string, name string) *ConnectedAccountGuildResponse {
	this := ConnectedAccountGuildResponse{}
	this.Id = id
	this.Name = name
	return &this
}

// NewConnectedAccountGuildResponseWithDefaults instantiates a new ConnectedAccountGuildResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectedAccountGuildResponseWithDefaults() *ConnectedAccountGuildResponse {
	this := ConnectedAccountGuildResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ConnectedAccountGuildResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConnectedAccountGuildResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConnectedAccountGuildResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ConnectedAccountGuildResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConnectedAccountGuildResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConnectedAccountGuildResponse) SetName(v string) {
	o.Name = v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectedAccountGuildResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectedAccountGuildResponse) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *ConnectedAccountGuildResponse) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *ConnectedAccountGuildResponse) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *ConnectedAccountGuildResponse) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *ConnectedAccountGuildResponse) UnsetIcon() {
	o.Icon.Unset()
}

func (o ConnectedAccountGuildResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectedAccountGuildResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	return toSerialize, nil
}

func (o *ConnectedAccountGuildResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varConnectedAccountGuildResponse := _ConnectedAccountGuildResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConnectedAccountGuildResponse)

	if err != nil {
		return err
	}

	*o = ConnectedAccountGuildResponse(varConnectedAccountGuildResponse)

	return err
}

type NullableConnectedAccountGuildResponse struct {
	value *ConnectedAccountGuildResponse
	isSet bool
}

func (v NullableConnectedAccountGuildResponse) Get() *ConnectedAccountGuildResponse {
	return v.value
}

func (v *NullableConnectedAccountGuildResponse) Set(val *ConnectedAccountGuildResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectedAccountGuildResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectedAccountGuildResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectedAccountGuildResponse(val *ConnectedAccountGuildResponse) *NullableConnectedAccountGuildResponse {
	return &NullableConnectedAccountGuildResponse{value: val, isSet: true}
}

func (v NullableConnectedAccountGuildResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectedAccountGuildResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

