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

// GetSticker200Response - struct for GetSticker200Response
type GetSticker200Response struct {
	GuildStickerResponse *GuildStickerResponse
	StandardStickerResponse *StandardStickerResponse
}

// GuildStickerResponseAsGetSticker200Response is a convenience function that returns GuildStickerResponse wrapped in GetSticker200Response
func GuildStickerResponseAsGetSticker200Response(v *GuildStickerResponse) GetSticker200Response {
	return GetSticker200Response{
		GuildStickerResponse: v,
	}
}

// StandardStickerResponseAsGetSticker200Response is a convenience function that returns StandardStickerResponse wrapped in GetSticker200Response
func StandardStickerResponseAsGetSticker200Response(v *StandardStickerResponse) GetSticker200Response {
	return GetSticker200Response{
		StandardStickerResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetSticker200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GuildStickerResponse
	err = newStrictDecoder(data).Decode(&dst.GuildStickerResponse)
	if err == nil {
		jsonGuildStickerResponse, _ := json.Marshal(dst.GuildStickerResponse)
		if string(jsonGuildStickerResponse) == "{}" { // empty struct
			dst.GuildStickerResponse = nil
		} else {
			if err = validator.Validate(dst.GuildStickerResponse); err != nil {
				dst.GuildStickerResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.GuildStickerResponse = nil
	}

	// try to unmarshal data into StandardStickerResponse
	err = newStrictDecoder(data).Decode(&dst.StandardStickerResponse)
	if err == nil {
		jsonStandardStickerResponse, _ := json.Marshal(dst.StandardStickerResponse)
		if string(jsonStandardStickerResponse) == "{}" { // empty struct
			dst.StandardStickerResponse = nil
		} else {
			if err = validator.Validate(dst.StandardStickerResponse); err != nil {
				dst.StandardStickerResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.StandardStickerResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GuildStickerResponse = nil
		dst.StandardStickerResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetSticker200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetSticker200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetSticker200Response) MarshalJSON() ([]byte, error) {
	if src.GuildStickerResponse != nil {
		return json.Marshal(&src.GuildStickerResponse)
	}

	if src.StandardStickerResponse != nil {
		return json.Marshal(&src.StandardStickerResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetSticker200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GuildStickerResponse != nil {
		return obj.GuildStickerResponse
	}

	if obj.StandardStickerResponse != nil {
		return obj.StandardStickerResponse
	}

	// all schemas are nil
	return nil
}

type NullableGetSticker200Response struct {
	value *GetSticker200Response
	isSet bool
}

func (v NullableGetSticker200Response) Get() *GetSticker200Response {
	return v.value
}

func (v *NullableGetSticker200Response) Set(val *GetSticker200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSticker200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSticker200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSticker200Response(val *GetSticker200Response) *NullableGetSticker200Response {
	return &NullableGetSticker200Response{value: val, isSet: true}
}

func (v NullableGetSticker200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSticker200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


