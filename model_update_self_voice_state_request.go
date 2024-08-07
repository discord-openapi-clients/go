/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discord

import (
	"encoding/json"
	"time"
)

// checks if the UpdateSelfVoiceStateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSelfVoiceStateRequest{}

// UpdateSelfVoiceStateRequest struct for UpdateSelfVoiceStateRequest
type UpdateSelfVoiceStateRequest struct {
	RequestToSpeakTimestamp NullableTime `json:"request_to_speak_timestamp,omitempty"`
	Suppress NullableBool `json:"suppress,omitempty"`
	ChannelId *string `json:"channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

// NewUpdateSelfVoiceStateRequest instantiates a new UpdateSelfVoiceStateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSelfVoiceStateRequest() *UpdateSelfVoiceStateRequest {
	this := UpdateSelfVoiceStateRequest{}
	return &this
}

// NewUpdateSelfVoiceStateRequestWithDefaults instantiates a new UpdateSelfVoiceStateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSelfVoiceStateRequestWithDefaults() *UpdateSelfVoiceStateRequest {
	this := UpdateSelfVoiceStateRequest{}
	return &this
}

// GetRequestToSpeakTimestamp returns the RequestToSpeakTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSelfVoiceStateRequest) GetRequestToSpeakTimestamp() time.Time {
	if o == nil || IsNil(o.RequestToSpeakTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RequestToSpeakTimestamp.Get()
}

// GetRequestToSpeakTimestampOk returns a tuple with the RequestToSpeakTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSelfVoiceStateRequest) GetRequestToSpeakTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestToSpeakTimestamp.Get(), o.RequestToSpeakTimestamp.IsSet()
}

// HasRequestToSpeakTimestamp returns a boolean if a field has been set.
func (o *UpdateSelfVoiceStateRequest) HasRequestToSpeakTimestamp() bool {
	if o != nil && o.RequestToSpeakTimestamp.IsSet() {
		return true
	}

	return false
}

// SetRequestToSpeakTimestamp gets a reference to the given NullableTime and assigns it to the RequestToSpeakTimestamp field.
func (o *UpdateSelfVoiceStateRequest) SetRequestToSpeakTimestamp(v time.Time) {
	o.RequestToSpeakTimestamp.Set(&v)
}
// SetRequestToSpeakTimestampNil sets the value for RequestToSpeakTimestamp to be an explicit nil
func (o *UpdateSelfVoiceStateRequest) SetRequestToSpeakTimestampNil() {
	o.RequestToSpeakTimestamp.Set(nil)
}

// UnsetRequestToSpeakTimestamp ensures that no value is present for RequestToSpeakTimestamp, not even an explicit nil
func (o *UpdateSelfVoiceStateRequest) UnsetRequestToSpeakTimestamp() {
	o.RequestToSpeakTimestamp.Unset()
}

// GetSuppress returns the Suppress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSelfVoiceStateRequest) GetSuppress() bool {
	if o == nil || IsNil(o.Suppress.Get()) {
		var ret bool
		return ret
	}
	return *o.Suppress.Get()
}

// GetSuppressOk returns a tuple with the Suppress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSelfVoiceStateRequest) GetSuppressOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Suppress.Get(), o.Suppress.IsSet()
}

// HasSuppress returns a boolean if a field has been set.
func (o *UpdateSelfVoiceStateRequest) HasSuppress() bool {
	if o != nil && o.Suppress.IsSet() {
		return true
	}

	return false
}

// SetSuppress gets a reference to the given NullableBool and assigns it to the Suppress field.
func (o *UpdateSelfVoiceStateRequest) SetSuppress(v bool) {
	o.Suppress.Set(&v)
}
// SetSuppressNil sets the value for Suppress to be an explicit nil
func (o *UpdateSelfVoiceStateRequest) SetSuppressNil() {
	o.Suppress.Set(nil)
}

// UnsetSuppress ensures that no value is present for Suppress, not even an explicit nil
func (o *UpdateSelfVoiceStateRequest) UnsetSuppress() {
	o.Suppress.Unset()
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *UpdateSelfVoiceStateRequest) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSelfVoiceStateRequest) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *UpdateSelfVoiceStateRequest) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *UpdateSelfVoiceStateRequest) SetChannelId(v string) {
	o.ChannelId = &v
}

func (o UpdateSelfVoiceStateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSelfVoiceStateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestToSpeakTimestamp.IsSet() {
		toSerialize["request_to_speak_timestamp"] = o.RequestToSpeakTimestamp.Get()
	}
	if o.Suppress.IsSet() {
		toSerialize["suppress"] = o.Suppress.Get()
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	return toSerialize, nil
}

type NullableUpdateSelfVoiceStateRequest struct {
	value *UpdateSelfVoiceStateRequest
	isSet bool
}

func (v NullableUpdateSelfVoiceStateRequest) Get() *UpdateSelfVoiceStateRequest {
	return v.value
}

func (v *NullableUpdateSelfVoiceStateRequest) Set(val *UpdateSelfVoiceStateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSelfVoiceStateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSelfVoiceStateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSelfVoiceStateRequest(val *UpdateSelfVoiceStateRequest) *NullableUpdateSelfVoiceStateRequest {
	return &NullableUpdateSelfVoiceStateRequest{value: val, isSet: true}
}

func (v NullableUpdateSelfVoiceStateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSelfVoiceStateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


