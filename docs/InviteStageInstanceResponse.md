# InviteStageInstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | **string** |  | 
**ParticipantCount** | Pointer to **NullableInt32** |  | [optional] 
**SpeakerCount** | Pointer to **NullableInt32** |  | [optional] 
**Members** | Pointer to [**[]GuildMemberResponse**](GuildMemberResponse.md) |  | [optional] 

## Methods

### NewInviteStageInstanceResponse

`func NewInviteStageInstanceResponse(topic string, ) *InviteStageInstanceResponse`

NewInviteStageInstanceResponse instantiates a new InviteStageInstanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteStageInstanceResponseWithDefaults

`func NewInviteStageInstanceResponseWithDefaults() *InviteStageInstanceResponse`

NewInviteStageInstanceResponseWithDefaults instantiates a new InviteStageInstanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *InviteStageInstanceResponse) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *InviteStageInstanceResponse) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *InviteStageInstanceResponse) SetTopic(v string)`

SetTopic sets Topic field to given value.


### GetParticipantCount

`func (o *InviteStageInstanceResponse) GetParticipantCount() int32`

GetParticipantCount returns the ParticipantCount field if non-nil, zero value otherwise.

### GetParticipantCountOk

`func (o *InviteStageInstanceResponse) GetParticipantCountOk() (*int32, bool)`

GetParticipantCountOk returns a tuple with the ParticipantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantCount

`func (o *InviteStageInstanceResponse) SetParticipantCount(v int32)`

SetParticipantCount sets ParticipantCount field to given value.

### HasParticipantCount

`func (o *InviteStageInstanceResponse) HasParticipantCount() bool`

HasParticipantCount returns a boolean if a field has been set.

### SetParticipantCountNil

`func (o *InviteStageInstanceResponse) SetParticipantCountNil(b bool)`

 SetParticipantCountNil sets the value for ParticipantCount to be an explicit nil

### UnsetParticipantCount
`func (o *InviteStageInstanceResponse) UnsetParticipantCount()`

UnsetParticipantCount ensures that no value is present for ParticipantCount, not even an explicit nil
### GetSpeakerCount

`func (o *InviteStageInstanceResponse) GetSpeakerCount() int32`

GetSpeakerCount returns the SpeakerCount field if non-nil, zero value otherwise.

### GetSpeakerCountOk

`func (o *InviteStageInstanceResponse) GetSpeakerCountOk() (*int32, bool)`

GetSpeakerCountOk returns a tuple with the SpeakerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeakerCount

`func (o *InviteStageInstanceResponse) SetSpeakerCount(v int32)`

SetSpeakerCount sets SpeakerCount field to given value.

### HasSpeakerCount

`func (o *InviteStageInstanceResponse) HasSpeakerCount() bool`

HasSpeakerCount returns a boolean if a field has been set.

### SetSpeakerCountNil

`func (o *InviteStageInstanceResponse) SetSpeakerCountNil(b bool)`

 SetSpeakerCountNil sets the value for SpeakerCount to be an explicit nil

### UnsetSpeakerCount
`func (o *InviteStageInstanceResponse) UnsetSpeakerCount()`

UnsetSpeakerCount ensures that no value is present for SpeakerCount, not even an explicit nil
### GetMembers

`func (o *InviteStageInstanceResponse) GetMembers() []GuildMemberResponse`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *InviteStageInstanceResponse) GetMembersOk() (*[]GuildMemberResponse, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *InviteStageInstanceResponse) SetMembers(v []GuildMemberResponse)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *InviteStageInstanceResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *InviteStageInstanceResponse) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *InviteStageInstanceResponse) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


