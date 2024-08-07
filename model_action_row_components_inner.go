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

// ActionRowComponentsInner - struct for ActionRowComponentsInner
type ActionRowComponentsInner struct {
	Button *Button
	ChannelSelect *ChannelSelect
	InputText *InputText
	MentionableSelect *MentionableSelect
	RoleSelect *RoleSelect
	StringSelect *StringSelect
	UserSelect *UserSelect
}

// ButtonAsActionRowComponentsInner is a convenience function that returns Button wrapped in ActionRowComponentsInner
func ButtonAsActionRowComponentsInner(v *Button) ActionRowComponentsInner {
	return ActionRowComponentsInner{
		Button: v,
	}
}

// ChannelSelectAsActionRowComponentsInner is a convenience function that returns ChannelSelect wrapped in ActionRowComponentsInner
func ChannelSelectAsActionRowComponentsInner(v *ChannelSelect) ActionRowComponentsInner {
	return ActionRowComponentsInner{
		ChannelSelect: v,
	}
}

// InputTextAsActionRowComponentsInner is a convenience function that returns InputText wrapped in ActionRowComponentsInner
func InputTextAsActionRowComponentsInner(v *InputText) ActionRowComponentsInner {
	return ActionRowComponentsInner{
		InputText: v,
	}
}

// MentionableSelectAsActionRowComponentsInner is a convenience function that returns MentionableSelect wrapped in ActionRowComponentsInner
func MentionableSelectAsActionRowComponentsInner(v *MentionableSelect) ActionRowComponentsInner {
	return ActionRowComponentsInner{
		MentionableSelect: v,
	}
}

// RoleSelectAsActionRowComponentsInner is a convenience function that returns RoleSelect wrapped in ActionRowComponentsInner
func RoleSelectAsActionRowComponentsInner(v *RoleSelect) ActionRowComponentsInner {
	return ActionRowComponentsInner{
		RoleSelect: v,
	}
}

// StringSelectAsActionRowComponentsInner is a convenience function that returns StringSelect wrapped in ActionRowComponentsInner
func StringSelectAsActionRowComponentsInner(v *StringSelect) ActionRowComponentsInner {
	return ActionRowComponentsInner{
		StringSelect: v,
	}
}

// UserSelectAsActionRowComponentsInner is a convenience function that returns UserSelect wrapped in ActionRowComponentsInner
func UserSelectAsActionRowComponentsInner(v *UserSelect) ActionRowComponentsInner {
	return ActionRowComponentsInner{
		UserSelect: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ActionRowComponentsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Button
	err = newStrictDecoder(data).Decode(&dst.Button)
	if err == nil {
		jsonButton, _ := json.Marshal(dst.Button)
		if string(jsonButton) == "{}" { // empty struct
			dst.Button = nil
		} else {
			if err = validator.Validate(dst.Button); err != nil {
				dst.Button = nil
			} else {
				match++
			}
		}
	} else {
		dst.Button = nil
	}

	// try to unmarshal data into ChannelSelect
	err = newStrictDecoder(data).Decode(&dst.ChannelSelect)
	if err == nil {
		jsonChannelSelect, _ := json.Marshal(dst.ChannelSelect)
		if string(jsonChannelSelect) == "{}" { // empty struct
			dst.ChannelSelect = nil
		} else {
			if err = validator.Validate(dst.ChannelSelect); err != nil {
				dst.ChannelSelect = nil
			} else {
				match++
			}
		}
	} else {
		dst.ChannelSelect = nil
	}

	// try to unmarshal data into InputText
	err = newStrictDecoder(data).Decode(&dst.InputText)
	if err == nil {
		jsonInputText, _ := json.Marshal(dst.InputText)
		if string(jsonInputText) == "{}" { // empty struct
			dst.InputText = nil
		} else {
			if err = validator.Validate(dst.InputText); err != nil {
				dst.InputText = nil
			} else {
				match++
			}
		}
	} else {
		dst.InputText = nil
	}

	// try to unmarshal data into MentionableSelect
	err = newStrictDecoder(data).Decode(&dst.MentionableSelect)
	if err == nil {
		jsonMentionableSelect, _ := json.Marshal(dst.MentionableSelect)
		if string(jsonMentionableSelect) == "{}" { // empty struct
			dst.MentionableSelect = nil
		} else {
			if err = validator.Validate(dst.MentionableSelect); err != nil {
				dst.MentionableSelect = nil
			} else {
				match++
			}
		}
	} else {
		dst.MentionableSelect = nil
	}

	// try to unmarshal data into RoleSelect
	err = newStrictDecoder(data).Decode(&dst.RoleSelect)
	if err == nil {
		jsonRoleSelect, _ := json.Marshal(dst.RoleSelect)
		if string(jsonRoleSelect) == "{}" { // empty struct
			dst.RoleSelect = nil
		} else {
			if err = validator.Validate(dst.RoleSelect); err != nil {
				dst.RoleSelect = nil
			} else {
				match++
			}
		}
	} else {
		dst.RoleSelect = nil
	}

	// try to unmarshal data into StringSelect
	err = newStrictDecoder(data).Decode(&dst.StringSelect)
	if err == nil {
		jsonStringSelect, _ := json.Marshal(dst.StringSelect)
		if string(jsonStringSelect) == "{}" { // empty struct
			dst.StringSelect = nil
		} else {
			if err = validator.Validate(dst.StringSelect); err != nil {
				dst.StringSelect = nil
			} else {
				match++
			}
		}
	} else {
		dst.StringSelect = nil
	}

	// try to unmarshal data into UserSelect
	err = newStrictDecoder(data).Decode(&dst.UserSelect)
	if err == nil {
		jsonUserSelect, _ := json.Marshal(dst.UserSelect)
		if string(jsonUserSelect) == "{}" { // empty struct
			dst.UserSelect = nil
		} else {
			if err = validator.Validate(dst.UserSelect); err != nil {
				dst.UserSelect = nil
			} else {
				match++
			}
		}
	} else {
		dst.UserSelect = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Button = nil
		dst.ChannelSelect = nil
		dst.InputText = nil
		dst.MentionableSelect = nil
		dst.RoleSelect = nil
		dst.StringSelect = nil
		dst.UserSelect = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ActionRowComponentsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ActionRowComponentsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ActionRowComponentsInner) MarshalJSON() ([]byte, error) {
	if src.Button != nil {
		return json.Marshal(&src.Button)
	}

	if src.ChannelSelect != nil {
		return json.Marshal(&src.ChannelSelect)
	}

	if src.InputText != nil {
		return json.Marshal(&src.InputText)
	}

	if src.MentionableSelect != nil {
		return json.Marshal(&src.MentionableSelect)
	}

	if src.RoleSelect != nil {
		return json.Marshal(&src.RoleSelect)
	}

	if src.StringSelect != nil {
		return json.Marshal(&src.StringSelect)
	}

	if src.UserSelect != nil {
		return json.Marshal(&src.UserSelect)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ActionRowComponentsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Button != nil {
		return obj.Button
	}

	if obj.ChannelSelect != nil {
		return obj.ChannelSelect
	}

	if obj.InputText != nil {
		return obj.InputText
	}

	if obj.MentionableSelect != nil {
		return obj.MentionableSelect
	}

	if obj.RoleSelect != nil {
		return obj.RoleSelect
	}

	if obj.StringSelect != nil {
		return obj.StringSelect
	}

	if obj.UserSelect != nil {
		return obj.UserSelect
	}

	// all schemas are nil
	return nil
}

type NullableActionRowComponentsInner struct {
	value *ActionRowComponentsInner
	isSet bool
}

func (v NullableActionRowComponentsInner) Get() *ActionRowComponentsInner {
	return v.value
}

func (v *NullableActionRowComponentsInner) Set(val *ActionRowComponentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableActionRowComponentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableActionRowComponentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionRowComponentsInner(val *ActionRowComponentsInner) *NullableActionRowComponentsInner {
	return &NullableActionRowComponentsInner{value: val, isSet: true}
}

func (v NullableActionRowComponentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionRowComponentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


