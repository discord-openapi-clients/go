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

// checks if the MessageComponentStringSelectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageComponentStringSelectResponse{}

// MessageComponentStringSelectResponse struct for MessageComponentStringSelectResponse
type MessageComponentStringSelectResponse struct {
	Type MessageComponentTypes `json:"type"`
	Id int32 `json:"id"`
	CustomId string `json:"custom_id"`
	Placeholder NullableString `json:"placeholder,omitempty"`
	MinValues NullableInt32 `json:"min_values,omitempty"`
	MaxValues NullableInt32 `json:"max_values,omitempty"`
	Disabled NullableBool `json:"disabled,omitempty"`
	Options []MessageComponentStringSelectResponseOptionsInner `json:"options,omitempty"`
}

type _MessageComponentStringSelectResponse MessageComponentStringSelectResponse

// NewMessageComponentStringSelectResponse instantiates a new MessageComponentStringSelectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageComponentStringSelectResponse(type_ MessageComponentTypes, id int32, customId string) *MessageComponentStringSelectResponse {
	this := MessageComponentStringSelectResponse{}
	this.Type = type_
	this.Id = id
	this.CustomId = customId
	return &this
}

// NewMessageComponentStringSelectResponseWithDefaults instantiates a new MessageComponentStringSelectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageComponentStringSelectResponseWithDefaults() *MessageComponentStringSelectResponse {
	this := MessageComponentStringSelectResponse{}
	return &this
}

// GetType returns the Type field value
func (o *MessageComponentStringSelectResponse) GetType() MessageComponentTypes {
	if o == nil {
		var ret MessageComponentTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MessageComponentStringSelectResponse) GetTypeOk() (*MessageComponentTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MessageComponentStringSelectResponse) SetType(v MessageComponentTypes) {
	o.Type = v
}

// GetId returns the Id field value
func (o *MessageComponentStringSelectResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MessageComponentStringSelectResponse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MessageComponentStringSelectResponse) SetId(v int32) {
	o.Id = v
}

// GetCustomId returns the CustomId field value
func (o *MessageComponentStringSelectResponse) GetCustomId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomId
}

// GetCustomIdOk returns a tuple with the CustomId field value
// and a boolean to check if the value has been set.
func (o *MessageComponentStringSelectResponse) GetCustomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomId, true
}

// SetCustomId sets field value
func (o *MessageComponentStringSelectResponse) SetCustomId(v string) {
	o.CustomId = v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageComponentStringSelectResponse) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder.Get()) {
		var ret string
		return ret
	}
	return *o.Placeholder.Get()
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageComponentStringSelectResponse) GetPlaceholderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Placeholder.Get(), o.Placeholder.IsSet()
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *MessageComponentStringSelectResponse) HasPlaceholder() bool {
	if o != nil && o.Placeholder.IsSet() {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given NullableString and assigns it to the Placeholder field.
func (o *MessageComponentStringSelectResponse) SetPlaceholder(v string) {
	o.Placeholder.Set(&v)
}
// SetPlaceholderNil sets the value for Placeholder to be an explicit nil
func (o *MessageComponentStringSelectResponse) SetPlaceholderNil() {
	o.Placeholder.Set(nil)
}

// UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
func (o *MessageComponentStringSelectResponse) UnsetPlaceholder() {
	o.Placeholder.Unset()
}

// GetMinValues returns the MinValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageComponentStringSelectResponse) GetMinValues() int32 {
	if o == nil || IsNil(o.MinValues.Get()) {
		var ret int32
		return ret
	}
	return *o.MinValues.Get()
}

// GetMinValuesOk returns a tuple with the MinValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageComponentStringSelectResponse) GetMinValuesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinValues.Get(), o.MinValues.IsSet()
}

// HasMinValues returns a boolean if a field has been set.
func (o *MessageComponentStringSelectResponse) HasMinValues() bool {
	if o != nil && o.MinValues.IsSet() {
		return true
	}

	return false
}

// SetMinValues gets a reference to the given NullableInt32 and assigns it to the MinValues field.
func (o *MessageComponentStringSelectResponse) SetMinValues(v int32) {
	o.MinValues.Set(&v)
}
// SetMinValuesNil sets the value for MinValues to be an explicit nil
func (o *MessageComponentStringSelectResponse) SetMinValuesNil() {
	o.MinValues.Set(nil)
}

// UnsetMinValues ensures that no value is present for MinValues, not even an explicit nil
func (o *MessageComponentStringSelectResponse) UnsetMinValues() {
	o.MinValues.Unset()
}

// GetMaxValues returns the MaxValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageComponentStringSelectResponse) GetMaxValues() int32 {
	if o == nil || IsNil(o.MaxValues.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxValues.Get()
}

// GetMaxValuesOk returns a tuple with the MaxValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageComponentStringSelectResponse) GetMaxValuesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxValues.Get(), o.MaxValues.IsSet()
}

// HasMaxValues returns a boolean if a field has been set.
func (o *MessageComponentStringSelectResponse) HasMaxValues() bool {
	if o != nil && o.MaxValues.IsSet() {
		return true
	}

	return false
}

// SetMaxValues gets a reference to the given NullableInt32 and assigns it to the MaxValues field.
func (o *MessageComponentStringSelectResponse) SetMaxValues(v int32) {
	o.MaxValues.Set(&v)
}
// SetMaxValuesNil sets the value for MaxValues to be an explicit nil
func (o *MessageComponentStringSelectResponse) SetMaxValuesNil() {
	o.MaxValues.Set(nil)
}

// UnsetMaxValues ensures that no value is present for MaxValues, not even an explicit nil
func (o *MessageComponentStringSelectResponse) UnsetMaxValues() {
	o.MaxValues.Unset()
}

// GetDisabled returns the Disabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageComponentStringSelectResponse) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Disabled.Get()
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageComponentStringSelectResponse) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disabled.Get(), o.Disabled.IsSet()
}

// HasDisabled returns a boolean if a field has been set.
func (o *MessageComponentStringSelectResponse) HasDisabled() bool {
	if o != nil && o.Disabled.IsSet() {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given NullableBool and assigns it to the Disabled field.
func (o *MessageComponentStringSelectResponse) SetDisabled(v bool) {
	o.Disabled.Set(&v)
}
// SetDisabledNil sets the value for Disabled to be an explicit nil
func (o *MessageComponentStringSelectResponse) SetDisabledNil() {
	o.Disabled.Set(nil)
}

// UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
func (o *MessageComponentStringSelectResponse) UnsetDisabled() {
	o.Disabled.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageComponentStringSelectResponse) GetOptions() []MessageComponentStringSelectResponseOptionsInner {
	if o == nil {
		var ret []MessageComponentStringSelectResponseOptionsInner
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageComponentStringSelectResponse) GetOptionsOk() ([]MessageComponentStringSelectResponseOptionsInner, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *MessageComponentStringSelectResponse) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []MessageComponentStringSelectResponseOptionsInner and assigns it to the Options field.
func (o *MessageComponentStringSelectResponse) SetOptions(v []MessageComponentStringSelectResponseOptionsInner) {
	o.Options = v
}

func (o MessageComponentStringSelectResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageComponentStringSelectResponse) ToMap() (map[string]interface{}, error) {
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
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

func (o *MessageComponentStringSelectResponse) UnmarshalJSON(data []byte) (err error) {
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

	varMessageComponentStringSelectResponse := _MessageComponentStringSelectResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageComponentStringSelectResponse)

	if err != nil {
		return err
	}

	*o = MessageComponentStringSelectResponse(varMessageComponentStringSelectResponse)

	return err
}

type NullableMessageComponentStringSelectResponse struct {
	value *MessageComponentStringSelectResponse
	isSet bool
}

func (v NullableMessageComponentStringSelectResponse) Get() *MessageComponentStringSelectResponse {
	return v.value
}

func (v *NullableMessageComponentStringSelectResponse) Set(val *MessageComponentStringSelectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageComponentStringSelectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageComponentStringSelectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageComponentStringSelectResponse(val *MessageComponentStringSelectResponse) *NullableMessageComponentStringSelectResponse {
	return &NullableMessageComponentStringSelectResponse{value: val, isSet: true}
}

func (v NullableMessageComponentStringSelectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageComponentStringSelectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


