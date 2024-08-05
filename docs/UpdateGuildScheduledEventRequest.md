# UpdateGuildScheduledEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**GuildScheduledEventStatuses**](GuildScheduledEventStatuses.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Image** | Pointer to **NullableString** |  | [optional] 
**ScheduledStartTime** | Pointer to **time.Time** |  | [optional] 
**ScheduledEndTime** | Pointer to **NullableTime** |  | [optional] 
**EntityType** | Pointer to [**NullableGuildScheduledEventEntityTypes**](GuildScheduledEventEntityTypes.md) |  | [optional] 
**PrivacyLevel** | Pointer to [**GuildScheduledEventPrivacyLevels**](GuildScheduledEventPrivacyLevels.md) |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**EntityMetadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateGuildScheduledEventRequest

`func NewUpdateGuildScheduledEventRequest() *UpdateGuildScheduledEventRequest`

NewUpdateGuildScheduledEventRequest instantiates a new UpdateGuildScheduledEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGuildScheduledEventRequestWithDefaults

`func NewUpdateGuildScheduledEventRequestWithDefaults() *UpdateGuildScheduledEventRequest`

NewUpdateGuildScheduledEventRequestWithDefaults instantiates a new UpdateGuildScheduledEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateGuildScheduledEventRequest) GetStatus() GuildScheduledEventStatuses`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateGuildScheduledEventRequest) GetStatusOk() (*GuildScheduledEventStatuses, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateGuildScheduledEventRequest) SetStatus(v GuildScheduledEventStatuses)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateGuildScheduledEventRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetName

`func (o *UpdateGuildScheduledEventRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGuildScheduledEventRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGuildScheduledEventRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGuildScheduledEventRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateGuildScheduledEventRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateGuildScheduledEventRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateGuildScheduledEventRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateGuildScheduledEventRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateGuildScheduledEventRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateGuildScheduledEventRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImage

`func (o *UpdateGuildScheduledEventRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *UpdateGuildScheduledEventRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *UpdateGuildScheduledEventRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *UpdateGuildScheduledEventRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *UpdateGuildScheduledEventRequest) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *UpdateGuildScheduledEventRequest) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetScheduledStartTime

`func (o *UpdateGuildScheduledEventRequest) GetScheduledStartTime() time.Time`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *UpdateGuildScheduledEventRequest) GetScheduledStartTimeOk() (*time.Time, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *UpdateGuildScheduledEventRequest) SetScheduledStartTime(v time.Time)`

SetScheduledStartTime sets ScheduledStartTime field to given value.

### HasScheduledStartTime

`func (o *UpdateGuildScheduledEventRequest) HasScheduledStartTime() bool`

HasScheduledStartTime returns a boolean if a field has been set.

### GetScheduledEndTime

`func (o *UpdateGuildScheduledEventRequest) GetScheduledEndTime() time.Time`

GetScheduledEndTime returns the ScheduledEndTime field if non-nil, zero value otherwise.

### GetScheduledEndTimeOk

`func (o *UpdateGuildScheduledEventRequest) GetScheduledEndTimeOk() (*time.Time, bool)`

GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEndTime

`func (o *UpdateGuildScheduledEventRequest) SetScheduledEndTime(v time.Time)`

SetScheduledEndTime sets ScheduledEndTime field to given value.

### HasScheduledEndTime

`func (o *UpdateGuildScheduledEventRequest) HasScheduledEndTime() bool`

HasScheduledEndTime returns a boolean if a field has been set.

### SetScheduledEndTimeNil

`func (o *UpdateGuildScheduledEventRequest) SetScheduledEndTimeNil(b bool)`

 SetScheduledEndTimeNil sets the value for ScheduledEndTime to be an explicit nil

### UnsetScheduledEndTime
`func (o *UpdateGuildScheduledEventRequest) UnsetScheduledEndTime()`

UnsetScheduledEndTime ensures that no value is present for ScheduledEndTime, not even an explicit nil
### GetEntityType

`func (o *UpdateGuildScheduledEventRequest) GetEntityType() GuildScheduledEventEntityTypes`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *UpdateGuildScheduledEventRequest) GetEntityTypeOk() (*GuildScheduledEventEntityTypes, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *UpdateGuildScheduledEventRequest) SetEntityType(v GuildScheduledEventEntityTypes)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *UpdateGuildScheduledEventRequest) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### SetEntityTypeNil

`func (o *UpdateGuildScheduledEventRequest) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *UpdateGuildScheduledEventRequest) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetPrivacyLevel

`func (o *UpdateGuildScheduledEventRequest) GetPrivacyLevel() GuildScheduledEventPrivacyLevels`

GetPrivacyLevel returns the PrivacyLevel field if non-nil, zero value otherwise.

### GetPrivacyLevelOk

`func (o *UpdateGuildScheduledEventRequest) GetPrivacyLevelOk() (*GuildScheduledEventPrivacyLevels, bool)`

GetPrivacyLevelOk returns a tuple with the PrivacyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevel

`func (o *UpdateGuildScheduledEventRequest) SetPrivacyLevel(v GuildScheduledEventPrivacyLevels)`

SetPrivacyLevel sets PrivacyLevel field to given value.

### HasPrivacyLevel

`func (o *UpdateGuildScheduledEventRequest) HasPrivacyLevel() bool`

HasPrivacyLevel returns a boolean if a field has been set.

### GetChannelId

`func (o *UpdateGuildScheduledEventRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *UpdateGuildScheduledEventRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *UpdateGuildScheduledEventRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *UpdateGuildScheduledEventRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetEntityMetadata

`func (o *UpdateGuildScheduledEventRequest) GetEntityMetadata() map[string]interface{}`

GetEntityMetadata returns the EntityMetadata field if non-nil, zero value otherwise.

### GetEntityMetadataOk

`func (o *UpdateGuildScheduledEventRequest) GetEntityMetadataOk() (*map[string]interface{}, bool)`

GetEntityMetadataOk returns a tuple with the EntityMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityMetadata

`func (o *UpdateGuildScheduledEventRequest) SetEntityMetadata(v map[string]interface{})`

SetEntityMetadata sets EntityMetadata field to given value.

### HasEntityMetadata

`func (o *UpdateGuildScheduledEventRequest) HasEntityMetadata() bool`

HasEntityMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


