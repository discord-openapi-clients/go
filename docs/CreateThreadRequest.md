# CreateThreadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AutoArchiveDuration** | Pointer to [**ThreadAutoArchiveDuration**](ThreadAutoArchiveDuration.md) |  | [optional] 
**RateLimitPerUser** | Pointer to **NullableInt32** |  | [optional] 
**AppliedTags** | Pointer to **[]string** |  | [optional] 
**Message** | [**BaseCreateMessageCreateRequest**](BaseCreateMessageCreateRequest.md) |  | 
**Type** | Pointer to [**NullableChannelTypes**](ChannelTypes.md) |  | [optional] 
**Invitable** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCreateThreadRequest

`func NewCreateThreadRequest(name string, message BaseCreateMessageCreateRequest, ) *CreateThreadRequest`

NewCreateThreadRequest instantiates a new CreateThreadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateThreadRequestWithDefaults

`func NewCreateThreadRequestWithDefaults() *CreateThreadRequest`

NewCreateThreadRequestWithDefaults instantiates a new CreateThreadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateThreadRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateThreadRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateThreadRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAutoArchiveDuration

`func (o *CreateThreadRequest) GetAutoArchiveDuration() ThreadAutoArchiveDuration`

GetAutoArchiveDuration returns the AutoArchiveDuration field if non-nil, zero value otherwise.

### GetAutoArchiveDurationOk

`func (o *CreateThreadRequest) GetAutoArchiveDurationOk() (*ThreadAutoArchiveDuration, bool)`

GetAutoArchiveDurationOk returns a tuple with the AutoArchiveDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoArchiveDuration

`func (o *CreateThreadRequest) SetAutoArchiveDuration(v ThreadAutoArchiveDuration)`

SetAutoArchiveDuration sets AutoArchiveDuration field to given value.

### HasAutoArchiveDuration

`func (o *CreateThreadRequest) HasAutoArchiveDuration() bool`

HasAutoArchiveDuration returns a boolean if a field has been set.

### GetRateLimitPerUser

`func (o *CreateThreadRequest) GetRateLimitPerUser() int32`

GetRateLimitPerUser returns the RateLimitPerUser field if non-nil, zero value otherwise.

### GetRateLimitPerUserOk

`func (o *CreateThreadRequest) GetRateLimitPerUserOk() (*int32, bool)`

GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerUser

`func (o *CreateThreadRequest) SetRateLimitPerUser(v int32)`

SetRateLimitPerUser sets RateLimitPerUser field to given value.

### HasRateLimitPerUser

`func (o *CreateThreadRequest) HasRateLimitPerUser() bool`

HasRateLimitPerUser returns a boolean if a field has been set.

### SetRateLimitPerUserNil

`func (o *CreateThreadRequest) SetRateLimitPerUserNil(b bool)`

 SetRateLimitPerUserNil sets the value for RateLimitPerUser to be an explicit nil

### UnsetRateLimitPerUser
`func (o *CreateThreadRequest) UnsetRateLimitPerUser()`

UnsetRateLimitPerUser ensures that no value is present for RateLimitPerUser, not even an explicit nil
### GetAppliedTags

`func (o *CreateThreadRequest) GetAppliedTags() []string`

GetAppliedTags returns the AppliedTags field if non-nil, zero value otherwise.

### GetAppliedTagsOk

`func (o *CreateThreadRequest) GetAppliedTagsOk() (*[]string, bool)`

GetAppliedTagsOk returns a tuple with the AppliedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTags

`func (o *CreateThreadRequest) SetAppliedTags(v []string)`

SetAppliedTags sets AppliedTags field to given value.

### HasAppliedTags

`func (o *CreateThreadRequest) HasAppliedTags() bool`

HasAppliedTags returns a boolean if a field has been set.

### SetAppliedTagsNil

`func (o *CreateThreadRequest) SetAppliedTagsNil(b bool)`

 SetAppliedTagsNil sets the value for AppliedTags to be an explicit nil

### UnsetAppliedTags
`func (o *CreateThreadRequest) UnsetAppliedTags()`

UnsetAppliedTags ensures that no value is present for AppliedTags, not even an explicit nil
### GetMessage

`func (o *CreateThreadRequest) GetMessage() BaseCreateMessageCreateRequest`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateThreadRequest) GetMessageOk() (*BaseCreateMessageCreateRequest, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateThreadRequest) SetMessage(v BaseCreateMessageCreateRequest)`

SetMessage sets Message field to given value.


### GetType

`func (o *CreateThreadRequest) GetType() ChannelTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateThreadRequest) GetTypeOk() (*ChannelTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateThreadRequest) SetType(v ChannelTypes)`

SetType sets Type field to given value.

### HasType

`func (o *CreateThreadRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CreateThreadRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CreateThreadRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetInvitable

`func (o *CreateThreadRequest) GetInvitable() bool`

GetInvitable returns the Invitable field if non-nil, zero value otherwise.

### GetInvitableOk

`func (o *CreateThreadRequest) GetInvitableOk() (*bool, bool)`

GetInvitableOk returns a tuple with the Invitable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitable

`func (o *CreateThreadRequest) SetInvitable(v bool)`

SetInvitable sets Invitable field to given value.

### HasInvitable

`func (o *CreateThreadRequest) HasInvitable() bool`

HasInvitable returns a boolean if a field has been set.

### SetInvitableNil

`func (o *CreateThreadRequest) SetInvitableNil(b bool)`

 SetInvitableNil sets the value for Invitable to be an explicit nil

### UnsetInvitable
`func (o *CreateThreadRequest) UnsetInvitable()`

UnsetInvitable ensures that no value is present for Invitable, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


