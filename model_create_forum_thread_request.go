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

// checks if the CreateForumThreadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateForumThreadRequest{}

// CreateForumThreadRequest struct for CreateForumThreadRequest
type CreateForumThreadRequest struct {
	Name string `json:"name"`
	AutoArchiveDuration NullableThreadAutoArchiveDuration `json:"auto_archive_duration,omitempty"`
	RateLimitPerUser NullableInt32 `json:"rate_limit_per_user,omitempty"`
	AppliedTags []string `json:"applied_tags,omitempty"`
	Message BaseCreateMessageCreateRequest `json:"message"`
}

type _CreateForumThreadRequest CreateForumThreadRequest

// NewCreateForumThreadRequest instantiates a new CreateForumThreadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateForumThreadRequest(name string, message BaseCreateMessageCreateRequest) *CreateForumThreadRequest {
	this := CreateForumThreadRequest{}
	this.Name = name
	this.Message = message
	return &this
}

// NewCreateForumThreadRequestWithDefaults instantiates a new CreateForumThreadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateForumThreadRequestWithDefaults() *CreateForumThreadRequest {
	this := CreateForumThreadRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateForumThreadRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateForumThreadRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateForumThreadRequest) SetName(v string) {
	o.Name = v
}

// GetAutoArchiveDuration returns the AutoArchiveDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateForumThreadRequest) GetAutoArchiveDuration() ThreadAutoArchiveDuration {
	if o == nil || IsNil(o.AutoArchiveDuration.Get()) {
		var ret ThreadAutoArchiveDuration
		return ret
	}
	return *o.AutoArchiveDuration.Get()
}

// GetAutoArchiveDurationOk returns a tuple with the AutoArchiveDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateForumThreadRequest) GetAutoArchiveDurationOk() (*ThreadAutoArchiveDuration, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoArchiveDuration.Get(), o.AutoArchiveDuration.IsSet()
}

// HasAutoArchiveDuration returns a boolean if a field has been set.
func (o *CreateForumThreadRequest) HasAutoArchiveDuration() bool {
	if o != nil && o.AutoArchiveDuration.IsSet() {
		return true
	}

	return false
}

// SetAutoArchiveDuration gets a reference to the given NullableThreadAutoArchiveDuration and assigns it to the AutoArchiveDuration field.
func (o *CreateForumThreadRequest) SetAutoArchiveDuration(v ThreadAutoArchiveDuration) {
	o.AutoArchiveDuration.Set(&v)
}
// SetAutoArchiveDurationNil sets the value for AutoArchiveDuration to be an explicit nil
func (o *CreateForumThreadRequest) SetAutoArchiveDurationNil() {
	o.AutoArchiveDuration.Set(nil)
}

// UnsetAutoArchiveDuration ensures that no value is present for AutoArchiveDuration, not even an explicit nil
func (o *CreateForumThreadRequest) UnsetAutoArchiveDuration() {
	o.AutoArchiveDuration.Unset()
}

// GetRateLimitPerUser returns the RateLimitPerUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateForumThreadRequest) GetRateLimitPerUser() int32 {
	if o == nil || IsNil(o.RateLimitPerUser.Get()) {
		var ret int32
		return ret
	}
	return *o.RateLimitPerUser.Get()
}

// GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateForumThreadRequest) GetRateLimitPerUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RateLimitPerUser.Get(), o.RateLimitPerUser.IsSet()
}

// HasRateLimitPerUser returns a boolean if a field has been set.
func (o *CreateForumThreadRequest) HasRateLimitPerUser() bool {
	if o != nil && o.RateLimitPerUser.IsSet() {
		return true
	}

	return false
}

// SetRateLimitPerUser gets a reference to the given NullableInt32 and assigns it to the RateLimitPerUser field.
func (o *CreateForumThreadRequest) SetRateLimitPerUser(v int32) {
	o.RateLimitPerUser.Set(&v)
}
// SetRateLimitPerUserNil sets the value for RateLimitPerUser to be an explicit nil
func (o *CreateForumThreadRequest) SetRateLimitPerUserNil() {
	o.RateLimitPerUser.Set(nil)
}

// UnsetRateLimitPerUser ensures that no value is present for RateLimitPerUser, not even an explicit nil
func (o *CreateForumThreadRequest) UnsetRateLimitPerUser() {
	o.RateLimitPerUser.Unset()
}

// GetAppliedTags returns the AppliedTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateForumThreadRequest) GetAppliedTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AppliedTags
}

// GetAppliedTagsOk returns a tuple with the AppliedTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateForumThreadRequest) GetAppliedTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.AppliedTags) {
		return nil, false
	}
	return o.AppliedTags, true
}

// HasAppliedTags returns a boolean if a field has been set.
func (o *CreateForumThreadRequest) HasAppliedTags() bool {
	if o != nil && !IsNil(o.AppliedTags) {
		return true
	}

	return false
}

// SetAppliedTags gets a reference to the given []string and assigns it to the AppliedTags field.
func (o *CreateForumThreadRequest) SetAppliedTags(v []string) {
	o.AppliedTags = v
}

// GetMessage returns the Message field value
func (o *CreateForumThreadRequest) GetMessage() BaseCreateMessageCreateRequest {
	if o == nil {
		var ret BaseCreateMessageCreateRequest
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CreateForumThreadRequest) GetMessageOk() (*BaseCreateMessageCreateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CreateForumThreadRequest) SetMessage(v BaseCreateMessageCreateRequest) {
	o.Message = v
}

func (o CreateForumThreadRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateForumThreadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.AutoArchiveDuration.IsSet() {
		toSerialize["auto_archive_duration"] = o.AutoArchiveDuration.Get()
	}
	if o.RateLimitPerUser.IsSet() {
		toSerialize["rate_limit_per_user"] = o.RateLimitPerUser.Get()
	}
	if o.AppliedTags != nil {
		toSerialize["applied_tags"] = o.AppliedTags
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *CreateForumThreadRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"message",
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

	varCreateForumThreadRequest := _CreateForumThreadRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateForumThreadRequest)

	if err != nil {
		return err
	}

	*o = CreateForumThreadRequest(varCreateForumThreadRequest)

	return err
}

type NullableCreateForumThreadRequest struct {
	value *CreateForumThreadRequest
	isSet bool
}

func (v NullableCreateForumThreadRequest) Get() *CreateForumThreadRequest {
	return v.value
}

func (v *NullableCreateForumThreadRequest) Set(val *CreateForumThreadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateForumThreadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateForumThreadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateForumThreadRequest(val *CreateForumThreadRequest) *NullableCreateForumThreadRequest {
	return &NullableCreateForumThreadRequest{value: val, isSet: true}
}

func (v NullableCreateForumThreadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateForumThreadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


