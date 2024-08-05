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

// checks if the GuildHomeSettingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildHomeSettingsResponse{}

// GuildHomeSettingsResponse struct for GuildHomeSettingsResponse
type GuildHomeSettingsResponse struct {
	GuildId string `json:"guild_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Enabled bool `json:"enabled"`
	WelcomeMessage NullableWelcomeMessageResponse `json:"welcome_message,omitempty"`
	NewMemberActions []GuildHomeSettingsResponseNewMemberActionsInner `json:"new_member_actions,omitempty"`
	ResourceChannels []GuildHomeSettingsResponseResourceChannelsInner `json:"resource_channels,omitempty"`
}

type _GuildHomeSettingsResponse GuildHomeSettingsResponse

// NewGuildHomeSettingsResponse instantiates a new GuildHomeSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildHomeSettingsResponse(guildId string, enabled bool) *GuildHomeSettingsResponse {
	this := GuildHomeSettingsResponse{}
	this.GuildId = guildId
	this.Enabled = enabled
	return &this
}

// NewGuildHomeSettingsResponseWithDefaults instantiates a new GuildHomeSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildHomeSettingsResponseWithDefaults() *GuildHomeSettingsResponse {
	this := GuildHomeSettingsResponse{}
	return &this
}

// GetGuildId returns the GuildId field value
func (o *GuildHomeSettingsResponse) GetGuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value
// and a boolean to check if the value has been set.
func (o *GuildHomeSettingsResponse) GetGuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildId, true
}

// SetGuildId sets field value
func (o *GuildHomeSettingsResponse) SetGuildId(v string) {
	o.GuildId = v
}

// GetEnabled returns the Enabled field value
func (o *GuildHomeSettingsResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GuildHomeSettingsResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GuildHomeSettingsResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetWelcomeMessage returns the WelcomeMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildHomeSettingsResponse) GetWelcomeMessage() WelcomeMessageResponse {
	if o == nil || IsNil(o.WelcomeMessage.Get()) {
		var ret WelcomeMessageResponse
		return ret
	}
	return *o.WelcomeMessage.Get()
}

// GetWelcomeMessageOk returns a tuple with the WelcomeMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildHomeSettingsResponse) GetWelcomeMessageOk() (*WelcomeMessageResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.WelcomeMessage.Get(), o.WelcomeMessage.IsSet()
}

// HasWelcomeMessage returns a boolean if a field has been set.
func (o *GuildHomeSettingsResponse) HasWelcomeMessage() bool {
	if o != nil && o.WelcomeMessage.IsSet() {
		return true
	}

	return false
}

// SetWelcomeMessage gets a reference to the given NullableWelcomeMessageResponse and assigns it to the WelcomeMessage field.
func (o *GuildHomeSettingsResponse) SetWelcomeMessage(v WelcomeMessageResponse) {
	o.WelcomeMessage.Set(&v)
}
// SetWelcomeMessageNil sets the value for WelcomeMessage to be an explicit nil
func (o *GuildHomeSettingsResponse) SetWelcomeMessageNil() {
	o.WelcomeMessage.Set(nil)
}

// UnsetWelcomeMessage ensures that no value is present for WelcomeMessage, not even an explicit nil
func (o *GuildHomeSettingsResponse) UnsetWelcomeMessage() {
	o.WelcomeMessage.Unset()
}

// GetNewMemberActions returns the NewMemberActions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildHomeSettingsResponse) GetNewMemberActions() []GuildHomeSettingsResponseNewMemberActionsInner {
	if o == nil {
		var ret []GuildHomeSettingsResponseNewMemberActionsInner
		return ret
	}
	return o.NewMemberActions
}

// GetNewMemberActionsOk returns a tuple with the NewMemberActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildHomeSettingsResponse) GetNewMemberActionsOk() ([]GuildHomeSettingsResponseNewMemberActionsInner, bool) {
	if o == nil || IsNil(o.NewMemberActions) {
		return nil, false
	}
	return o.NewMemberActions, true
}

// HasNewMemberActions returns a boolean if a field has been set.
func (o *GuildHomeSettingsResponse) HasNewMemberActions() bool {
	if o != nil && !IsNil(o.NewMemberActions) {
		return true
	}

	return false
}

// SetNewMemberActions gets a reference to the given []GuildHomeSettingsResponseNewMemberActionsInner and assigns it to the NewMemberActions field.
func (o *GuildHomeSettingsResponse) SetNewMemberActions(v []GuildHomeSettingsResponseNewMemberActionsInner) {
	o.NewMemberActions = v
}

// GetResourceChannels returns the ResourceChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuildHomeSettingsResponse) GetResourceChannels() []GuildHomeSettingsResponseResourceChannelsInner {
	if o == nil {
		var ret []GuildHomeSettingsResponseResourceChannelsInner
		return ret
	}
	return o.ResourceChannels
}

// GetResourceChannelsOk returns a tuple with the ResourceChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuildHomeSettingsResponse) GetResourceChannelsOk() ([]GuildHomeSettingsResponseResourceChannelsInner, bool) {
	if o == nil || IsNil(o.ResourceChannels) {
		return nil, false
	}
	return o.ResourceChannels, true
}

// HasResourceChannels returns a boolean if a field has been set.
func (o *GuildHomeSettingsResponse) HasResourceChannels() bool {
	if o != nil && !IsNil(o.ResourceChannels) {
		return true
	}

	return false
}

// SetResourceChannels gets a reference to the given []GuildHomeSettingsResponseResourceChannelsInner and assigns it to the ResourceChannels field.
func (o *GuildHomeSettingsResponse) SetResourceChannels(v []GuildHomeSettingsResponseResourceChannelsInner) {
	o.ResourceChannels = v
}

func (o GuildHomeSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildHomeSettingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["guild_id"] = o.GuildId
	toSerialize["enabled"] = o.Enabled
	if o.WelcomeMessage.IsSet() {
		toSerialize["welcome_message"] = o.WelcomeMessage.Get()
	}
	if o.NewMemberActions != nil {
		toSerialize["new_member_actions"] = o.NewMemberActions
	}
	if o.ResourceChannels != nil {
		toSerialize["resource_channels"] = o.ResourceChannels
	}
	return toSerialize, nil
}

func (o *GuildHomeSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"guild_id",
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

	varGuildHomeSettingsResponse := _GuildHomeSettingsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildHomeSettingsResponse)

	if err != nil {
		return err
	}

	*o = GuildHomeSettingsResponse(varGuildHomeSettingsResponse)

	return err
}

type NullableGuildHomeSettingsResponse struct {
	value *GuildHomeSettingsResponse
	isSet bool
}

func (v NullableGuildHomeSettingsResponse) Get() *GuildHomeSettingsResponse {
	return v.value
}

func (v *NullableGuildHomeSettingsResponse) Set(val *GuildHomeSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildHomeSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildHomeSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildHomeSettingsResponse(val *GuildHomeSettingsResponse) *NullableGuildHomeSettingsResponse {
	return &NullableGuildHomeSettingsResponse{value: val, isSet: true}
}

func (v NullableGuildHomeSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildHomeSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


