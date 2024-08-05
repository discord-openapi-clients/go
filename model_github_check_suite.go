/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GithubCheckSuite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubCheckSuite{}

// GithubCheckSuite struct for GithubCheckSuite
type GithubCheckSuite struct {
	Conclusion NullableString `json:"conclusion,omitempty"`
	HeadBranch NullableString `json:"head_branch,omitempty"`
	HeadSha string `json:"head_sha"`
	PullRequests []GithubCheckPullRequest `json:"pull_requests,omitempty"`
	App GithubCheckApp `json:"app"`
}

type _GithubCheckSuite GithubCheckSuite

// NewGithubCheckSuite instantiates a new GithubCheckSuite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubCheckSuite(headSha string, app GithubCheckApp) *GithubCheckSuite {
	this := GithubCheckSuite{}
	this.HeadSha = headSha
	this.App = app
	return &this
}

// NewGithubCheckSuiteWithDefaults instantiates a new GithubCheckSuite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubCheckSuiteWithDefaults() *GithubCheckSuite {
	this := GithubCheckSuite{}
	return &this
}

// GetConclusion returns the Conclusion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubCheckSuite) GetConclusion() string {
	if o == nil || IsNil(o.Conclusion.Get()) {
		var ret string
		return ret
	}
	return *o.Conclusion.Get()
}

// GetConclusionOk returns a tuple with the Conclusion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubCheckSuite) GetConclusionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conclusion.Get(), o.Conclusion.IsSet()
}

// HasConclusion returns a boolean if a field has been set.
func (o *GithubCheckSuite) HasConclusion() bool {
	if o != nil && o.Conclusion.IsSet() {
		return true
	}

	return false
}

// SetConclusion gets a reference to the given NullableString and assigns it to the Conclusion field.
func (o *GithubCheckSuite) SetConclusion(v string) {
	o.Conclusion.Set(&v)
}
// SetConclusionNil sets the value for Conclusion to be an explicit nil
func (o *GithubCheckSuite) SetConclusionNil() {
	o.Conclusion.Set(nil)
}

// UnsetConclusion ensures that no value is present for Conclusion, not even an explicit nil
func (o *GithubCheckSuite) UnsetConclusion() {
	o.Conclusion.Unset()
}

// GetHeadBranch returns the HeadBranch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubCheckSuite) GetHeadBranch() string {
	if o == nil || IsNil(o.HeadBranch.Get()) {
		var ret string
		return ret
	}
	return *o.HeadBranch.Get()
}

// GetHeadBranchOk returns a tuple with the HeadBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubCheckSuite) GetHeadBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeadBranch.Get(), o.HeadBranch.IsSet()
}

// HasHeadBranch returns a boolean if a field has been set.
func (o *GithubCheckSuite) HasHeadBranch() bool {
	if o != nil && o.HeadBranch.IsSet() {
		return true
	}

	return false
}

// SetHeadBranch gets a reference to the given NullableString and assigns it to the HeadBranch field.
func (o *GithubCheckSuite) SetHeadBranch(v string) {
	o.HeadBranch.Set(&v)
}
// SetHeadBranchNil sets the value for HeadBranch to be an explicit nil
func (o *GithubCheckSuite) SetHeadBranchNil() {
	o.HeadBranch.Set(nil)
}

// UnsetHeadBranch ensures that no value is present for HeadBranch, not even an explicit nil
func (o *GithubCheckSuite) UnsetHeadBranch() {
	o.HeadBranch.Unset()
}

// GetHeadSha returns the HeadSha field value
func (o *GithubCheckSuite) GetHeadSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HeadSha
}

// GetHeadShaOk returns a tuple with the HeadSha field value
// and a boolean to check if the value has been set.
func (o *GithubCheckSuite) GetHeadShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeadSha, true
}

// SetHeadSha sets field value
func (o *GithubCheckSuite) SetHeadSha(v string) {
	o.HeadSha = v
}

// GetPullRequests returns the PullRequests field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubCheckSuite) GetPullRequests() []GithubCheckPullRequest {
	if o == nil {
		var ret []GithubCheckPullRequest
		return ret
	}
	return o.PullRequests
}

// GetPullRequestsOk returns a tuple with the PullRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubCheckSuite) GetPullRequestsOk() ([]GithubCheckPullRequest, bool) {
	if o == nil || IsNil(o.PullRequests) {
		return nil, false
	}
	return o.PullRequests, true
}

// HasPullRequests returns a boolean if a field has been set.
func (o *GithubCheckSuite) HasPullRequests() bool {
	if o != nil && !IsNil(o.PullRequests) {
		return true
	}

	return false
}

// SetPullRequests gets a reference to the given []GithubCheckPullRequest and assigns it to the PullRequests field.
func (o *GithubCheckSuite) SetPullRequests(v []GithubCheckPullRequest) {
	o.PullRequests = v
}

// GetApp returns the App field value
func (o *GithubCheckSuite) GetApp() GithubCheckApp {
	if o == nil {
		var ret GithubCheckApp
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *GithubCheckSuite) GetAppOk() (*GithubCheckApp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *GithubCheckSuite) SetApp(v GithubCheckApp) {
	o.App = v
}

func (o GithubCheckSuite) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubCheckSuite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Conclusion.IsSet() {
		toSerialize["conclusion"] = o.Conclusion.Get()
	}
	if o.HeadBranch.IsSet() {
		toSerialize["head_branch"] = o.HeadBranch.Get()
	}
	toSerialize["head_sha"] = o.HeadSha
	if o.PullRequests != nil {
		toSerialize["pull_requests"] = o.PullRequests
	}
	toSerialize["app"] = o.App
	return toSerialize, nil
}

func (o *GithubCheckSuite) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"head_sha",
		"app",
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

	varGithubCheckSuite := _GithubCheckSuite{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGithubCheckSuite)

	if err != nil {
		return err
	}

	*o = GithubCheckSuite(varGithubCheckSuite)

	return err
}

type NullableGithubCheckSuite struct {
	value *GithubCheckSuite
	isSet bool
}

func (v NullableGithubCheckSuite) Get() *GithubCheckSuite {
	return v.value
}

func (v *NullableGithubCheckSuite) Set(val *GithubCheckSuite) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubCheckSuite) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubCheckSuite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubCheckSuite(val *GithubCheckSuite) *NullableGithubCheckSuite {
	return &NullableGithubCheckSuite{value: val, isSet: true}
}

func (v NullableGithubCheckSuite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubCheckSuite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


