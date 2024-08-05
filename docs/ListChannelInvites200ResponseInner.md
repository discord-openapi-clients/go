# ListChannelInvites200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**NullableInviteTypes**](InviteTypes.md) |  | [optional] 
**Code** | **string** |  | 
**Inviter** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**MaxAge** | Pointer to **NullableInt32** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**FriendsCount** | Pointer to **NullableInt32** |  | [optional] 
**Channel** | Pointer to [**InviteChannelResponse**](InviteChannelResponse.md) |  | [optional] 
**IsContact** | Pointer to **NullableBool** |  | [optional] 
**Uses** | Pointer to **NullableInt32** |  | [optional] 
**MaxUses** | Pointer to **NullableInt32** |  | [optional] 
**Flags** | Pointer to **NullableInt32** |  | [optional] 
**ApproximateMemberCount** | Pointer to **NullableInt32** |  | [optional] 
**Guild** | Pointer to [**InviteGuildResponse**](InviteGuildResponse.md) |  | [optional] 
**GuildId** | Pointer to **string** |  | [optional] 
**StageInstance** | Pointer to [**InviteStageInstanceResponse**](InviteStageInstanceResponse.md) |  | [optional] 
**TargetType** | Pointer to [**InviteTargetTypes**](InviteTargetTypes.md) |  | [optional] 
**TargetUser** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**TargetApplication** | Pointer to [**InviteApplicationResponse**](InviteApplicationResponse.md) |  | [optional] 
**GuildScheduledEvent** | Pointer to [**ScheduledEventResponse**](ScheduledEventResponse.md) |  | [optional] 
**Temporary** | Pointer to **NullableBool** |  | [optional] 
**ApproximatePresenceCount** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewListChannelInvites200ResponseInner

`func NewListChannelInvites200ResponseInner(code string, ) *ListChannelInvites200ResponseInner`

NewListChannelInvites200ResponseInner instantiates a new ListChannelInvites200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListChannelInvites200ResponseInnerWithDefaults

`func NewListChannelInvites200ResponseInnerWithDefaults() *ListChannelInvites200ResponseInner`

NewListChannelInvites200ResponseInnerWithDefaults instantiates a new ListChannelInvites200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ListChannelInvites200ResponseInner) GetType() InviteTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListChannelInvites200ResponseInner) GetTypeOk() (*InviteTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListChannelInvites200ResponseInner) SetType(v InviteTypes)`

SetType sets Type field to given value.

### HasType

`func (o *ListChannelInvites200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ListChannelInvites200ResponseInner) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ListChannelInvites200ResponseInner) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCode

`func (o *ListChannelInvites200ResponseInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListChannelInvites200ResponseInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListChannelInvites200ResponseInner) SetCode(v string)`

SetCode sets Code field to given value.


### GetInviter

`func (o *ListChannelInvites200ResponseInner) GetInviter() UserResponse`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *ListChannelInvites200ResponseInner) GetInviterOk() (*UserResponse, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *ListChannelInvites200ResponseInner) SetInviter(v UserResponse)`

SetInviter sets Inviter field to given value.

### HasInviter

`func (o *ListChannelInvites200ResponseInner) HasInviter() bool`

HasInviter returns a boolean if a field has been set.

### GetMaxAge

`func (o *ListChannelInvites200ResponseInner) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *ListChannelInvites200ResponseInner) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *ListChannelInvites200ResponseInner) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *ListChannelInvites200ResponseInner) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### SetMaxAgeNil

`func (o *ListChannelInvites200ResponseInner) SetMaxAgeNil(b bool)`

 SetMaxAgeNil sets the value for MaxAge to be an explicit nil

### UnsetMaxAge
`func (o *ListChannelInvites200ResponseInner) UnsetMaxAge()`

UnsetMaxAge ensures that no value is present for MaxAge, not even an explicit nil
### GetCreatedAt

`func (o *ListChannelInvites200ResponseInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListChannelInvites200ResponseInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListChannelInvites200ResponseInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListChannelInvites200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ListChannelInvites200ResponseInner) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ListChannelInvites200ResponseInner) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetExpiresAt

`func (o *ListChannelInvites200ResponseInner) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ListChannelInvites200ResponseInner) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ListChannelInvites200ResponseInner) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ListChannelInvites200ResponseInner) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *ListChannelInvites200ResponseInner) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *ListChannelInvites200ResponseInner) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetFriendsCount

`func (o *ListChannelInvites200ResponseInner) GetFriendsCount() int32`

GetFriendsCount returns the FriendsCount field if non-nil, zero value otherwise.

### GetFriendsCountOk

`func (o *ListChannelInvites200ResponseInner) GetFriendsCountOk() (*int32, bool)`

GetFriendsCountOk returns a tuple with the FriendsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendsCount

`func (o *ListChannelInvites200ResponseInner) SetFriendsCount(v int32)`

SetFriendsCount sets FriendsCount field to given value.

### HasFriendsCount

`func (o *ListChannelInvites200ResponseInner) HasFriendsCount() bool`

HasFriendsCount returns a boolean if a field has been set.

### SetFriendsCountNil

`func (o *ListChannelInvites200ResponseInner) SetFriendsCountNil(b bool)`

 SetFriendsCountNil sets the value for FriendsCount to be an explicit nil

### UnsetFriendsCount
`func (o *ListChannelInvites200ResponseInner) UnsetFriendsCount()`

UnsetFriendsCount ensures that no value is present for FriendsCount, not even an explicit nil
### GetChannel

`func (o *ListChannelInvites200ResponseInner) GetChannel() InviteChannelResponse`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ListChannelInvites200ResponseInner) GetChannelOk() (*InviteChannelResponse, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ListChannelInvites200ResponseInner) SetChannel(v InviteChannelResponse)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ListChannelInvites200ResponseInner) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetIsContact

`func (o *ListChannelInvites200ResponseInner) GetIsContact() bool`

GetIsContact returns the IsContact field if non-nil, zero value otherwise.

### GetIsContactOk

`func (o *ListChannelInvites200ResponseInner) GetIsContactOk() (*bool, bool)`

GetIsContactOk returns a tuple with the IsContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContact

`func (o *ListChannelInvites200ResponseInner) SetIsContact(v bool)`

SetIsContact sets IsContact field to given value.

### HasIsContact

`func (o *ListChannelInvites200ResponseInner) HasIsContact() bool`

HasIsContact returns a boolean if a field has been set.

### SetIsContactNil

`func (o *ListChannelInvites200ResponseInner) SetIsContactNil(b bool)`

 SetIsContactNil sets the value for IsContact to be an explicit nil

### UnsetIsContact
`func (o *ListChannelInvites200ResponseInner) UnsetIsContact()`

UnsetIsContact ensures that no value is present for IsContact, not even an explicit nil
### GetUses

`func (o *ListChannelInvites200ResponseInner) GetUses() int32`

GetUses returns the Uses field if non-nil, zero value otherwise.

### GetUsesOk

`func (o *ListChannelInvites200ResponseInner) GetUsesOk() (*int32, bool)`

GetUsesOk returns a tuple with the Uses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUses

`func (o *ListChannelInvites200ResponseInner) SetUses(v int32)`

SetUses sets Uses field to given value.

### HasUses

`func (o *ListChannelInvites200ResponseInner) HasUses() bool`

HasUses returns a boolean if a field has been set.

### SetUsesNil

`func (o *ListChannelInvites200ResponseInner) SetUsesNil(b bool)`

 SetUsesNil sets the value for Uses to be an explicit nil

### UnsetUses
`func (o *ListChannelInvites200ResponseInner) UnsetUses()`

UnsetUses ensures that no value is present for Uses, not even an explicit nil
### GetMaxUses

`func (o *ListChannelInvites200ResponseInner) GetMaxUses() int32`

GetMaxUses returns the MaxUses field if non-nil, zero value otherwise.

### GetMaxUsesOk

`func (o *ListChannelInvites200ResponseInner) GetMaxUsesOk() (*int32, bool)`

GetMaxUsesOk returns a tuple with the MaxUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUses

`func (o *ListChannelInvites200ResponseInner) SetMaxUses(v int32)`

SetMaxUses sets MaxUses field to given value.

### HasMaxUses

`func (o *ListChannelInvites200ResponseInner) HasMaxUses() bool`

HasMaxUses returns a boolean if a field has been set.

### SetMaxUsesNil

`func (o *ListChannelInvites200ResponseInner) SetMaxUsesNil(b bool)`

 SetMaxUsesNil sets the value for MaxUses to be an explicit nil

### UnsetMaxUses
`func (o *ListChannelInvites200ResponseInner) UnsetMaxUses()`

UnsetMaxUses ensures that no value is present for MaxUses, not even an explicit nil
### GetFlags

`func (o *ListChannelInvites200ResponseInner) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ListChannelInvites200ResponseInner) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ListChannelInvites200ResponseInner) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *ListChannelInvites200ResponseInner) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### SetFlagsNil

`func (o *ListChannelInvites200ResponseInner) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *ListChannelInvites200ResponseInner) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil
### GetApproximateMemberCount

`func (o *ListChannelInvites200ResponseInner) GetApproximateMemberCount() int32`

GetApproximateMemberCount returns the ApproximateMemberCount field if non-nil, zero value otherwise.

### GetApproximateMemberCountOk

`func (o *ListChannelInvites200ResponseInner) GetApproximateMemberCountOk() (*int32, bool)`

GetApproximateMemberCountOk returns a tuple with the ApproximateMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateMemberCount

`func (o *ListChannelInvites200ResponseInner) SetApproximateMemberCount(v int32)`

SetApproximateMemberCount sets ApproximateMemberCount field to given value.

### HasApproximateMemberCount

`func (o *ListChannelInvites200ResponseInner) HasApproximateMemberCount() bool`

HasApproximateMemberCount returns a boolean if a field has been set.

### SetApproximateMemberCountNil

`func (o *ListChannelInvites200ResponseInner) SetApproximateMemberCountNil(b bool)`

 SetApproximateMemberCountNil sets the value for ApproximateMemberCount to be an explicit nil

### UnsetApproximateMemberCount
`func (o *ListChannelInvites200ResponseInner) UnsetApproximateMemberCount()`

UnsetApproximateMemberCount ensures that no value is present for ApproximateMemberCount, not even an explicit nil
### GetGuild

`func (o *ListChannelInvites200ResponseInner) GetGuild() InviteGuildResponse`

GetGuild returns the Guild field if non-nil, zero value otherwise.

### GetGuildOk

`func (o *ListChannelInvites200ResponseInner) GetGuildOk() (*InviteGuildResponse, bool)`

GetGuildOk returns a tuple with the Guild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuild

`func (o *ListChannelInvites200ResponseInner) SetGuild(v InviteGuildResponse)`

SetGuild sets Guild field to given value.

### HasGuild

`func (o *ListChannelInvites200ResponseInner) HasGuild() bool`

HasGuild returns a boolean if a field has been set.

### GetGuildId

`func (o *ListChannelInvites200ResponseInner) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *ListChannelInvites200ResponseInner) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *ListChannelInvites200ResponseInner) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.

### HasGuildId

`func (o *ListChannelInvites200ResponseInner) HasGuildId() bool`

HasGuildId returns a boolean if a field has been set.

### GetStageInstance

`func (o *ListChannelInvites200ResponseInner) GetStageInstance() InviteStageInstanceResponse`

GetStageInstance returns the StageInstance field if non-nil, zero value otherwise.

### GetStageInstanceOk

`func (o *ListChannelInvites200ResponseInner) GetStageInstanceOk() (*InviteStageInstanceResponse, bool)`

GetStageInstanceOk returns a tuple with the StageInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageInstance

`func (o *ListChannelInvites200ResponseInner) SetStageInstance(v InviteStageInstanceResponse)`

SetStageInstance sets StageInstance field to given value.

### HasStageInstance

`func (o *ListChannelInvites200ResponseInner) HasStageInstance() bool`

HasStageInstance returns a boolean if a field has been set.

### GetTargetType

`func (o *ListChannelInvites200ResponseInner) GetTargetType() InviteTargetTypes`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ListChannelInvites200ResponseInner) GetTargetTypeOk() (*InviteTargetTypes, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ListChannelInvites200ResponseInner) SetTargetType(v InviteTargetTypes)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ListChannelInvites200ResponseInner) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTargetUser

`func (o *ListChannelInvites200ResponseInner) GetTargetUser() UserResponse`

GetTargetUser returns the TargetUser field if non-nil, zero value otherwise.

### GetTargetUserOk

`func (o *ListChannelInvites200ResponseInner) GetTargetUserOk() (*UserResponse, bool)`

GetTargetUserOk returns a tuple with the TargetUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUser

`func (o *ListChannelInvites200ResponseInner) SetTargetUser(v UserResponse)`

SetTargetUser sets TargetUser field to given value.

### HasTargetUser

`func (o *ListChannelInvites200ResponseInner) HasTargetUser() bool`

HasTargetUser returns a boolean if a field has been set.

### GetTargetApplication

`func (o *ListChannelInvites200ResponseInner) GetTargetApplication() InviteApplicationResponse`

GetTargetApplication returns the TargetApplication field if non-nil, zero value otherwise.

### GetTargetApplicationOk

`func (o *ListChannelInvites200ResponseInner) GetTargetApplicationOk() (*InviteApplicationResponse, bool)`

GetTargetApplicationOk returns a tuple with the TargetApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApplication

`func (o *ListChannelInvites200ResponseInner) SetTargetApplication(v InviteApplicationResponse)`

SetTargetApplication sets TargetApplication field to given value.

### HasTargetApplication

`func (o *ListChannelInvites200ResponseInner) HasTargetApplication() bool`

HasTargetApplication returns a boolean if a field has been set.

### GetGuildScheduledEvent

`func (o *ListChannelInvites200ResponseInner) GetGuildScheduledEvent() ScheduledEventResponse`

GetGuildScheduledEvent returns the GuildScheduledEvent field if non-nil, zero value otherwise.

### GetGuildScheduledEventOk

`func (o *ListChannelInvites200ResponseInner) GetGuildScheduledEventOk() (*ScheduledEventResponse, bool)`

GetGuildScheduledEventOk returns a tuple with the GuildScheduledEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildScheduledEvent

`func (o *ListChannelInvites200ResponseInner) SetGuildScheduledEvent(v ScheduledEventResponse)`

SetGuildScheduledEvent sets GuildScheduledEvent field to given value.

### HasGuildScheduledEvent

`func (o *ListChannelInvites200ResponseInner) HasGuildScheduledEvent() bool`

HasGuildScheduledEvent returns a boolean if a field has been set.

### GetTemporary

`func (o *ListChannelInvites200ResponseInner) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *ListChannelInvites200ResponseInner) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *ListChannelInvites200ResponseInner) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.

### HasTemporary

`func (o *ListChannelInvites200ResponseInner) HasTemporary() bool`

HasTemporary returns a boolean if a field has been set.

### SetTemporaryNil

`func (o *ListChannelInvites200ResponseInner) SetTemporaryNil(b bool)`

 SetTemporaryNil sets the value for Temporary to be an explicit nil

### UnsetTemporary
`func (o *ListChannelInvites200ResponseInner) UnsetTemporary()`

UnsetTemporary ensures that no value is present for Temporary, not even an explicit nil
### GetApproximatePresenceCount

`func (o *ListChannelInvites200ResponseInner) GetApproximatePresenceCount() int32`

GetApproximatePresenceCount returns the ApproximatePresenceCount field if non-nil, zero value otherwise.

### GetApproximatePresenceCountOk

`func (o *ListChannelInvites200ResponseInner) GetApproximatePresenceCountOk() (*int32, bool)`

GetApproximatePresenceCountOk returns a tuple with the ApproximatePresenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximatePresenceCount

`func (o *ListChannelInvites200ResponseInner) SetApproximatePresenceCount(v int32)`

SetApproximatePresenceCount sets ApproximatePresenceCount field to given value.

### HasApproximatePresenceCount

`func (o *ListChannelInvites200ResponseInner) HasApproximatePresenceCount() bool`

HasApproximatePresenceCount returns a boolean if a field has been set.

### SetApproximatePresenceCountNil

`func (o *ListChannelInvites200ResponseInner) SetApproximatePresenceCountNil(b bool)`

 SetApproximatePresenceCountNil sets the value for ApproximatePresenceCount to be an explicit nil

### UnsetApproximatePresenceCount
`func (o *ListChannelInvites200ResponseInner) UnsetApproximatePresenceCount()`

UnsetApproximatePresenceCount ensures that no value is present for ApproximatePresenceCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


