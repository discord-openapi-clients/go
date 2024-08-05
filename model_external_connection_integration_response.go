/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discord

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ExternalConnectionIntegrationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalConnectionIntegrationResponse{}

// ExternalConnectionIntegrationResponse struct for ExternalConnectionIntegrationResponse
type ExternalConnectionIntegrationResponse struct {
	Type IntegrationTypes `json:"type"`
	Name NullableString `json:"name,omitempty"`
	Account NullableAccountResponse `json:"account,omitempty"`
	Enabled NullableBool `json:"enabled,omitempty"`
	Id string `json:"id"`
	User UserResponse `json:"user"`
	Revoked NullableBool `json:"revoked,omitempty"`
	ExpireBehavior NullableIntegrationExpireBehaviorTypes `json:"expire_behavior,omitempty"`
	ExpireGracePeriod NullableIntegrationExpireGracePeriodTypes `json:"expire_grace_period,omitempty"`
	SubscriberCount NullableInt32 `json:"subscriber_count,omitempty"`
	SyncedAt NullableTime `json:"synced_at,omitempty"`
	RoleId *string `json:"role_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Syncing NullableBool `json:"syncing,omitempty"`
	EnableEmoticons NullableBool `json:"enable_emoticons,omitempty"`
}

type _ExternalConnectionIntegrationResponse ExternalConnectionIntegrationResponse

// NewExternalConnectionIntegrationResponse instantiates a new ExternalConnectionIntegrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalConnectionIntegrationResponse(type_ IntegrationTypes, id string, user UserResponse) *ExternalConnectionIntegrationResponse {
	this := ExternalConnectionIntegrationResponse{}
	this.Type = type_
	this.Id = id
	this.User = user
	return &this
}

// NewExternalConnectionIntegrationResponseWithDefaults instantiates a new ExternalConnectionIntegrationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalConnectionIntegrationResponseWithDefaults() *ExternalConnectionIntegrationResponse {
	this := ExternalConnectionIntegrationResponse{}
	return &this
}

// GetType returns the Type field value
func (o *ExternalConnectionIntegrationResponse) GetType() IntegrationTypes {
	if o == nil {
		var ret IntegrationTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExternalConnectionIntegrationResponse) GetTypeOk() (*IntegrationTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalConnectionIntegrationResponse) SetType(v IntegrationTypes) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalConnectionIntegrationResponse) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalConnectionIntegrationResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ExternalConnectionIntegrationResponse) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ExternalConnectionIntegrationResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ExternalConnectionIntegrationResponse) UnsetName() {
	o.Name.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalConnectionIntegrationResponse) GetAccount() AccountResponse {
	if o == nil || IsNil(o.Account.Get()) {
		var ret AccountResponse
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalConnectionIntegrationResponse) GetAccountOk() (*AccountResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableAccountResponse and assigns it to the Account field.
func (o *ExternalConnectionIntegrationResponse) SetAccount(v AccountResponse) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *ExternalConnectionIntegrationResponse) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *ExternalConnectionIntegrationResponse) UnsetAccount() {
	o.Account.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalConnectionIntegrationResponse) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalConnectionIntegrationResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *ExternalConnectionIntegrationResponse) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *ExternalConnectionIntegrationResponse) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *ExternalConnectionIntegrationResponse) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetId returns the Id field value
func (o *ExternalConnectionIntegrationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalConnectionIntegrationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalConnectionIntegrationResponse) SetId(v string) {
	o.Id = v
}

// GetUser returns the User field value
func (o *ExternalConnectionIntegrationResponse) GetUser() UserResponse {
	if o == nil {
		var ret UserResponse
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ExternalConnectionIntegrationResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ExternalConnectionIntegrationResponse) SetUser(v UserResponse) {
	o.User = v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalConnectionIntegrationResponse) GetRevoked() bool {
	if o == nil || IsNil(o.Revoked.Get()) {
		var ret bool
		return ret
	}
	return *o.Revoked.Get()
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalConnectionIntegrationResponse) GetRevokedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Revoked.Get(), o.Revoked.IsSet()
}

// HasRevoked returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasRevoked() bool {
	if o != nil && o.Revoked.IsSet() {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given NullableBool and assigns it to the Revoked field.
func (o *ExternalConnectionIntegrationResponse) SetRevoked(v bool) {
	o.Revoked.Set(&v)
}
// SetRevokedNil sets the value for Revoked to be an explicit nil
func (o *ExternalConnectionIntegrationResponse) SetRevokedNil() {
	o.Revoked.Set(nil)
}

// UnsetRevoked ensures that no value is present for Revoked, not even an explicit nil
func (o *ExternalConnectionIntegrationResponse) UnsetRevoked() {
	o.Revoked.Unset()
}

// GetExpireBehavior returns the ExpireBehavior field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalConnectionIntegrationResponse) GetExpireBehavior() IntegrationExpireBehaviorTypes {
	if o == nil || IsNil(o.ExpireBehavior.Get()) {
		var ret IntegrationExpireBehaviorTypes
		return ret
	}
	return *o.ExpireBehavior.Get()
}

// GetExpireBehaviorOk returns a tuple with the ExpireBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalConnectionIntegrationResponse) GetExpireBehaviorOk() (*IntegrationExpireBehaviorTypes, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpireBehavior.Get(), o.ExpireBehavior.IsSet()
}

// HasExpireBehavior returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasExpireBehavior() bool {
	if o != nil && o.ExpireBehavior.IsSet() {
		return true
	}

	return false
}

// SetExpireBehavior gets a reference to the given NullableIntegrationExpireBehaviorTypes and assigns it to the ExpireBehavior field.
func (o *ExternalConnectionIntegrationResponse) SetExpireBehavior(v IntegrationExpireBehaviorTypes) {
	o.ExpireBehavior.Set(&v)
}
// SetExpireBehaviorNil sets the value for ExpireBehavior to be an explicit nil
func (o *ExternalConnectionIntegrationResponse) SetExpireBehaviorNil() {
	o.ExpireBehavior.Set(nil)
}

// UnsetExpireBehavior ensures that no value is present for ExpireBehavior, not even an explicit nil
func (o *ExternalConnectionIntegrationResponse) UnsetExpireBehavior() {
	o.ExpireBehavior.Unset()
}

// GetExpireGracePeriod returns the ExpireGracePeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalConnectionIntegrationResponse) GetExpireGracePeriod() IntegrationExpireGracePeriodTypes {
	if o == nil || IsNil(o.ExpireGracePeriod.Get()) {
		var ret IntegrationExpireGracePeriodTypes
		return ret
	}
	return *o.ExpireGracePeriod.Get()
}

// GetExpireGracePeriodOk returns a tuple with the ExpireGracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalConnectionIntegrationResponse) GetExpireGracePeriodOk() (*IntegrationExpireGracePeriodTypes, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpireGracePeriod.Get(), o.ExpireGracePeriod.IsSet()
}

// HasExpireGracePeriod returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasExpireGracePeriod() bool {
	if o != nil && o.ExpireGracePeriod.IsSet() {
		return true
	}

	return false
}

// SetExpireGracePeriod gets a reference to the given NullableIntegrationExpireGracePeriodTypes and assigns it to the ExpireGracePeriod field.
func (o *ExternalConnectionIntegrationResponse) SetExpireGracePeriod(v IntegrationExpireGracePeriodTypes) {
	o.ExpireGracePeriod.Set(&v)
}
// SetExpireGracePeriodNil sets the value for ExpireGracePeriod to be an explicit nil
func (o *ExternalConnectionIntegrationResponse) SetExpireGracePeriodNil() {
	o.ExpireGracePeriod.Set(nil)
}

// UnsetExpireGracePeriod ensures that no value is present for ExpireGracePeriod, not even an explicit nil
func (o *ExternalConnectionIntegrationResponse) UnsetExpireGracePeriod() {
	o.ExpireGracePeriod.Unset()
}

// GetSubscriberCount returns the SubscriberCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalConnectionIntegrationResponse) GetSubscriberCount() int32 {
	if o == nil || IsNil(o.SubscriberCount.Get()) {
		var ret int32
		return ret
	}
	return *o.SubscriberCount.Get()
}

// GetSubscriberCountOk returns a tuple with the SubscriberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalConnectionIntegrationResponse) GetSubscriberCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriberCount.Get(), o.SubscriberCount.IsSet()
}

// HasSubscriberCount returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasSubscriberCount() bool {
	if o != nil && o.SubscriberCount.IsSet() {
		return true
	}

	return false
}

// SetSubscriberCount gets a reference to the given NullableInt32 and assigns it to the SubscriberCount field.
func (o *ExternalConnectionIntegrationResponse) SetSubscriberCount(v int32) {
	o.SubscriberCount.Set(&v)
}
// SetSubscriberCountNil sets the value for SubscriberCount to be an explicit nil
func (o *ExternalConnectionIntegrationResponse) SetSubscriberCountNil() {
	o.SubscriberCount.Set(nil)
}

// UnsetSubscriberCount ensures that no value is present for SubscriberCount, not even an explicit nil
func (o *ExternalConnectionIntegrationResponse) UnsetSubscriberCount() {
	o.SubscriberCount.Unset()
}

// GetSyncedAt returns the SyncedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalConnectionIntegrationResponse) GetSyncedAt() time.Time {
	if o == nil || IsNil(o.SyncedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SyncedAt.Get()
}

// GetSyncedAtOk returns a tuple with the SyncedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalConnectionIntegrationResponse) GetSyncedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyncedAt.Get(), o.SyncedAt.IsSet()
}

// HasSyncedAt returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasSyncedAt() bool {
	if o != nil && o.SyncedAt.IsSet() {
		return true
	}

	return false
}

// SetSyncedAt gets a reference to the given NullableTime and assigns it to the SyncedAt field.
func (o *ExternalConnectionIntegrationResponse) SetSyncedAt(v time.Time) {
	o.SyncedAt.Set(&v)
}
// SetSyncedAtNil sets the value for SyncedAt to be an explicit nil
func (o *ExternalConnectionIntegrationResponse) SetSyncedAtNil() {
	o.SyncedAt.Set(nil)
}

// UnsetSyncedAt ensures that no value is present for SyncedAt, not even an explicit nil
func (o *ExternalConnectionIntegrationResponse) UnsetSyncedAt() {
	o.SyncedAt.Unset()
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *ExternalConnectionIntegrationResponse) GetRoleId() string {
	if o == nil || IsNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalConnectionIntegrationResponse) GetRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *ExternalConnectionIntegrationResponse) SetRoleId(v string) {
	o.RoleId = &v
}

// GetSyncing returns the Syncing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalConnectionIntegrationResponse) GetSyncing() bool {
	if o == nil || IsNil(o.Syncing.Get()) {
		var ret bool
		return ret
	}
	return *o.Syncing.Get()
}

// GetSyncingOk returns a tuple with the Syncing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalConnectionIntegrationResponse) GetSyncingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Syncing.Get(), o.Syncing.IsSet()
}

// HasSyncing returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasSyncing() bool {
	if o != nil && o.Syncing.IsSet() {
		return true
	}

	return false
}

// SetSyncing gets a reference to the given NullableBool and assigns it to the Syncing field.
func (o *ExternalConnectionIntegrationResponse) SetSyncing(v bool) {
	o.Syncing.Set(&v)
}
// SetSyncingNil sets the value for Syncing to be an explicit nil
func (o *ExternalConnectionIntegrationResponse) SetSyncingNil() {
	o.Syncing.Set(nil)
}

// UnsetSyncing ensures that no value is present for Syncing, not even an explicit nil
func (o *ExternalConnectionIntegrationResponse) UnsetSyncing() {
	o.Syncing.Unset()
}

// GetEnableEmoticons returns the EnableEmoticons field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalConnectionIntegrationResponse) GetEnableEmoticons() bool {
	if o == nil || IsNil(o.EnableEmoticons.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableEmoticons.Get()
}

// GetEnableEmoticonsOk returns a tuple with the EnableEmoticons field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalConnectionIntegrationResponse) GetEnableEmoticonsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableEmoticons.Get(), o.EnableEmoticons.IsSet()
}

// HasEnableEmoticons returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasEnableEmoticons() bool {
	if o != nil && o.EnableEmoticons.IsSet() {
		return true
	}

	return false
}

// SetEnableEmoticons gets a reference to the given NullableBool and assigns it to the EnableEmoticons field.
func (o *ExternalConnectionIntegrationResponse) SetEnableEmoticons(v bool) {
	o.EnableEmoticons.Set(&v)
}
// SetEnableEmoticonsNil sets the value for EnableEmoticons to be an explicit nil
func (o *ExternalConnectionIntegrationResponse) SetEnableEmoticonsNil() {
	o.EnableEmoticons.Set(nil)
}

// UnsetEnableEmoticons ensures that no value is present for EnableEmoticons, not even an explicit nil
func (o *ExternalConnectionIntegrationResponse) UnsetEnableEmoticons() {
	o.EnableEmoticons.Unset()
}

func (o ExternalConnectionIntegrationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalConnectionIntegrationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["user"] = o.User
	if o.Revoked.IsSet() {
		toSerialize["revoked"] = o.Revoked.Get()
	}
	if o.ExpireBehavior.IsSet() {
		toSerialize["expire_behavior"] = o.ExpireBehavior.Get()
	}
	if o.ExpireGracePeriod.IsSet() {
		toSerialize["expire_grace_period"] = o.ExpireGracePeriod.Get()
	}
	if o.SubscriberCount.IsSet() {
		toSerialize["subscriber_count"] = o.SubscriberCount.Get()
	}
	if o.SyncedAt.IsSet() {
		toSerialize["synced_at"] = o.SyncedAt.Get()
	}
	if !IsNil(o.RoleId) {
		toSerialize["role_id"] = o.RoleId
	}
	if o.Syncing.IsSet() {
		toSerialize["syncing"] = o.Syncing.Get()
	}
	if o.EnableEmoticons.IsSet() {
		toSerialize["enable_emoticons"] = o.EnableEmoticons.Get()
	}
	return toSerialize, nil
}

func (o *ExternalConnectionIntegrationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"user",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varExternalConnectionIntegrationResponse := _ExternalConnectionIntegrationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalConnectionIntegrationResponse)

	if err != nil {
		return err
	}

	*o = ExternalConnectionIntegrationResponse(varExternalConnectionIntegrationResponse)

	return err
}

type NullableExternalConnectionIntegrationResponse struct {
	value *ExternalConnectionIntegrationResponse
	isSet bool
}

func (v NullableExternalConnectionIntegrationResponse) Get() *ExternalConnectionIntegrationResponse {
	return v.value
}

func (v *NullableExternalConnectionIntegrationResponse) Set(val *ExternalConnectionIntegrationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalConnectionIntegrationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalConnectionIntegrationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalConnectionIntegrationResponse(val *ExternalConnectionIntegrationResponse) *NullableExternalConnectionIntegrationResponse {
	return &NullableExternalConnectionIntegrationResponse{value: val, isSet: true}
}

func (v NullableExternalConnectionIntegrationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalConnectionIntegrationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


