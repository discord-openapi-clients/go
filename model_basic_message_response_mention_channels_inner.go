/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discord

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// BasicMessageResponseMentionChannelsInner - struct for BasicMessageResponseMentionChannelsInner
type BasicMessageResponseMentionChannelsInner struct {
	MessageMentionChannelResponse *MessageMentionChannelResponse
}

// MessageMentionChannelResponseAsBasicMessageResponseMentionChannelsInner is a convenience function that returns MessageMentionChannelResponse wrapped in BasicMessageResponseMentionChannelsInner
func MessageMentionChannelResponseAsBasicMessageResponseMentionChannelsInner(v *MessageMentionChannelResponse) BasicMessageResponseMentionChannelsInner {
	return BasicMessageResponseMentionChannelsInner{
		MessageMentionChannelResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *BasicMessageResponseMentionChannelsInner) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into MessageMentionChannelResponse
	err = newStrictDecoder(data).Decode(&dst.MessageMentionChannelResponse)
	if err == nil {
		jsonMessageMentionChannelResponse, _ := json.Marshal(dst.MessageMentionChannelResponse)
		if string(jsonMessageMentionChannelResponse) == "{}" { // empty struct
			dst.MessageMentionChannelResponse = nil
		} else {
			if err = validator.Validate(dst.MessageMentionChannelResponse); err != nil {
				dst.MessageMentionChannelResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.MessageMentionChannelResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MessageMentionChannelResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BasicMessageResponseMentionChannelsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BasicMessageResponseMentionChannelsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BasicMessageResponseMentionChannelsInner) MarshalJSON() ([]byte, error) {
	if src.MessageMentionChannelResponse != nil {
		return json.Marshal(&src.MessageMentionChannelResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BasicMessageResponseMentionChannelsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MessageMentionChannelResponse != nil {
		return obj.MessageMentionChannelResponse
	}

	// all schemas are nil
	return nil
}

type NullableBasicMessageResponseMentionChannelsInner struct {
	value *BasicMessageResponseMentionChannelsInner
	isSet bool
}

func (v NullableBasicMessageResponseMentionChannelsInner) Get() *BasicMessageResponseMentionChannelsInner {
	return v.value
}

func (v *NullableBasicMessageResponseMentionChannelsInner) Set(val *BasicMessageResponseMentionChannelsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicMessageResponseMentionChannelsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicMessageResponseMentionChannelsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicMessageResponseMentionChannelsInner(val *BasicMessageResponseMentionChannelsInner) *NullableBasicMessageResponseMentionChannelsInner {
	return &NullableBasicMessageResponseMentionChannelsInner{value: val, isSet: true}
}

func (v NullableBasicMessageResponseMentionChannelsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicMessageResponseMentionChannelsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


