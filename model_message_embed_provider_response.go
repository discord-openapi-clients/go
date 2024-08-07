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

// checks if the MessageEmbedProviderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageEmbedProviderResponse{}

// MessageEmbedProviderResponse struct for MessageEmbedProviderResponse
type MessageEmbedProviderResponse struct {
	Name string `json:"name"`
	Url NullableString `json:"url,omitempty"`
}

type _MessageEmbedProviderResponse MessageEmbedProviderResponse

// NewMessageEmbedProviderResponse instantiates a new MessageEmbedProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageEmbedProviderResponse(name string) *MessageEmbedProviderResponse {
	this := MessageEmbedProviderResponse{}
	this.Name = name
	return &this
}

// NewMessageEmbedProviderResponseWithDefaults instantiates a new MessageEmbedProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageEmbedProviderResponseWithDefaults() *MessageEmbedProviderResponse {
	this := MessageEmbedProviderResponse{}
	return &this
}

// GetName returns the Name field value
func (o *MessageEmbedProviderResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MessageEmbedProviderResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MessageEmbedProviderResponse) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEmbedProviderResponse) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEmbedProviderResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *MessageEmbedProviderResponse) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *MessageEmbedProviderResponse) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *MessageEmbedProviderResponse) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *MessageEmbedProviderResponse) UnsetUrl() {
	o.Url.Unset()
}

func (o MessageEmbedProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageEmbedProviderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return toSerialize, nil
}

func (o *MessageEmbedProviderResponse) UnmarshalJSON(data []byte) (err error) {
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

	varMessageEmbedProviderResponse := _MessageEmbedProviderResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageEmbedProviderResponse)

	if err != nil {
		return err
	}

	*o = MessageEmbedProviderResponse(varMessageEmbedProviderResponse)

	return err
}

type NullableMessageEmbedProviderResponse struct {
	value *MessageEmbedProviderResponse
	isSet bool
}

func (v NullableMessageEmbedProviderResponse) Get() *MessageEmbedProviderResponse {
	return v.value
}

func (v *NullableMessageEmbedProviderResponse) Set(val *MessageEmbedProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageEmbedProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageEmbedProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageEmbedProviderResponse(val *MessageEmbedProviderResponse) *NullableMessageEmbedProviderResponse {
	return &NullableMessageEmbedProviderResponse{value: val, isSet: true}
}

func (v NullableMessageEmbedProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageEmbedProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


