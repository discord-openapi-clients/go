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

// InteractionCallbackTypes - struct for InteractionCallbackTypes
type InteractionCallbackTypes struct {
	Float32 *float32
}

// float32AsInteractionCallbackTypes is a convenience function that returns float32 wrapped in InteractionCallbackTypes
func Float32AsInteractionCallbackTypes(v *float32) InteractionCallbackTypes {
	return InteractionCallbackTypes{
		Float32: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *InteractionCallbackTypes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Float32
	err = newStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			if err = validator.Validate(dst.Float32); err != nil {
				dst.Float32 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Float32 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Float32 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(InteractionCallbackTypes)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(InteractionCallbackTypes)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InteractionCallbackTypes) MarshalJSON() ([]byte, error) {
	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InteractionCallbackTypes) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Float32 != nil {
		return obj.Float32
	}

	// all schemas are nil
	return nil
}

type NullableInteractionCallbackTypes struct {
	value *InteractionCallbackTypes
	isSet bool
}

func (v NullableInteractionCallbackTypes) Get() *InteractionCallbackTypes {
	return v.value
}

func (v *NullableInteractionCallbackTypes) Set(val *InteractionCallbackTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableInteractionCallbackTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableInteractionCallbackTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInteractionCallbackTypes(val *InteractionCallbackTypes) *NullableInteractionCallbackTypes {
	return &NullableInteractionCallbackTypes{value: val, isSet: true}
}

func (v NullableInteractionCallbackTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInteractionCallbackTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


