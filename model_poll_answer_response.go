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

// checks if the PollAnswerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PollAnswerResponse{}

// PollAnswerResponse struct for PollAnswerResponse
type PollAnswerResponse struct {
	AnswerId int32 `json:"answer_id"`
	PollMedia PollMediaResponse `json:"poll_media"`
}

type _PollAnswerResponse PollAnswerResponse

// NewPollAnswerResponse instantiates a new PollAnswerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPollAnswerResponse(answerId int32, pollMedia PollMediaResponse) *PollAnswerResponse {
	this := PollAnswerResponse{}
	this.AnswerId = answerId
	this.PollMedia = pollMedia
	return &this
}

// NewPollAnswerResponseWithDefaults instantiates a new PollAnswerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPollAnswerResponseWithDefaults() *PollAnswerResponse {
	this := PollAnswerResponse{}
	return &this
}

// GetAnswerId returns the AnswerId field value
func (o *PollAnswerResponse) GetAnswerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AnswerId
}

// GetAnswerIdOk returns a tuple with the AnswerId field value
// and a boolean to check if the value has been set.
func (o *PollAnswerResponse) GetAnswerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnswerId, true
}

// SetAnswerId sets field value
func (o *PollAnswerResponse) SetAnswerId(v int32) {
	o.AnswerId = v
}

// GetPollMedia returns the PollMedia field value
func (o *PollAnswerResponse) GetPollMedia() PollMediaResponse {
	if o == nil {
		var ret PollMediaResponse
		return ret
	}

	return o.PollMedia
}

// GetPollMediaOk returns a tuple with the PollMedia field value
// and a boolean to check if the value has been set.
func (o *PollAnswerResponse) GetPollMediaOk() (*PollMediaResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PollMedia, true
}

// SetPollMedia sets field value
func (o *PollAnswerResponse) SetPollMedia(v PollMediaResponse) {
	o.PollMedia = v
}

func (o PollAnswerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PollAnswerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["answer_id"] = o.AnswerId
	toSerialize["poll_media"] = o.PollMedia
	return toSerialize, nil
}

func (o *PollAnswerResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"answer_id",
		"poll_media",
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

	varPollAnswerResponse := _PollAnswerResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPollAnswerResponse)

	if err != nil {
		return err
	}

	*o = PollAnswerResponse(varPollAnswerResponse)

	return err
}

type NullablePollAnswerResponse struct {
	value *PollAnswerResponse
	isSet bool
}

func (v NullablePollAnswerResponse) Get() *PollAnswerResponse {
	return v.value
}

func (v *NullablePollAnswerResponse) Set(val *PollAnswerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePollAnswerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePollAnswerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePollAnswerResponse(val *PollAnswerResponse) *NullablePollAnswerResponse {
	return &NullablePollAnswerResponse{value: val, isSet: true}
}

func (v NullablePollAnswerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePollAnswerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


