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

// ListAutoModerationRules200ResponseInner - struct for ListAutoModerationRules200ResponseInner
type ListAutoModerationRules200ResponseInner struct {
	DefaultKeywordRuleResponse *DefaultKeywordRuleResponse
	KeywordRuleResponse *KeywordRuleResponse
	MLSpamRuleResponse *MLSpamRuleResponse
	MentionSpamRuleResponse *MentionSpamRuleResponse
	SpamLinkRuleResponse *SpamLinkRuleResponse
}

// DefaultKeywordRuleResponseAsListAutoModerationRules200ResponseInner is a convenience function that returns DefaultKeywordRuleResponse wrapped in ListAutoModerationRules200ResponseInner
func DefaultKeywordRuleResponseAsListAutoModerationRules200ResponseInner(v *DefaultKeywordRuleResponse) ListAutoModerationRules200ResponseInner {
	return ListAutoModerationRules200ResponseInner{
		DefaultKeywordRuleResponse: v,
	}
}

// KeywordRuleResponseAsListAutoModerationRules200ResponseInner is a convenience function that returns KeywordRuleResponse wrapped in ListAutoModerationRules200ResponseInner
func KeywordRuleResponseAsListAutoModerationRules200ResponseInner(v *KeywordRuleResponse) ListAutoModerationRules200ResponseInner {
	return ListAutoModerationRules200ResponseInner{
		KeywordRuleResponse: v,
	}
}

// MLSpamRuleResponseAsListAutoModerationRules200ResponseInner is a convenience function that returns MLSpamRuleResponse wrapped in ListAutoModerationRules200ResponseInner
func MLSpamRuleResponseAsListAutoModerationRules200ResponseInner(v *MLSpamRuleResponse) ListAutoModerationRules200ResponseInner {
	return ListAutoModerationRules200ResponseInner{
		MLSpamRuleResponse: v,
	}
}

// MentionSpamRuleResponseAsListAutoModerationRules200ResponseInner is a convenience function that returns MentionSpamRuleResponse wrapped in ListAutoModerationRules200ResponseInner
func MentionSpamRuleResponseAsListAutoModerationRules200ResponseInner(v *MentionSpamRuleResponse) ListAutoModerationRules200ResponseInner {
	return ListAutoModerationRules200ResponseInner{
		MentionSpamRuleResponse: v,
	}
}

// SpamLinkRuleResponseAsListAutoModerationRules200ResponseInner is a convenience function that returns SpamLinkRuleResponse wrapped in ListAutoModerationRules200ResponseInner
func SpamLinkRuleResponseAsListAutoModerationRules200ResponseInner(v *SpamLinkRuleResponse) ListAutoModerationRules200ResponseInner {
	return ListAutoModerationRules200ResponseInner{
		SpamLinkRuleResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListAutoModerationRules200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into DefaultKeywordRuleResponse
	err = newStrictDecoder(data).Decode(&dst.DefaultKeywordRuleResponse)
	if err == nil {
		jsonDefaultKeywordRuleResponse, _ := json.Marshal(dst.DefaultKeywordRuleResponse)
		if string(jsonDefaultKeywordRuleResponse) == "{}" { // empty struct
			dst.DefaultKeywordRuleResponse = nil
		} else {
			if err = validator.Validate(dst.DefaultKeywordRuleResponse); err != nil {
				dst.DefaultKeywordRuleResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.DefaultKeywordRuleResponse = nil
	}

	// try to unmarshal data into KeywordRuleResponse
	err = newStrictDecoder(data).Decode(&dst.KeywordRuleResponse)
	if err == nil {
		jsonKeywordRuleResponse, _ := json.Marshal(dst.KeywordRuleResponse)
		if string(jsonKeywordRuleResponse) == "{}" { // empty struct
			dst.KeywordRuleResponse = nil
		} else {
			if err = validator.Validate(dst.KeywordRuleResponse); err != nil {
				dst.KeywordRuleResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.KeywordRuleResponse = nil
	}

	// try to unmarshal data into MLSpamRuleResponse
	err = newStrictDecoder(data).Decode(&dst.MLSpamRuleResponse)
	if err == nil {
		jsonMLSpamRuleResponse, _ := json.Marshal(dst.MLSpamRuleResponse)
		if string(jsonMLSpamRuleResponse) == "{}" { // empty struct
			dst.MLSpamRuleResponse = nil
		} else {
			if err = validator.Validate(dst.MLSpamRuleResponse); err != nil {
				dst.MLSpamRuleResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.MLSpamRuleResponse = nil
	}

	// try to unmarshal data into MentionSpamRuleResponse
	err = newStrictDecoder(data).Decode(&dst.MentionSpamRuleResponse)
	if err == nil {
		jsonMentionSpamRuleResponse, _ := json.Marshal(dst.MentionSpamRuleResponse)
		if string(jsonMentionSpamRuleResponse) == "{}" { // empty struct
			dst.MentionSpamRuleResponse = nil
		} else {
			if err = validator.Validate(dst.MentionSpamRuleResponse); err != nil {
				dst.MentionSpamRuleResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.MentionSpamRuleResponse = nil
	}

	// try to unmarshal data into SpamLinkRuleResponse
	err = newStrictDecoder(data).Decode(&dst.SpamLinkRuleResponse)
	if err == nil {
		jsonSpamLinkRuleResponse, _ := json.Marshal(dst.SpamLinkRuleResponse)
		if string(jsonSpamLinkRuleResponse) == "{}" { // empty struct
			dst.SpamLinkRuleResponse = nil
		} else {
			if err = validator.Validate(dst.SpamLinkRuleResponse); err != nil {
				dst.SpamLinkRuleResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.SpamLinkRuleResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DefaultKeywordRuleResponse = nil
		dst.KeywordRuleResponse = nil
		dst.MLSpamRuleResponse = nil
		dst.MentionSpamRuleResponse = nil
		dst.SpamLinkRuleResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListAutoModerationRules200ResponseInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListAutoModerationRules200ResponseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListAutoModerationRules200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.DefaultKeywordRuleResponse != nil {
		return json.Marshal(&src.DefaultKeywordRuleResponse)
	}

	if src.KeywordRuleResponse != nil {
		return json.Marshal(&src.KeywordRuleResponse)
	}

	if src.MLSpamRuleResponse != nil {
		return json.Marshal(&src.MLSpamRuleResponse)
	}

	if src.MentionSpamRuleResponse != nil {
		return json.Marshal(&src.MentionSpamRuleResponse)
	}

	if src.SpamLinkRuleResponse != nil {
		return json.Marshal(&src.SpamLinkRuleResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListAutoModerationRules200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DefaultKeywordRuleResponse != nil {
		return obj.DefaultKeywordRuleResponse
	}

	if obj.KeywordRuleResponse != nil {
		return obj.KeywordRuleResponse
	}

	if obj.MLSpamRuleResponse != nil {
		return obj.MLSpamRuleResponse
	}

	if obj.MentionSpamRuleResponse != nil {
		return obj.MentionSpamRuleResponse
	}

	if obj.SpamLinkRuleResponse != nil {
		return obj.SpamLinkRuleResponse
	}

	// all schemas are nil
	return nil
}

type NullableListAutoModerationRules200ResponseInner struct {
	value *ListAutoModerationRules200ResponseInner
	isSet bool
}

func (v NullableListAutoModerationRules200ResponseInner) Get() *ListAutoModerationRules200ResponseInner {
	return v.value
}

func (v *NullableListAutoModerationRules200ResponseInner) Set(val *ListAutoModerationRules200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListAutoModerationRules200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListAutoModerationRules200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAutoModerationRules200ResponseInner(val *ListAutoModerationRules200ResponseInner) *NullableListAutoModerationRules200ResponseInner {
	return &NullableListAutoModerationRules200ResponseInner{value: val, isSet: true}
}

func (v NullableListAutoModerationRules200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAutoModerationRules200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


