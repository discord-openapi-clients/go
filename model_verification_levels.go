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

// VerificationLevels - struct for VerificationLevels
type VerificationLevels struct {
	Float32 *float32
}

// float32AsVerificationLevels is a convenience function that returns float32 wrapped in VerificationLevels
func Float32AsVerificationLevels(v *float32) VerificationLevels {
	return VerificationLevels{
		Float32: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *VerificationLevels) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(VerificationLevels)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(VerificationLevels)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VerificationLevels) MarshalJSON() ([]byte, error) {
	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VerificationLevels) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Float32 != nil {
		return obj.Float32
	}

	// all schemas are nil
	return nil
}

type NullableVerificationLevels struct {
	value *VerificationLevels
	isSet bool
}

func (v NullableVerificationLevels) Get() *VerificationLevels {
	return v.value
}

func (v *NullableVerificationLevels) Set(val *VerificationLevels) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationLevels) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationLevels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationLevels(val *VerificationLevels) *NullableVerificationLevels {
	return &NullableVerificationLevels{value: val, isSet: true}
}

func (v NullableVerificationLevels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationLevels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


