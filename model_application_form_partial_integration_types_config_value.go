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

// ApplicationFormPartialIntegrationTypesConfigValue - struct for ApplicationFormPartialIntegrationTypesConfigValue
type ApplicationFormPartialIntegrationTypesConfigValue struct {
	ApplicationIntegrationTypeConfiguration *ApplicationIntegrationTypeConfiguration
}

// ApplicationIntegrationTypeConfigurationAsApplicationFormPartialIntegrationTypesConfigValue is a convenience function that returns ApplicationIntegrationTypeConfiguration wrapped in ApplicationFormPartialIntegrationTypesConfigValue
func ApplicationIntegrationTypeConfigurationAsApplicationFormPartialIntegrationTypesConfigValue(v *ApplicationIntegrationTypeConfiguration) ApplicationFormPartialIntegrationTypesConfigValue {
	return ApplicationFormPartialIntegrationTypesConfigValue{
		ApplicationIntegrationTypeConfiguration: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApplicationFormPartialIntegrationTypesConfigValue) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into ApplicationIntegrationTypeConfiguration
	err = newStrictDecoder(data).Decode(&dst.ApplicationIntegrationTypeConfiguration)
	if err == nil {
		jsonApplicationIntegrationTypeConfiguration, _ := json.Marshal(dst.ApplicationIntegrationTypeConfiguration)
		if string(jsonApplicationIntegrationTypeConfiguration) == "{}" { // empty struct
			dst.ApplicationIntegrationTypeConfiguration = nil
		} else {
			if err = validator.Validate(dst.ApplicationIntegrationTypeConfiguration); err != nil {
				dst.ApplicationIntegrationTypeConfiguration = nil
			} else {
				match++
			}
		}
	} else {
		dst.ApplicationIntegrationTypeConfiguration = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApplicationIntegrationTypeConfiguration = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ApplicationFormPartialIntegrationTypesConfigValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ApplicationFormPartialIntegrationTypesConfigValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApplicationFormPartialIntegrationTypesConfigValue) MarshalJSON() ([]byte, error) {
	if src.ApplicationIntegrationTypeConfiguration != nil {
		return json.Marshal(&src.ApplicationIntegrationTypeConfiguration)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ApplicationFormPartialIntegrationTypesConfigValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ApplicationIntegrationTypeConfiguration != nil {
		return obj.ApplicationIntegrationTypeConfiguration
	}

	// all schemas are nil
	return nil
}

type NullableApplicationFormPartialIntegrationTypesConfigValue struct {
	value *ApplicationFormPartialIntegrationTypesConfigValue
	isSet bool
}

func (v NullableApplicationFormPartialIntegrationTypesConfigValue) Get() *ApplicationFormPartialIntegrationTypesConfigValue {
	return v.value
}

func (v *NullableApplicationFormPartialIntegrationTypesConfigValue) Set(val *ApplicationFormPartialIntegrationTypesConfigValue) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationFormPartialIntegrationTypesConfigValue) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationFormPartialIntegrationTypesConfigValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationFormPartialIntegrationTypesConfigValue(val *ApplicationFormPartialIntegrationTypesConfigValue) *NullableApplicationFormPartialIntegrationTypesConfigValue {
	return &NullableApplicationFormPartialIntegrationTypesConfigValue{value: val, isSet: true}
}

func (v NullableApplicationFormPartialIntegrationTypesConfigValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationFormPartialIntegrationTypesConfigValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


