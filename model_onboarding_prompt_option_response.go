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

// checks if the OnboardingPromptOptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnboardingPromptOptionResponse{}

// OnboardingPromptOptionResponse struct for OnboardingPromptOptionResponse
type OnboardingPromptOptionResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Title string `json:"title"`
	Description string `json:"description"`
	Emoji SettingsEmojiResponse `json:"emoji"`
	RoleIds []string `json:"role_ids"`
	ChannelIds []string `json:"channel_ids"`
}

type _OnboardingPromptOptionResponse OnboardingPromptOptionResponse

// NewOnboardingPromptOptionResponse instantiates a new OnboardingPromptOptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnboardingPromptOptionResponse(id string, title string, description string, emoji SettingsEmojiResponse, roleIds []string, channelIds []string) *OnboardingPromptOptionResponse {
	this := OnboardingPromptOptionResponse{}
	this.Id = id
	this.Title = title
	this.Description = description
	this.Emoji = emoji
	this.RoleIds = roleIds
	this.ChannelIds = channelIds
	return &this
}

// NewOnboardingPromptOptionResponseWithDefaults instantiates a new OnboardingPromptOptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnboardingPromptOptionResponseWithDefaults() *OnboardingPromptOptionResponse {
	this := OnboardingPromptOptionResponse{}
	return &this
}

// GetId returns the Id field value
func (o *OnboardingPromptOptionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OnboardingPromptOptionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OnboardingPromptOptionResponse) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *OnboardingPromptOptionResponse) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *OnboardingPromptOptionResponse) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *OnboardingPromptOptionResponse) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *OnboardingPromptOptionResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *OnboardingPromptOptionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *OnboardingPromptOptionResponse) SetDescription(v string) {
	o.Description = v
}

// GetEmoji returns the Emoji field value
func (o *OnboardingPromptOptionResponse) GetEmoji() SettingsEmojiResponse {
	if o == nil {
		var ret SettingsEmojiResponse
		return ret
	}

	return o.Emoji
}

// GetEmojiOk returns a tuple with the Emoji field value
// and a boolean to check if the value has been set.
func (o *OnboardingPromptOptionResponse) GetEmojiOk() (*SettingsEmojiResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Emoji, true
}

// SetEmoji sets field value
func (o *OnboardingPromptOptionResponse) SetEmoji(v SettingsEmojiResponse) {
	o.Emoji = v
}

// GetRoleIds returns the RoleIds field value
func (o *OnboardingPromptOptionResponse) GetRoleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value
// and a boolean to check if the value has been set.
func (o *OnboardingPromptOptionResponse) GetRoleIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleIds, true
}

// SetRoleIds sets field value
func (o *OnboardingPromptOptionResponse) SetRoleIds(v []string) {
	o.RoleIds = v
}

// GetChannelIds returns the ChannelIds field value
func (o *OnboardingPromptOptionResponse) GetChannelIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ChannelIds
}

// GetChannelIdsOk returns a tuple with the ChannelIds field value
// and a boolean to check if the value has been set.
func (o *OnboardingPromptOptionResponse) GetChannelIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelIds, true
}

// SetChannelIds sets field value
func (o *OnboardingPromptOptionResponse) SetChannelIds(v []string) {
	o.ChannelIds = v
}

func (o OnboardingPromptOptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnboardingPromptOptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	toSerialize["description"] = o.Description
	toSerialize["emoji"] = o.Emoji
	toSerialize["role_ids"] = o.RoleIds
	toSerialize["channel_ids"] = o.ChannelIds
	return toSerialize, nil
}

func (o *OnboardingPromptOptionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"title",
		"description",
		"emoji",
		"role_ids",
		"channel_ids",
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

	varOnboardingPromptOptionResponse := _OnboardingPromptOptionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOnboardingPromptOptionResponse)

	if err != nil {
		return err
	}

	*o = OnboardingPromptOptionResponse(varOnboardingPromptOptionResponse)

	return err
}

type NullableOnboardingPromptOptionResponse struct {
	value *OnboardingPromptOptionResponse
	isSet bool
}

func (v NullableOnboardingPromptOptionResponse) Get() *OnboardingPromptOptionResponse {
	return v.value
}

func (v *NullableOnboardingPromptOptionResponse) Set(val *OnboardingPromptOptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOnboardingPromptOptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOnboardingPromptOptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnboardingPromptOptionResponse(val *OnboardingPromptOptionResponse) *NullableOnboardingPromptOptionResponse {
	return &NullableOnboardingPromptOptionResponse{value: val, isSet: true}
}

func (v NullableOnboardingPromptOptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnboardingPromptOptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


