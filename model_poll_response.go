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
	"bytes"
	"fmt"
)

// checks if the PollResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PollResponse{}

// PollResponse struct for PollResponse
type PollResponse struct {
	Question PollMediaResponse `json:"question"`
	Answers []PollAnswerResponse `json:"answers"`
	Expiry time.Time `json:"expiry"`
	AllowMultiselect bool `json:"allow_multiselect"`
	LayoutType PollLayoutTypes `json:"layout_type"`
	Results PollResultsResponse `json:"results"`
}

type _PollResponse PollResponse

// NewPollResponse instantiates a new PollResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPollResponse(question PollMediaResponse, answers []PollAnswerResponse, expiry time.Time, allowMultiselect bool, layoutType PollLayoutTypes, results PollResultsResponse) *PollResponse {
	this := PollResponse{}
	this.Question = question
	this.Answers = answers
	this.Expiry = expiry
	this.AllowMultiselect = allowMultiselect
	this.LayoutType = layoutType
	this.Results = results
	return &this
}

// NewPollResponseWithDefaults instantiates a new PollResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPollResponseWithDefaults() *PollResponse {
	this := PollResponse{}
	return &this
}

// GetQuestion returns the Question field value
func (o *PollResponse) GetQuestion() PollMediaResponse {
	if o == nil {
		var ret PollMediaResponse
		return ret
	}

	return o.Question
}

// GetQuestionOk returns a tuple with the Question field value
// and a boolean to check if the value has been set.
func (o *PollResponse) GetQuestionOk() (*PollMediaResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Question, true
}

// SetQuestion sets field value
func (o *PollResponse) SetQuestion(v PollMediaResponse) {
	o.Question = v
}

// GetAnswers returns the Answers field value
func (o *PollResponse) GetAnswers() []PollAnswerResponse {
	if o == nil {
		var ret []PollAnswerResponse
		return ret
	}

	return o.Answers
}

// GetAnswersOk returns a tuple with the Answers field value
// and a boolean to check if the value has been set.
func (o *PollResponse) GetAnswersOk() ([]PollAnswerResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Answers, true
}

// SetAnswers sets field value
func (o *PollResponse) SetAnswers(v []PollAnswerResponse) {
	o.Answers = v
}

// GetExpiry returns the Expiry field value
func (o *PollResponse) GetExpiry() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value
// and a boolean to check if the value has been set.
func (o *PollResponse) GetExpiryOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiry, true
}

// SetExpiry sets field value
func (o *PollResponse) SetExpiry(v time.Time) {
	o.Expiry = v
}

// GetAllowMultiselect returns the AllowMultiselect field value
func (o *PollResponse) GetAllowMultiselect() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowMultiselect
}

// GetAllowMultiselectOk returns a tuple with the AllowMultiselect field value
// and a boolean to check if the value has been set.
func (o *PollResponse) GetAllowMultiselectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowMultiselect, true
}

// SetAllowMultiselect sets field value
func (o *PollResponse) SetAllowMultiselect(v bool) {
	o.AllowMultiselect = v
}

// GetLayoutType returns the LayoutType field value
func (o *PollResponse) GetLayoutType() PollLayoutTypes {
	if o == nil {
		var ret PollLayoutTypes
		return ret
	}

	return o.LayoutType
}

// GetLayoutTypeOk returns a tuple with the LayoutType field value
// and a boolean to check if the value has been set.
func (o *PollResponse) GetLayoutTypeOk() (*PollLayoutTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LayoutType, true
}

// SetLayoutType sets field value
func (o *PollResponse) SetLayoutType(v PollLayoutTypes) {
	o.LayoutType = v
}

// GetResults returns the Results field value
func (o *PollResponse) GetResults() PollResultsResponse {
	if o == nil {
		var ret PollResultsResponse
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PollResponse) GetResultsOk() (*PollResultsResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PollResponse) SetResults(v PollResultsResponse) {
	o.Results = v
}

func (o PollResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PollResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["question"] = o.Question
	toSerialize["answers"] = o.Answers
	toSerialize["expiry"] = o.Expiry
	toSerialize["allow_multiselect"] = o.AllowMultiselect
	toSerialize["layout_type"] = o.LayoutType
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *PollResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"question",
		"answers",
		"expiry",
		"allow_multiselect",
		"layout_type",
		"results",
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

	varPollResponse := _PollResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPollResponse)

	if err != nil {
		return err
	}

	*o = PollResponse(varPollResponse)

	return err
}

type NullablePollResponse struct {
	value *PollResponse
	isSet bool
}

func (v NullablePollResponse) Get() *PollResponse {
	return v.value
}

func (v *NullablePollResponse) Set(val *PollResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePollResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePollResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePollResponse(val *PollResponse) *NullablePollResponse {
	return &NullablePollResponse{value: val, isSet: true}
}

func (v NullablePollResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePollResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

