/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discord

import (
	"encoding/json"
)

// checks if the MessageReactionEmojiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageReactionEmojiResponse{}

// MessageReactionEmojiResponse struct for MessageReactionEmojiResponse
type MessageReactionEmojiResponse struct {
	Id *string `json:"id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name NullableString `json:"name,omitempty"`
	Animated NullableBool `json:"animated,omitempty"`
}

// NewMessageReactionEmojiResponse instantiates a new MessageReactionEmojiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageReactionEmojiResponse() *MessageReactionEmojiResponse {
	this := MessageReactionEmojiResponse{}
	return &this
}

// NewMessageReactionEmojiResponseWithDefaults instantiates a new MessageReactionEmojiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageReactionEmojiResponseWithDefaults() *MessageReactionEmojiResponse {
	this := MessageReactionEmojiResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MessageReactionEmojiResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageReactionEmojiResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MessageReactionEmojiResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MessageReactionEmojiResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageReactionEmojiResponse) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageReactionEmojiResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MessageReactionEmojiResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MessageReactionEmojiResponse) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MessageReactionEmojiResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MessageReactionEmojiResponse) UnsetName() {
	o.Name.Unset()
}

// GetAnimated returns the Animated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageReactionEmojiResponse) GetAnimated() bool {
	if o == nil || IsNil(o.Animated.Get()) {
		var ret bool
		return ret
	}
	return *o.Animated.Get()
}

// GetAnimatedOk returns a tuple with the Animated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageReactionEmojiResponse) GetAnimatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Animated.Get(), o.Animated.IsSet()
}

// HasAnimated returns a boolean if a field has been set.
func (o *MessageReactionEmojiResponse) HasAnimated() bool {
	if o != nil && o.Animated.IsSet() {
		return true
	}

	return false
}

// SetAnimated gets a reference to the given NullableBool and assigns it to the Animated field.
func (o *MessageReactionEmojiResponse) SetAnimated(v bool) {
	o.Animated.Set(&v)
}
// SetAnimatedNil sets the value for Animated to be an explicit nil
func (o *MessageReactionEmojiResponse) SetAnimatedNil() {
	o.Animated.Set(nil)
}

// UnsetAnimated ensures that no value is present for Animated, not even an explicit nil
func (o *MessageReactionEmojiResponse) UnsetAnimated() {
	o.Animated.Unset()
}

func (o MessageReactionEmojiResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageReactionEmojiResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Animated.IsSet() {
		toSerialize["animated"] = o.Animated.Get()
	}
	return toSerialize, nil
}

type NullableMessageReactionEmojiResponse struct {
	value *MessageReactionEmojiResponse
	isSet bool
}

func (v NullableMessageReactionEmojiResponse) Get() *MessageReactionEmojiResponse {
	return v.value
}

func (v *NullableMessageReactionEmojiResponse) Set(val *MessageReactionEmojiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageReactionEmojiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageReactionEmojiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageReactionEmojiResponse(val *MessageReactionEmojiResponse) *NullableMessageReactionEmojiResponse {
	return &NullableMessageReactionEmojiResponse{value: val, isSet: true}
}

func (v NullableMessageReactionEmojiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageReactionEmojiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


