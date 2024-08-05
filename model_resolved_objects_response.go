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

// checks if the ResolvedObjectsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResolvedObjectsResponse{}

// ResolvedObjectsResponse struct for ResolvedObjectsResponse
type ResolvedObjectsResponse struct {
	Users map[string]UserResponse `json:"users"`
	Members map[string]GuildMemberResponse `json:"members"`
	Channels map[string]ListGuildChannels200ResponseInner `json:"channels"`
	Roles map[string]GuildRoleResponse `json:"roles"`
}

type _ResolvedObjectsResponse ResolvedObjectsResponse

// NewResolvedObjectsResponse instantiates a new ResolvedObjectsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResolvedObjectsResponse(users map[string]UserResponse, members map[string]GuildMemberResponse, channels map[string]ListGuildChannels200ResponseInner, roles map[string]GuildRoleResponse) *ResolvedObjectsResponse {
	this := ResolvedObjectsResponse{}
	this.Users = users
	this.Members = members
	this.Channels = channels
	this.Roles = roles
	return &this
}

// NewResolvedObjectsResponseWithDefaults instantiates a new ResolvedObjectsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResolvedObjectsResponseWithDefaults() *ResolvedObjectsResponse {
	this := ResolvedObjectsResponse{}
	return &this
}

// GetUsers returns the Users field value
func (o *ResolvedObjectsResponse) GetUsers() map[string]UserResponse {
	if o == nil {
		var ret map[string]UserResponse
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *ResolvedObjectsResponse) GetUsersOk() (*map[string]UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Users, true
}

// SetUsers sets field value
func (o *ResolvedObjectsResponse) SetUsers(v map[string]UserResponse) {
	o.Users = v
}

// GetMembers returns the Members field value
func (o *ResolvedObjectsResponse) GetMembers() map[string]GuildMemberResponse {
	if o == nil {
		var ret map[string]GuildMemberResponse
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *ResolvedObjectsResponse) GetMembersOk() (*map[string]GuildMemberResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Members, true
}

// SetMembers sets field value
func (o *ResolvedObjectsResponse) SetMembers(v map[string]GuildMemberResponse) {
	o.Members = v
}

// GetChannels returns the Channels field value
func (o *ResolvedObjectsResponse) GetChannels() map[string]ListGuildChannels200ResponseInner {
	if o == nil {
		var ret map[string]ListGuildChannels200ResponseInner
		return ret
	}

	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value
// and a boolean to check if the value has been set.
func (o *ResolvedObjectsResponse) GetChannelsOk() (*map[string]ListGuildChannels200ResponseInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channels, true
}

// SetChannels sets field value
func (o *ResolvedObjectsResponse) SetChannels(v map[string]ListGuildChannels200ResponseInner) {
	o.Channels = v
}

// GetRoles returns the Roles field value
func (o *ResolvedObjectsResponse) GetRoles() map[string]GuildRoleResponse {
	if o == nil {
		var ret map[string]GuildRoleResponse
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *ResolvedObjectsResponse) GetRolesOk() (*map[string]GuildRoleResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *ResolvedObjectsResponse) SetRoles(v map[string]GuildRoleResponse) {
	o.Roles = v
}

func (o ResolvedObjectsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResolvedObjectsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	toSerialize["members"] = o.Members
	toSerialize["channels"] = o.Channels
	toSerialize["roles"] = o.Roles
	return toSerialize, nil
}

func (o *ResolvedObjectsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"users",
		"members",
		"channels",
		"roles",
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

	varResolvedObjectsResponse := _ResolvedObjectsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResolvedObjectsResponse)

	if err != nil {
		return err
	}

	*o = ResolvedObjectsResponse(varResolvedObjectsResponse)

	return err
}

type NullableResolvedObjectsResponse struct {
	value *ResolvedObjectsResponse
	isSet bool
}

func (v NullableResolvedObjectsResponse) Get() *ResolvedObjectsResponse {
	return v.value
}

func (v *NullableResolvedObjectsResponse) Set(val *ResolvedObjectsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResolvedObjectsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResolvedObjectsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResolvedObjectsResponse(val *ResolvedObjectsResponse) *NullableResolvedObjectsResponse {
	return &NullableResolvedObjectsResponse{value: val, isSet: true}
}

func (v NullableResolvedObjectsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResolvedObjectsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

