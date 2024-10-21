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


// UpdateChannelRequest struct for UpdateChannelRequest
type UpdateChannelRequest struct {
	PrivateChannelRequestPartial *PrivateChannelRequestPartial
	UpdateGuildChannelRequestPartial *UpdateGuildChannelRequestPartial
	UpdateThreadRequestPartial *UpdateThreadRequestPartial
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UpdateChannelRequest) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PrivateChannelRequestPartial
	err = json.Unmarshal(data, &dst.PrivateChannelRequestPartial);
	if err == nil {
		jsonPrivateChannelRequestPartial, _ := json.Marshal(dst.PrivateChannelRequestPartial)
		if string(jsonPrivateChannelRequestPartial) == "{}" { // empty struct
			dst.PrivateChannelRequestPartial = nil
		} else {
			return nil // data stored in dst.PrivateChannelRequestPartial, return on the first match
		}
	} else {
		dst.PrivateChannelRequestPartial = nil
	}

	// try to unmarshal JSON data into UpdateGuildChannelRequestPartial
	err = json.Unmarshal(data, &dst.UpdateGuildChannelRequestPartial);
	if err == nil {
		jsonUpdateGuildChannelRequestPartial, _ := json.Marshal(dst.UpdateGuildChannelRequestPartial)
		if string(jsonUpdateGuildChannelRequestPartial) == "{}" { // empty struct
			dst.UpdateGuildChannelRequestPartial = nil
		} else {
			return nil // data stored in dst.UpdateGuildChannelRequestPartial, return on the first match
		}
	} else {
		dst.UpdateGuildChannelRequestPartial = nil
	}

	// try to unmarshal JSON data into UpdateThreadRequestPartial
	err = json.Unmarshal(data, &dst.UpdateThreadRequestPartial);
	if err == nil {
		jsonUpdateThreadRequestPartial, _ := json.Marshal(dst.UpdateThreadRequestPartial)
		if string(jsonUpdateThreadRequestPartial) == "{}" { // empty struct
			dst.UpdateThreadRequestPartial = nil
		} else {
			return nil // data stored in dst.UpdateThreadRequestPartial, return on the first match
		}
	} else {
		dst.UpdateThreadRequestPartial = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(UpdateChannelRequest)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UpdateChannelRequest) MarshalJSON() ([]byte, error) {
	if src.PrivateChannelRequestPartial != nil {
		return json.Marshal(&src.PrivateChannelRequestPartial)
	}

	if src.UpdateGuildChannelRequestPartial != nil {
		return json.Marshal(&src.UpdateGuildChannelRequestPartial)
	}

	if src.UpdateThreadRequestPartial != nil {
		return json.Marshal(&src.UpdateThreadRequestPartial)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableUpdateChannelRequest struct {
	value *UpdateChannelRequest
	isSet bool
}

func (v NullableUpdateChannelRequest) Get() *UpdateChannelRequest {
	return v.value
}

func (v *NullableUpdateChannelRequest) Set(val *UpdateChannelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateChannelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateChannelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateChannelRequest(val *UpdateChannelRequest) *NullableUpdateChannelRequest {
	return &NullableUpdateChannelRequest{value: val, isSet: true}
}

func (v NullableUpdateChannelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateChannelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


