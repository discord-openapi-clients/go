# ExecuteWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **NullableString** |  | [optional] 
**Embeds** | Pointer to [**[]RichEmbed**](RichEmbed.md) |  | [optional] 
**AllowedMentions** | Pointer to [**MessageAllowedMentionsRequest**](MessageAllowedMentionsRequest.md) |  | [optional] 
**Components** | Pointer to [**[]ActionRow**](ActionRow.md) |  | [optional] 
**Attachments** | Pointer to [**[]MessageAttachmentRequest**](MessageAttachmentRequest.md) |  | [optional] 
**Tts** | Pointer to **NullableBool** |  | [optional] 
**Flags** | Pointer to **NullableInt32** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**AvatarUrl** | Pointer to **NullableString** |  | [optional] 
**ThreadName** | Pointer to **NullableString** |  | [optional] 
**AppliedTags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewExecuteWebhookRequest

`func NewExecuteWebhookRequest() *ExecuteWebhookRequest`

NewExecuteWebhookRequest instantiates a new ExecuteWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteWebhookRequestWithDefaults

`func NewExecuteWebhookRequestWithDefaults() *ExecuteWebhookRequest`

NewExecuteWebhookRequestWithDefaults instantiates a new ExecuteWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ExecuteWebhookRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ExecuteWebhookRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ExecuteWebhookRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ExecuteWebhookRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *ExecuteWebhookRequest) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *ExecuteWebhookRequest) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetEmbeds

`func (o *ExecuteWebhookRequest) GetEmbeds() []RichEmbed`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *ExecuteWebhookRequest) GetEmbedsOk() (*[]RichEmbed, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *ExecuteWebhookRequest) SetEmbeds(v []RichEmbed)`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *ExecuteWebhookRequest) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.

### SetEmbedsNil

`func (o *ExecuteWebhookRequest) SetEmbedsNil(b bool)`

 SetEmbedsNil sets the value for Embeds to be an explicit nil

### UnsetEmbeds
`func (o *ExecuteWebhookRequest) UnsetEmbeds()`

UnsetEmbeds ensures that no value is present for Embeds, not even an explicit nil
### GetAllowedMentions

`func (o *ExecuteWebhookRequest) GetAllowedMentions() MessageAllowedMentionsRequest`

GetAllowedMentions returns the AllowedMentions field if non-nil, zero value otherwise.

### GetAllowedMentionsOk

`func (o *ExecuteWebhookRequest) GetAllowedMentionsOk() (*MessageAllowedMentionsRequest, bool)`

GetAllowedMentionsOk returns a tuple with the AllowedMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMentions

`func (o *ExecuteWebhookRequest) SetAllowedMentions(v MessageAllowedMentionsRequest)`

SetAllowedMentions sets AllowedMentions field to given value.

### HasAllowedMentions

`func (o *ExecuteWebhookRequest) HasAllowedMentions() bool`

HasAllowedMentions returns a boolean if a field has been set.

### GetComponents

`func (o *ExecuteWebhookRequest) GetComponents() []ActionRow`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ExecuteWebhookRequest) GetComponentsOk() (*[]ActionRow, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ExecuteWebhookRequest) SetComponents(v []ActionRow)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ExecuteWebhookRequest) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### SetComponentsNil

`func (o *ExecuteWebhookRequest) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *ExecuteWebhookRequest) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil
### GetAttachments

`func (o *ExecuteWebhookRequest) GetAttachments() []MessageAttachmentRequest`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ExecuteWebhookRequest) GetAttachmentsOk() (*[]MessageAttachmentRequest, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ExecuteWebhookRequest) SetAttachments(v []MessageAttachmentRequest)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ExecuteWebhookRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *ExecuteWebhookRequest) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *ExecuteWebhookRequest) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetTts

`func (o *ExecuteWebhookRequest) GetTts() bool`

GetTts returns the Tts field if non-nil, zero value otherwise.

### GetTtsOk

`func (o *ExecuteWebhookRequest) GetTtsOk() (*bool, bool)`

GetTtsOk returns a tuple with the Tts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTts

`func (o *ExecuteWebhookRequest) SetTts(v bool)`

SetTts sets Tts field to given value.

### HasTts

`func (o *ExecuteWebhookRequest) HasTts() bool`

HasTts returns a boolean if a field has been set.

### SetTtsNil

`func (o *ExecuteWebhookRequest) SetTtsNil(b bool)`

 SetTtsNil sets the value for Tts to be an explicit nil

### UnsetTts
`func (o *ExecuteWebhookRequest) UnsetTts()`

UnsetTts ensures that no value is present for Tts, not even an explicit nil
### GetFlags

`func (o *ExecuteWebhookRequest) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ExecuteWebhookRequest) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ExecuteWebhookRequest) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *ExecuteWebhookRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### SetFlagsNil

`func (o *ExecuteWebhookRequest) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *ExecuteWebhookRequest) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil
### GetUsername

`func (o *ExecuteWebhookRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ExecuteWebhookRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ExecuteWebhookRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ExecuteWebhookRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ExecuteWebhookRequest) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ExecuteWebhookRequest) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetAvatarUrl

`func (o *ExecuteWebhookRequest) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *ExecuteWebhookRequest) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *ExecuteWebhookRequest) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *ExecuteWebhookRequest) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *ExecuteWebhookRequest) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *ExecuteWebhookRequest) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetThreadName

`func (o *ExecuteWebhookRequest) GetThreadName() string`

GetThreadName returns the ThreadName field if non-nil, zero value otherwise.

### GetThreadNameOk

`func (o *ExecuteWebhookRequest) GetThreadNameOk() (*string, bool)`

GetThreadNameOk returns a tuple with the ThreadName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadName

`func (o *ExecuteWebhookRequest) SetThreadName(v string)`

SetThreadName sets ThreadName field to given value.

### HasThreadName

`func (o *ExecuteWebhookRequest) HasThreadName() bool`

HasThreadName returns a boolean if a field has been set.

### SetThreadNameNil

`func (o *ExecuteWebhookRequest) SetThreadNameNil(b bool)`

 SetThreadNameNil sets the value for ThreadName to be an explicit nil

### UnsetThreadName
`func (o *ExecuteWebhookRequest) UnsetThreadName()`

UnsetThreadName ensures that no value is present for ThreadName, not even an explicit nil
### GetAppliedTags

`func (o *ExecuteWebhookRequest) GetAppliedTags() []string`

GetAppliedTags returns the AppliedTags field if non-nil, zero value otherwise.

### GetAppliedTagsOk

`func (o *ExecuteWebhookRequest) GetAppliedTagsOk() (*[]string, bool)`

GetAppliedTagsOk returns a tuple with the AppliedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTags

`func (o *ExecuteWebhookRequest) SetAppliedTags(v []string)`

SetAppliedTags sets AppliedTags field to given value.

### HasAppliedTags

`func (o *ExecuteWebhookRequest) HasAppliedTags() bool`

HasAppliedTags returns a boolean if a field has been set.

### SetAppliedTagsNil

`func (o *ExecuteWebhookRequest) SetAppliedTagsNil(b bool)`

 SetAppliedTagsNil sets the value for AppliedTags to be an explicit nil

### UnsetAppliedTags
`func (o *ExecuteWebhookRequest) UnsetAppliedTags()`

UnsetAppliedTags ensures that no value is present for AppliedTags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


