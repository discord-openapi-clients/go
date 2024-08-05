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

// checks if the WebhookSourceGuildResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookSourceGuildResponse{}

// WebhookSourceGuildResponse struct for WebhookSourceGuildResponse
type WebhookSourceGuildResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Icon NullableString `json:"icon,omitempty"`
	Name string `json:"name"`
}

type _WebhookSourceGuildResponse WebhookSourceGuildResponse

// NewWebhookSourceGuildResponse instantiates a new WebhookSourceGuildResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookSourceGuildResponse(id string, name string) *WebhookSourceGuildResponse {
	this := WebhookSourceGuildResponse{}
	this.Id = id
	this.Name = name
	return &this
}

// NewWebhookSourceGuildResponseWithDefaults instantiates a new WebhookSourceGuildResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookSourceGuildResponseWithDefaults() *WebhookSourceGuildResponse {
	this := WebhookSourceGuildResponse{}
	return &this
}

// GetId returns the Id field value
func (o *WebhookSourceGuildResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhookSourceGuildResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhookSourceGuildResponse) SetId(v string) {
	o.Id = v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSourceGuildResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSourceGuildResponse) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *WebhookSourceGuildResponse) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *WebhookSourceGuildResponse) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *WebhookSourceGuildResponse) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *WebhookSourceGuildResponse) UnsetIcon() {
	o.Icon.Unset()
}

// GetName returns the Name field value
func (o *WebhookSourceGuildResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebhookSourceGuildResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WebhookSourceGuildResponse) SetName(v string) {
	o.Name = v
}

func (o WebhookSourceGuildResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookSourceGuildResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *WebhookSourceGuildResponse) UnmarshalJSON(data []byte) (err error) {
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

	varWebhookSourceGuildResponse := _WebhookSourceGuildResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookSourceGuildResponse)

	if err != nil {
		return err
	}

	*o = WebhookSourceGuildResponse(varWebhookSourceGuildResponse)

	return err
}

type NullableWebhookSourceGuildResponse struct {
	value *WebhookSourceGuildResponse
	isSet bool
}

func (v NullableWebhookSourceGuildResponse) Get() *WebhookSourceGuildResponse {
	return v.value
}

func (v *NullableWebhookSourceGuildResponse) Set(val *WebhookSourceGuildResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookSourceGuildResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookSourceGuildResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookSourceGuildResponse(val *WebhookSourceGuildResponse) *NullableWebhookSourceGuildResponse {
	return &NullableWebhookSourceGuildResponse{value: val, isSet: true}
}

func (v NullableWebhookSourceGuildResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookSourceGuildResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

