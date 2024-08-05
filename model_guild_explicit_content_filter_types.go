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

// GuildExplicitContentFilterTypes - struct for GuildExplicitContentFilterTypes
type GuildExplicitContentFilterTypes struct {
	Float32 *float32
}

// float32AsGuildExplicitContentFilterTypes is a convenience function that returns float32 wrapped in GuildExplicitContentFilterTypes
func Float32AsGuildExplicitContentFilterTypes(v *float32) GuildExplicitContentFilterTypes {
	return GuildExplicitContentFilterTypes{
		Float32: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GuildExplicitContentFilterTypes) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(GuildExplicitContentFilterTypes)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GuildExplicitContentFilterTypes)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GuildExplicitContentFilterTypes) MarshalJSON() ([]byte, error) {
	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GuildExplicitContentFilterTypes) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Float32 != nil {
		return obj.Float32
	}

	// all schemas are nil
	return nil
}

type NullableGuildExplicitContentFilterTypes struct {
	value *GuildExplicitContentFilterTypes
	isSet bool
}

func (v NullableGuildExplicitContentFilterTypes) Get() *GuildExplicitContentFilterTypes {
	return v.value
}

func (v *NullableGuildExplicitContentFilterTypes) Set(val *GuildExplicitContentFilterTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildExplicitContentFilterTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildExplicitContentFilterTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildExplicitContentFilterTypes(val *GuildExplicitContentFilterTypes) *NullableGuildExplicitContentFilterTypes {
	return &NullableGuildExplicitContentFilterTypes{value: val, isSet: true}
}

func (v NullableGuildExplicitContentFilterTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildExplicitContentFilterTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


