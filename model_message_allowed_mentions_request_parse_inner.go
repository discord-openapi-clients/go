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

// MessageAllowedMentionsRequestParseInner - struct for MessageAllowedMentionsRequestParseInner
type MessageAllowedMentionsRequestParseInner struct {
	AllowedMentionTypes *AllowedMentionTypes
}

// AllowedMentionTypesAsMessageAllowedMentionsRequestParseInner is a convenience function that returns AllowedMentionTypes wrapped in MessageAllowedMentionsRequestParseInner
func AllowedMentionTypesAsMessageAllowedMentionsRequestParseInner(v *AllowedMentionTypes) MessageAllowedMentionsRequestParseInner {
	return MessageAllowedMentionsRequestParseInner{
		AllowedMentionTypes: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MessageAllowedMentionsRequestParseInner) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into AllowedMentionTypes
	err = newStrictDecoder(data).Decode(&dst.AllowedMentionTypes)
	if err == nil {
		jsonAllowedMentionTypes, _ := json.Marshal(dst.AllowedMentionTypes)
		if string(jsonAllowedMentionTypes) == "{}" { // empty struct
			dst.AllowedMentionTypes = nil
		} else {
			if err = validator.Validate(dst.AllowedMentionTypes); err != nil {
				dst.AllowedMentionTypes = nil
			} else {
				match++
			}
		}
	} else {
		dst.AllowedMentionTypes = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AllowedMentionTypes = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MessageAllowedMentionsRequestParseInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MessageAllowedMentionsRequestParseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MessageAllowedMentionsRequestParseInner) MarshalJSON() ([]byte, error) {
	if src.AllowedMentionTypes != nil {
		return json.Marshal(&src.AllowedMentionTypes)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MessageAllowedMentionsRequestParseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AllowedMentionTypes != nil {
		return obj.AllowedMentionTypes
	}

	// all schemas are nil
	return nil
}

type NullableMessageAllowedMentionsRequestParseInner struct {
	value *MessageAllowedMentionsRequestParseInner
	isSet bool
}

func (v NullableMessageAllowedMentionsRequestParseInner) Get() *MessageAllowedMentionsRequestParseInner {
	return v.value
}

func (v *NullableMessageAllowedMentionsRequestParseInner) Set(val *MessageAllowedMentionsRequestParseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageAllowedMentionsRequestParseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageAllowedMentionsRequestParseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageAllowedMentionsRequestParseInner(val *MessageAllowedMentionsRequestParseInner) *NullableMessageAllowedMentionsRequestParseInner {
	return &NullableMessageAllowedMentionsRequestParseInner{value: val, isSet: true}
}

func (v NullableMessageAllowedMentionsRequestParseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageAllowedMentionsRequestParseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


