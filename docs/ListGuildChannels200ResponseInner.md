# ListGuildChannels200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**ChannelTypes**](ChannelTypes.md) |  | 
**LastMessageId** | Pointer to **string** |  | [optional] 
**Flags** | **int32** |  | 
**LastPinTimestamp** | Pointer to **NullableTime** |  | [optional] 
**GuildId** | **string** |  | 
**Name** | **string** |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**RateLimitPerUser** | Pointer to **NullableInt32** |  | [optional] 
**Bitrate** | Pointer to **NullableInt32** |  | [optional] 
**UserLimit** | Pointer to **NullableInt32** |  | [optional] 
**RtcRegion** | Pointer to **NullableString** |  | [optional] 
**VideoQualityMode** | Pointer to [**VideoQualityModes**](VideoQualityModes.md) |  | [optional] 
**Permissions** | Pointer to **NullableString** |  | [optional] 
**Topic** | Pointer to **NullableString** |  | [optional] 
**DefaultAutoArchiveDuration** | Pointer to [**ThreadAutoArchiveDuration**](ThreadAutoArchiveDuration.md) |  | [optional] 
**DefaultThreadRateLimitPerUser** | Pointer to **NullableInt32** |  | [optional] 
**Position** | **int32** |  | 
**PermissionOverwrites** | Pointer to [**[]ChannelPermissionOverwriteResponse**](ChannelPermissionOverwriteResponse.md) |  | [optional] 
**Nsfw** | Pointer to **NullableBool** |  | [optional] 
**AvailableTags** | Pointer to [**[]ForumTagResponse**](ForumTagResponse.md) |  | [optional] 
**DefaultReactionEmoji** | Pointer to [**DefaultReactionEmojiResponse**](DefaultReactionEmojiResponse.md) |  | [optional] 
**DefaultSortOrder** | Pointer to [**ThreadSortOrder**](ThreadSortOrder.md) |  | [optional] 
**DefaultForumLayout** | Pointer to [**ForumLayout**](ForumLayout.md) |  | [optional] 
**Recipients** | [**[]UserResponse**](UserResponse.md) |  | 
**Icon** | Pointer to **NullableString** |  | [optional] 
**OwnerId** | **string** |  | 
**Managed** | Pointer to **NullableBool** |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**ThreadMetadata** | Pointer to [**ThreadMetadataResponse**](ThreadMetadataResponse.md) |  | [optional] 
**MessageCount** | **int32** |  | 
**MemberCount** | **int32** |  | 
**TotalMessageSent** | **int32** |  | 
**AppliedTags** | Pointer to **[]string** |  | [optional] 
**Member** | Pointer to [**ThreadMemberResponse**](ThreadMemberResponse.md) |  | [optional] 

## Methods

### NewListGuildChannels200ResponseInner

`func NewListGuildChannels200ResponseInner(id string, type_ ChannelTypes, flags int32, guildId string, name string, position int32, recipients []UserResponse, ownerId string, messageCount int32, memberCount int32, totalMessageSent int32, ) *ListGuildChannels200ResponseInner`

NewListGuildChannels200ResponseInner instantiates a new ListGuildChannels200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGuildChannels200ResponseInnerWithDefaults

`func NewListGuildChannels200ResponseInnerWithDefaults() *ListGuildChannels200ResponseInner`

NewListGuildChannels200ResponseInnerWithDefaults instantiates a new ListGuildChannels200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListGuildChannels200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListGuildChannels200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListGuildChannels200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ListGuildChannels200ResponseInner) GetType() ChannelTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListGuildChannels200ResponseInner) GetTypeOk() (*ChannelTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListGuildChannels200ResponseInner) SetType(v ChannelTypes)`

SetType sets Type field to given value.


### GetLastMessageId

`func (o *ListGuildChannels200ResponseInner) GetLastMessageId() string`

GetLastMessageId returns the LastMessageId field if non-nil, zero value otherwise.

### GetLastMessageIdOk

`func (o *ListGuildChannels200ResponseInner) GetLastMessageIdOk() (*string, bool)`

GetLastMessageIdOk returns a tuple with the LastMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageId

`func (o *ListGuildChannels200ResponseInner) SetLastMessageId(v string)`

SetLastMessageId sets LastMessageId field to given value.

### HasLastMessageId

`func (o *ListGuildChannels200ResponseInner) HasLastMessageId() bool`

HasLastMessageId returns a boolean if a field has been set.

### GetFlags

`func (o *ListGuildChannels200ResponseInner) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ListGuildChannels200ResponseInner) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ListGuildChannels200ResponseInner) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetLastPinTimestamp

`func (o *ListGuildChannels200ResponseInner) GetLastPinTimestamp() time.Time`

GetLastPinTimestamp returns the LastPinTimestamp field if non-nil, zero value otherwise.

### GetLastPinTimestampOk

`func (o *ListGuildChannels200ResponseInner) GetLastPinTimestampOk() (*time.Time, bool)`

GetLastPinTimestampOk returns a tuple with the LastPinTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPinTimestamp

`func (o *ListGuildChannels200ResponseInner) SetLastPinTimestamp(v time.Time)`

SetLastPinTimestamp sets LastPinTimestamp field to given value.

### HasLastPinTimestamp

`func (o *ListGuildChannels200ResponseInner) HasLastPinTimestamp() bool`

HasLastPinTimestamp returns a boolean if a field has been set.

### SetLastPinTimestampNil

`func (o *ListGuildChannels200ResponseInner) SetLastPinTimestampNil(b bool)`

 SetLastPinTimestampNil sets the value for LastPinTimestamp to be an explicit nil

### UnsetLastPinTimestamp
`func (o *ListGuildChannels200ResponseInner) UnsetLastPinTimestamp()`

UnsetLastPinTimestamp ensures that no value is present for LastPinTimestamp, not even an explicit nil
### GetGuildId

`func (o *ListGuildChannels200ResponseInner) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *ListGuildChannels200ResponseInner) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *ListGuildChannels200ResponseInner) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetName

`func (o *ListGuildChannels200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListGuildChannels200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListGuildChannels200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *ListGuildChannels200ResponseInner) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ListGuildChannels200ResponseInner) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ListGuildChannels200ResponseInner) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ListGuildChannels200ResponseInner) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetRateLimitPerUser

`func (o *ListGuildChannels200ResponseInner) GetRateLimitPerUser() int32`

GetRateLimitPerUser returns the RateLimitPerUser field if non-nil, zero value otherwise.

### GetRateLimitPerUserOk

`func (o *ListGuildChannels200ResponseInner) GetRateLimitPerUserOk() (*int32, bool)`

GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerUser

`func (o *ListGuildChannels200ResponseInner) SetRateLimitPerUser(v int32)`

SetRateLimitPerUser sets RateLimitPerUser field to given value.

### HasRateLimitPerUser

`func (o *ListGuildChannels200ResponseInner) HasRateLimitPerUser() bool`

HasRateLimitPerUser returns a boolean if a field has been set.

### SetRateLimitPerUserNil

`func (o *ListGuildChannels200ResponseInner) SetRateLimitPerUserNil(b bool)`

 SetRateLimitPerUserNil sets the value for RateLimitPerUser to be an explicit nil

### UnsetRateLimitPerUser
`func (o *ListGuildChannels200ResponseInner) UnsetRateLimitPerUser()`

UnsetRateLimitPerUser ensures that no value is present for RateLimitPerUser, not even an explicit nil
### GetBitrate

`func (o *ListGuildChannels200ResponseInner) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *ListGuildChannels200ResponseInner) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *ListGuildChannels200ResponseInner) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *ListGuildChannels200ResponseInner) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *ListGuildChannels200ResponseInner) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *ListGuildChannels200ResponseInner) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetUserLimit

`func (o *ListGuildChannels200ResponseInner) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *ListGuildChannels200ResponseInner) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *ListGuildChannels200ResponseInner) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *ListGuildChannels200ResponseInner) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### SetUserLimitNil

`func (o *ListGuildChannels200ResponseInner) SetUserLimitNil(b bool)`

 SetUserLimitNil sets the value for UserLimit to be an explicit nil

### UnsetUserLimit
`func (o *ListGuildChannels200ResponseInner) UnsetUserLimit()`

UnsetUserLimit ensures that no value is present for UserLimit, not even an explicit nil
### GetRtcRegion

`func (o *ListGuildChannels200ResponseInner) GetRtcRegion() string`

GetRtcRegion returns the RtcRegion field if non-nil, zero value otherwise.

### GetRtcRegionOk

`func (o *ListGuildChannels200ResponseInner) GetRtcRegionOk() (*string, bool)`

GetRtcRegionOk returns a tuple with the RtcRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtcRegion

`func (o *ListGuildChannels200ResponseInner) SetRtcRegion(v string)`

SetRtcRegion sets RtcRegion field to given value.

### HasRtcRegion

`func (o *ListGuildChannels200ResponseInner) HasRtcRegion() bool`

HasRtcRegion returns a boolean if a field has been set.

### SetRtcRegionNil

`func (o *ListGuildChannels200ResponseInner) SetRtcRegionNil(b bool)`

 SetRtcRegionNil sets the value for RtcRegion to be an explicit nil

### UnsetRtcRegion
`func (o *ListGuildChannels200ResponseInner) UnsetRtcRegion()`

UnsetRtcRegion ensures that no value is present for RtcRegion, not even an explicit nil
### GetVideoQualityMode

`func (o *ListGuildChannels200ResponseInner) GetVideoQualityMode() VideoQualityModes`

GetVideoQualityMode returns the VideoQualityMode field if non-nil, zero value otherwise.

### GetVideoQualityModeOk

`func (o *ListGuildChannels200ResponseInner) GetVideoQualityModeOk() (*VideoQualityModes, bool)`

GetVideoQualityModeOk returns a tuple with the VideoQualityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQualityMode

`func (o *ListGuildChannels200ResponseInner) SetVideoQualityMode(v VideoQualityModes)`

SetVideoQualityMode sets VideoQualityMode field to given value.

### HasVideoQualityMode

`func (o *ListGuildChannels200ResponseInner) HasVideoQualityMode() bool`

HasVideoQualityMode returns a boolean if a field has been set.

### GetPermissions

`func (o *ListGuildChannels200ResponseInner) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ListGuildChannels200ResponseInner) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ListGuildChannels200ResponseInner) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ListGuildChannels200ResponseInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *ListGuildChannels200ResponseInner) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *ListGuildChannels200ResponseInner) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetTopic

`func (o *ListGuildChannels200ResponseInner) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ListGuildChannels200ResponseInner) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ListGuildChannels200ResponseInner) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *ListGuildChannels200ResponseInner) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *ListGuildChannels200ResponseInner) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *ListGuildChannels200ResponseInner) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetDefaultAutoArchiveDuration

`func (o *ListGuildChannels200ResponseInner) GetDefaultAutoArchiveDuration() ThreadAutoArchiveDuration`

GetDefaultAutoArchiveDuration returns the DefaultAutoArchiveDuration field if non-nil, zero value otherwise.

### GetDefaultAutoArchiveDurationOk

`func (o *ListGuildChannels200ResponseInner) GetDefaultAutoArchiveDurationOk() (*ThreadAutoArchiveDuration, bool)`

GetDefaultAutoArchiveDurationOk returns a tuple with the DefaultAutoArchiveDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAutoArchiveDuration

`func (o *ListGuildChannels200ResponseInner) SetDefaultAutoArchiveDuration(v ThreadAutoArchiveDuration)`

SetDefaultAutoArchiveDuration sets DefaultAutoArchiveDuration field to given value.

### HasDefaultAutoArchiveDuration

`func (o *ListGuildChannels200ResponseInner) HasDefaultAutoArchiveDuration() bool`

HasDefaultAutoArchiveDuration returns a boolean if a field has been set.

### GetDefaultThreadRateLimitPerUser

`func (o *ListGuildChannels200ResponseInner) GetDefaultThreadRateLimitPerUser() int32`

GetDefaultThreadRateLimitPerUser returns the DefaultThreadRateLimitPerUser field if non-nil, zero value otherwise.

### GetDefaultThreadRateLimitPerUserOk

`func (o *ListGuildChannels200ResponseInner) GetDefaultThreadRateLimitPerUserOk() (*int32, bool)`

GetDefaultThreadRateLimitPerUserOk returns a tuple with the DefaultThreadRateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultThreadRateLimitPerUser

`func (o *ListGuildChannels200ResponseInner) SetDefaultThreadRateLimitPerUser(v int32)`

SetDefaultThreadRateLimitPerUser sets DefaultThreadRateLimitPerUser field to given value.

### HasDefaultThreadRateLimitPerUser

`func (o *ListGuildChannels200ResponseInner) HasDefaultThreadRateLimitPerUser() bool`

HasDefaultThreadRateLimitPerUser returns a boolean if a field has been set.

### SetDefaultThreadRateLimitPerUserNil

`func (o *ListGuildChannels200ResponseInner) SetDefaultThreadRateLimitPerUserNil(b bool)`

 SetDefaultThreadRateLimitPerUserNil sets the value for DefaultThreadRateLimitPerUser to be an explicit nil

### UnsetDefaultThreadRateLimitPerUser
`func (o *ListGuildChannels200ResponseInner) UnsetDefaultThreadRateLimitPerUser()`

UnsetDefaultThreadRateLimitPerUser ensures that no value is present for DefaultThreadRateLimitPerUser, not even an explicit nil
### GetPosition

`func (o *ListGuildChannels200ResponseInner) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ListGuildChannels200ResponseInner) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ListGuildChannels200ResponseInner) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetPermissionOverwrites

`func (o *ListGuildChannels200ResponseInner) GetPermissionOverwrites() []ChannelPermissionOverwriteResponse`

GetPermissionOverwrites returns the PermissionOverwrites field if non-nil, zero value otherwise.

### GetPermissionOverwritesOk

`func (o *ListGuildChannels200ResponseInner) GetPermissionOverwritesOk() (*[]ChannelPermissionOverwriteResponse, bool)`

GetPermissionOverwritesOk returns a tuple with the PermissionOverwrites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionOverwrites

`func (o *ListGuildChannels200ResponseInner) SetPermissionOverwrites(v []ChannelPermissionOverwriteResponse)`

SetPermissionOverwrites sets PermissionOverwrites field to given value.

### HasPermissionOverwrites

`func (o *ListGuildChannels200ResponseInner) HasPermissionOverwrites() bool`

HasPermissionOverwrites returns a boolean if a field has been set.

### SetPermissionOverwritesNil

`func (o *ListGuildChannels200ResponseInner) SetPermissionOverwritesNil(b bool)`

 SetPermissionOverwritesNil sets the value for PermissionOverwrites to be an explicit nil

### UnsetPermissionOverwrites
`func (o *ListGuildChannels200ResponseInner) UnsetPermissionOverwrites()`

UnsetPermissionOverwrites ensures that no value is present for PermissionOverwrites, not even an explicit nil
### GetNsfw

`func (o *ListGuildChannels200ResponseInner) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *ListGuildChannels200ResponseInner) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *ListGuildChannels200ResponseInner) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.

### HasNsfw

`func (o *ListGuildChannels200ResponseInner) HasNsfw() bool`

HasNsfw returns a boolean if a field has been set.

### SetNsfwNil

`func (o *ListGuildChannels200ResponseInner) SetNsfwNil(b bool)`

 SetNsfwNil sets the value for Nsfw to be an explicit nil

### UnsetNsfw
`func (o *ListGuildChannels200ResponseInner) UnsetNsfw()`

UnsetNsfw ensures that no value is present for Nsfw, not even an explicit nil
### GetAvailableTags

`func (o *ListGuildChannels200ResponseInner) GetAvailableTags() []ForumTagResponse`

GetAvailableTags returns the AvailableTags field if non-nil, zero value otherwise.

### GetAvailableTagsOk

`func (o *ListGuildChannels200ResponseInner) GetAvailableTagsOk() (*[]ForumTagResponse, bool)`

GetAvailableTagsOk returns a tuple with the AvailableTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTags

`func (o *ListGuildChannels200ResponseInner) SetAvailableTags(v []ForumTagResponse)`

SetAvailableTags sets AvailableTags field to given value.

### HasAvailableTags

`func (o *ListGuildChannels200ResponseInner) HasAvailableTags() bool`

HasAvailableTags returns a boolean if a field has been set.

### SetAvailableTagsNil

`func (o *ListGuildChannels200ResponseInner) SetAvailableTagsNil(b bool)`

 SetAvailableTagsNil sets the value for AvailableTags to be an explicit nil

### UnsetAvailableTags
`func (o *ListGuildChannels200ResponseInner) UnsetAvailableTags()`

UnsetAvailableTags ensures that no value is present for AvailableTags, not even an explicit nil
### GetDefaultReactionEmoji

`func (o *ListGuildChannels200ResponseInner) GetDefaultReactionEmoji() DefaultReactionEmojiResponse`

GetDefaultReactionEmoji returns the DefaultReactionEmoji field if non-nil, zero value otherwise.

### GetDefaultReactionEmojiOk

`func (o *ListGuildChannels200ResponseInner) GetDefaultReactionEmojiOk() (*DefaultReactionEmojiResponse, bool)`

GetDefaultReactionEmojiOk returns a tuple with the DefaultReactionEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultReactionEmoji

`func (o *ListGuildChannels200ResponseInner) SetDefaultReactionEmoji(v DefaultReactionEmojiResponse)`

SetDefaultReactionEmoji sets DefaultReactionEmoji field to given value.

### HasDefaultReactionEmoji

`func (o *ListGuildChannels200ResponseInner) HasDefaultReactionEmoji() bool`

HasDefaultReactionEmoji returns a boolean if a field has been set.

### GetDefaultSortOrder

`func (o *ListGuildChannels200ResponseInner) GetDefaultSortOrder() ThreadSortOrder`

GetDefaultSortOrder returns the DefaultSortOrder field if non-nil, zero value otherwise.

### GetDefaultSortOrderOk

`func (o *ListGuildChannels200ResponseInner) GetDefaultSortOrderOk() (*ThreadSortOrder, bool)`

GetDefaultSortOrderOk returns a tuple with the DefaultSortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSortOrder

`func (o *ListGuildChannels200ResponseInner) SetDefaultSortOrder(v ThreadSortOrder)`

SetDefaultSortOrder sets DefaultSortOrder field to given value.

### HasDefaultSortOrder

`func (o *ListGuildChannels200ResponseInner) HasDefaultSortOrder() bool`

HasDefaultSortOrder returns a boolean if a field has been set.

### GetDefaultForumLayout

`func (o *ListGuildChannels200ResponseInner) GetDefaultForumLayout() ForumLayout`

GetDefaultForumLayout returns the DefaultForumLayout field if non-nil, zero value otherwise.

### GetDefaultForumLayoutOk

`func (o *ListGuildChannels200ResponseInner) GetDefaultForumLayoutOk() (*ForumLayout, bool)`

GetDefaultForumLayoutOk returns a tuple with the DefaultForumLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForumLayout

`func (o *ListGuildChannels200ResponseInner) SetDefaultForumLayout(v ForumLayout)`

SetDefaultForumLayout sets DefaultForumLayout field to given value.

### HasDefaultForumLayout

`func (o *ListGuildChannels200ResponseInner) HasDefaultForumLayout() bool`

HasDefaultForumLayout returns a boolean if a field has been set.

### GetRecipients

`func (o *ListGuildChannels200ResponseInner) GetRecipients() []UserResponse`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ListGuildChannels200ResponseInner) GetRecipientsOk() (*[]UserResponse, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ListGuildChannels200ResponseInner) SetRecipients(v []UserResponse)`

SetRecipients sets Recipients field to given value.


### GetIcon

`func (o *ListGuildChannels200ResponseInner) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ListGuildChannels200ResponseInner) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ListGuildChannels200ResponseInner) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ListGuildChannels200ResponseInner) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *ListGuildChannels200ResponseInner) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *ListGuildChannels200ResponseInner) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetOwnerId

`func (o *ListGuildChannels200ResponseInner) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ListGuildChannels200ResponseInner) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ListGuildChannels200ResponseInner) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetManaged

`func (o *ListGuildChannels200ResponseInner) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ListGuildChannels200ResponseInner) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ListGuildChannels200ResponseInner) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ListGuildChannels200ResponseInner) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *ListGuildChannels200ResponseInner) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *ListGuildChannels200ResponseInner) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetApplicationId

`func (o *ListGuildChannels200ResponseInner) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ListGuildChannels200ResponseInner) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ListGuildChannels200ResponseInner) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ListGuildChannels200ResponseInner) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetThreadMetadata

`func (o *ListGuildChannels200ResponseInner) GetThreadMetadata() ThreadMetadataResponse`

GetThreadMetadata returns the ThreadMetadata field if non-nil, zero value otherwise.

### GetThreadMetadataOk

`func (o *ListGuildChannels200ResponseInner) GetThreadMetadataOk() (*ThreadMetadataResponse, bool)`

GetThreadMetadataOk returns a tuple with the ThreadMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadMetadata

`func (o *ListGuildChannels200ResponseInner) SetThreadMetadata(v ThreadMetadataResponse)`

SetThreadMetadata sets ThreadMetadata field to given value.

### HasThreadMetadata

`func (o *ListGuildChannels200ResponseInner) HasThreadMetadata() bool`

HasThreadMetadata returns a boolean if a field has been set.

### GetMessageCount

`func (o *ListGuildChannels200ResponseInner) GetMessageCount() int32`

GetMessageCount returns the MessageCount field if non-nil, zero value otherwise.

### GetMessageCountOk

`func (o *ListGuildChannels200ResponseInner) GetMessageCountOk() (*int32, bool)`

GetMessageCountOk returns a tuple with the MessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCount

`func (o *ListGuildChannels200ResponseInner) SetMessageCount(v int32)`

SetMessageCount sets MessageCount field to given value.


### GetMemberCount

`func (o *ListGuildChannels200ResponseInner) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *ListGuildChannels200ResponseInner) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *ListGuildChannels200ResponseInner) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.


### GetTotalMessageSent

`func (o *ListGuildChannels200ResponseInner) GetTotalMessageSent() int32`

GetTotalMessageSent returns the TotalMessageSent field if non-nil, zero value otherwise.

### GetTotalMessageSentOk

`func (o *ListGuildChannels200ResponseInner) GetTotalMessageSentOk() (*int32, bool)`

GetTotalMessageSentOk returns a tuple with the TotalMessageSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMessageSent

`func (o *ListGuildChannels200ResponseInner) SetTotalMessageSent(v int32)`

SetTotalMessageSent sets TotalMessageSent field to given value.


### GetAppliedTags

`func (o *ListGuildChannels200ResponseInner) GetAppliedTags() []string`

GetAppliedTags returns the AppliedTags field if non-nil, zero value otherwise.

### GetAppliedTagsOk

`func (o *ListGuildChannels200ResponseInner) GetAppliedTagsOk() (*[]string, bool)`

GetAppliedTagsOk returns a tuple with the AppliedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTags

`func (o *ListGuildChannels200ResponseInner) SetAppliedTags(v []string)`

SetAppliedTags sets AppliedTags field to given value.

### HasAppliedTags

`func (o *ListGuildChannels200ResponseInner) HasAppliedTags() bool`

HasAppliedTags returns a boolean if a field has been set.

### SetAppliedTagsNil

`func (o *ListGuildChannels200ResponseInner) SetAppliedTagsNil(b bool)`

 SetAppliedTagsNil sets the value for AppliedTags to be an explicit nil

### UnsetAppliedTags
`func (o *ListGuildChannels200ResponseInner) UnsetAppliedTags()`

UnsetAppliedTags ensures that no value is present for AppliedTags, not even an explicit nil
### GetMember

`func (o *ListGuildChannels200ResponseInner) GetMember() ThreadMemberResponse`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ListGuildChannels200ResponseInner) GetMemberOk() (*ThreadMemberResponse, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ListGuildChannels200ResponseInner) SetMember(v ThreadMemberResponse)`

SetMember sets Member field to given value.

### HasMember

`func (o *ListGuildChannels200ResponseInner) HasMember() bool`

HasMember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


