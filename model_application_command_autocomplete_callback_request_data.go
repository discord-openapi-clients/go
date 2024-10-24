/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discord

import (
	"encoding/json"
	"fmt"
)


// ApplicationCommandAutocompleteCallbackRequestData struct for ApplicationCommandAutocompleteCallbackRequestData
type ApplicationCommandAutocompleteCallbackRequestData struct {
	InteractionApplicationCommandAutocompleteCallbackIntegerData *InteractionApplicationCommandAutocompleteCallbackIntegerData
	InteractionApplicationCommandAutocompleteCallbackNumberData *InteractionApplicationCommandAutocompleteCallbackNumberData
	InteractionApplicationCommandAutocompleteCallbackStringData *InteractionApplicationCommandAutocompleteCallbackStringData
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ApplicationCommandAutocompleteCallbackRequestData) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into InteractionApplicationCommandAutocompleteCallbackIntegerData
	err = json.Unmarshal(data, &dst.InteractionApplicationCommandAutocompleteCallbackIntegerData);
	if err == nil {
		jsonInteractionApplicationCommandAutocompleteCallbackIntegerData, _ := json.Marshal(dst.InteractionApplicationCommandAutocompleteCallbackIntegerData)
		if string(jsonInteractionApplicationCommandAutocompleteCallbackIntegerData) == "{}" { // empty struct
			dst.InteractionApplicationCommandAutocompleteCallbackIntegerData = nil
		} else {
			return nil // data stored in dst.InteractionApplicationCommandAutocompleteCallbackIntegerData, return on the first match
		}
	} else {
		dst.InteractionApplicationCommandAutocompleteCallbackIntegerData = nil
	}

	// try to unmarshal JSON data into InteractionApplicationCommandAutocompleteCallbackNumberData
	err = json.Unmarshal(data, &dst.InteractionApplicationCommandAutocompleteCallbackNumberData);
	if err == nil {
		jsonInteractionApplicationCommandAutocompleteCallbackNumberData, _ := json.Marshal(dst.InteractionApplicationCommandAutocompleteCallbackNumberData)
		if string(jsonInteractionApplicationCommandAutocompleteCallbackNumberData) == "{}" { // empty struct
			dst.InteractionApplicationCommandAutocompleteCallbackNumberData = nil
		} else {
			return nil // data stored in dst.InteractionApplicationCommandAutocompleteCallbackNumberData, return on the first match
		}
	} else {
		dst.InteractionApplicationCommandAutocompleteCallbackNumberData = nil
	}

	// try to unmarshal JSON data into InteractionApplicationCommandAutocompleteCallbackStringData
	err = json.Unmarshal(data, &dst.InteractionApplicationCommandAutocompleteCallbackStringData);
	if err == nil {
		jsonInteractionApplicationCommandAutocompleteCallbackStringData, _ := json.Marshal(dst.InteractionApplicationCommandAutocompleteCallbackStringData)
		if string(jsonInteractionApplicationCommandAutocompleteCallbackStringData) == "{}" { // empty struct
			dst.InteractionApplicationCommandAutocompleteCallbackStringData = nil
		} else {
			return nil // data stored in dst.InteractionApplicationCommandAutocompleteCallbackStringData, return on the first match
		}
	} else {
		dst.InteractionApplicationCommandAutocompleteCallbackStringData = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ApplicationCommandAutocompleteCallbackRequestData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ApplicationCommandAutocompleteCallbackRequestData) MarshalJSON() ([]byte, error) {
	if src.InteractionApplicationCommandAutocompleteCallbackIntegerData != nil {
		return json.Marshal(&src.InteractionApplicationCommandAutocompleteCallbackIntegerData)
	}

	if src.InteractionApplicationCommandAutocompleteCallbackNumberData != nil {
		return json.Marshal(&src.InteractionApplicationCommandAutocompleteCallbackNumberData)
	}

	if src.InteractionApplicationCommandAutocompleteCallbackStringData != nil {
		return json.Marshal(&src.InteractionApplicationCommandAutocompleteCallbackStringData)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableApplicationCommandAutocompleteCallbackRequestData struct {
	value *ApplicationCommandAutocompleteCallbackRequestData
	isSet bool
}

func (v NullableApplicationCommandAutocompleteCallbackRequestData) Get() *ApplicationCommandAutocompleteCallbackRequestData {
	return v.value
}

func (v *NullableApplicationCommandAutocompleteCallbackRequestData) Set(val *ApplicationCommandAutocompleteCallbackRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCommandAutocompleteCallbackRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCommandAutocompleteCallbackRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCommandAutocompleteCallbackRequestData(val *ApplicationCommandAutocompleteCallbackRequestData) *NullableApplicationCommandAutocompleteCallbackRequestData {
	return &NullableApplicationCommandAutocompleteCallbackRequestData{value: val, isSet: true}
}

func (v NullableApplicationCommandAutocompleteCallbackRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCommandAutocompleteCallbackRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


