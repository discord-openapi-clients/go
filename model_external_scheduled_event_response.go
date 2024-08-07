/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discord

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ExternalScheduledEventResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalScheduledEventResponse{}

// ExternalScheduledEventResponse struct for ExternalScheduledEventResponse
type ExternalScheduledEventResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	GuildId string `json:"guild_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
	ChannelId *string `json:"channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	CreatorId *string `json:"creator_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Creator NullableUserResponse `json:"creator,omitempty"`
	Image NullableString `json:"image,omitempty"`
	ScheduledStartTime time.Time `json:"scheduled_start_time"`
	ScheduledEndTime NullableTime `json:"scheduled_end_time,omitempty"`
	Status GuildScheduledEventStatuses `json:"status"`
	EntityType GuildScheduledEventEntityTypes `json:"entity_type"`
	EntityId *string `json:"entity_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	UserCount NullableInt32 `json:"user_count,omitempty"`
	PrivacyLevel GuildScheduledEventPrivacyLevels `json:"privacy_level"`
	UserRsvp NullableScheduledEventUserResponse `json:"user_rsvp,omitempty"`
	EntityMetadata EntityMetadataExternalResponse `json:"entity_metadata"`
}

type _ExternalScheduledEventResponse ExternalScheduledEventResponse

// NewExternalScheduledEventResponse instantiates a new ExternalScheduledEventResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalScheduledEventResponse(id string, guildId string, name string, scheduledStartTime time.Time, status GuildScheduledEventStatuses, entityType GuildScheduledEventEntityTypes, privacyLevel GuildScheduledEventPrivacyLevels, entityMetadata EntityMetadataExternalResponse) *ExternalScheduledEventResponse {
	this := ExternalScheduledEventResponse{}
	this.Id = id
	this.GuildId = guildId
	this.Name = name
	this.ScheduledStartTime = scheduledStartTime
	this.Status = status
	this.EntityType = entityType
	this.PrivacyLevel = privacyLevel
	this.EntityMetadata = entityMetadata
	return &this
}

// NewExternalScheduledEventResponseWithDefaults instantiates a new ExternalScheduledEventResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalScheduledEventResponseWithDefaults() *ExternalScheduledEventResponse {
	this := ExternalScheduledEventResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ExternalScheduledEventResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalScheduledEventResponse) SetId(v string) {
	o.Id = v
}

// GetGuildId returns the GuildId field value
func (o *ExternalScheduledEventResponse) GetGuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetGuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildId, true
}

// SetGuildId sets field value
func (o *ExternalScheduledEventResponse) SetGuildId(v string) {
	o.GuildId = v
}

// GetName returns the Name field value
func (o *ExternalScheduledEventResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExternalScheduledEventResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalScheduledEventResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalScheduledEventResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExternalScheduledEventResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ExternalScheduledEventResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ExternalScheduledEventResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ExternalScheduledEventResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *ExternalScheduledEventResponse) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *ExternalScheduledEventResponse) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *ExternalScheduledEventResponse) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *ExternalScheduledEventResponse) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *ExternalScheduledEventResponse) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *ExternalScheduledEventResponse) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalScheduledEventResponse) GetCreator() UserResponse {
	if o == nil || IsNil(o.Creator.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.Creator.Get()
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalScheduledEventResponse) GetCreatorOk() (*UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Creator.Get(), o.Creator.IsSet()
}

// HasCreator returns a boolean if a field has been set.
func (o *ExternalScheduledEventResponse) HasCreator() bool {
	if o != nil && o.Creator.IsSet() {
		return true
	}

	return false
}

// SetCreator gets a reference to the given NullableUserResponse and assigns it to the Creator field.
func (o *ExternalScheduledEventResponse) SetCreator(v UserResponse) {
	o.Creator.Set(&v)
}
// SetCreatorNil sets the value for Creator to be an explicit nil
func (o *ExternalScheduledEventResponse) SetCreatorNil() {
	o.Creator.Set(nil)
}

// UnsetCreator ensures that no value is present for Creator, not even an explicit nil
func (o *ExternalScheduledEventResponse) UnsetCreator() {
	o.Creator.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalScheduledEventResponse) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalScheduledEventResponse) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *ExternalScheduledEventResponse) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *ExternalScheduledEventResponse) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *ExternalScheduledEventResponse) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *ExternalScheduledEventResponse) UnsetImage() {
	o.Image.Unset()
}

// GetScheduledStartTime returns the ScheduledStartTime field value
func (o *ExternalScheduledEventResponse) GetScheduledStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ScheduledStartTime
}

// GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetScheduledStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledStartTime, true
}

// SetScheduledStartTime sets field value
func (o *ExternalScheduledEventResponse) SetScheduledStartTime(v time.Time) {
	o.ScheduledStartTime = v
}

// GetScheduledEndTime returns the ScheduledEndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalScheduledEventResponse) GetScheduledEndTime() time.Time {
	if o == nil || IsNil(o.ScheduledEndTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledEndTime.Get()
}

// GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalScheduledEventResponse) GetScheduledEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledEndTime.Get(), o.ScheduledEndTime.IsSet()
}

// HasScheduledEndTime returns a boolean if a field has been set.
func (o *ExternalScheduledEventResponse) HasScheduledEndTime() bool {
	if o != nil && o.ScheduledEndTime.IsSet() {
		return true
	}

	return false
}

// SetScheduledEndTime gets a reference to the given NullableTime and assigns it to the ScheduledEndTime field.
func (o *ExternalScheduledEventResponse) SetScheduledEndTime(v time.Time) {
	o.ScheduledEndTime.Set(&v)
}
// SetScheduledEndTimeNil sets the value for ScheduledEndTime to be an explicit nil
func (o *ExternalScheduledEventResponse) SetScheduledEndTimeNil() {
	o.ScheduledEndTime.Set(nil)
}

// UnsetScheduledEndTime ensures that no value is present for ScheduledEndTime, not even an explicit nil
func (o *ExternalScheduledEventResponse) UnsetScheduledEndTime() {
	o.ScheduledEndTime.Unset()
}

// GetStatus returns the Status field value
func (o *ExternalScheduledEventResponse) GetStatus() GuildScheduledEventStatuses {
	if o == nil {
		var ret GuildScheduledEventStatuses
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetStatusOk() (*GuildScheduledEventStatuses, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ExternalScheduledEventResponse) SetStatus(v GuildScheduledEventStatuses) {
	o.Status = v
}

// GetEntityType returns the EntityType field value
func (o *ExternalScheduledEventResponse) GetEntityType() GuildScheduledEventEntityTypes {
	if o == nil {
		var ret GuildScheduledEventEntityTypes
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetEntityTypeOk() (*GuildScheduledEventEntityTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *ExternalScheduledEventResponse) SetEntityType(v GuildScheduledEventEntityTypes) {
	o.EntityType = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *ExternalScheduledEventResponse) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *ExternalScheduledEventResponse) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *ExternalScheduledEventResponse) SetEntityId(v string) {
	o.EntityId = &v
}

// GetUserCount returns the UserCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalScheduledEventResponse) GetUserCount() int32 {
	if o == nil || IsNil(o.UserCount.Get()) {
		var ret int32
		return ret
	}
	return *o.UserCount.Get()
}

// GetUserCountOk returns a tuple with the UserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalScheduledEventResponse) GetUserCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserCount.Get(), o.UserCount.IsSet()
}

// HasUserCount returns a boolean if a field has been set.
func (o *ExternalScheduledEventResponse) HasUserCount() bool {
	if o != nil && o.UserCount.IsSet() {
		return true
	}

	return false
}

// SetUserCount gets a reference to the given NullableInt32 and assigns it to the UserCount field.
func (o *ExternalScheduledEventResponse) SetUserCount(v int32) {
	o.UserCount.Set(&v)
}
// SetUserCountNil sets the value for UserCount to be an explicit nil
func (o *ExternalScheduledEventResponse) SetUserCountNil() {
	o.UserCount.Set(nil)
}

// UnsetUserCount ensures that no value is present for UserCount, not even an explicit nil
func (o *ExternalScheduledEventResponse) UnsetUserCount() {
	o.UserCount.Unset()
}

// GetPrivacyLevel returns the PrivacyLevel field value
func (o *ExternalScheduledEventResponse) GetPrivacyLevel() GuildScheduledEventPrivacyLevels {
	if o == nil {
		var ret GuildScheduledEventPrivacyLevels
		return ret
	}

	return o.PrivacyLevel
}

// GetPrivacyLevelOk returns a tuple with the PrivacyLevel field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetPrivacyLevelOk() (*GuildScheduledEventPrivacyLevels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivacyLevel, true
}

// SetPrivacyLevel sets field value
func (o *ExternalScheduledEventResponse) SetPrivacyLevel(v GuildScheduledEventPrivacyLevels) {
	o.PrivacyLevel = v
}

// GetUserRsvp returns the UserRsvp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalScheduledEventResponse) GetUserRsvp() ScheduledEventUserResponse {
	if o == nil || IsNil(o.UserRsvp.Get()) {
		var ret ScheduledEventUserResponse
		return ret
	}
	return *o.UserRsvp.Get()
}

// GetUserRsvpOk returns a tuple with the UserRsvp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalScheduledEventResponse) GetUserRsvpOk() (*ScheduledEventUserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserRsvp.Get(), o.UserRsvp.IsSet()
}

// HasUserRsvp returns a boolean if a field has been set.
func (o *ExternalScheduledEventResponse) HasUserRsvp() bool {
	if o != nil && o.UserRsvp.IsSet() {
		return true
	}

	return false
}

// SetUserRsvp gets a reference to the given NullableScheduledEventUserResponse and assigns it to the UserRsvp field.
func (o *ExternalScheduledEventResponse) SetUserRsvp(v ScheduledEventUserResponse) {
	o.UserRsvp.Set(&v)
}
// SetUserRsvpNil sets the value for UserRsvp to be an explicit nil
func (o *ExternalScheduledEventResponse) SetUserRsvpNil() {
	o.UserRsvp.Set(nil)
}

// UnsetUserRsvp ensures that no value is present for UserRsvp, not even an explicit nil
func (o *ExternalScheduledEventResponse) UnsetUserRsvp() {
	o.UserRsvp.Unset()
}

// GetEntityMetadata returns the EntityMetadata field value
func (o *ExternalScheduledEventResponse) GetEntityMetadata() EntityMetadataExternalResponse {
	if o == nil {
		var ret EntityMetadataExternalResponse
		return ret
	}

	return o.EntityMetadata
}

// GetEntityMetadataOk returns a tuple with the EntityMetadata field value
// and a boolean to check if the value has been set.
func (o *ExternalScheduledEventResponse) GetEntityMetadataOk() (*EntityMetadataExternalResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityMetadata, true
}

// SetEntityMetadata sets field value
func (o *ExternalScheduledEventResponse) SetEntityMetadata(v EntityMetadataExternalResponse) {
	o.EntityMetadata = v
}

func (o ExternalScheduledEventResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalScheduledEventResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["guild_id"] = o.GuildId
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.CreatorId) {
		toSerialize["creator_id"] = o.CreatorId
	}
	if o.Creator.IsSet() {
		toSerialize["creator"] = o.Creator.Get()
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	toSerialize["scheduled_start_time"] = o.ScheduledStartTime
	if o.ScheduledEndTime.IsSet() {
		toSerialize["scheduled_end_time"] = o.ScheduledEndTime.Get()
	}
	toSerialize["status"] = o.Status
	toSerialize["entity_type"] = o.EntityType
	if !IsNil(o.EntityId) {
		toSerialize["entity_id"] = o.EntityId
	}
	if o.UserCount.IsSet() {
		toSerialize["user_count"] = o.UserCount.Get()
	}
	toSerialize["privacy_level"] = o.PrivacyLevel
	if o.UserRsvp.IsSet() {
		toSerialize["user_rsvp"] = o.UserRsvp.Get()
	}
	toSerialize["entity_metadata"] = o.EntityMetadata
	return toSerialize, nil
}

func (o *ExternalScheduledEventResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"guild_id",
		"name",
		"scheduled_start_time",
		"status",
		"entity_type",
		"privacy_level",
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

	varExternalScheduledEventResponse := _ExternalScheduledEventResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalScheduledEventResponse)

	if err != nil {
		return err
	}

	*o = ExternalScheduledEventResponse(varExternalScheduledEventResponse)

	return err
}

type NullableExternalScheduledEventResponse struct {
	value *ExternalScheduledEventResponse
	isSet bool
}

func (v NullableExternalScheduledEventResponse) Get() *ExternalScheduledEventResponse {
	return v.value
}

func (v *NullableExternalScheduledEventResponse) Set(val *ExternalScheduledEventResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalScheduledEventResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalScheduledEventResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalScheduledEventResponse(val *ExternalScheduledEventResponse) *NullableExternalScheduledEventResponse {
	return &NullableExternalScheduledEventResponse{value: val, isSet: true}
}

func (v NullableExternalScheduledEventResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalScheduledEventResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


