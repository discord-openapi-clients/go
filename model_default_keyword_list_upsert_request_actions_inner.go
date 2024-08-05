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

// DefaultKeywordListUpsertRequestActionsInner - struct for DefaultKeywordListUpsertRequestActionsInner
type DefaultKeywordListUpsertRequestActionsInner struct {
	BlockMessageAction *BlockMessageAction
	FlagToChannelAction *FlagToChannelAction
	QuarantineUserAction *QuarantineUserAction
	UserCommunicationDisabledAction *UserCommunicationDisabledAction
}

// BlockMessageActionAsDefaultKeywordListUpsertRequestActionsInner is a convenience function that returns BlockMessageAction wrapped in DefaultKeywordListUpsertRequestActionsInner
func BlockMessageActionAsDefaultKeywordListUpsertRequestActionsInner(v *BlockMessageAction) DefaultKeywordListUpsertRequestActionsInner {
	return DefaultKeywordListUpsertRequestActionsInner{
		BlockMessageAction: v,
	}
}

// FlagToChannelActionAsDefaultKeywordListUpsertRequestActionsInner is a convenience function that returns FlagToChannelAction wrapped in DefaultKeywordListUpsertRequestActionsInner
func FlagToChannelActionAsDefaultKeywordListUpsertRequestActionsInner(v *FlagToChannelAction) DefaultKeywordListUpsertRequestActionsInner {
	return DefaultKeywordListUpsertRequestActionsInner{
		FlagToChannelAction: v,
	}
}

// QuarantineUserActionAsDefaultKeywordListUpsertRequestActionsInner is a convenience function that returns QuarantineUserAction wrapped in DefaultKeywordListUpsertRequestActionsInner
func QuarantineUserActionAsDefaultKeywordListUpsertRequestActionsInner(v *QuarantineUserAction) DefaultKeywordListUpsertRequestActionsInner {
	return DefaultKeywordListUpsertRequestActionsInner{
		QuarantineUserAction: v,
	}
}

// UserCommunicationDisabledActionAsDefaultKeywordListUpsertRequestActionsInner is a convenience function that returns UserCommunicationDisabledAction wrapped in DefaultKeywordListUpsertRequestActionsInner
func UserCommunicationDisabledActionAsDefaultKeywordListUpsertRequestActionsInner(v *UserCommunicationDisabledAction) DefaultKeywordListUpsertRequestActionsInner {
	return DefaultKeywordListUpsertRequestActionsInner{
		UserCommunicationDisabledAction: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DefaultKeywordListUpsertRequestActionsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BlockMessageAction
	err = newStrictDecoder(data).Decode(&dst.BlockMessageAction)
	if err == nil {
		jsonBlockMessageAction, _ := json.Marshal(dst.BlockMessageAction)
		if string(jsonBlockMessageAction) == "{}" { // empty struct
			dst.BlockMessageAction = nil
		} else {
			if err = validator.Validate(dst.BlockMessageAction); err != nil {
				dst.BlockMessageAction = nil
			} else {
				match++
			}
		}
	} else {
		dst.BlockMessageAction = nil
	}

	// try to unmarshal data into FlagToChannelAction
	err = newStrictDecoder(data).Decode(&dst.FlagToChannelAction)
	if err == nil {
		jsonFlagToChannelAction, _ := json.Marshal(dst.FlagToChannelAction)
		if string(jsonFlagToChannelAction) == "{}" { // empty struct
			dst.FlagToChannelAction = nil
		} else {
			if err = validator.Validate(dst.FlagToChannelAction); err != nil {
				dst.FlagToChannelAction = nil
			} else {
				match++
			}
		}
	} else {
		dst.FlagToChannelAction = nil
	}

	// try to unmarshal data into QuarantineUserAction
	err = newStrictDecoder(data).Decode(&dst.QuarantineUserAction)
	if err == nil {
		jsonQuarantineUserAction, _ := json.Marshal(dst.QuarantineUserAction)
		if string(jsonQuarantineUserAction) == "{}" { // empty struct
			dst.QuarantineUserAction = nil
		} else {
			if err = validator.Validate(dst.QuarantineUserAction); err != nil {
				dst.QuarantineUserAction = nil
			} else {
				match++
			}
		}
	} else {
		dst.QuarantineUserAction = nil
	}

	// try to unmarshal data into UserCommunicationDisabledAction
	err = newStrictDecoder(data).Decode(&dst.UserCommunicationDisabledAction)
	if err == nil {
		jsonUserCommunicationDisabledAction, _ := json.Marshal(dst.UserCommunicationDisabledAction)
		if string(jsonUserCommunicationDisabledAction) == "{}" { // empty struct
			dst.UserCommunicationDisabledAction = nil
		} else {
			if err = validator.Validate(dst.UserCommunicationDisabledAction); err != nil {
				dst.UserCommunicationDisabledAction = nil
			} else {
				match++
			}
		}
	} else {
		dst.UserCommunicationDisabledAction = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlockMessageAction = nil
		dst.FlagToChannelAction = nil
		dst.QuarantineUserAction = nil
		dst.UserCommunicationDisabledAction = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DefaultKeywordListUpsertRequestActionsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DefaultKeywordListUpsertRequestActionsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DefaultKeywordListUpsertRequestActionsInner) MarshalJSON() ([]byte, error) {
	if src.BlockMessageAction != nil {
		return json.Marshal(&src.BlockMessageAction)
	}

	if src.FlagToChannelAction != nil {
		return json.Marshal(&src.FlagToChannelAction)
	}

	if src.QuarantineUserAction != nil {
		return json.Marshal(&src.QuarantineUserAction)
	}

	if src.UserCommunicationDisabledAction != nil {
		return json.Marshal(&src.UserCommunicationDisabledAction)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DefaultKeywordListUpsertRequestActionsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlockMessageAction != nil {
		return obj.BlockMessageAction
	}

	if obj.FlagToChannelAction != nil {
		return obj.FlagToChannelAction
	}

	if obj.QuarantineUserAction != nil {
		return obj.QuarantineUserAction
	}

	if obj.UserCommunicationDisabledAction != nil {
		return obj.UserCommunicationDisabledAction
	}

	// all schemas are nil
	return nil
}

type NullableDefaultKeywordListUpsertRequestActionsInner struct {
	value *DefaultKeywordListUpsertRequestActionsInner
	isSet bool
}

func (v NullableDefaultKeywordListUpsertRequestActionsInner) Get() *DefaultKeywordListUpsertRequestActionsInner {
	return v.value
}

func (v *NullableDefaultKeywordListUpsertRequestActionsInner) Set(val *DefaultKeywordListUpsertRequestActionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultKeywordListUpsertRequestActionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultKeywordListUpsertRequestActionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultKeywordListUpsertRequestActionsInner(val *DefaultKeywordListUpsertRequestActionsInner) *NullableDefaultKeywordListUpsertRequestActionsInner {
	return &NullableDefaultKeywordListUpsertRequestActionsInner{value: val, isSet: true}
}

func (v NullableDefaultKeywordListUpsertRequestActionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultKeywordListUpsertRequestActionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


