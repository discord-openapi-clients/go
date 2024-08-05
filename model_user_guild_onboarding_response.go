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

// checks if the UserGuildOnboardingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGuildOnboardingResponse{}

// UserGuildOnboardingResponse struct for UserGuildOnboardingResponse
type UserGuildOnboardingResponse struct {
	GuildId string `json:"guild_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Prompts []OnboardingPromptResponse `json:"prompts"`
	DefaultChannelIds []string `json:"default_channel_ids"`
	Enabled bool `json:"enabled"`
}

type _UserGuildOnboardingResponse UserGuildOnboardingResponse

// NewUserGuildOnboardingResponse instantiates a new UserGuildOnboardingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGuildOnboardingResponse(guildId string, prompts []OnboardingPromptResponse, defaultChannelIds []string, enabled bool) *UserGuildOnboardingResponse {
	this := UserGuildOnboardingResponse{}
	this.GuildId = guildId
	this.Prompts = prompts
	this.DefaultChannelIds = defaultChannelIds
	this.Enabled = enabled
	return &this
}

// NewUserGuildOnboardingResponseWithDefaults instantiates a new UserGuildOnboardingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGuildOnboardingResponseWithDefaults() *UserGuildOnboardingResponse {
	this := UserGuildOnboardingResponse{}
	return &this
}

// GetGuildId returns the GuildId field value
func (o *UserGuildOnboardingResponse) GetGuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value
// and a boolean to check if the value has been set.
func (o *UserGuildOnboardingResponse) GetGuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildId, true
}

// SetGuildId sets field value
func (o *UserGuildOnboardingResponse) SetGuildId(v string) {
	o.GuildId = v
}

// GetPrompts returns the Prompts field value
func (o *UserGuildOnboardingResponse) GetPrompts() []OnboardingPromptResponse {
	if o == nil {
		var ret []OnboardingPromptResponse
		return ret
	}

	return o.Prompts
}

// GetPromptsOk returns a tuple with the Prompts field value
// and a boolean to check if the value has been set.
func (o *UserGuildOnboardingResponse) GetPromptsOk() ([]OnboardingPromptResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prompts, true
}

// SetPrompts sets field value
func (o *UserGuildOnboardingResponse) SetPrompts(v []OnboardingPromptResponse) {
	o.Prompts = v
}

// GetDefaultChannelIds returns the DefaultChannelIds field value
func (o *UserGuildOnboardingResponse) GetDefaultChannelIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DefaultChannelIds
}

// GetDefaultChannelIdsOk returns a tuple with the DefaultChannelIds field value
// and a boolean to check if the value has been set.
func (o *UserGuildOnboardingResponse) GetDefaultChannelIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultChannelIds, true
}

// SetDefaultChannelIds sets field value
func (o *UserGuildOnboardingResponse) SetDefaultChannelIds(v []string) {
	o.DefaultChannelIds = v
}

// GetEnabled returns the Enabled field value
func (o *UserGuildOnboardingResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UserGuildOnboardingResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *UserGuildOnboardingResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o UserGuildOnboardingResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGuildOnboardingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["guild_id"] = o.GuildId
	toSerialize["prompts"] = o.Prompts
	toSerialize["default_channel_ids"] = o.DefaultChannelIds
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

func (o *UserGuildOnboardingResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"guild_id",
		"prompts",
		"default_channel_ids",
		"enabled",
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

	varUserGuildOnboardingResponse := _UserGuildOnboardingResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserGuildOnboardingResponse)

	if err != nil {
		return err
	}

	*o = UserGuildOnboardingResponse(varUserGuildOnboardingResponse)

	return err
}

type NullableUserGuildOnboardingResponse struct {
	value *UserGuildOnboardingResponse
	isSet bool
}

func (v NullableUserGuildOnboardingResponse) Get() *UserGuildOnboardingResponse {
	return v.value
}

func (v *NullableUserGuildOnboardingResponse) Set(val *UserGuildOnboardingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGuildOnboardingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGuildOnboardingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGuildOnboardingResponse(val *UserGuildOnboardingResponse) *NullableUserGuildOnboardingResponse {
	return &NullableUserGuildOnboardingResponse{value: val, isSet: true}
}

func (v NullableUserGuildOnboardingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGuildOnboardingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


