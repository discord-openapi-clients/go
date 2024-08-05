/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// GetGuildScheduledEvent200Response - struct for GetGuildScheduledEvent200Response
type GetGuildScheduledEvent200Response struct {
	ExternalScheduledEventResponse *ExternalScheduledEventResponse
	StageScheduledEventResponse *StageScheduledEventResponse
	VoiceScheduledEventResponse *VoiceScheduledEventResponse
}

// ExternalScheduledEventResponseAsGetGuildScheduledEvent200Response is a convenience function that returns ExternalScheduledEventResponse wrapped in GetGuildScheduledEvent200Response
func ExternalScheduledEventResponseAsGetGuildScheduledEvent200Response(v *ExternalScheduledEventResponse) GetGuildScheduledEvent200Response {
	return GetGuildScheduledEvent200Response{
		ExternalScheduledEventResponse: v,
	}
}

// StageScheduledEventResponseAsGetGuildScheduledEvent200Response is a convenience function that returns StageScheduledEventResponse wrapped in GetGuildScheduledEvent200Response
func StageScheduledEventResponseAsGetGuildScheduledEvent200Response(v *StageScheduledEventResponse) GetGuildScheduledEvent200Response {
	return GetGuildScheduledEvent200Response{
		StageScheduledEventResponse: v,
	}
}

// VoiceScheduledEventResponseAsGetGuildScheduledEvent200Response is a convenience function that returns VoiceScheduledEventResponse wrapped in GetGuildScheduledEvent200Response
func VoiceScheduledEventResponseAsGetGuildScheduledEvent200Response(v *VoiceScheduledEventResponse) GetGuildScheduledEvent200Response {
	return GetGuildScheduledEvent200Response{
		VoiceScheduledEventResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetGuildScheduledEvent200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ExternalScheduledEventResponse
	err = newStrictDecoder(data).Decode(&dst.ExternalScheduledEventResponse)
	if err == nil {
		jsonExternalScheduledEventResponse, _ := json.Marshal(dst.ExternalScheduledEventResponse)
		if string(jsonExternalScheduledEventResponse) == "{}" { // empty struct
			dst.ExternalScheduledEventResponse = nil
		} else {
			if err = validator.Validate(dst.ExternalScheduledEventResponse); err != nil {
				dst.ExternalScheduledEventResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ExternalScheduledEventResponse = nil
	}

	// try to unmarshal data into StageScheduledEventResponse
	err = newStrictDecoder(data).Decode(&dst.StageScheduledEventResponse)
	if err == nil {
		jsonStageScheduledEventResponse, _ := json.Marshal(dst.StageScheduledEventResponse)
		if string(jsonStageScheduledEventResponse) == "{}" { // empty struct
			dst.StageScheduledEventResponse = nil
		} else {
			if err = validator.Validate(dst.StageScheduledEventResponse); err != nil {
				dst.StageScheduledEventResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.StageScheduledEventResponse = nil
	}

	// try to unmarshal data into VoiceScheduledEventResponse
	err = newStrictDecoder(data).Decode(&dst.VoiceScheduledEventResponse)
	if err == nil {
		jsonVoiceScheduledEventResponse, _ := json.Marshal(dst.VoiceScheduledEventResponse)
		if string(jsonVoiceScheduledEventResponse) == "{}" { // empty struct
			dst.VoiceScheduledEventResponse = nil
		} else {
			if err = validator.Validate(dst.VoiceScheduledEventResponse); err != nil {
				dst.VoiceScheduledEventResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.VoiceScheduledEventResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ExternalScheduledEventResponse = nil
		dst.StageScheduledEventResponse = nil
		dst.VoiceScheduledEventResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetGuildScheduledEvent200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetGuildScheduledEvent200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetGuildScheduledEvent200Response) MarshalJSON() ([]byte, error) {
	if src.ExternalScheduledEventResponse != nil {
		return json.Marshal(&src.ExternalScheduledEventResponse)
	}

	if src.StageScheduledEventResponse != nil {
		return json.Marshal(&src.StageScheduledEventResponse)
	}

	if src.VoiceScheduledEventResponse != nil {
		return json.Marshal(&src.VoiceScheduledEventResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetGuildScheduledEvent200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ExternalScheduledEventResponse != nil {
		return obj.ExternalScheduledEventResponse
	}

	if obj.StageScheduledEventResponse != nil {
		return obj.StageScheduledEventResponse
	}

	if obj.VoiceScheduledEventResponse != nil {
		return obj.VoiceScheduledEventResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetGuildScheduledEvent200Response struct {
	value *GetGuildScheduledEvent200Response
	isSet bool
}

func (v NullableGetGuildScheduledEvent200Response) Get() *GetGuildScheduledEvent200Response {
	return v.value
}

func (v *NullableGetGuildScheduledEvent200Response) Set(val *GetGuildScheduledEvent200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGuildScheduledEvent200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGuildScheduledEvent200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGuildScheduledEvent200Response(val *GetGuildScheduledEvent200Response) *NullableGetGuildScheduledEvent200Response {
	return &NullableGetGuildScheduledEvent200Response{value: val, isSet: true}
}

func (v NullableGetGuildScheduledEvent200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGuildScheduledEvent200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


