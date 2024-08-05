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

// CreateGuildChannelRequestAvailableTagsInner - struct for CreateGuildChannelRequestAvailableTagsInner
type CreateGuildChannelRequestAvailableTagsInner struct {
	CreateOrUpdateThreadTagRequest *CreateOrUpdateThreadTagRequest
}

// CreateOrUpdateThreadTagRequestAsCreateGuildChannelRequestAvailableTagsInner is a convenience function that returns CreateOrUpdateThreadTagRequest wrapped in CreateGuildChannelRequestAvailableTagsInner
func CreateOrUpdateThreadTagRequestAsCreateGuildChannelRequestAvailableTagsInner(v *CreateOrUpdateThreadTagRequest) CreateGuildChannelRequestAvailableTagsInner {
	return CreateGuildChannelRequestAvailableTagsInner{
		CreateOrUpdateThreadTagRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateGuildChannelRequestAvailableTagsInner) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into CreateOrUpdateThreadTagRequest
	err = newStrictDecoder(data).Decode(&dst.CreateOrUpdateThreadTagRequest)
	if err == nil {
		jsonCreateOrUpdateThreadTagRequest, _ := json.Marshal(dst.CreateOrUpdateThreadTagRequest)
		if string(jsonCreateOrUpdateThreadTagRequest) == "{}" { // empty struct
			dst.CreateOrUpdateThreadTagRequest = nil
		} else {
			if err = validator.Validate(dst.CreateOrUpdateThreadTagRequest); err != nil {
				dst.CreateOrUpdateThreadTagRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.CreateOrUpdateThreadTagRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CreateOrUpdateThreadTagRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateGuildChannelRequestAvailableTagsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateGuildChannelRequestAvailableTagsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateGuildChannelRequestAvailableTagsInner) MarshalJSON() ([]byte, error) {
	if src.CreateOrUpdateThreadTagRequest != nil {
		return json.Marshal(&src.CreateOrUpdateThreadTagRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateGuildChannelRequestAvailableTagsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CreateOrUpdateThreadTagRequest != nil {
		return obj.CreateOrUpdateThreadTagRequest
	}

	// all schemas are nil
	return nil
}

type NullableCreateGuildChannelRequestAvailableTagsInner struct {
	value *CreateGuildChannelRequestAvailableTagsInner
	isSet bool
}

func (v NullableCreateGuildChannelRequestAvailableTagsInner) Get() *CreateGuildChannelRequestAvailableTagsInner {
	return v.value
}

func (v *NullableCreateGuildChannelRequestAvailableTagsInner) Set(val *CreateGuildChannelRequestAvailableTagsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGuildChannelRequestAvailableTagsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGuildChannelRequestAvailableTagsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGuildChannelRequestAvailableTagsInner(val *CreateGuildChannelRequestAvailableTagsInner) *NullableCreateGuildChannelRequestAvailableTagsInner {
	return &NullableCreateGuildChannelRequestAvailableTagsInner{value: val, isSet: true}
}

func (v NullableCreateGuildChannelRequestAvailableTagsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGuildChannelRequestAvailableTagsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

