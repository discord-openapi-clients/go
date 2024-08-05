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

// checks if the CreateGuildFromTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGuildFromTemplateRequest{}

// CreateGuildFromTemplateRequest struct for CreateGuildFromTemplateRequest
type CreateGuildFromTemplateRequest struct {
	Name string `json:"name"`
	Icon NullableString `json:"icon,omitempty"`
}

type _CreateGuildFromTemplateRequest CreateGuildFromTemplateRequest

// NewCreateGuildFromTemplateRequest instantiates a new CreateGuildFromTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGuildFromTemplateRequest(name string) *CreateGuildFromTemplateRequest {
	this := CreateGuildFromTemplateRequest{}
	this.Name = name
	return &this
}

// NewCreateGuildFromTemplateRequestWithDefaults instantiates a new CreateGuildFromTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGuildFromTemplateRequestWithDefaults() *CreateGuildFromTemplateRequest {
	this := CreateGuildFromTemplateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateGuildFromTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateGuildFromTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateGuildFromTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateGuildFromTemplateRequest) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateGuildFromTemplateRequest) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *CreateGuildFromTemplateRequest) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *CreateGuildFromTemplateRequest) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *CreateGuildFromTemplateRequest) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *CreateGuildFromTemplateRequest) UnsetIcon() {
	o.Icon.Unset()
}

func (o CreateGuildFromTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGuildFromTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	return toSerialize, nil
}

func (o *CreateGuildFromTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varCreateGuildFromTemplateRequest := _CreateGuildFromTemplateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateGuildFromTemplateRequest)

	if err != nil {
		return err
	}

	*o = CreateGuildFromTemplateRequest(varCreateGuildFromTemplateRequest)

	return err
}

type NullableCreateGuildFromTemplateRequest struct {
	value *CreateGuildFromTemplateRequest
	isSet bool
}

func (v NullableCreateGuildFromTemplateRequest) Get() *CreateGuildFromTemplateRequest {
	return v.value
}

func (v *NullableCreateGuildFromTemplateRequest) Set(val *CreateGuildFromTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGuildFromTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGuildFromTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGuildFromTemplateRequest(val *CreateGuildFromTemplateRequest) *NullableCreateGuildFromTemplateRequest {
	return &NullableCreateGuildFromTemplateRequest{value: val, isSet: true}
}

func (v NullableCreateGuildFromTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGuildFromTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

