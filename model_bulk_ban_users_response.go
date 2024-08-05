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

// checks if the BulkBanUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkBanUsersResponse{}

// BulkBanUsersResponse struct for BulkBanUsersResponse
type BulkBanUsersResponse struct {
	BannedUsers []string `json:"banned_users"`
	FailedUsers []string `json:"failed_users"`
}

type _BulkBanUsersResponse BulkBanUsersResponse

// NewBulkBanUsersResponse instantiates a new BulkBanUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkBanUsersResponse(bannedUsers []string, failedUsers []string) *BulkBanUsersResponse {
	this := BulkBanUsersResponse{}
	this.BannedUsers = bannedUsers
	this.FailedUsers = failedUsers
	return &this
}

// NewBulkBanUsersResponseWithDefaults instantiates a new BulkBanUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkBanUsersResponseWithDefaults() *BulkBanUsersResponse {
	this := BulkBanUsersResponse{}
	return &this
}

// GetBannedUsers returns the BannedUsers field value
func (o *BulkBanUsersResponse) GetBannedUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BannedUsers
}

// GetBannedUsersOk returns a tuple with the BannedUsers field value
// and a boolean to check if the value has been set.
func (o *BulkBanUsersResponse) GetBannedUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BannedUsers, true
}

// SetBannedUsers sets field value
func (o *BulkBanUsersResponse) SetBannedUsers(v []string) {
	o.BannedUsers = v
}

// GetFailedUsers returns the FailedUsers field value
func (o *BulkBanUsersResponse) GetFailedUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FailedUsers
}

// GetFailedUsersOk returns a tuple with the FailedUsers field value
// and a boolean to check if the value has been set.
func (o *BulkBanUsersResponse) GetFailedUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailedUsers, true
}

// SetFailedUsers sets field value
func (o *BulkBanUsersResponse) SetFailedUsers(v []string) {
	o.FailedUsers = v
}

func (o BulkBanUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkBanUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["banned_users"] = o.BannedUsers
	toSerialize["failed_users"] = o.FailedUsers
	return toSerialize, nil
}

func (o *BulkBanUsersResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"banned_users",
		"failed_users",
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

	varBulkBanUsersResponse := _BulkBanUsersResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBulkBanUsersResponse)

	if err != nil {
		return err
	}

	*o = BulkBanUsersResponse(varBulkBanUsersResponse)

	return err
}

type NullableBulkBanUsersResponse struct {
	value *BulkBanUsersResponse
	isSet bool
}

func (v NullableBulkBanUsersResponse) Get() *BulkBanUsersResponse {
	return v.value
}

func (v *NullableBulkBanUsersResponse) Set(val *BulkBanUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkBanUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkBanUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkBanUsersResponse(val *BulkBanUsersResponse) *NullableBulkBanUsersResponse {
	return &NullableBulkBanUsersResponse{value: val, isSet: true}
}

func (v NullableBulkBanUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkBanUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


