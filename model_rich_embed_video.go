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

// checks if the RichEmbedVideo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RichEmbedVideo{}

// RichEmbedVideo struct for RichEmbedVideo
type RichEmbedVideo struct {
	Url NullableString `json:"url,omitempty"`
	Width NullableInt32 `json:"width,omitempty"`
	Height NullableInt32 `json:"height,omitempty"`
	Placeholder NullableString `json:"placeholder,omitempty"`
	PlaceholderVersion NullableInt32 `json:"placeholder_version,omitempty"`
}

// NewRichEmbedVideo instantiates a new RichEmbedVideo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRichEmbedVideo() *RichEmbedVideo {
	this := RichEmbedVideo{}
	return &this
}

// NewRichEmbedVideoWithDefaults instantiates a new RichEmbedVideo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRichEmbedVideoWithDefaults() *RichEmbedVideo {
	this := RichEmbedVideo{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RichEmbedVideo) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RichEmbedVideo) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *RichEmbedVideo) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *RichEmbedVideo) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *RichEmbedVideo) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *RichEmbedVideo) UnsetUrl() {
	o.Url.Unset()
}

// GetWidth returns the Width field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RichEmbedVideo) GetWidth() int32 {
	if o == nil || IsNil(o.Width.Get()) {
		var ret int32
		return ret
	}
	return *o.Width.Get()
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RichEmbedVideo) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Width.Get(), o.Width.IsSet()
}

// HasWidth returns a boolean if a field has been set.
func (o *RichEmbedVideo) HasWidth() bool {
	if o != nil && o.Width.IsSet() {
		return true
	}

	return false
}

// SetWidth gets a reference to the given NullableInt32 and assigns it to the Width field.
func (o *RichEmbedVideo) SetWidth(v int32) {
	o.Width.Set(&v)
}
// SetWidthNil sets the value for Width to be an explicit nil
func (o *RichEmbedVideo) SetWidthNil() {
	o.Width.Set(nil)
}

// UnsetWidth ensures that no value is present for Width, not even an explicit nil
func (o *RichEmbedVideo) UnsetWidth() {
	o.Width.Unset()
}

// GetHeight returns the Height field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RichEmbedVideo) GetHeight() int32 {
	if o == nil || IsNil(o.Height.Get()) {
		var ret int32
		return ret
	}
	return *o.Height.Get()
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RichEmbedVideo) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Height.Get(), o.Height.IsSet()
}

// HasHeight returns a boolean if a field has been set.
func (o *RichEmbedVideo) HasHeight() bool {
	if o != nil && o.Height.IsSet() {
		return true
	}

	return false
}

// SetHeight gets a reference to the given NullableInt32 and assigns it to the Height field.
func (o *RichEmbedVideo) SetHeight(v int32) {
	o.Height.Set(&v)
}
// SetHeightNil sets the value for Height to be an explicit nil
func (o *RichEmbedVideo) SetHeightNil() {
	o.Height.Set(nil)
}

// UnsetHeight ensures that no value is present for Height, not even an explicit nil
func (o *RichEmbedVideo) UnsetHeight() {
	o.Height.Unset()
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RichEmbedVideo) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder.Get()) {
		var ret string
		return ret
	}
	return *o.Placeholder.Get()
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RichEmbedVideo) GetPlaceholderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Placeholder.Get(), o.Placeholder.IsSet()
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *RichEmbedVideo) HasPlaceholder() bool {
	if o != nil && o.Placeholder.IsSet() {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given NullableString and assigns it to the Placeholder field.
func (o *RichEmbedVideo) SetPlaceholder(v string) {
	o.Placeholder.Set(&v)
}
// SetPlaceholderNil sets the value for Placeholder to be an explicit nil
func (o *RichEmbedVideo) SetPlaceholderNil() {
	o.Placeholder.Set(nil)
}

// UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
func (o *RichEmbedVideo) UnsetPlaceholder() {
	o.Placeholder.Unset()
}

// GetPlaceholderVersion returns the PlaceholderVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RichEmbedVideo) GetPlaceholderVersion() int32 {
	if o == nil || IsNil(o.PlaceholderVersion.Get()) {
		var ret int32
		return ret
	}
	return *o.PlaceholderVersion.Get()
}

// GetPlaceholderVersionOk returns a tuple with the PlaceholderVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RichEmbedVideo) GetPlaceholderVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlaceholderVersion.Get(), o.PlaceholderVersion.IsSet()
}

// HasPlaceholderVersion returns a boolean if a field has been set.
func (o *RichEmbedVideo) HasPlaceholderVersion() bool {
	if o != nil && o.PlaceholderVersion.IsSet() {
		return true
	}

	return false
}

// SetPlaceholderVersion gets a reference to the given NullableInt32 and assigns it to the PlaceholderVersion field.
func (o *RichEmbedVideo) SetPlaceholderVersion(v int32) {
	o.PlaceholderVersion.Set(&v)
}
// SetPlaceholderVersionNil sets the value for PlaceholderVersion to be an explicit nil
func (o *RichEmbedVideo) SetPlaceholderVersionNil() {
	o.PlaceholderVersion.Set(nil)
}

// UnsetPlaceholderVersion ensures that no value is present for PlaceholderVersion, not even an explicit nil
func (o *RichEmbedVideo) UnsetPlaceholderVersion() {
	o.PlaceholderVersion.Unset()
}

func (o RichEmbedVideo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RichEmbedVideo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Width.IsSet() {
		toSerialize["width"] = o.Width.Get()
	}
	if o.Height.IsSet() {
		toSerialize["height"] = o.Height.Get()
	}
	if o.Placeholder.IsSet() {
		toSerialize["placeholder"] = o.Placeholder.Get()
	}
	if o.PlaceholderVersion.IsSet() {
		toSerialize["placeholder_version"] = o.PlaceholderVersion.Get()
	}
	return toSerialize, nil
}

type NullableRichEmbedVideo struct {
	value *RichEmbedVideo
	isSet bool
}

func (v NullableRichEmbedVideo) Get() *RichEmbedVideo {
	return v.value
}

func (v *NullableRichEmbedVideo) Set(val *RichEmbedVideo) {
	v.value = val
	v.isSet = true
}

func (v NullableRichEmbedVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableRichEmbedVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRichEmbedVideo(val *RichEmbedVideo) *NullableRichEmbedVideo {
	return &NullableRichEmbedVideo{value: val, isSet: true}
}

func (v NullableRichEmbedVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRichEmbedVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


