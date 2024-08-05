# ListGuildIntegrations200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**IntegrationTypes**](IntegrationTypes.md) |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Account** | Pointer to [**AccountResponse**](AccountResponse.md) |  | [optional] 
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**Id** | **string** |  | 
**Application** | [**IntegrationApplicationResponse**](IntegrationApplicationResponse.md) |  | 
**Scopes** | [**[]OAuth2Scopes**](OAuth2Scopes.md) |  | 
**User** | [**UserResponse**](UserResponse.md) |  | 
**Revoked** | Pointer to **NullableBool** |  | [optional] 
**ExpireBehavior** | Pointer to [**IntegrationExpireBehaviorTypes**](IntegrationExpireBehaviorTypes.md) |  | [optional] 
**ExpireGracePeriod** | Pointer to [**IntegrationExpireGracePeriodTypes**](IntegrationExpireGracePeriodTypes.md) |  | [optional] 
**SubscriberCount** | Pointer to **NullableInt32** |  | [optional] 
**SyncedAt** | Pointer to **NullableTime** |  | [optional] 
**RoleId** | Pointer to **string** |  | [optional] 
**Syncing** | Pointer to **NullableBool** |  | [optional] 
**EnableEmoticons** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewListGuildIntegrations200ResponseInner

`func NewListGuildIntegrations200ResponseInner(type_ IntegrationTypes, id string, application IntegrationApplicationResponse, scopes []OAuth2Scopes, user UserResponse, ) *ListGuildIntegrations200ResponseInner`

NewListGuildIntegrations200ResponseInner instantiates a new ListGuildIntegrations200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGuildIntegrations200ResponseInnerWithDefaults

`func NewListGuildIntegrations200ResponseInnerWithDefaults() *ListGuildIntegrations200ResponseInner`

NewListGuildIntegrations200ResponseInnerWithDefaults instantiates a new ListGuildIntegrations200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ListGuildIntegrations200ResponseInner) GetType() IntegrationTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListGuildIntegrations200ResponseInner) GetTypeOk() (*IntegrationTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListGuildIntegrations200ResponseInner) SetType(v IntegrationTypes)`

SetType sets Type field to given value.


### GetName

`func (o *ListGuildIntegrations200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListGuildIntegrations200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListGuildIntegrations200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListGuildIntegrations200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ListGuildIntegrations200ResponseInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ListGuildIntegrations200ResponseInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAccount

`func (o *ListGuildIntegrations200ResponseInner) GetAccount() AccountResponse`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListGuildIntegrations200ResponseInner) GetAccountOk() (*AccountResponse, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListGuildIntegrations200ResponseInner) SetAccount(v AccountResponse)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListGuildIntegrations200ResponseInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetEnabled

`func (o *ListGuildIntegrations200ResponseInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListGuildIntegrations200ResponseInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListGuildIntegrations200ResponseInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListGuildIntegrations200ResponseInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *ListGuildIntegrations200ResponseInner) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *ListGuildIntegrations200ResponseInner) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetId

`func (o *ListGuildIntegrations200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListGuildIntegrations200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListGuildIntegrations200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetApplication

`func (o *ListGuildIntegrations200ResponseInner) GetApplication() IntegrationApplicationResponse`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ListGuildIntegrations200ResponseInner) GetApplicationOk() (*IntegrationApplicationResponse, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ListGuildIntegrations200ResponseInner) SetApplication(v IntegrationApplicationResponse)`

SetApplication sets Application field to given value.


### GetScopes

`func (o *ListGuildIntegrations200ResponseInner) GetScopes() []OAuth2Scopes`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ListGuildIntegrations200ResponseInner) GetScopesOk() (*[]OAuth2Scopes, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ListGuildIntegrations200ResponseInner) SetScopes(v []OAuth2Scopes)`

SetScopes sets Scopes field to given value.


### GetUser

`func (o *ListGuildIntegrations200ResponseInner) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ListGuildIntegrations200ResponseInner) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ListGuildIntegrations200ResponseInner) SetUser(v UserResponse)`

SetUser sets User field to given value.


### GetRevoked

`func (o *ListGuildIntegrations200ResponseInner) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *ListGuildIntegrations200ResponseInner) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *ListGuildIntegrations200ResponseInner) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *ListGuildIntegrations200ResponseInner) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### SetRevokedNil

`func (o *ListGuildIntegrations200ResponseInner) SetRevokedNil(b bool)`

 SetRevokedNil sets the value for Revoked to be an explicit nil

### UnsetRevoked
`func (o *ListGuildIntegrations200ResponseInner) UnsetRevoked()`

UnsetRevoked ensures that no value is present for Revoked, not even an explicit nil
### GetExpireBehavior

`func (o *ListGuildIntegrations200ResponseInner) GetExpireBehavior() IntegrationExpireBehaviorTypes`

GetExpireBehavior returns the ExpireBehavior field if non-nil, zero value otherwise.

### GetExpireBehaviorOk

`func (o *ListGuildIntegrations200ResponseInner) GetExpireBehaviorOk() (*IntegrationExpireBehaviorTypes, bool)`

GetExpireBehaviorOk returns a tuple with the ExpireBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireBehavior

`func (o *ListGuildIntegrations200ResponseInner) SetExpireBehavior(v IntegrationExpireBehaviorTypes)`

SetExpireBehavior sets ExpireBehavior field to given value.

### HasExpireBehavior

`func (o *ListGuildIntegrations200ResponseInner) HasExpireBehavior() bool`

HasExpireBehavior returns a boolean if a field has been set.

### GetExpireGracePeriod

`func (o *ListGuildIntegrations200ResponseInner) GetExpireGracePeriod() IntegrationExpireGracePeriodTypes`

GetExpireGracePeriod returns the ExpireGracePeriod field if non-nil, zero value otherwise.

### GetExpireGracePeriodOk

`func (o *ListGuildIntegrations200ResponseInner) GetExpireGracePeriodOk() (*IntegrationExpireGracePeriodTypes, bool)`

GetExpireGracePeriodOk returns a tuple with the ExpireGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireGracePeriod

`func (o *ListGuildIntegrations200ResponseInner) SetExpireGracePeriod(v IntegrationExpireGracePeriodTypes)`

SetExpireGracePeriod sets ExpireGracePeriod field to given value.

### HasExpireGracePeriod

`func (o *ListGuildIntegrations200ResponseInner) HasExpireGracePeriod() bool`

HasExpireGracePeriod returns a boolean if a field has been set.

### GetSubscriberCount

`func (o *ListGuildIntegrations200ResponseInner) GetSubscriberCount() int32`

GetSubscriberCount returns the SubscriberCount field if non-nil, zero value otherwise.

### GetSubscriberCountOk

`func (o *ListGuildIntegrations200ResponseInner) GetSubscriberCountOk() (*int32, bool)`

GetSubscriberCountOk returns a tuple with the SubscriberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberCount

`func (o *ListGuildIntegrations200ResponseInner) SetSubscriberCount(v int32)`

SetSubscriberCount sets SubscriberCount field to given value.

### HasSubscriberCount

`func (o *ListGuildIntegrations200ResponseInner) HasSubscriberCount() bool`

HasSubscriberCount returns a boolean if a field has been set.

### SetSubscriberCountNil

`func (o *ListGuildIntegrations200ResponseInner) SetSubscriberCountNil(b bool)`

 SetSubscriberCountNil sets the value for SubscriberCount to be an explicit nil

### UnsetSubscriberCount
`func (o *ListGuildIntegrations200ResponseInner) UnsetSubscriberCount()`

UnsetSubscriberCount ensures that no value is present for SubscriberCount, not even an explicit nil
### GetSyncedAt

`func (o *ListGuildIntegrations200ResponseInner) GetSyncedAt() time.Time`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *ListGuildIntegrations200ResponseInner) GetSyncedAtOk() (*time.Time, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *ListGuildIntegrations200ResponseInner) SetSyncedAt(v time.Time)`

SetSyncedAt sets SyncedAt field to given value.

### HasSyncedAt

`func (o *ListGuildIntegrations200ResponseInner) HasSyncedAt() bool`

HasSyncedAt returns a boolean if a field has been set.

### SetSyncedAtNil

`func (o *ListGuildIntegrations200ResponseInner) SetSyncedAtNil(b bool)`

 SetSyncedAtNil sets the value for SyncedAt to be an explicit nil

### UnsetSyncedAt
`func (o *ListGuildIntegrations200ResponseInner) UnsetSyncedAt()`

UnsetSyncedAt ensures that no value is present for SyncedAt, not even an explicit nil
### GetRoleId

`func (o *ListGuildIntegrations200ResponseInner) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ListGuildIntegrations200ResponseInner) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ListGuildIntegrations200ResponseInner) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *ListGuildIntegrations200ResponseInner) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetSyncing

`func (o *ListGuildIntegrations200ResponseInner) GetSyncing() bool`

GetSyncing returns the Syncing field if non-nil, zero value otherwise.

### GetSyncingOk

`func (o *ListGuildIntegrations200ResponseInner) GetSyncingOk() (*bool, bool)`

GetSyncingOk returns a tuple with the Syncing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncing

`func (o *ListGuildIntegrations200ResponseInner) SetSyncing(v bool)`

SetSyncing sets Syncing field to given value.

### HasSyncing

`func (o *ListGuildIntegrations200ResponseInner) HasSyncing() bool`

HasSyncing returns a boolean if a field has been set.

### SetSyncingNil

`func (o *ListGuildIntegrations200ResponseInner) SetSyncingNil(b bool)`

 SetSyncingNil sets the value for Syncing to be an explicit nil

### UnsetSyncing
`func (o *ListGuildIntegrations200ResponseInner) UnsetSyncing()`

UnsetSyncing ensures that no value is present for Syncing, not even an explicit nil
### GetEnableEmoticons

`func (o *ListGuildIntegrations200ResponseInner) GetEnableEmoticons() bool`

GetEnableEmoticons returns the EnableEmoticons field if non-nil, zero value otherwise.

### GetEnableEmoticonsOk

`func (o *ListGuildIntegrations200ResponseInner) GetEnableEmoticonsOk() (*bool, bool)`

GetEnableEmoticonsOk returns a tuple with the EnableEmoticons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmoticons

`func (o *ListGuildIntegrations200ResponseInner) SetEnableEmoticons(v bool)`

SetEnableEmoticons sets EnableEmoticons field to given value.

### HasEnableEmoticons

`func (o *ListGuildIntegrations200ResponseInner) HasEnableEmoticons() bool`

HasEnableEmoticons returns a boolean if a field has been set.

### SetEnableEmoticonsNil

`func (o *ListGuildIntegrations200ResponseInner) SetEnableEmoticonsNil(b bool)`

 SetEnableEmoticonsNil sets the value for EnableEmoticons to be an explicit nil

### UnsetEnableEmoticons
`func (o *ListGuildIntegrations200ResponseInner) UnsetEnableEmoticons()`

UnsetEnableEmoticons ensures that no value is present for EnableEmoticons, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


