/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discord

import (
	"encoding/json"
)

// checks if the WebhookSlackEmbedField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookSlackEmbedField{}

// WebhookSlackEmbedField struct for WebhookSlackEmbedField
type WebhookSlackEmbedField struct {
	Name NullableString `json:"name,omitempty"`
	Value NullableString `json:"value,omitempty"`
	Inline NullableBool `json:"inline,omitempty"`
}

// NewWebhookSlackEmbedField instantiates a new WebhookSlackEmbedField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookSlackEmbedField() *WebhookSlackEmbedField {
	this := WebhookSlackEmbedField{}
	return &this
}

// NewWebhookSlackEmbedFieldWithDefaults instantiates a new WebhookSlackEmbedField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookSlackEmbedFieldWithDefaults() *WebhookSlackEmbedField {
	this := WebhookSlackEmbedField{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSlackEmbedField) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSlackEmbedField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WebhookSlackEmbedField) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WebhookSlackEmbedField) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *WebhookSlackEmbedField) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WebhookSlackEmbedField) UnsetName() {
	o.Name.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSlackEmbedField) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSlackEmbedField) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *WebhookSlackEmbedField) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *WebhookSlackEmbedField) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *WebhookSlackEmbedField) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *WebhookSlackEmbedField) UnsetValue() {
	o.Value.Unset()
}

// GetInline returns the Inline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSlackEmbedField) GetInline() bool {
	if o == nil || IsNil(o.Inline.Get()) {
		var ret bool
		return ret
	}
	return *o.Inline.Get()
}

// GetInlineOk returns a tuple with the Inline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSlackEmbedField) GetInlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inline.Get(), o.Inline.IsSet()
}

// HasInline returns a boolean if a field has been set.
func (o *WebhookSlackEmbedField) HasInline() bool {
	if o != nil && o.Inline.IsSet() {
		return true
	}

	return false
}

// SetInline gets a reference to the given NullableBool and assigns it to the Inline field.
func (o *WebhookSlackEmbedField) SetInline(v bool) {
	o.Inline.Set(&v)
}
// SetInlineNil sets the value for Inline to be an explicit nil
func (o *WebhookSlackEmbedField) SetInlineNil() {
	o.Inline.Set(nil)
}

// UnsetInline ensures that no value is present for Inline, not even an explicit nil
func (o *WebhookSlackEmbedField) UnsetInline() {
	o.Inline.Unset()
}

func (o WebhookSlackEmbedField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookSlackEmbedField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.Inline.IsSet() {
		toSerialize["inline"] = o.Inline.Get()
	}
	return toSerialize, nil
}

type NullableWebhookSlackEmbedField struct {
	value *WebhookSlackEmbedField
	isSet bool
}

func (v NullableWebhookSlackEmbedField) Get() *WebhookSlackEmbedField {
	return v.value
}

func (v *NullableWebhookSlackEmbedField) Set(val *WebhookSlackEmbedField) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookSlackEmbedField) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookSlackEmbedField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookSlackEmbedField(val *WebhookSlackEmbedField) *NullableWebhookSlackEmbedField {
	return &NullableWebhookSlackEmbedField{value: val, isSet: true}
}

func (v NullableWebhookSlackEmbedField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookSlackEmbedField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


