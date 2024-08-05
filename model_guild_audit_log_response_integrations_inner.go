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

// GuildAuditLogResponseIntegrationsInner - struct for GuildAuditLogResponseIntegrationsInner
type GuildAuditLogResponseIntegrationsInner struct {
	PartialDiscordIntegrationResponse *PartialDiscordIntegrationResponse
	PartialExternalConnectionIntegrationResponse *PartialExternalConnectionIntegrationResponse
	PartialGuildSubscriptionIntegrationResponse *PartialGuildSubscriptionIntegrationResponse
}

// PartialDiscordIntegrationResponseAsGuildAuditLogResponseIntegrationsInner is a convenience function that returns PartialDiscordIntegrationResponse wrapped in GuildAuditLogResponseIntegrationsInner
func PartialDiscordIntegrationResponseAsGuildAuditLogResponseIntegrationsInner(v *PartialDiscordIntegrationResponse) GuildAuditLogResponseIntegrationsInner {
	return GuildAuditLogResponseIntegrationsInner{
		PartialDiscordIntegrationResponse: v,
	}
}

// PartialExternalConnectionIntegrationResponseAsGuildAuditLogResponseIntegrationsInner is a convenience function that returns PartialExternalConnectionIntegrationResponse wrapped in GuildAuditLogResponseIntegrationsInner
func PartialExternalConnectionIntegrationResponseAsGuildAuditLogResponseIntegrationsInner(v *PartialExternalConnectionIntegrationResponse) GuildAuditLogResponseIntegrationsInner {
	return GuildAuditLogResponseIntegrationsInner{
		PartialExternalConnectionIntegrationResponse: v,
	}
}

// PartialGuildSubscriptionIntegrationResponseAsGuildAuditLogResponseIntegrationsInner is a convenience function that returns PartialGuildSubscriptionIntegrationResponse wrapped in GuildAuditLogResponseIntegrationsInner
func PartialGuildSubscriptionIntegrationResponseAsGuildAuditLogResponseIntegrationsInner(v *PartialGuildSubscriptionIntegrationResponse) GuildAuditLogResponseIntegrationsInner {
	return GuildAuditLogResponseIntegrationsInner{
		PartialGuildSubscriptionIntegrationResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GuildAuditLogResponseIntegrationsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PartialDiscordIntegrationResponse
	err = newStrictDecoder(data).Decode(&dst.PartialDiscordIntegrationResponse)
	if err == nil {
		jsonPartialDiscordIntegrationResponse, _ := json.Marshal(dst.PartialDiscordIntegrationResponse)
		if string(jsonPartialDiscordIntegrationResponse) == "{}" { // empty struct
			dst.PartialDiscordIntegrationResponse = nil
		} else {
			if err = validator.Validate(dst.PartialDiscordIntegrationResponse); err != nil {
				dst.PartialDiscordIntegrationResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.PartialDiscordIntegrationResponse = nil
	}

	// try to unmarshal data into PartialExternalConnectionIntegrationResponse
	err = newStrictDecoder(data).Decode(&dst.PartialExternalConnectionIntegrationResponse)
	if err == nil {
		jsonPartialExternalConnectionIntegrationResponse, _ := json.Marshal(dst.PartialExternalConnectionIntegrationResponse)
		if string(jsonPartialExternalConnectionIntegrationResponse) == "{}" { // empty struct
			dst.PartialExternalConnectionIntegrationResponse = nil
		} else {
			if err = validator.Validate(dst.PartialExternalConnectionIntegrationResponse); err != nil {
				dst.PartialExternalConnectionIntegrationResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.PartialExternalConnectionIntegrationResponse = nil
	}

	// try to unmarshal data into PartialGuildSubscriptionIntegrationResponse
	err = newStrictDecoder(data).Decode(&dst.PartialGuildSubscriptionIntegrationResponse)
	if err == nil {
		jsonPartialGuildSubscriptionIntegrationResponse, _ := json.Marshal(dst.PartialGuildSubscriptionIntegrationResponse)
		if string(jsonPartialGuildSubscriptionIntegrationResponse) == "{}" { // empty struct
			dst.PartialGuildSubscriptionIntegrationResponse = nil
		} else {
			if err = validator.Validate(dst.PartialGuildSubscriptionIntegrationResponse); err != nil {
				dst.PartialGuildSubscriptionIntegrationResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.PartialGuildSubscriptionIntegrationResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PartialDiscordIntegrationResponse = nil
		dst.PartialExternalConnectionIntegrationResponse = nil
		dst.PartialGuildSubscriptionIntegrationResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GuildAuditLogResponseIntegrationsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GuildAuditLogResponseIntegrationsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GuildAuditLogResponseIntegrationsInner) MarshalJSON() ([]byte, error) {
	if src.PartialDiscordIntegrationResponse != nil {
		return json.Marshal(&src.PartialDiscordIntegrationResponse)
	}

	if src.PartialExternalConnectionIntegrationResponse != nil {
		return json.Marshal(&src.PartialExternalConnectionIntegrationResponse)
	}

	if src.PartialGuildSubscriptionIntegrationResponse != nil {
		return json.Marshal(&src.PartialGuildSubscriptionIntegrationResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GuildAuditLogResponseIntegrationsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PartialDiscordIntegrationResponse != nil {
		return obj.PartialDiscordIntegrationResponse
	}

	if obj.PartialExternalConnectionIntegrationResponse != nil {
		return obj.PartialExternalConnectionIntegrationResponse
	}

	if obj.PartialGuildSubscriptionIntegrationResponse != nil {
		return obj.PartialGuildSubscriptionIntegrationResponse
	}

	// all schemas are nil
	return nil
}

type NullableGuildAuditLogResponseIntegrationsInner struct {
	value *GuildAuditLogResponseIntegrationsInner
	isSet bool
}

func (v NullableGuildAuditLogResponseIntegrationsInner) Get() *GuildAuditLogResponseIntegrationsInner {
	return v.value
}

func (v *NullableGuildAuditLogResponseIntegrationsInner) Set(val *GuildAuditLogResponseIntegrationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildAuditLogResponseIntegrationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildAuditLogResponseIntegrationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildAuditLogResponseIntegrationsInner(val *GuildAuditLogResponseIntegrationsInner) *NullableGuildAuditLogResponseIntegrationsInner {
	return &NullableGuildAuditLogResponseIntegrationsInner{value: val, isSet: true}
}

func (v NullableGuildAuditLogResponseIntegrationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildAuditLogResponseIntegrationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


