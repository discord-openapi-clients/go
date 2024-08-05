# SetGuildMfaLevelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | [**GuildMFALevel**](GuildMFALevel.md) |  | 

## Methods

### NewSetGuildMfaLevelRequest

`func NewSetGuildMfaLevelRequest(level GuildMFALevel, ) *SetGuildMfaLevelRequest`

NewSetGuildMfaLevelRequest instantiates a new SetGuildMfaLevelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetGuildMfaLevelRequestWithDefaults

`func NewSetGuildMfaLevelRequestWithDefaults() *SetGuildMfaLevelRequest`

NewSetGuildMfaLevelRequestWithDefaults instantiates a new SetGuildMfaLevelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *SetGuildMfaLevelRequest) GetLevel() GuildMFALevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *SetGuildMfaLevelRequest) GetLevelOk() (*GuildMFALevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *SetGuildMfaLevelRequest) SetLevel(v GuildMFALevel)`

SetLevel sets Level field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


