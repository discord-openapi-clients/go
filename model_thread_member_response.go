/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ThreadMemberResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreadMemberResponse{}

// ThreadMemberResponse struct for ThreadMemberResponse
type ThreadMemberResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	UserId string `json:"user_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	JoinTimestamp time.Time `json:"join_timestamp"`
	Flags int32 `json:"flags"`
	Member NullableGuildMemberResponse `json:"member,omitempty"`
}

type _ThreadMemberResponse ThreadMemberResponse

// NewThreadMemberResponse instantiates a new ThreadMemberResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreadMemberResponse(id string, userId string, joinTimestamp time.Time, flags int32) *ThreadMemberResponse {
	this := ThreadMemberResponse{}
	this.Id = id
	this.UserId = userId
	this.JoinTimestamp = joinTimestamp
	this.Flags = flags
	return &this
}

// NewThreadMemberResponseWithDefaults instantiates a new ThreadMemberResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreadMemberResponseWithDefaults() *ThreadMemberResponse {
	this := ThreadMemberResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ThreadMemberResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ThreadMemberResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ThreadMemberResponse) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *ThreadMemberResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ThreadMemberResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ThreadMemberResponse) SetUserId(v string) {
	o.UserId = v
}

// GetJoinTimestamp returns the JoinTimestamp field value
func (o *ThreadMemberResponse) GetJoinTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.JoinTimestamp
}

// GetJoinTimestampOk returns a tuple with the JoinTimestamp field value
// and a boolean to check if the value has been set.
func (o *ThreadMemberResponse) GetJoinTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinTimestamp, true
}

// SetJoinTimestamp sets field value
func (o *ThreadMemberResponse) SetJoinTimestamp(v time.Time) {
	o.JoinTimestamp = v
}

// GetFlags returns the Flags field value
func (o *ThreadMemberResponse) GetFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *ThreadMemberResponse) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *ThreadMemberResponse) SetFlags(v int32) {
	o.Flags = v
}

// GetMember returns the Member field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreadMemberResponse) GetMember() GuildMemberResponse {
	if o == nil || IsNil(o.Member.Get()) {
		var ret GuildMemberResponse
		return ret
	}
	return *o.Member.Get()
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadMemberResponse) GetMemberOk() (*GuildMemberResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Member.Get(), o.Member.IsSet()
}

// HasMember returns a boolean if a field has been set.
func (o *ThreadMemberResponse) HasMember() bool {
	if o != nil && o.Member.IsSet() {
		return true
	}

	return false
}

// SetMember gets a reference to the given NullableGuildMemberResponse and assigns it to the Member field.
func (o *ThreadMemberResponse) SetMember(v GuildMemberResponse) {
	o.Member.Set(&v)
}
// SetMemberNil sets the value for Member to be an explicit nil
func (o *ThreadMemberResponse) SetMemberNil() {
	o.Member.Set(nil)
}

// UnsetMember ensures that no value is present for Member, not even an explicit nil
func (o *ThreadMemberResponse) UnsetMember() {
	o.Member.Unset()
}

func (o ThreadMemberResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreadMemberResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["user_id"] = o.UserId
	toSerialize["join_timestamp"] = o.JoinTimestamp
	toSerialize["flags"] = o.Flags
	if o.Member.IsSet() {
		toSerialize["member"] = o.Member.Get()
	}
	return toSerialize, nil
}

func (o *ThreadMemberResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"user_id",
		"join_timestamp",
		"flags",
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

	varThreadMemberResponse := _ThreadMemberResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varThreadMemberResponse)

	if err != nil {
		return err
	}

	*o = ThreadMemberResponse(varThreadMemberResponse)

	return err
}

type NullableThreadMemberResponse struct {
	value *ThreadMemberResponse
	isSet bool
}

func (v NullableThreadMemberResponse) Get() *ThreadMemberResponse {
	return v.value
}

func (v *NullableThreadMemberResponse) Set(val *ThreadMemberResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThreadMemberResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThreadMemberResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreadMemberResponse(val *ThreadMemberResponse) *NullableThreadMemberResponse {
	return &NullableThreadMemberResponse{value: val, isSet: true}
}

func (v NullableThreadMemberResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreadMemberResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


