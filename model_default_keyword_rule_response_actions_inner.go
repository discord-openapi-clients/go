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

// DefaultKeywordRuleResponseActionsInner - struct for DefaultKeywordRuleResponseActionsInner
type DefaultKeywordRuleResponseActionsInner struct {
	BlockMessageActionResponse *BlockMessageActionResponse
	FlagToChannelActionResponse *FlagToChannelActionResponse
	QuarantineUserActionResponse *QuarantineUserActionResponse
	UserCommunicationDisabledActionResponse *UserCommunicationDisabledActionResponse
}

// BlockMessageActionResponseAsDefaultKeywordRuleResponseActionsInner is a convenience function that returns BlockMessageActionResponse wrapped in DefaultKeywordRuleResponseActionsInner
func BlockMessageActionResponseAsDefaultKeywordRuleResponseActionsInner(v *BlockMessageActionResponse) DefaultKeywordRuleResponseActionsInner {
	return DefaultKeywordRuleResponseActionsInner{
		BlockMessageActionResponse: v,
	}
}

// FlagToChannelActionResponseAsDefaultKeywordRuleResponseActionsInner is a convenience function that returns FlagToChannelActionResponse wrapped in DefaultKeywordRuleResponseActionsInner
func FlagToChannelActionResponseAsDefaultKeywordRuleResponseActionsInner(v *FlagToChannelActionResponse) DefaultKeywordRuleResponseActionsInner {
	return DefaultKeywordRuleResponseActionsInner{
		FlagToChannelActionResponse: v,
	}
}

// QuarantineUserActionResponseAsDefaultKeywordRuleResponseActionsInner is a convenience function that returns QuarantineUserActionResponse wrapped in DefaultKeywordRuleResponseActionsInner
func QuarantineUserActionResponseAsDefaultKeywordRuleResponseActionsInner(v *QuarantineUserActionResponse) DefaultKeywordRuleResponseActionsInner {
	return DefaultKeywordRuleResponseActionsInner{
		QuarantineUserActionResponse: v,
	}
}

// UserCommunicationDisabledActionResponseAsDefaultKeywordRuleResponseActionsInner is a convenience function that returns UserCommunicationDisabledActionResponse wrapped in DefaultKeywordRuleResponseActionsInner
func UserCommunicationDisabledActionResponseAsDefaultKeywordRuleResponseActionsInner(v *UserCommunicationDisabledActionResponse) DefaultKeywordRuleResponseActionsInner {
	return DefaultKeywordRuleResponseActionsInner{
		UserCommunicationDisabledActionResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DefaultKeywordRuleResponseActionsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BlockMessageActionResponse
	err = newStrictDecoder(data).Decode(&dst.BlockMessageActionResponse)
	if err == nil {
		jsonBlockMessageActionResponse, _ := json.Marshal(dst.BlockMessageActionResponse)
		if string(jsonBlockMessageActionResponse) == "{}" { // empty struct
			dst.BlockMessageActionResponse = nil
		} else {
			if err = validator.Validate(dst.BlockMessageActionResponse); err != nil {
				dst.BlockMessageActionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.BlockMessageActionResponse = nil
	}

	// try to unmarshal data into FlagToChannelActionResponse
	err = newStrictDecoder(data).Decode(&dst.FlagToChannelActionResponse)
	if err == nil {
		jsonFlagToChannelActionResponse, _ := json.Marshal(dst.FlagToChannelActionResponse)
		if string(jsonFlagToChannelActionResponse) == "{}" { // empty struct
			dst.FlagToChannelActionResponse = nil
		} else {
			if err = validator.Validate(dst.FlagToChannelActionResponse); err != nil {
				dst.FlagToChannelActionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.FlagToChannelActionResponse = nil
	}

	// try to unmarshal data into QuarantineUserActionResponse
	err = newStrictDecoder(data).Decode(&dst.QuarantineUserActionResponse)
	if err == nil {
		jsonQuarantineUserActionResponse, _ := json.Marshal(dst.QuarantineUserActionResponse)
		if string(jsonQuarantineUserActionResponse) == "{}" { // empty struct
			dst.QuarantineUserActionResponse = nil
		} else {
			if err = validator.Validate(dst.QuarantineUserActionResponse); err != nil {
				dst.QuarantineUserActionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.QuarantineUserActionResponse = nil
	}

	// try to unmarshal data into UserCommunicationDisabledActionResponse
	err = newStrictDecoder(data).Decode(&dst.UserCommunicationDisabledActionResponse)
	if err == nil {
		jsonUserCommunicationDisabledActionResponse, _ := json.Marshal(dst.UserCommunicationDisabledActionResponse)
		if string(jsonUserCommunicationDisabledActionResponse) == "{}" { // empty struct
			dst.UserCommunicationDisabledActionResponse = nil
		} else {
			if err = validator.Validate(dst.UserCommunicationDisabledActionResponse); err != nil {
				dst.UserCommunicationDisabledActionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.UserCommunicationDisabledActionResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlockMessageActionResponse = nil
		dst.FlagToChannelActionResponse = nil
		dst.QuarantineUserActionResponse = nil
		dst.UserCommunicationDisabledActionResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DefaultKeywordRuleResponseActionsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DefaultKeywordRuleResponseActionsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DefaultKeywordRuleResponseActionsInner) MarshalJSON() ([]byte, error) {
	if src.BlockMessageActionResponse != nil {
		return json.Marshal(&src.BlockMessageActionResponse)
	}

	if src.FlagToChannelActionResponse != nil {
		return json.Marshal(&src.FlagToChannelActionResponse)
	}

	if src.QuarantineUserActionResponse != nil {
		return json.Marshal(&src.QuarantineUserActionResponse)
	}

	if src.UserCommunicationDisabledActionResponse != nil {
		return json.Marshal(&src.UserCommunicationDisabledActionResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DefaultKeywordRuleResponseActionsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlockMessageActionResponse != nil {
		return obj.BlockMessageActionResponse
	}

	if obj.FlagToChannelActionResponse != nil {
		return obj.FlagToChannelActionResponse
	}

	if obj.QuarantineUserActionResponse != nil {
		return obj.QuarantineUserActionResponse
	}

	if obj.UserCommunicationDisabledActionResponse != nil {
		return obj.UserCommunicationDisabledActionResponse
	}

	// all schemas are nil
	return nil
}

type NullableDefaultKeywordRuleResponseActionsInner struct {
	value *DefaultKeywordRuleResponseActionsInner
	isSet bool
}

func (v NullableDefaultKeywordRuleResponseActionsInner) Get() *DefaultKeywordRuleResponseActionsInner {
	return v.value
}

func (v *NullableDefaultKeywordRuleResponseActionsInner) Set(val *DefaultKeywordRuleResponseActionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultKeywordRuleResponseActionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultKeywordRuleResponseActionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultKeywordRuleResponseActionsInner(val *DefaultKeywordRuleResponseActionsInner) *NullableDefaultKeywordRuleResponseActionsInner {
	return &NullableDefaultKeywordRuleResponseActionsInner{value: val, isSet: true}
}

func (v NullableDefaultKeywordRuleResponseActionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultKeywordRuleResponseActionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


