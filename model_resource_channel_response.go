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

// checks if the ResourceChannelResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceChannelResponse{}

// ResourceChannelResponse struct for ResourceChannelResponse
type ResourceChannelResponse struct {
	ChannelId string `json:"channel_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Title string `json:"title"`
	Emoji NullableSettingsEmojiResponse `json:"emoji,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
	Description string `json:"description"`
}

type _ResourceChannelResponse ResourceChannelResponse

// NewResourceChannelResponse instantiates a new ResourceChannelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceChannelResponse(channelId string, title string, description string) *ResourceChannelResponse {
	this := ResourceChannelResponse{}
	this.ChannelId = channelId
	this.Title = title
	this.Description = description
	return &this
}

// NewResourceChannelResponseWithDefaults instantiates a new ResourceChannelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceChannelResponseWithDefaults() *ResourceChannelResponse {
	this := ResourceChannelResponse{}
	return &this
}

// GetChannelId returns the ChannelId field value
func (o *ResourceChannelResponse) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *ResourceChannelResponse) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *ResourceChannelResponse) SetChannelId(v string) {
	o.ChannelId = v
}

// GetTitle returns the Title field value
func (o *ResourceChannelResponse) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ResourceChannelResponse) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ResourceChannelResponse) SetTitle(v string) {
	o.Title = v
}

// GetEmoji returns the Emoji field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceChannelResponse) GetEmoji() SettingsEmojiResponse {
	if o == nil || IsNil(o.Emoji.Get()) {
		var ret SettingsEmojiResponse
		return ret
	}
	return *o.Emoji.Get()
}

// GetEmojiOk returns a tuple with the Emoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceChannelResponse) GetEmojiOk() (*SettingsEmojiResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emoji.Get(), o.Emoji.IsSet()
}

// HasEmoji returns a boolean if a field has been set.
func (o *ResourceChannelResponse) HasEmoji() bool {
	if o != nil && o.Emoji.IsSet() {
		return true
	}

	return false
}

// SetEmoji gets a reference to the given NullableSettingsEmojiResponse and assigns it to the Emoji field.
func (o *ResourceChannelResponse) SetEmoji(v SettingsEmojiResponse) {
	o.Emoji.Set(&v)
}
// SetEmojiNil sets the value for Emoji to be an explicit nil
func (o *ResourceChannelResponse) SetEmojiNil() {
	o.Emoji.Set(nil)
}

// UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
func (o *ResourceChannelResponse) UnsetEmoji() {
	o.Emoji.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceChannelResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceChannelResponse) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *ResourceChannelResponse) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *ResourceChannelResponse) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *ResourceChannelResponse) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *ResourceChannelResponse) UnsetIcon() {
	o.Icon.Unset()
}

// GetDescription returns the Description field value
func (o *ResourceChannelResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ResourceChannelResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ResourceChannelResponse) SetDescription(v string) {
	o.Description = v
}

func (o ResourceChannelResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceChannelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel_id"] = o.ChannelId
	toSerialize["title"] = o.Title
	if o.Emoji.IsSet() {
		toSerialize["emoji"] = o.Emoji.Get()
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

func (o *ResourceChannelResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"channel_id",
		"title",
		"description",
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

	varResourceChannelResponse := _ResourceChannelResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceChannelResponse)

	if err != nil {
		return err
	}

	*o = ResourceChannelResponse(varResourceChannelResponse)

	return err
}

type NullableResourceChannelResponse struct {
	value *ResourceChannelResponse
	isSet bool
}

func (v NullableResourceChannelResponse) Get() *ResourceChannelResponse {
	return v.value
}

func (v *NullableResourceChannelResponse) Set(val *ResourceChannelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceChannelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceChannelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceChannelResponse(val *ResourceChannelResponse) *NullableResourceChannelResponse {
	return &NullableResourceChannelResponse{value: val, isSet: true}
}

func (v NullableResourceChannelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceChannelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


