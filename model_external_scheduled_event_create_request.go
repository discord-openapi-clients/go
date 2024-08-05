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
	"bytes"
	"fmt"
)

// checks if the ExternalScheduledEventCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalScheduledEventCreateRequest{}

// ExternalScheduledEventCreateRequest struct for ExternalScheduledEventCreateRequest
type ExternalScheduledEventCreateRequest struct {
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
	Image NullableString `json:"image,omitempty"`
	ScheduledStartTime time.Time `json:"scheduled_start_time"`
	ScheduledEndTime NullableTime `json:"scheduled_end_time,omitempty"`
	PrivacyLevel GuildScheduledEventPrivacyLevels `json:"privacy_level"`
	EntityType GuildScheduledEventEntityTypes `json:"entity_type"`
	ChannelId *string `json:"channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	EntityMetadata EntityMetadataExternal `json:"entity_metadata"`
}

type _ExternalScheduledEventCreateRequest ExternalScheduledEventCreateRequest

// NewExternalScheduledEventCreateRequest instantiates a new ExternalScheduledEventCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalScheduledEventCreateRequest(name string, scheduledStartTime time.Time, privacyLevel GuildScheduledEventPrivacyLevels, entityType GuildScheduledEventEntityTypes, entityMetadata EntityMetadataExternal) *ExternalScheduledEventCreateRequest {
	this := ExternalScheduledEventCreateRequest{}
	this.Name = name
	this.ScheduledStartTime = scheduledStartTime
	this.PrivacyLevel = privacyLevel
	this.EntityType = entityType
	this.EntityMetadata = entityMetadata
	return &this
}

// NewExternalScheduledEventCreateRequestWithDefaults instantiates a new ExternalScheduledEventCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalScheduledEventCreateRequestWithDefaults() *ExternalScheduledEventCreateRequest {
	this := ExternalScheduledEventCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ExternalScheduledEventCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExternalScheduledEventCreateRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalScheduledEventCreateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalScheduledEventCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExternalScheduledEventCreateRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ExternalScheduledEventCreateRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ExternalScheduledEventCreateRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ExternalScheduledEventCreateRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalScheduledEventCreateRequest) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalScheduledEventCreateRequest) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *ExternalScheduledEventCreateRequest) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *ExternalScheduledEventCreateRequest) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *ExternalScheduledEventCreateRequest) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *ExternalScheduledEventCreateRequest) UnsetImage() {
	o.Image.Unset()
}

// GetScheduledStartTime returns the ScheduledStartTime field value
func (o *ExternalScheduledEventCreateRequest) GetScheduledStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ScheduledStartTime
}

// GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventCreateRequest) GetScheduledStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledStartTime, true
}

// SetScheduledStartTime sets field value
func (o *ExternalScheduledEventCreateRequest) SetScheduledStartTime(v time.Time) {
	o.ScheduledStartTime = v
}

// GetScheduledEndTime returns the ScheduledEndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalScheduledEventCreateRequest) GetScheduledEndTime() time.Time {
	if o == nil || IsNil(o.ScheduledEndTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledEndTime.Get()
}

// GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalScheduledEventCreateRequest) GetScheduledEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledEndTime.Get(), o.ScheduledEndTime.IsSet()
}

// HasScheduledEndTime returns a boolean if a field has been set.
func (o *ExternalScheduledEventCreateRequest) HasScheduledEndTime() bool {
	if o != nil && o.ScheduledEndTime.IsSet() {
		return true
	}

	return false
}

// SetScheduledEndTime gets a reference to the given NullableTime and assigns it to the ScheduledEndTime field.
func (o *ExternalScheduledEventCreateRequest) SetScheduledEndTime(v time.Time) {
	o.ScheduledEndTime.Set(&v)
}
// SetScheduledEndTimeNil sets the value for ScheduledEndTime to be an explicit nil
func (o *ExternalScheduledEventCreateRequest) SetScheduledEndTimeNil() {
	o.ScheduledEndTime.Set(nil)
}

// UnsetScheduledEndTime ensures that no value is present for ScheduledEndTime, not even an explicit nil
func (o *ExternalScheduledEventCreateRequest) UnsetScheduledEndTime() {
	o.ScheduledEndTime.Unset()
}

// GetPrivacyLevel returns the PrivacyLevel field value
func (o *ExternalScheduledEventCreateRequest) GetPrivacyLevel() GuildScheduledEventPrivacyLevels {
	if o == nil {
		var ret GuildScheduledEventPrivacyLevels
		return ret
	}

	return o.PrivacyLevel
}

// GetPrivacyLevelOk returns a tuple with the PrivacyLevel field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventCreateRequest) GetPrivacyLevelOk() (*GuildScheduledEventPrivacyLevels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivacyLevel, true
}

// SetPrivacyLevel sets field value
func (o *ExternalScheduledEventCreateRequest) SetPrivacyLevel(v GuildScheduledEventPrivacyLevels) {
	o.PrivacyLevel = v
}

// GetEntityType returns the EntityType field value
func (o *ExternalScheduledEventCreateRequest) GetEntityType() GuildScheduledEventEntityTypes {
	if o == nil {
		var ret GuildScheduledEventEntityTypes
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventCreateRequest) GetEntityTypeOk() (*GuildScheduledEventEntityTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *ExternalScheduledEventCreateRequest) SetEntityType(v GuildScheduledEventEntityTypes) {
	o.EntityType = v
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *ExternalScheduledEventCreateRequest) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventCreateRequest) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *ExternalScheduledEventCreateRequest) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *ExternalScheduledEventCreateRequest) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetEntityMetadata returns the EntityMetadata field value
func (o *ExternalScheduledEventCreateRequest) GetEntityMetadata() EntityMetadataExternal {
	if o == nil {
		var ret EntityMetadataExternal
		return ret
	}

	return o.EntityMetadata
}

// GetEntityMetadataOk returns a tuple with the EntityMetadata field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventCreateRequest) GetEntityMetadataOk() (*EntityMetadataExternal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityMetadata, true
}

// SetEntityMetadata sets field value
func (o *ExternalScheduledEventCreateRequest) SetEntityMetadata(v EntityMetadataExternal) {
	o.EntityMetadata = v
}

func (o ExternalScheduledEventCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalScheduledEventCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	toSerialize["scheduled_start_time"] = o.ScheduledStartTime
	if o.ScheduledEndTime.IsSet() {
		toSerialize["scheduled_end_time"] = o.ScheduledEndTime.Get()
	}
	toSerialize["privacy_level"] = o.PrivacyLevel
	toSerialize["entity_type"] = o.EntityType
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	toSerialize["entity_metadata"] = o.EntityMetadata
	return toSerialize, nil
}

func (o *ExternalScheduledEventCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"scheduled_start_time",
		"privacy_level",
		"entity_type",
		"entity_metadata",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varExternalScheduledEventCreateRequest := _ExternalScheduledEventCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalScheduledEventCreateRequest)

	if err != nil {
		return err
	}

	*o = ExternalScheduledEventCreateRequest(varExternalScheduledEventCreateRequest)

	return err
}

type NullableExternalScheduledEventCreateRequest struct {
	value *ExternalScheduledEventCreateRequest
	isSet bool
}

func (v NullableExternalScheduledEventCreateRequest) Get() *ExternalScheduledEventCreateRequest {
	return v.value
}

func (v *NullableExternalScheduledEventCreateRequest) Set(val *ExternalScheduledEventCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalScheduledEventCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalScheduledEventCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalScheduledEventCreateRequest(val *ExternalScheduledEventCreateRequest) *NullableExternalScheduledEventCreateRequest {
	return &NullableExternalScheduledEventCreateRequest{value: val, isSet: true}
}

func (v NullableExternalScheduledEventCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalScheduledEventCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

