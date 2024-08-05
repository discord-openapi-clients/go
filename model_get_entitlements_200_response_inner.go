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

// GetEntitlements200ResponseInner - struct for GetEntitlements200ResponseInner
type GetEntitlements200ResponseInner struct {
	EntitlementResponse *EntitlementResponse
}

// EntitlementResponseAsGetEntitlements200ResponseInner is a convenience function that returns EntitlementResponse wrapped in GetEntitlements200ResponseInner
func EntitlementResponseAsGetEntitlements200ResponseInner(v *EntitlementResponse) GetEntitlements200ResponseInner {
	return GetEntitlements200ResponseInner{
		EntitlementResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetEntitlements200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into EntitlementResponse
	err = newStrictDecoder(data).Decode(&dst.EntitlementResponse)
	if err == nil {
		jsonEntitlementResponse, _ := json.Marshal(dst.EntitlementResponse)
		if string(jsonEntitlementResponse) == "{}" { // empty struct
			dst.EntitlementResponse = nil
		} else {
			if err = validator.Validate(dst.EntitlementResponse); err != nil {
				dst.EntitlementResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.EntitlementResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.EntitlementResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetEntitlements200ResponseInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetEntitlements200ResponseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetEntitlements200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.EntitlementResponse != nil {
		return json.Marshal(&src.EntitlementResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetEntitlements200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.EntitlementResponse != nil {
		return obj.EntitlementResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetEntitlements200ResponseInner struct {
	value *GetEntitlements200ResponseInner
	isSet bool
}

func (v NullableGetEntitlements200ResponseInner) Get() *GetEntitlements200ResponseInner {
	return v.value
}

func (v *NullableGetEntitlements200ResponseInner) Set(val *GetEntitlements200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEntitlements200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEntitlements200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEntitlements200ResponseInner(val *GetEntitlements200ResponseInner) *NullableGetEntitlements200ResponseInner {
	return &NullableGetEntitlements200ResponseInner{value: val, isSet: true}
}

func (v NullableGetEntitlements200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEntitlements200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


