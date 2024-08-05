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

// OAuth2Scopes - struct for OAuth2Scopes
type OAuth2Scopes struct {
	String *string
}

// stringAsOAuth2Scopes is a convenience function that returns string wrapped in OAuth2Scopes
func StringAsOAuth2Scopes(v *string) OAuth2Scopes {
	return OAuth2Scopes{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *OAuth2Scopes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(OAuth2Scopes)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(OAuth2Scopes)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OAuth2Scopes) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OAuth2Scopes) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableOAuth2Scopes struct {
	value *OAuth2Scopes
	isSet bool
}

func (v NullableOAuth2Scopes) Get() *OAuth2Scopes {
	return v.value
}

func (v *NullableOAuth2Scopes) Set(val *OAuth2Scopes) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2Scopes) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2Scopes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2Scopes(val *OAuth2Scopes) *NullableOAuth2Scopes {
	return &NullableOAuth2Scopes{value: val, isSet: true}
}

func (v NullableOAuth2Scopes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2Scopes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

