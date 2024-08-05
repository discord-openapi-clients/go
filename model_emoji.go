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

// checks if the Emoji type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Emoji{}

// Emoji struct for Emoji
type Emoji struct {
	Id *string `json:"id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Animated NullableBool `json:"animated,omitempty"`
}

type _Emoji Emoji

// NewEmoji instantiates a new Emoji object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmoji(name string) *Emoji {
	this := Emoji{}
	this.Name = name
	return &this
}

// NewEmojiWithDefaults instantiates a new Emoji object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmojiWithDefaults() *Emoji {
	this := Emoji{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Emoji) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Emoji) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Emoji) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Emoji) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Emoji) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Emoji) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Emoji) SetName(v string) {
	o.Name = v
}

// GetAnimated returns the Animated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Emoji) GetAnimated() bool {
	if o == nil || IsNil(o.Animated.Get()) {
		var ret bool
		return ret
	}
	return *o.Animated.Get()
}

// GetAnimatedOk returns a tuple with the Animated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Emoji) GetAnimatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Animated.Get(), o.Animated.IsSet()
}

// HasAnimated returns a boolean if a field has been set.
func (o *Emoji) HasAnimated() bool {
	if o != nil && o.Animated.IsSet() {
		return true
	}

	return false
}

// SetAnimated gets a reference to the given NullableBool and assigns it to the Animated field.
func (o *Emoji) SetAnimated(v bool) {
	o.Animated.Set(&v)
}
// SetAnimatedNil sets the value for Animated to be an explicit nil
func (o *Emoji) SetAnimatedNil() {
	o.Animated.Set(nil)
}

// UnsetAnimated ensures that no value is present for Animated, not even an explicit nil
func (o *Emoji) UnsetAnimated() {
	o.Animated.Unset()
}

func (o Emoji) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Emoji) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if o.Animated.IsSet() {
		toSerialize["animated"] = o.Animated.Get()
	}
	return toSerialize, nil
}

func (o *Emoji) UnmarshalJSON(data []byte) (err error) {
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

	varEmoji := _Emoji{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmoji)

	if err != nil {
		return err
	}

	*o = Emoji(varEmoji)

	return err
}

type NullableEmoji struct {
	value *Emoji
	isSet bool
}

func (v NullableEmoji) Get() *Emoji {
	return v.value
}

func (v *NullableEmoji) Set(val *Emoji) {
	v.value = val
	v.isSet = true
}

func (v NullableEmoji) IsSet() bool {
	return v.isSet
}

func (v *NullableEmoji) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmoji(val *Emoji) *NullableEmoji {
	return &NullableEmoji{value: val, isSet: true}
}

func (v NullableEmoji) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmoji) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


