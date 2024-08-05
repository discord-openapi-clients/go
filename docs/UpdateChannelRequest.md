# UpdateChannelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**NullableChannelTypes**](ChannelTypes.md) |  | [optional] 
**Position** | Pointer to **NullableInt32** |  | [optional] 
**Topic** | Pointer to **NullableString** |  | [optional] 
**Bitrate** | Pointer to **NullableInt32** |  | [optional] 
**UserLimit** | Pointer to **NullableInt32** |  | [optional] 
**Nsfw** | Pointer to **NullableBool** |  | [optional] 
**RateLimitPerUser** | Pointer to **NullableInt32** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**PermissionOverwrites** | Pointer to [**[]ChannelPermissionOverwriteRequest**](ChannelPermissionOverwriteRequest.md) |  | [optional] 
**RtcRegion** | Pointer to **NullableString** |  | [optional] 
**VideoQualityMode** | Pointer to [**VideoQualityModes**](VideoQualityModes.md) |  | [optional] 
**DefaultAutoArchiveDuration** | Pointer to [**ThreadAutoArchiveDuration**](ThreadAutoArchiveDuration.md) |  | [optional] 
**DefaultReactionEmoji** | Pointer to [**UpdateDefaultReactionEmojiRequest**](UpdateDefaultReactionEmojiRequest.md) |  | [optional] 
**DefaultThreadRateLimitPerUser** | Pointer to **NullableInt32** |  | [optional] 
**DefaultSortOrder** | Pointer to [**ThreadSortOrder**](ThreadSortOrder.md) |  | [optional] 
**DefaultForumLayout** | Pointer to [**ForumLayout**](ForumLayout.md) |  | [optional] 
**Flags** | Pointer to **NullableInt32** |  | [optional] 
**AvailableTags** | Pointer to [**[]UpdateThreadTagRequest**](UpdateThreadTagRequest.md) |  | [optional] 
**Archived** | Pointer to **NullableBool** |  | [optional] 
**Locked** | Pointer to **NullableBool** |  | [optional] 
**Invitable** | Pointer to **NullableBool** |  | [optional] 
**AutoArchiveDuration** | Pointer to [**ThreadAutoArchiveDuration**](ThreadAutoArchiveDuration.md) |  | [optional] 
**AppliedTags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateChannelRequest

`func NewUpdateChannelRequest() *UpdateChannelRequest`

NewUpdateChannelRequest instantiates a new UpdateChannelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateChannelRequestWithDefaults

`func NewUpdateChannelRequestWithDefaults() *UpdateChannelRequest`

NewUpdateChannelRequestWithDefaults instantiates a new UpdateChannelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateChannelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateChannelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateChannelRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateChannelRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateChannelRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateChannelRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIcon

`func (o *UpdateChannelRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *UpdateChannelRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *UpdateChannelRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *UpdateChannelRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *UpdateChannelRequest) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *UpdateChannelRequest) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetType

`func (o *UpdateChannelRequest) GetType() ChannelTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateChannelRequest) GetTypeOk() (*ChannelTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateChannelRequest) SetType(v ChannelTypes)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateChannelRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *UpdateChannelRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *UpdateChannelRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetPosition

`func (o *UpdateChannelRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *UpdateChannelRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *UpdateChannelRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *UpdateChannelRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *UpdateChannelRequest) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *UpdateChannelRequest) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetTopic

`func (o *UpdateChannelRequest) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *UpdateChannelRequest) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *UpdateChannelRequest) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *UpdateChannelRequest) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *UpdateChannelRequest) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *UpdateChannelRequest) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetBitrate

`func (o *UpdateChannelRequest) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *UpdateChannelRequest) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *UpdateChannelRequest) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *UpdateChannelRequest) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *UpdateChannelRequest) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *UpdateChannelRequest) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetUserLimit

`func (o *UpdateChannelRequest) GetUserLimit() int32`

GetUserLimit returns the UserLimit field if non-nil, zero value otherwise.

### GetUserLimitOk

`func (o *UpdateChannelRequest) GetUserLimitOk() (*int32, bool)`

GetUserLimitOk returns a tuple with the UserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimit

`func (o *UpdateChannelRequest) SetUserLimit(v int32)`

SetUserLimit sets UserLimit field to given value.

### HasUserLimit

`func (o *UpdateChannelRequest) HasUserLimit() bool`

HasUserLimit returns a boolean if a field has been set.

### SetUserLimitNil

`func (o *UpdateChannelRequest) SetUserLimitNil(b bool)`

 SetUserLimitNil sets the value for UserLimit to be an explicit nil

### UnsetUserLimit
`func (o *UpdateChannelRequest) UnsetUserLimit()`

UnsetUserLimit ensures that no value is present for UserLimit, not even an explicit nil
### GetNsfw

`func (o *UpdateChannelRequest) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *UpdateChannelRequest) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *UpdateChannelRequest) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.

### HasNsfw

`func (o *UpdateChannelRequest) HasNsfw() bool`

HasNsfw returns a boolean if a field has been set.

### SetNsfwNil

`func (o *UpdateChannelRequest) SetNsfwNil(b bool)`

 SetNsfwNil sets the value for Nsfw to be an explicit nil

### UnsetNsfw
`func (o *UpdateChannelRequest) UnsetNsfw()`

UnsetNsfw ensures that no value is present for Nsfw, not even an explicit nil
### GetRateLimitPerUser

`func (o *UpdateChannelRequest) GetRateLimitPerUser() int32`

GetRateLimitPerUser returns the RateLimitPerUser field if non-nil, zero value otherwise.

### GetRateLimitPerUserOk

`func (o *UpdateChannelRequest) GetRateLimitPerUserOk() (*int32, bool)`

GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerUser

`func (o *UpdateChannelRequest) SetRateLimitPerUser(v int32)`

SetRateLimitPerUser sets RateLimitPerUser field to given value.

### HasRateLimitPerUser

`func (o *UpdateChannelRequest) HasRateLimitPerUser() bool`

HasRateLimitPerUser returns a boolean if a field has been set.

### SetRateLimitPerUserNil

`func (o *UpdateChannelRequest) SetRateLimitPerUserNil(b bool)`

 SetRateLimitPerUserNil sets the value for RateLimitPerUser to be an explicit nil

### UnsetRateLimitPerUser
`func (o *UpdateChannelRequest) UnsetRateLimitPerUser()`

UnsetRateLimitPerUser ensures that no value is present for RateLimitPerUser, not even an explicit nil
### GetParentId

`func (o *UpdateChannelRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *UpdateChannelRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *UpdateChannelRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *UpdateChannelRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPermissionOverwrites

`func (o *UpdateChannelRequest) GetPermissionOverwrites() []ChannelPermissionOverwriteRequest`

GetPermissionOverwrites returns the PermissionOverwrites field if non-nil, zero value otherwise.

### GetPermissionOverwritesOk

`func (o *UpdateChannelRequest) GetPermissionOverwritesOk() (*[]ChannelPermissionOverwriteRequest, bool)`

GetPermissionOverwritesOk returns a tuple with the PermissionOverwrites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionOverwrites

`func (o *UpdateChannelRequest) SetPermissionOverwrites(v []ChannelPermissionOverwriteRequest)`

SetPermissionOverwrites sets PermissionOverwrites field to given value.

### HasPermissionOverwrites

`func (o *UpdateChannelRequest) HasPermissionOverwrites() bool`

HasPermissionOverwrites returns a boolean if a field has been set.

### SetPermissionOverwritesNil

`func (o *UpdateChannelRequest) SetPermissionOverwritesNil(b bool)`

 SetPermissionOverwritesNil sets the value for PermissionOverwrites to be an explicit nil

### UnsetPermissionOverwrites
`func (o *UpdateChannelRequest) UnsetPermissionOverwrites()`

UnsetPermissionOverwrites ensures that no value is present for PermissionOverwrites, not even an explicit nil
### GetRtcRegion

`func (o *UpdateChannelRequest) GetRtcRegion() string`

GetRtcRegion returns the RtcRegion field if non-nil, zero value otherwise.

### GetRtcRegionOk

`func (o *UpdateChannelRequest) GetRtcRegionOk() (*string, bool)`

GetRtcRegionOk returns a tuple with the RtcRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtcRegion

`func (o *UpdateChannelRequest) SetRtcRegion(v string)`

SetRtcRegion sets RtcRegion field to given value.

### HasRtcRegion

`func (o *UpdateChannelRequest) HasRtcRegion() bool`

HasRtcRegion returns a boolean if a field has been set.

### SetRtcRegionNil

`func (o *UpdateChannelRequest) SetRtcRegionNil(b bool)`

 SetRtcRegionNil sets the value for RtcRegion to be an explicit nil

### UnsetRtcRegion
`func (o *UpdateChannelRequest) UnsetRtcRegion()`

UnsetRtcRegion ensures that no value is present for RtcRegion, not even an explicit nil
### GetVideoQualityMode

`func (o *UpdateChannelRequest) GetVideoQualityMode() VideoQualityModes`

GetVideoQualityMode returns the VideoQualityMode field if non-nil, zero value otherwise.

### GetVideoQualityModeOk

`func (o *UpdateChannelRequest) GetVideoQualityModeOk() (*VideoQualityModes, bool)`

GetVideoQualityModeOk returns a tuple with the VideoQualityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQualityMode

`func (o *UpdateChannelRequest) SetVideoQualityMode(v VideoQualityModes)`

SetVideoQualityMode sets VideoQualityMode field to given value.

### HasVideoQualityMode

`func (o *UpdateChannelRequest) HasVideoQualityMode() bool`

HasVideoQualityMode returns a boolean if a field has been set.

### GetDefaultAutoArchiveDuration

`func (o *UpdateChannelRequest) GetDefaultAutoArchiveDuration() ThreadAutoArchiveDuration`

GetDefaultAutoArchiveDuration returns the DefaultAutoArchiveDuration field if non-nil, zero value otherwise.

### GetDefaultAutoArchiveDurationOk

`func (o *UpdateChannelRequest) GetDefaultAutoArchiveDurationOk() (*ThreadAutoArchiveDuration, bool)`

GetDefaultAutoArchiveDurationOk returns a tuple with the DefaultAutoArchiveDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAutoArchiveDuration

`func (o *UpdateChannelRequest) SetDefaultAutoArchiveDuration(v ThreadAutoArchiveDuration)`

SetDefaultAutoArchiveDuration sets DefaultAutoArchiveDuration field to given value.

### HasDefaultAutoArchiveDuration

`func (o *UpdateChannelRequest) HasDefaultAutoArchiveDuration() bool`

HasDefaultAutoArchiveDuration returns a boolean if a field has been set.

### GetDefaultReactionEmoji

`func (o *UpdateChannelRequest) GetDefaultReactionEmoji() UpdateDefaultReactionEmojiRequest`

GetDefaultReactionEmoji returns the DefaultReactionEmoji field if non-nil, zero value otherwise.

### GetDefaultReactionEmojiOk

`func (o *UpdateChannelRequest) GetDefaultReactionEmojiOk() (*UpdateDefaultReactionEmojiRequest, bool)`

GetDefaultReactionEmojiOk returns a tuple with the DefaultReactionEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultReactionEmoji

`func (o *UpdateChannelRequest) SetDefaultReactionEmoji(v UpdateDefaultReactionEmojiRequest)`

SetDefaultReactionEmoji sets DefaultReactionEmoji field to given value.

### HasDefaultReactionEmoji

`func (o *UpdateChannelRequest) HasDefaultReactionEmoji() bool`

HasDefaultReactionEmoji returns a boolean if a field has been set.

### GetDefaultThreadRateLimitPerUser

`func (o *UpdateChannelRequest) GetDefaultThreadRateLimitPerUser() int32`

GetDefaultThreadRateLimitPerUser returns the DefaultThreadRateLimitPerUser field if non-nil, zero value otherwise.

### GetDefaultThreadRateLimitPerUserOk

`func (o *UpdateChannelRequest) GetDefaultThreadRateLimitPerUserOk() (*int32, bool)`

GetDefaultThreadRateLimitPerUserOk returns a tuple with the DefaultThreadRateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultThreadRateLimitPerUser

`func (o *UpdateChannelRequest) SetDefaultThreadRateLimitPerUser(v int32)`

SetDefaultThreadRateLimitPerUser sets DefaultThreadRateLimitPerUser field to given value.

### HasDefaultThreadRateLimitPerUser

`func (o *UpdateChannelRequest) HasDefaultThreadRateLimitPerUser() bool`

HasDefaultThreadRateLimitPerUser returns a boolean if a field has been set.

### SetDefaultThreadRateLimitPerUserNil

`func (o *UpdateChannelRequest) SetDefaultThreadRateLimitPerUserNil(b bool)`

 SetDefaultThreadRateLimitPerUserNil sets the value for DefaultThreadRateLimitPerUser to be an explicit nil

### UnsetDefaultThreadRateLimitPerUser
`func (o *UpdateChannelRequest) UnsetDefaultThreadRateLimitPerUser()`

UnsetDefaultThreadRateLimitPerUser ensures that no value is present for DefaultThreadRateLimitPerUser, not even an explicit nil
### GetDefaultSortOrder

`func (o *UpdateChannelRequest) GetDefaultSortOrder() ThreadSortOrder`

GetDefaultSortOrder returns the DefaultSortOrder field if non-nil, zero value otherwise.

### GetDefaultSortOrderOk

`func (o *UpdateChannelRequest) GetDefaultSortOrderOk() (*ThreadSortOrder, bool)`

GetDefaultSortOrderOk returns a tuple with the DefaultSortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSortOrder

`func (o *UpdateChannelRequest) SetDefaultSortOrder(v ThreadSortOrder)`

SetDefaultSortOrder sets DefaultSortOrder field to given value.

### HasDefaultSortOrder

`func (o *UpdateChannelRequest) HasDefaultSortOrder() bool`

HasDefaultSortOrder returns a boolean if a field has been set.

### GetDefaultForumLayout

`func (o *UpdateChannelRequest) GetDefaultForumLayout() ForumLayout`

GetDefaultForumLayout returns the DefaultForumLayout field if non-nil, zero value otherwise.

### GetDefaultForumLayoutOk

`func (o *UpdateChannelRequest) GetDefaultForumLayoutOk() (*ForumLayout, bool)`

GetDefaultForumLayoutOk returns a tuple with the DefaultForumLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForumLayout

`func (o *UpdateChannelRequest) SetDefaultForumLayout(v ForumLayout)`

SetDefaultForumLayout sets DefaultForumLayout field to given value.

### HasDefaultForumLayout

`func (o *UpdateChannelRequest) HasDefaultForumLayout() bool`

HasDefaultForumLayout returns a boolean if a field has been set.

### GetFlags

`func (o *UpdateChannelRequest) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *UpdateChannelRequest) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *UpdateChannelRequest) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *UpdateChannelRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### SetFlagsNil

`func (o *UpdateChannelRequest) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *UpdateChannelRequest) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil
### GetAvailableTags

`func (o *UpdateChannelRequest) GetAvailableTags() []UpdateThreadTagRequest`

GetAvailableTags returns the AvailableTags field if non-nil, zero value otherwise.

### GetAvailableTagsOk

`func (o *UpdateChannelRequest) GetAvailableTagsOk() (*[]UpdateThreadTagRequest, bool)`

GetAvailableTagsOk returns a tuple with the AvailableTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTags

`func (o *UpdateChannelRequest) SetAvailableTags(v []UpdateThreadTagRequest)`

SetAvailableTags sets AvailableTags field to given value.

### HasAvailableTags

`func (o *UpdateChannelRequest) HasAvailableTags() bool`

HasAvailableTags returns a boolean if a field has been set.

### SetAvailableTagsNil

`func (o *UpdateChannelRequest) SetAvailableTagsNil(b bool)`

 SetAvailableTagsNil sets the value for AvailableTags to be an explicit nil

### UnsetAvailableTags
`func (o *UpdateChannelRequest) UnsetAvailableTags()`

UnsetAvailableTags ensures that no value is present for AvailableTags, not even an explicit nil
### GetArchived

`func (o *UpdateChannelRequest) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UpdateChannelRequest) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UpdateChannelRequest) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UpdateChannelRequest) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### SetArchivedNil

`func (o *UpdateChannelRequest) SetArchivedNil(b bool)`

 SetArchivedNil sets the value for Archived to be an explicit nil

### UnsetArchived
`func (o *UpdateChannelRequest) UnsetArchived()`

UnsetArchived ensures that no value is present for Archived, not even an explicit nil
### GetLocked

`func (o *UpdateChannelRequest) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *UpdateChannelRequest) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *UpdateChannelRequest) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *UpdateChannelRequest) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLockedNil

`func (o *UpdateChannelRequest) SetLockedNil(b bool)`

 SetLockedNil sets the value for Locked to be an explicit nil

### UnsetLocked
`func (o *UpdateChannelRequest) UnsetLocked()`

UnsetLocked ensures that no value is present for Locked, not even an explicit nil
### GetInvitable

`func (o *UpdateChannelRequest) GetInvitable() bool`

GetInvitable returns the Invitable field if non-nil, zero value otherwise.

### GetInvitableOk

`func (o *UpdateChannelRequest) GetInvitableOk() (*bool, bool)`

GetInvitableOk returns a tuple with the Invitable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitable

`func (o *UpdateChannelRequest) SetInvitable(v bool)`

SetInvitable sets Invitable field to given value.

### HasInvitable

`func (o *UpdateChannelRequest) HasInvitable() bool`

HasInvitable returns a boolean if a field has been set.

### SetInvitableNil

`func (o *UpdateChannelRequest) SetInvitableNil(b bool)`

 SetInvitableNil sets the value for Invitable to be an explicit nil

### UnsetInvitable
`func (o *UpdateChannelRequest) UnsetInvitable()`

UnsetInvitable ensures that no value is present for Invitable, not even an explicit nil
### GetAutoArchiveDuration

`func (o *UpdateChannelRequest) GetAutoArchiveDuration() ThreadAutoArchiveDuration`

GetAutoArchiveDuration returns the AutoArchiveDuration field if non-nil, zero value otherwise.

### GetAutoArchiveDurationOk

`func (o *UpdateChannelRequest) GetAutoArchiveDurationOk() (*ThreadAutoArchiveDuration, bool)`

GetAutoArchiveDurationOk returns a tuple with the AutoArchiveDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoArchiveDuration

`func (o *UpdateChannelRequest) SetAutoArchiveDuration(v ThreadAutoArchiveDuration)`

SetAutoArchiveDuration sets AutoArchiveDuration field to given value.

### HasAutoArchiveDuration

`func (o *UpdateChannelRequest) HasAutoArchiveDuration() bool`

HasAutoArchiveDuration returns a boolean if a field has been set.

### GetAppliedTags

`func (o *UpdateChannelRequest) GetAppliedTags() []string`

GetAppliedTags returns the AppliedTags field if non-nil, zero value otherwise.

### GetAppliedTagsOk

`func (o *UpdateChannelRequest) GetAppliedTagsOk() (*[]string, bool)`

GetAppliedTagsOk returns a tuple with the AppliedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTags

`func (o *UpdateChannelRequest) SetAppliedTags(v []string)`

SetAppliedTags sets AppliedTags field to given value.

### HasAppliedTags

`func (o *UpdateChannelRequest) HasAppliedTags() bool`

HasAppliedTags returns a boolean if a field has been set.

### SetAppliedTagsNil

`func (o *UpdateChannelRequest) SetAppliedTagsNil(b bool)`

 SetAppliedTagsNil sets the value for AppliedTags to be an explicit nil

### UnsetAppliedTags
`func (o *UpdateChannelRequest) UnsetAppliedTags()`

UnsetAppliedTags ensures that no value is present for AppliedTags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


