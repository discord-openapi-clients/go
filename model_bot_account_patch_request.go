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

// checks if the BotAccountPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BotAccountPatchRequest{}

// BotAccountPatchRequest struct for BotAccountPatchRequest
type BotAccountPatchRequest struct {
	Username string `json:"username"`
	Avatar NullableString `json:"avatar,omitempty"`
	Banner NullableString `json:"banner,omitempty"`
}

type _BotAccountPatchRequest BotAccountPatchRequest

// NewBotAccountPatchRequest instantiates a new BotAccountPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotAccountPatchRequest(username string) *BotAccountPatchRequest {
	this := BotAccountPatchRequest{}
	this.Username = username
	return &this
}

// NewBotAccountPatchRequestWithDefaults instantiates a new BotAccountPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotAccountPatchRequestWithDefaults() *BotAccountPatchRequest {
	this := BotAccountPatchRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *BotAccountPatchRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *BotAccountPatchRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *BotAccountPatchRequest) SetUsername(v string) {
	o.Username = v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BotAccountPatchRequest) GetAvatar() string {
	if o == nil || IsNil(o.Avatar.Get()) {
		var ret string
		return ret
	}
	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BotAccountPatchRequest) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// HasAvatar returns a boolean if a field has been set.
func (o *BotAccountPatchRequest) HasAvatar() bool {
	if o != nil && o.Avatar.IsSet() {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given NullableString and assigns it to the Avatar field.
func (o *BotAccountPatchRequest) SetAvatar(v string) {
	o.Avatar.Set(&v)
}
// SetAvatarNil sets the value for Avatar to be an explicit nil
func (o *BotAccountPatchRequest) SetAvatarNil() {
	o.Avatar.Set(nil)
}

// UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
func (o *BotAccountPatchRequest) UnsetAvatar() {
	o.Avatar.Unset()
}

// GetBanner returns the Banner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BotAccountPatchRequest) GetBanner() string {
	if o == nil || IsNil(o.Banner.Get()) {
		var ret string
		return ret
	}
	return *o.Banner.Get()
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BotAccountPatchRequest) GetBannerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Banner.Get(), o.Banner.IsSet()
}

// HasBanner returns a boolean if a field has been set.
func (o *BotAccountPatchRequest) HasBanner() bool {
	if o != nil && o.Banner.IsSet() {
		return true
	}

	return false
}

// SetBanner gets a reference to the given NullableString and assigns it to the Banner field.
func (o *BotAccountPatchRequest) SetBanner(v string) {
	o.Banner.Set(&v)
}
// SetBannerNil sets the value for Banner to be an explicit nil
func (o *BotAccountPatchRequest) SetBannerNil() {
	o.Banner.Set(nil)
}

// UnsetBanner ensures that no value is present for Banner, not even an explicit nil
func (o *BotAccountPatchRequest) UnsetBanner() {
	o.Banner.Unset()
}

func (o BotAccountPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BotAccountPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	if o.Avatar.IsSet() {
		toSerialize["avatar"] = o.Avatar.Get()
	}
	if o.Banner.IsSet() {
		toSerialize["banner"] = o.Banner.Get()
	}
	return toSerialize, nil
}

func (o *BotAccountPatchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
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

	varBotAccountPatchRequest := _BotAccountPatchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBotAccountPatchRequest)

	if err != nil {
		return err
	}

	*o = BotAccountPatchRequest(varBotAccountPatchRequest)

	return err
}

type NullableBotAccountPatchRequest struct {
	value *BotAccountPatchRequest
	isSet bool
}

func (v NullableBotAccountPatchRequest) Get() *BotAccountPatchRequest {
	return v.value
}

func (v *NullableBotAccountPatchRequest) Set(val *BotAccountPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBotAccountPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBotAccountPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBotAccountPatchRequest(val *BotAccountPatchRequest) *NullableBotAccountPatchRequest {
	return &NullableBotAccountPatchRequest{value: val, isSet: true}
}

func (v NullableBotAccountPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBotAccountPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


