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

// checks if the MessageCallResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageCallResponse{}

// MessageCallResponse struct for MessageCallResponse
type MessageCallResponse struct {
	EndedTimestamp NullableTime `json:"ended_timestamp,omitempty"`
	Participants []string `json:"participants"`
}

type _MessageCallResponse MessageCallResponse

// NewMessageCallResponse instantiates a new MessageCallResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageCallResponse(participants []string) *MessageCallResponse {
	this := MessageCallResponse{}
	this.Participants = participants
	return &this
}

// NewMessageCallResponseWithDefaults instantiates a new MessageCallResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageCallResponseWithDefaults() *MessageCallResponse {
	this := MessageCallResponse{}
	return &this
}

// GetEndedTimestamp returns the EndedTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageCallResponse) GetEndedTimestamp() time.Time {
	if o == nil || IsNil(o.EndedTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EndedTimestamp.Get()
}

// GetEndedTimestampOk returns a tuple with the EndedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageCallResponse) GetEndedTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndedTimestamp.Get(), o.EndedTimestamp.IsSet()
}

// HasEndedTimestamp returns a boolean if a field has been set.
func (o *MessageCallResponse) HasEndedTimestamp() bool {
	if o != nil && o.EndedTimestamp.IsSet() {
		return true
	}

	return false
}

// SetEndedTimestamp gets a reference to the given NullableTime and assigns it to the EndedTimestamp field.
func (o *MessageCallResponse) SetEndedTimestamp(v time.Time) {
	o.EndedTimestamp.Set(&v)
}
// SetEndedTimestampNil sets the value for EndedTimestamp to be an explicit nil
func (o *MessageCallResponse) SetEndedTimestampNil() {
	o.EndedTimestamp.Set(nil)
}

// UnsetEndedTimestamp ensures that no value is present for EndedTimestamp, not even an explicit nil
func (o *MessageCallResponse) UnsetEndedTimestamp() {
	o.EndedTimestamp.Unset()
}

// GetParticipants returns the Participants field value
func (o *MessageCallResponse) GetParticipants() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value
// and a boolean to check if the value has been set.
func (o *MessageCallResponse) GetParticipantsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Participants, true
}

// SetParticipants sets field value
func (o *MessageCallResponse) SetParticipants(v []string) {
	o.Participants = v
}

func (o MessageCallResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageCallResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EndedTimestamp.IsSet() {
		toSerialize["ended_timestamp"] = o.EndedTimestamp.Get()
	}
	toSerialize["participants"] = o.Participants
	return toSerialize, nil
}

func (o *MessageCallResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"participants",
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

	varMessageCallResponse := _MessageCallResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageCallResponse)

	if err != nil {
		return err
	}

	*o = MessageCallResponse(varMessageCallResponse)

	return err
}

type NullableMessageCallResponse struct {
	value *MessageCallResponse
	isSet bool
}

func (v NullableMessageCallResponse) Get() *MessageCallResponse {
	return v.value
}

func (v *NullableMessageCallResponse) Set(val *MessageCallResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageCallResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageCallResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageCallResponse(val *MessageCallResponse) *NullableMessageCallResponse {
	return &NullableMessageCallResponse{value: val, isSet: true}
}

func (v NullableMessageCallResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageCallResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


