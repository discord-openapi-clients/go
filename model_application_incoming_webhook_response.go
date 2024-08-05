/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApplicationIncomingWebhookResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationIncomingWebhookResponse{}

// ApplicationIncomingWebhookResponse struct for ApplicationIncomingWebhookResponse
type ApplicationIncomingWebhookResponse struct {
	ApplicationId *string `json:"application_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Avatar NullableString `json:"avatar,omitempty"`
	ChannelId *string `json:"channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	GuildId *string `json:"guild_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Type WebhookTypes `json:"type"`
	User NullableUserResponse `json:"user,omitempty"`
}

type _ApplicationIncomingWebhookResponse ApplicationIncomingWebhookResponse

// NewApplicationIncomingWebhookResponse instantiates a new ApplicationIncomingWebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationIncomingWebhookResponse(id string, name string, type_ WebhookTypes) *ApplicationIncomingWebhookResponse {
	this := ApplicationIncomingWebhookResponse{}
	this.Id = id
	this.Name = name
	this.Type = type_
	return &this
}

// NewApplicationIncomingWebhookResponseWithDefaults instantiates a new ApplicationIncomingWebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationIncomingWebhookResponseWithDefaults() *ApplicationIncomingWebhookResponse {
	this := ApplicationIncomingWebhookResponse{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *ApplicationIncomingWebhookResponse) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationIncomingWebhookResponse) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *ApplicationIncomingWebhookResponse) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *ApplicationIncomingWebhookResponse) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationIncomingWebhookResponse) GetAvatar() string {
	if o == nil || IsNil(o.Avatar.Get()) {
		var ret string
		return ret
	}
	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationIncomingWebhookResponse) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// HasAvatar returns a boolean if a field has been set.
func (o *ApplicationIncomingWebhookResponse) HasAvatar() bool {
	if o != nil && o.Avatar.IsSet() {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given NullableString and assigns it to the Avatar field.
func (o *ApplicationIncomingWebhookResponse) SetAvatar(v string) {
	o.Avatar.Set(&v)
}
// SetAvatarNil sets the value for Avatar to be an explicit nil
func (o *ApplicationIncomingWebhookResponse) SetAvatarNil() {
	o.Avatar.Set(nil)
}

// UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
func (o *ApplicationIncomingWebhookResponse) UnsetAvatar() {
	o.Avatar.Unset()
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *ApplicationIncomingWebhookResponse) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationIncomingWebhookResponse) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *ApplicationIncomingWebhookResponse) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *ApplicationIncomingWebhookResponse) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetGuildId returns the GuildId field value if set, zero value otherwise.
func (o *ApplicationIncomingWebhookResponse) GetGuildId() string {
	if o == nil || IsNil(o.GuildId) {
		var ret string
		return ret
	}
	return *o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationIncomingWebhookResponse) GetGuildIdOk() (*string, bool) {
	if o == nil || IsNil(o.GuildId) {
		return nil, false
	}
	return o.GuildId, true
}

// HasGuildId returns a boolean if a field has been set.
func (o *ApplicationIncomingWebhookResponse) HasGuildId() bool {
	if o != nil && !IsNil(o.GuildId) {
		return true
	}

	return false
}

// SetGuildId gets a reference to the given string and assigns it to the GuildId field.
func (o *ApplicationIncomingWebhookResponse) SetGuildId(v string) {
	o.GuildId = &v
}

// GetId returns the Id field value
func (o *ApplicationIncomingWebhookResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationIncomingWebhookResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationIncomingWebhookResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ApplicationIncomingWebhookResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationIncomingWebhookResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationIncomingWebhookResponse) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ApplicationIncomingWebhookResponse) GetType() WebhookTypes {
	if o == nil {
		var ret WebhookTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationIncomingWebhookResponse) GetTypeOk() (*WebhookTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationIncomingWebhookResponse) SetType(v WebhookTypes) {
	o.Type = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationIncomingWebhookResponse) GetUser() UserResponse {
	if o == nil || IsNil(o.User.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationIncomingWebhookResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *ApplicationIncomingWebhookResponse) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUserResponse and assigns it to the User field.
func (o *ApplicationIncomingWebhookResponse) SetUser(v UserResponse) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *ApplicationIncomingWebhookResponse) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *ApplicationIncomingWebhookResponse) UnsetUser() {
	o.User.Unset()
}

func (o ApplicationIncomingWebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationIncomingWebhookResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationId) {
		toSerialize["application_id"] = o.ApplicationId
	}
	if o.Avatar.IsSet() {
		toSerialize["avatar"] = o.Avatar.Get()
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.GuildId) {
		toSerialize["guild_id"] = o.GuildId
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	return toSerialize, nil
}

func (o *ApplicationIncomingWebhookResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"type",
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

	varApplicationIncomingWebhookResponse := _ApplicationIncomingWebhookResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationIncomingWebhookResponse)

	if err != nil {
		return err
	}

	*o = ApplicationIncomingWebhookResponse(varApplicationIncomingWebhookResponse)

	return err
}

type NullableApplicationIncomingWebhookResponse struct {
	value *ApplicationIncomingWebhookResponse
	isSet bool
}

func (v NullableApplicationIncomingWebhookResponse) Get() *ApplicationIncomingWebhookResponse {
	return v.value
}

func (v *NullableApplicationIncomingWebhookResponse) Set(val *ApplicationIncomingWebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationIncomingWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationIncomingWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationIncomingWebhookResponse(val *ApplicationIncomingWebhookResponse) *NullableApplicationIncomingWebhookResponse {
	return &NullableApplicationIncomingWebhookResponse{value: val, isSet: true}
}

func (v NullableApplicationIncomingWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationIncomingWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


