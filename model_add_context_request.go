/*
Landscape omnikeeper REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package okclient

import (
	"encoding/json"
)

// AddContextRequest struct for AddContextRequest
type AddContextRequest struct {
	Name NullableString `json:"name,omitempty"`
	SpeakingName NullableString `json:"speakingName,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Configuration *GridViewConfiguration `json:"configuration,omitempty"`
}

// NewAddContextRequest instantiates a new AddContextRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddContextRequest() *AddContextRequest {
	this := AddContextRequest{}
	return &this
}

// NewAddContextRequestWithDefaults instantiates a new AddContextRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddContextRequestWithDefaults() *AddContextRequest {
	this := AddContextRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddContextRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddContextRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AddContextRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AddContextRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AddContextRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AddContextRequest) UnsetName() {
	o.Name.Unset()
}

// GetSpeakingName returns the SpeakingName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddContextRequest) GetSpeakingName() string {
	if o == nil || o.SpeakingName.Get() == nil {
		var ret string
		return ret
	}
	return *o.SpeakingName.Get()
}

// GetSpeakingNameOk returns a tuple with the SpeakingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddContextRequest) GetSpeakingNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SpeakingName.Get(), o.SpeakingName.IsSet()
}

// HasSpeakingName returns a boolean if a field has been set.
func (o *AddContextRequest) HasSpeakingName() bool {
	if o != nil && o.SpeakingName.IsSet() {
		return true
	}

	return false
}

// SetSpeakingName gets a reference to the given NullableString and assigns it to the SpeakingName field.
func (o *AddContextRequest) SetSpeakingName(v string) {
	o.SpeakingName.Set(&v)
}
// SetSpeakingNameNil sets the value for SpeakingName to be an explicit nil
func (o *AddContextRequest) SetSpeakingNameNil() {
	o.SpeakingName.Set(nil)
}

// UnsetSpeakingName ensures that no value is present for SpeakingName, not even an explicit nil
func (o *AddContextRequest) UnsetSpeakingName() {
	o.SpeakingName.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddContextRequest) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddContextRequest) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AddContextRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AddContextRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AddContextRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AddContextRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *AddContextRequest) GetConfiguration() GridViewConfiguration {
	if o == nil || o.Configuration == nil {
		var ret GridViewConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddContextRequest) GetConfigurationOk() (*GridViewConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *AddContextRequest) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given GridViewConfiguration and assigns it to the Configuration field.
func (o *AddContextRequest) SetConfiguration(v GridViewConfiguration) {
	o.Configuration = &v
}

func (o AddContextRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.SpeakingName.IsSet() {
		toSerialize["speakingName"] = o.SpeakingName.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	return json.Marshal(toSerialize)
}

type NullableAddContextRequest struct {
	value *AddContextRequest
	isSet bool
}

func (v NullableAddContextRequest) Get() *AddContextRequest {
	return v.value
}

func (v *NullableAddContextRequest) Set(val *AddContextRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddContextRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddContextRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddContextRequest(val *AddContextRequest) *NullableAddContextRequest {
	return &NullableAddContextRequest{value: val, isSet: true}
}

func (v NullableAddContextRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddContextRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


