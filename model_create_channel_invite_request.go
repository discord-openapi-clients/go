/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// CreateChannelInviteRequest struct for CreateChannelInviteRequest
type CreateChannelInviteRequest struct {
	CreateGroupDMInviteRequest *CreateGroupDMInviteRequest
	CreateGuildInviteRequest *CreateGuildInviteRequest
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CreateChannelInviteRequest) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CreateGroupDMInviteRequest
	err = json.Unmarshal(data, &dst.CreateGroupDMInviteRequest);
	if err == nil {
		jsonCreateGroupDMInviteRequest, _ := json.Marshal(dst.CreateGroupDMInviteRequest)
		if string(jsonCreateGroupDMInviteRequest) == "{}" { // empty struct
			dst.CreateGroupDMInviteRequest = nil
		} else {
			return nil // data stored in dst.CreateGroupDMInviteRequest, return on the first match
		}
	} else {
		dst.CreateGroupDMInviteRequest = nil
	}

	// try to unmarshal JSON data into CreateGuildInviteRequest
	err = json.Unmarshal(data, &dst.CreateGuildInviteRequest);
	if err == nil {
		jsonCreateGuildInviteRequest, _ := json.Marshal(dst.CreateGuildInviteRequest)
		if string(jsonCreateGuildInviteRequest) == "{}" { // empty struct
			dst.CreateGuildInviteRequest = nil
		} else {
			return nil // data stored in dst.CreateGuildInviteRequest, return on the first match
		}
	} else {
		dst.CreateGuildInviteRequest = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CreateChannelInviteRequest)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CreateChannelInviteRequest) MarshalJSON() ([]byte, error) {
	if src.CreateGroupDMInviteRequest != nil {
		return json.Marshal(&src.CreateGroupDMInviteRequest)
	}

	if src.CreateGuildInviteRequest != nil {
		return json.Marshal(&src.CreateGuildInviteRequest)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCreateChannelInviteRequest struct {
	value *CreateChannelInviteRequest
	isSet bool
}

func (v NullableCreateChannelInviteRequest) Get() *CreateChannelInviteRequest {
	return v.value
}

func (v *NullableCreateChannelInviteRequest) Set(val *CreateChannelInviteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateChannelInviteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateChannelInviteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateChannelInviteRequest(val *CreateChannelInviteRequest) *NullableCreateChannelInviteRequest {
	return &NullableCreateChannelInviteRequest{value: val, isSet: true}
}

func (v NullableCreateChannelInviteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateChannelInviteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


