/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the ExternalScheduledEventPatchRequestPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalScheduledEventPatchRequestPartial{}

// ExternalScheduledEventPatchRequestPartial struct for ExternalScheduledEventPatchRequestPartial
type ExternalScheduledEventPatchRequestPartial struct {
	Status NullableGuildScheduledEventStatuses `json:"status,omitempty"`
	Name *string `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Image NullableString `json:"image,omitempty"`
	ScheduledStartTime *time.Time `json:"scheduled_start_time,omitempty"`
	ScheduledEndTime NullableTime `json:"scheduled_end_time,omitempty"`
	EntityType NullableGuildScheduledEventEntityTypes `json:"entity_type,omitempty"`
	PrivacyLevel *GuildScheduledEventPrivacyLevels `json:"privacy_level,omitempty"`
	ChannelId *string `json:"channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	EntityMetadata *EntityMetadataExternal `json:"entity_metadata,omitempty"`
}

// NewExternalScheduledEventPatchRequestPartial instantiates a new ExternalScheduledEventPatchRequestPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalScheduledEventPatchRequestPartial() *ExternalScheduledEventPatchRequestPartial {
	this := ExternalScheduledEventPatchRequestPartial{}
	return &this
}

// NewExternalScheduledEventPatchRequestPartialWithDefaults instantiates a new ExternalScheduledEventPatchRequestPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalScheduledEventPatchRequestPartialWithDefaults() *ExternalScheduledEventPatchRequestPartial {
	this := ExternalScheduledEventPatchRequestPartial{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalScheduledEventPatchRequestPartial) GetStatus() GuildScheduledEventStatuses {
	if o == nil || IsNil(o.Status.Get()) {
		var ret GuildScheduledEventStatuses
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalScheduledEventPatchRequestPartial) GetStatusOk() (*GuildScheduledEventStatuses, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableGuildScheduledEventStatuses and assigns it to the Status field.
func (o *ExternalScheduledEventPatchRequestPartial) SetStatus(v GuildScheduledEventStatuses) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *ExternalScheduledEventPatchRequestPartial) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ExternalScheduledEventPatchRequestPartial) UnsetStatus() {
	o.Status.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExternalScheduledEventPatchRequestPartial) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventPatchRequestPartial) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExternalScheduledEventPatchRequestPartial) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalScheduledEventPatchRequestPartial) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalScheduledEventPatchRequestPartial) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ExternalScheduledEventPatchRequestPartial) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ExternalScheduledEventPatchRequestPartial) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ExternalScheduledEventPatchRequestPartial) UnsetDescription() {
	o.Description.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalScheduledEventPatchRequestPartial) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalScheduledEventPatchRequestPartial) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *ExternalScheduledEventPatchRequestPartial) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *ExternalScheduledEventPatchRequestPartial) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *ExternalScheduledEventPatchRequestPartial) UnsetImage() {
	o.Image.Unset()
}

// GetScheduledStartTime returns the ScheduledStartTime field value if set, zero value otherwise.
func (o *ExternalScheduledEventPatchRequestPartial) GetScheduledStartTime() time.Time {
	if o == nil || IsNil(o.ScheduledStartTime) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledStartTime
}

// GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventPatchRequestPartial) GetScheduledStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduledStartTime) {
		return nil, false
	}
	return o.ScheduledStartTime, true
}

// HasScheduledStartTime returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasScheduledStartTime() bool {
	if o != nil && !IsNil(o.ScheduledStartTime) {
		return true
	}

	return false
}

// SetScheduledStartTime gets a reference to the given time.Time and assigns it to the ScheduledStartTime field.
func (o *ExternalScheduledEventPatchRequestPartial) SetScheduledStartTime(v time.Time) {
	o.ScheduledStartTime = &v
}

// GetScheduledEndTime returns the ScheduledEndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalScheduledEventPatchRequestPartial) GetScheduledEndTime() time.Time {
	if o == nil || IsNil(o.ScheduledEndTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledEndTime.Get()
}

// GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalScheduledEventPatchRequestPartial) GetScheduledEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledEndTime.Get(), o.ScheduledEndTime.IsSet()
}

// HasScheduledEndTime returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasScheduledEndTime() bool {
	if o != nil && o.ScheduledEndTime.IsSet() {
		return true
	}

	return false
}

// SetScheduledEndTime gets a reference to the given NullableTime and assigns it to the ScheduledEndTime field.
func (o *ExternalScheduledEventPatchRequestPartial) SetScheduledEndTime(v time.Time) {
	o.ScheduledEndTime.Set(&v)
}
// SetScheduledEndTimeNil sets the value for ScheduledEndTime to be an explicit nil
func (o *ExternalScheduledEventPatchRequestPartial) SetScheduledEndTimeNil() {
	o.ScheduledEndTime.Set(nil)
}

// UnsetScheduledEndTime ensures that no value is present for ScheduledEndTime, not even an explicit nil
func (o *ExternalScheduledEventPatchRequestPartial) UnsetScheduledEndTime() {
	o.ScheduledEndTime.Unset()
}

// GetEntityType returns the EntityType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalScheduledEventPatchRequestPartial) GetEntityType() GuildScheduledEventEntityTypes {
	if o == nil || IsNil(o.EntityType.Get()) {
		var ret GuildScheduledEventEntityTypes
		return ret
	}
	return *o.EntityType.Get()
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalScheduledEventPatchRequestPartial) GetEntityTypeOk() (*GuildScheduledEventEntityTypes, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntityType.Get(), o.EntityType.IsSet()
}

// HasEntityType returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasEntityType() bool {
	if o != nil && o.EntityType.IsSet() {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given NullableGuildScheduledEventEntityTypes and assigns it to the EntityType field.
func (o *ExternalScheduledEventPatchRequestPartial) SetEntityType(v GuildScheduledEventEntityTypes) {
	o.EntityType.Set(&v)
}
// SetEntityTypeNil sets the value for EntityType to be an explicit nil
func (o *ExternalScheduledEventPatchRequestPartial) SetEntityTypeNil() {
	o.EntityType.Set(nil)
}

// UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
func (o *ExternalScheduledEventPatchRequestPartial) UnsetEntityType() {
	o.EntityType.Unset()
}

// GetPrivacyLevel returns the PrivacyLevel field value if set, zero value otherwise.
func (o *ExternalScheduledEventPatchRequestPartial) GetPrivacyLevel() GuildScheduledEventPrivacyLevels {
	if o == nil || IsNil(o.PrivacyLevel) {
		var ret GuildScheduledEventPrivacyLevels
		return ret
	}
	return *o.PrivacyLevel
}

// GetPrivacyLevelOk returns a tuple with the PrivacyLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventPatchRequestPartial) GetPrivacyLevelOk() (*GuildScheduledEventPrivacyLevels, bool) {
	if o == nil || IsNil(o.PrivacyLevel) {
		return nil, false
	}
	return o.PrivacyLevel, true
}

// HasPrivacyLevel returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasPrivacyLevel() bool {
	if o != nil && !IsNil(o.PrivacyLevel) {
		return true
	}

	return false
}

// SetPrivacyLevel gets a reference to the given GuildScheduledEventPrivacyLevels and assigns it to the PrivacyLevel field.
func (o *ExternalScheduledEventPatchRequestPartial) SetPrivacyLevel(v GuildScheduledEventPrivacyLevels) {
	o.PrivacyLevel = &v
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *ExternalScheduledEventPatchRequestPartial) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventPatchRequestPartial) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *ExternalScheduledEventPatchRequestPartial) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetEntityMetadata returns the EntityMetadata field value if set, zero value otherwise.
func (o *ExternalScheduledEventPatchRequestPartial) GetEntityMetadata() EntityMetadataExternal {
	if o == nil || IsNil(o.EntityMetadata) {
		var ret EntityMetadataExternal
		return ret
	}
	return *o.EntityMetadata
}

// GetEntityMetadataOk returns a tuple with the EntityMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventPatchRequestPartial) GetEntityMetadataOk() (*EntityMetadataExternal, bool) {
	if o == nil || IsNil(o.EntityMetadata) {
		return nil, false
	}
	return o.EntityMetadata, true
}

// HasEntityMetadata returns a boolean if a field has been set.
func (o *ExternalScheduledEventPatchRequestPartial) HasEntityMetadata() bool {
	if o != nil && !IsNil(o.EntityMetadata) {
		return true
	}

	return false
}

// SetEntityMetadata gets a reference to the given EntityMetadataExternal and assigns it to the EntityMetadata field.
func (o *ExternalScheduledEventPatchRequestPartial) SetEntityMetadata(v EntityMetadataExternal) {
	o.EntityMetadata = &v
}

func (o ExternalScheduledEventPatchRequestPartial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalScheduledEventPatchRequestPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	if !IsNil(o.ScheduledStartTime) {
		toSerialize["scheduled_start_time"] = o.ScheduledStartTime
	}
	if o.ScheduledEndTime.IsSet() {
		toSerialize["scheduled_end_time"] = o.ScheduledEndTime.Get()
	}
	if o.EntityType.IsSet() {
		toSerialize["entity_type"] = o.EntityType.Get()
	}
	if !IsNil(o.PrivacyLevel) {
		toSerialize["privacy_level"] = o.PrivacyLevel
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.EntityMetadata) {
		toSerialize["entity_metadata"] = o.EntityMetadata
	}
	return toSerialize, nil
}

type NullableExternalScheduledEventPatchRequestPartial struct {
	value *ExternalScheduledEventPatchRequestPartial
	isSet bool
}

func (v NullableExternalScheduledEventPatchRequestPartial) Get() *ExternalScheduledEventPatchRequestPartial {
	return v.value
}

func (v *NullableExternalScheduledEventPatchRequestPartial) Set(val *ExternalScheduledEventPatchRequestPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalScheduledEventPatchRequestPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalScheduledEventPatchRequestPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalScheduledEventPatchRequestPartial(val *ExternalScheduledEventPatchRequestPartial) *NullableExternalScheduledEventPatchRequestPartial {
	return &NullableExternalScheduledEventPatchRequestPartial{value: val, isSet: true}
}

func (v NullableExternalScheduledEventPatchRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalScheduledEventPatchRequestPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


