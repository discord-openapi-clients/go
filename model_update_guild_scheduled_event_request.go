/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discord

import (
	"encoding/json"
	"fmt"
)


// UpdateGuildScheduledEventRequest struct for UpdateGuildScheduledEventRequest
type UpdateGuildScheduledEventRequest struct {
	ExternalScheduledEventPatchRequestPartial *ExternalScheduledEventPatchRequestPartial
	StageScheduledEventPatchRequestPartial *StageScheduledEventPatchRequestPartial
	VoiceScheduledEventPatchRequestPartial *VoiceScheduledEventPatchRequestPartial
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UpdateGuildScheduledEventRequest) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ExternalScheduledEventPatchRequestPartial
	err = json.Unmarshal(data, &dst.ExternalScheduledEventPatchRequestPartial);
	if err == nil {
		jsonExternalScheduledEventPatchRequestPartial, _ := json.Marshal(dst.ExternalScheduledEventPatchRequestPartial)
		if string(jsonExternalScheduledEventPatchRequestPartial) == "{}" { // empty struct
			dst.ExternalScheduledEventPatchRequestPartial = nil
		} else {
			return nil // data stored in dst.ExternalScheduledEventPatchRequestPartial, return on the first match
		}
	} else {
		dst.ExternalScheduledEventPatchRequestPartial = nil
	}

	// try to unmarshal JSON data into StageScheduledEventPatchRequestPartial
	err = json.Unmarshal(data, &dst.StageScheduledEventPatchRequestPartial);
	if err == nil {
		jsonStageScheduledEventPatchRequestPartial, _ := json.Marshal(dst.StageScheduledEventPatchRequestPartial)
		if string(jsonStageScheduledEventPatchRequestPartial) == "{}" { // empty struct
			dst.StageScheduledEventPatchRequestPartial = nil
		} else {
			return nil // data stored in dst.StageScheduledEventPatchRequestPartial, return on the first match
		}
	} else {
		dst.StageScheduledEventPatchRequestPartial = nil
	}

	// try to unmarshal JSON data into VoiceScheduledEventPatchRequestPartial
	err = json.Unmarshal(data, &dst.VoiceScheduledEventPatchRequestPartial);
	if err == nil {
		jsonVoiceScheduledEventPatchRequestPartial, _ := json.Marshal(dst.VoiceScheduledEventPatchRequestPartial)
		if string(jsonVoiceScheduledEventPatchRequestPartial) == "{}" { // empty struct
			dst.VoiceScheduledEventPatchRequestPartial = nil
		} else {
			return nil // data stored in dst.VoiceScheduledEventPatchRequestPartial, return on the first match
		}
	} else {
		dst.VoiceScheduledEventPatchRequestPartial = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(UpdateGuildScheduledEventRequest)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UpdateGuildScheduledEventRequest) MarshalJSON() ([]byte, error) {
	if src.ExternalScheduledEventPatchRequestPartial != nil {
		return json.Marshal(&src.ExternalScheduledEventPatchRequestPartial)
	}

	if src.StageScheduledEventPatchRequestPartial != nil {
		return json.Marshal(&src.StageScheduledEventPatchRequestPartial)
	}

	if src.VoiceScheduledEventPatchRequestPartial != nil {
		return json.Marshal(&src.VoiceScheduledEventPatchRequestPartial)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableUpdateGuildScheduledEventRequest struct {
	value *UpdateGuildScheduledEventRequest
	isSet bool
}

func (v NullableUpdateGuildScheduledEventRequest) Get() *UpdateGuildScheduledEventRequest {
	return v.value
}

func (v *NullableUpdateGuildScheduledEventRequest) Set(val *UpdateGuildScheduledEventRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGuildScheduledEventRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGuildScheduledEventRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGuildScheduledEventRequest(val *UpdateGuildScheduledEventRequest) *NullableUpdateGuildScheduledEventRequest {
	return &NullableUpdateGuildScheduledEventRequest{value: val, isSet: true}
}

func (v NullableUpdateGuildScheduledEventRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGuildScheduledEventRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


