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

// WidgetImageStyles - struct for WidgetImageStyles
type WidgetImageStyles struct {
	String *string
}

// stringAsWidgetImageStyles is a convenience function that returns string wrapped in WidgetImageStyles
func StringAsWidgetImageStyles(v *string) WidgetImageStyles {
	return WidgetImageStyles{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *WidgetImageStyles) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(WidgetImageStyles)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(WidgetImageStyles)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WidgetImageStyles) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WidgetImageStyles) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableWidgetImageStyles struct {
	value *WidgetImageStyles
	isSet bool
}

func (v NullableWidgetImageStyles) Get() *WidgetImageStyles {
	return v.value
}

func (v *NullableWidgetImageStyles) Set(val *WidgetImageStyles) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetImageStyles) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetImageStyles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetImageStyles(val *WidgetImageStyles) *NullableWidgetImageStyles {
	return &NullableWidgetImageStyles{value: val, isSet: true}
}

func (v NullableWidgetImageStyles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetImageStyles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


