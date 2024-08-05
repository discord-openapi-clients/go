# GetGuildScheduledEvent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**GuildId** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**Image** | Pointer to **NullableString** |  | [optional] 
**ScheduledStartTime** | **time.Time** |  | 
**ScheduledEndTime** | Pointer to **NullableTime** |  | [optional] 
**Status** | [**GuildScheduledEventStatuses**](GuildScheduledEventStatuses.md) |  | 
**EntityType** | [**GuildScheduledEventEntityTypes**](GuildScheduledEventEntityTypes.md) |  | 
**EntityId** | Pointer to **string** |  | [optional] 
**UserCount** | Pointer to **NullableInt32** |  | [optional] 
**PrivacyLevel** | [**GuildScheduledEventPrivacyLevels**](GuildScheduledEventPrivacyLevels.md) |  | 
**UserRsvp** | Pointer to [**ScheduledEventUserResponse**](ScheduledEventUserResponse.md) |  | [optional] 
**EntityMetadata** | **map[string]interface{}** |  | 

## Methods

### NewGetGuildScheduledEvent200Response

`func NewGetGuildScheduledEvent200Response(id string, guildId string, name string, scheduledStartTime time.Time, status GuildScheduledEventStatuses, entityType GuildScheduledEventEntityTypes, privacyLevel GuildScheduledEventPrivacyLevels, entityMetadata map[string]interface{}, ) *GetGuildScheduledEvent200Response`

NewGetGuildScheduledEvent200Response instantiates a new GetGuildScheduledEvent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGuildScheduledEvent200ResponseWithDefaults

`func NewGetGuildScheduledEvent200ResponseWithDefaults() *GetGuildScheduledEvent200Response`

NewGetGuildScheduledEvent200ResponseWithDefaults instantiates a new GetGuildScheduledEvent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetGuildScheduledEvent200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetGuildScheduledEvent200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetGuildScheduledEvent200Response) SetId(v string)`

SetId sets Id field to given value.


### GetGuildId

`func (o *GetGuildScheduledEvent200Response) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *GetGuildScheduledEvent200Response) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *GetGuildScheduledEvent200Response) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetName

`func (o *GetGuildScheduledEvent200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetGuildScheduledEvent200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetGuildScheduledEvent200Response) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GetGuildScheduledEvent200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetGuildScheduledEvent200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetGuildScheduledEvent200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetGuildScheduledEvent200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetGuildScheduledEvent200Response) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetGuildScheduledEvent200Response) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetChannelId

`func (o *GetGuildScheduledEvent200Response) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *GetGuildScheduledEvent200Response) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *GetGuildScheduledEvent200Response) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *GetGuildScheduledEvent200Response) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetCreatorId

`func (o *GetGuildScheduledEvent200Response) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *GetGuildScheduledEvent200Response) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *GetGuildScheduledEvent200Response) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *GetGuildScheduledEvent200Response) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetCreator

`func (o *GetGuildScheduledEvent200Response) GetCreator() UserResponse`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *GetGuildScheduledEvent200Response) GetCreatorOk() (*UserResponse, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *GetGuildScheduledEvent200Response) SetCreator(v UserResponse)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *GetGuildScheduledEvent200Response) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetImage

`func (o *GetGuildScheduledEvent200Response) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *GetGuildScheduledEvent200Response) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *GetGuildScheduledEvent200Response) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *GetGuildScheduledEvent200Response) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *GetGuildScheduledEvent200Response) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *GetGuildScheduledEvent200Response) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetScheduledStartTime

`func (o *GetGuildScheduledEvent200Response) GetScheduledStartTime() time.Time`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *GetGuildScheduledEvent200Response) GetScheduledStartTimeOk() (*time.Time, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *GetGuildScheduledEvent200Response) SetScheduledStartTime(v time.Time)`

SetScheduledStartTime sets ScheduledStartTime field to given value.


### GetScheduledEndTime

`func (o *GetGuildScheduledEvent200Response) GetScheduledEndTime() time.Time`

GetScheduledEndTime returns the ScheduledEndTime field if non-nil, zero value otherwise.

### GetScheduledEndTimeOk

`func (o *GetGuildScheduledEvent200Response) GetScheduledEndTimeOk() (*time.Time, bool)`

GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEndTime

`func (o *GetGuildScheduledEvent200Response) SetScheduledEndTime(v time.Time)`

SetScheduledEndTime sets ScheduledEndTime field to given value.

### HasScheduledEndTime

`func (o *GetGuildScheduledEvent200Response) HasScheduledEndTime() bool`

HasScheduledEndTime returns a boolean if a field has been set.

### SetScheduledEndTimeNil

`func (o *GetGuildScheduledEvent200Response) SetScheduledEndTimeNil(b bool)`

 SetScheduledEndTimeNil sets the value for ScheduledEndTime to be an explicit nil

### UnsetScheduledEndTime
`func (o *GetGuildScheduledEvent200Response) UnsetScheduledEndTime()`

UnsetScheduledEndTime ensures that no value is present for ScheduledEndTime, not even an explicit nil
### GetStatus

`func (o *GetGuildScheduledEvent200Response) GetStatus() GuildScheduledEventStatuses`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetGuildScheduledEvent200Response) GetStatusOk() (*GuildScheduledEventStatuses, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetGuildScheduledEvent200Response) SetStatus(v GuildScheduledEventStatuses)`

SetStatus sets Status field to given value.


### GetEntityType

`func (o *GetGuildScheduledEvent200Response) GetEntityType() GuildScheduledEventEntityTypes`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *GetGuildScheduledEvent200Response) GetEntityTypeOk() (*GuildScheduledEventEntityTypes, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *GetGuildScheduledEvent200Response) SetEntityType(v GuildScheduledEventEntityTypes)`

SetEntityType sets EntityType field to given value.


### GetEntityId

`func (o *GetGuildScheduledEvent200Response) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *GetGuildScheduledEvent200Response) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *GetGuildScheduledEvent200Response) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *GetGuildScheduledEvent200Response) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetUserCount

`func (o *GetGuildScheduledEvent200Response) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *GetGuildScheduledEvent200Response) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *GetGuildScheduledEvent200Response) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *GetGuildScheduledEvent200Response) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### SetUserCountNil

`func (o *GetGuildScheduledEvent200Response) SetUserCountNil(b bool)`

 SetUserCountNil sets the value for UserCount to be an explicit nil

### UnsetUserCount
`func (o *GetGuildScheduledEvent200Response) UnsetUserCount()`

UnsetUserCount ensures that no value is present for UserCount, not even an explicit nil
### GetPrivacyLevel

`func (o *GetGuildScheduledEvent200Response) GetPrivacyLevel() GuildScheduledEventPrivacyLevels`

GetPrivacyLevel returns the PrivacyLevel field if non-nil, zero value otherwise.

### GetPrivacyLevelOk

`func (o *GetGuildScheduledEvent200Response) GetPrivacyLevelOk() (*GuildScheduledEventPrivacyLevels, bool)`

GetPrivacyLevelOk returns a tuple with the PrivacyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevel

`func (o *GetGuildScheduledEvent200Response) SetPrivacyLevel(v GuildScheduledEventPrivacyLevels)`

SetPrivacyLevel sets PrivacyLevel field to given value.


### GetUserRsvp

`func (o *GetGuildScheduledEvent200Response) GetUserRsvp() ScheduledEventUserResponse`

GetUserRsvp returns the UserRsvp field if non-nil, zero value otherwise.

### GetUserRsvpOk

`func (o *GetGuildScheduledEvent200Response) GetUserRsvpOk() (*ScheduledEventUserResponse, bool)`

GetUserRsvpOk returns a tuple with the UserRsvp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRsvp

`func (o *GetGuildScheduledEvent200Response) SetUserRsvp(v ScheduledEventUserResponse)`

SetUserRsvp sets UserRsvp field to given value.

### HasUserRsvp

`func (o *GetGuildScheduledEvent200Response) HasUserRsvp() bool`

HasUserRsvp returns a boolean if a field has been set.

### GetEntityMetadata

`func (o *GetGuildScheduledEvent200Response) GetEntityMetadata() map[string]interface{}`

GetEntityMetadata returns the EntityMetadata field if non-nil, zero value otherwise.

### GetEntityMetadataOk

`func (o *GetGuildScheduledEvent200Response) GetEntityMetadataOk() (*map[string]interface{}, bool)`

GetEntityMetadataOk returns a tuple with the EntityMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityMetadata

`func (o *GetGuildScheduledEvent200Response) SetEntityMetadata(v map[string]interface{})`

SetEntityMetadata sets EntityMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


