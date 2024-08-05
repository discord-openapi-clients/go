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

// checks if the MessageComponentMentionableSelectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageComponentMentionableSelectResponse{}

// MessageComponentMentionableSelectResponse struct for MessageComponentMentionableSelectResponse
type MessageComponentMentionableSelectResponse struct {
	Type MessageComponentTypes `json:"type"`
	Id int32 `json:"id"`
	CustomId string `json:"custom_id"`
	Placeholder NullableString `json:"placeholder,omitempty"`
	MinValues NullableInt32 `json:"min_values,omitempty"`
	MaxValues NullableInt32 `json:"max_values,omitempty"`
	Disabled NullableBool `json:"disabled,omitempty"`
}

type _MessageComponentMentionableSelectResponse MessageComponentMentionableSelectResponse

// NewMessageComponentMentionableSelectResponse instantiates a new MessageComponentMentionableSelectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageComponentMentionableSelectResponse(type_ MessageComponentTypes, id int32, customId string) *MessageComponentMentionableSelectResponse {
	this := MessageComponentMentionableSelectResponse{}
	this.Type = type_
	this.Id = id
	this.CustomId = customId
	return &this
}

// NewMessageComponentMentionableSelectResponseWithDefaults instantiates a new MessageComponentMentionableSelectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageComponentMentionableSelectResponseWithDefaults() *MessageComponentMentionableSelectResponse {
	this := MessageComponentMentionableSelectResponse{}
	return &this
}

// GetType returns the Type field value
func (o *MessageComponentMentionableSelectResponse) GetType() MessageComponentTypes {
	if o == nil {
		var ret MessageComponentTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MessageComponentMentionableSelectResponse) GetTypeOk() (*MessageComponentTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MessageComponentMentionableSelectResponse) SetType(v MessageComponentTypes) {
	o.Type = v
}

// GetId returns the Id field value
func (o *MessageComponentMentionableSelectResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MessageComponentMentionableSelectResponse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MessageComponentMentionableSelectResponse) SetId(v int32) {
	o.Id = v
}

// GetCustomId returns the CustomId field value
func (o *MessageComponentMentionableSelectResponse) GetCustomId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomId
}

// GetCustomIdOk returns a tuple with the CustomId field value
// and a boolean to check if the value has been set.
func (o *MessageComponentMentionableSelectResponse) GetCustomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomId, true
}

// SetCustomId sets field value
func (o *MessageComponentMentionableSelectResponse) SetCustomId(v string) {
	o.CustomId = v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageComponentMentionableSelectResponse) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder.Get()) {
		var ret string
		return ret
	}
	return *o.Placeholder.Get()
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageComponentMentionableSelectResponse) GetPlaceholderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Placeholder.Get(), o.Placeholder.IsSet()
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *MessageComponentMentionableSelectResponse) HasPlaceholder() bool {
	if o != nil && o.Placeholder.IsSet() {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given NullableString and assigns it to the Placeholder field.
func (o *MessageComponentMentionableSelectResponse) SetPlaceholder(v string) {
	o.Placeholder.Set(&v)
}
// SetPlaceholderNil sets the value for Placeholder to be an explicit nil
func (o *MessageComponentMentionableSelectResponse) SetPlaceholderNil() {
	o.Placeholder.Set(nil)
}

// UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
func (o *MessageComponentMentionableSelectResponse) UnsetPlaceholder() {
	o.Placeholder.Unset()
}

// GetMinValues returns the MinValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageComponentMentionableSelectResponse) GetMinValues() int32 {
	if o == nil || IsNil(o.MinValues.Get()) {
		var ret int32
		return ret
	}
	return *o.MinValues.Get()
}

// GetMinValuesOk returns a tuple with the MinValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageComponentMentionableSelectResponse) GetMinValuesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinValues.Get(), o.MinValues.IsSet()
}

// HasMinValues returns a boolean if a field has been set.
func (o *MessageComponentMentionableSelectResponse) HasMinValues() bool {
	if o != nil && o.MinValues.IsSet() {
		return true
	}

	return false
}

// SetMinValues gets a reference to the given NullableInt32 and assigns it to the MinValues field.
func (o *MessageComponentMentionableSelectResponse) SetMinValues(v int32) {
	o.MinValues.Set(&v)
}
// SetMinValuesNil sets the value for MinValues to be an explicit nil
func (o *MessageComponentMentionableSelectResponse) SetMinValuesNil() {
	o.MinValues.Set(nil)
}

// UnsetMinValues ensures that no value is present for MinValues, not even an explicit nil
func (o *MessageComponentMentionableSelectResponse) UnsetMinValues() {
	o.MinValues.Unset()
}

// GetMaxValues returns the MaxValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageComponentMentionableSelectResponse) GetMaxValues() int32 {
	if o == nil || IsNil(o.MaxValues.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxValues.Get()
}

// GetMaxValuesOk returns a tuple with the MaxValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageComponentMentionableSelectResponse) GetMaxValuesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxValues.Get(), o.MaxValues.IsSet()
}

// HasMaxValues returns a boolean if a field has been set.
func (o *MessageComponentMentionableSelectResponse) HasMaxValues() bool {
	if o != nil && o.MaxValues.IsSet() {
		return true
	}

	return false
}

// SetMaxValues gets a reference to the given NullableInt32 and assigns it to the MaxValues field.
func (o *MessageComponentMentionableSelectResponse) SetMaxValues(v int32) {
	o.MaxValues.Set(&v)
}
// SetMaxValuesNil sets the value for MaxValues to be an explicit nil
func (o *MessageComponentMentionableSelectResponse) SetMaxValuesNil() {
	o.MaxValues.Set(nil)
}

// UnsetMaxValues ensures that no value is present for MaxValues, not even an explicit nil
func (o *MessageComponentMentionableSelectResponse) UnsetMaxValues() {
	o.MaxValues.Unset()
}

// GetDisabled returns the Disabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageComponentMentionableSelectResponse) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Disabled.Get()
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageComponentMentionableSelectResponse) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disabled.Get(), o.Disabled.IsSet()
}

// HasDisabled returns a boolean if a field has been set.
func (o *MessageComponentMentionableSelectResponse) HasDisabled() bool {
	if o != nil && o.Disabled.IsSet() {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given NullableBool and assigns it to the Disabled field.
func (o *MessageComponentMentionableSelectResponse) SetDisabled(v bool) {
	o.Disabled.Set(&v)
}
// SetDisabledNil sets the value for Disabled to be an explicit nil
func (o *MessageComponentMentionableSelectResponse) SetDisabledNil() {
	o.Disabled.Set(nil)
}

// UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
func (o *MessageComponentMentionableSelectResponse) UnsetDisabled() {
	o.Disabled.Unset()
}

func (o MessageComponentMentionableSelectResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageComponentMentionableSelectResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["custom_id"] = o.CustomId
	if o.Placeholder.IsSet() {
		toSerialize["placeholder"] = o.Placeholder.Get()
	}
	if o.MinValues.IsSet() {
		toSerialize["min_values"] = o.MinValues.Get()
	}
	if o.MaxValues.IsSet() {
		toSerialize["max_values"] = o.MaxValues.Get()
	}
	if o.Disabled.IsSet() {
		toSerialize["disabled"] = o.Disabled.Get()
	}
	return toSerialize, nil
}

func (o *MessageComponentMentionableSelectResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"custom_id",
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

	varMessageComponentMentionableSelectResponse := _MessageComponentMentionableSelectResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageComponentMentionableSelectResponse)

	if err != nil {
		return err
	}

	*o = MessageComponentMentionableSelectResponse(varMessageComponentMentionableSelectResponse)

	return err
}

type NullableMessageComponentMentionableSelectResponse struct {
	value *MessageComponentMentionableSelectResponse
	isSet bool
}

func (v NullableMessageComponentMentionableSelectResponse) Get() *MessageComponentMentionableSelectResponse {
	return v.value
}

func (v *NullableMessageComponentMentionableSelectResponse) Set(val *MessageComponentMentionableSelectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageComponentMentionableSelectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageComponentMentionableSelectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageComponentMentionableSelectResponse(val *MessageComponentMentionableSelectResponse) *NullableMessageComponentMentionableSelectResponse {
	return &NullableMessageComponentMentionableSelectResponse{value: val, isSet: true}
}

func (v NullableMessageComponentMentionableSelectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageComponentMentionableSelectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


