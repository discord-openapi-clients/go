# UpdateAutoModerationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to [**AutomodEventType**](AutomodEventType.md) |  | [optional] 
**Actions** | Pointer to [**[]DefaultKeywordListUpsertRequestActionsInner**](DefaultKeywordListUpsertRequestActionsInner.md) |  | [optional] 
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**ExemptRoles** | Pointer to **[]string** |  | [optional] 
**ExemptChannels** | Pointer to **[]string** |  | [optional] 
**TriggerType** | Pointer to [**AutomodTriggerType**](AutomodTriggerType.md) |  | [optional] 
**TriggerMetadata** | Pointer to [**MentionSpamTriggerMetadata**](MentionSpamTriggerMetadata.md) |  | [optional] 

## Methods

### NewUpdateAutoModerationRuleRequest

`func NewUpdateAutoModerationRuleRequest() *UpdateAutoModerationRuleRequest`

NewUpdateAutoModerationRuleRequest instantiates a new UpdateAutoModerationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAutoModerationRuleRequestWithDefaults

`func NewUpdateAutoModerationRuleRequestWithDefaults() *UpdateAutoModerationRuleRequest`

NewUpdateAutoModerationRuleRequestWithDefaults instantiates a new UpdateAutoModerationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAutoModerationRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAutoModerationRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAutoModerationRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAutoModerationRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEventType

`func (o *UpdateAutoModerationRuleRequest) GetEventType() AutomodEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *UpdateAutoModerationRuleRequest) GetEventTypeOk() (*AutomodEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *UpdateAutoModerationRuleRequest) SetEventType(v AutomodEventType)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *UpdateAutoModerationRuleRequest) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetActions

`func (o *UpdateAutoModerationRuleRequest) GetActions() []DefaultKeywordListUpsertRequestActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *UpdateAutoModerationRuleRequest) GetActionsOk() (*[]DefaultKeywordListUpsertRequestActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *UpdateAutoModerationRuleRequest) SetActions(v []DefaultKeywordListUpsertRequestActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *UpdateAutoModerationRuleRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *UpdateAutoModerationRuleRequest) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *UpdateAutoModerationRuleRequest) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetEnabled

`func (o *UpdateAutoModerationRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateAutoModerationRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateAutoModerationRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateAutoModerationRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *UpdateAutoModerationRuleRequest) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *UpdateAutoModerationRuleRequest) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetExemptRoles

`func (o *UpdateAutoModerationRuleRequest) GetExemptRoles() []string`

GetExemptRoles returns the ExemptRoles field if non-nil, zero value otherwise.

### GetExemptRolesOk

`func (o *UpdateAutoModerationRuleRequest) GetExemptRolesOk() (*[]string, bool)`

GetExemptRolesOk returns a tuple with the ExemptRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptRoles

`func (o *UpdateAutoModerationRuleRequest) SetExemptRoles(v []string)`

SetExemptRoles sets ExemptRoles field to given value.

### HasExemptRoles

`func (o *UpdateAutoModerationRuleRequest) HasExemptRoles() bool`

HasExemptRoles returns a boolean if a field has been set.

### SetExemptRolesNil

`func (o *UpdateAutoModerationRuleRequest) SetExemptRolesNil(b bool)`

 SetExemptRolesNil sets the value for ExemptRoles to be an explicit nil

### UnsetExemptRoles
`func (o *UpdateAutoModerationRuleRequest) UnsetExemptRoles()`

UnsetExemptRoles ensures that no value is present for ExemptRoles, not even an explicit nil
### GetExemptChannels

`func (o *UpdateAutoModerationRuleRequest) GetExemptChannels() []string`

GetExemptChannels returns the ExemptChannels field if non-nil, zero value otherwise.

### GetExemptChannelsOk

`func (o *UpdateAutoModerationRuleRequest) GetExemptChannelsOk() (*[]string, bool)`

GetExemptChannelsOk returns a tuple with the ExemptChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptChannels

`func (o *UpdateAutoModerationRuleRequest) SetExemptChannels(v []string)`

SetExemptChannels sets ExemptChannels field to given value.

### HasExemptChannels

`func (o *UpdateAutoModerationRuleRequest) HasExemptChannels() bool`

HasExemptChannels returns a boolean if a field has been set.

### SetExemptChannelsNil

`func (o *UpdateAutoModerationRuleRequest) SetExemptChannelsNil(b bool)`

 SetExemptChannelsNil sets the value for ExemptChannels to be an explicit nil

### UnsetExemptChannels
`func (o *UpdateAutoModerationRuleRequest) UnsetExemptChannels()`

UnsetExemptChannels ensures that no value is present for ExemptChannels, not even an explicit nil
### GetTriggerType

`func (o *UpdateAutoModerationRuleRequest) GetTriggerType() AutomodTriggerType`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *UpdateAutoModerationRuleRequest) GetTriggerTypeOk() (*AutomodTriggerType, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *UpdateAutoModerationRuleRequest) SetTriggerType(v AutomodTriggerType)`

SetTriggerType sets TriggerType field to given value.

### HasTriggerType

`func (o *UpdateAutoModerationRuleRequest) HasTriggerType() bool`

HasTriggerType returns a boolean if a field has been set.

### GetTriggerMetadata

`func (o *UpdateAutoModerationRuleRequest) GetTriggerMetadata() MentionSpamTriggerMetadata`

GetTriggerMetadata returns the TriggerMetadata field if non-nil, zero value otherwise.

### GetTriggerMetadataOk

`func (o *UpdateAutoModerationRuleRequest) GetTriggerMetadataOk() (*MentionSpamTriggerMetadata, bool)`

GetTriggerMetadataOk returns a tuple with the TriggerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerMetadata

`func (o *UpdateAutoModerationRuleRequest) SetTriggerMetadata(v MentionSpamTriggerMetadata)`

SetTriggerMetadata sets TriggerMetadata field to given value.

### HasTriggerMetadata

`func (o *UpdateAutoModerationRuleRequest) HasTriggerMetadata() bool`

HasTriggerMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


