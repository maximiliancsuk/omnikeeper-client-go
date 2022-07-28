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
	Id *string `json:"id,omitempty"`
	SpeakingName *string `json:"speakingName,omitempty"`
	Description *string `json:"description,omitempty"`
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

// GetId returns the Id field value if set, zero value otherwise.
func (o *AddContextRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddContextRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AddContextRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AddContextRequest) SetId(v string) {
	o.Id = &v
}

// GetSpeakingName returns the SpeakingName field value if set, zero value otherwise.
func (o *AddContextRequest) GetSpeakingName() string {
	if o == nil || o.SpeakingName == nil {
		var ret string
		return ret
	}
	return *o.SpeakingName
}

// GetSpeakingNameOk returns a tuple with the SpeakingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddContextRequest) GetSpeakingNameOk() (*string, bool) {
	if o == nil || o.SpeakingName == nil {
		return nil, false
	}
	return o.SpeakingName, true
}

// HasSpeakingName returns a boolean if a field has been set.
func (o *AddContextRequest) HasSpeakingName() bool {
	if o != nil && o.SpeakingName != nil {
		return true
	}

	return false
}

// SetSpeakingName gets a reference to the given string and assigns it to the SpeakingName field.
func (o *AddContextRequest) SetSpeakingName(v string) {
	o.SpeakingName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddContextRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddContextRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddContextRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddContextRequest) SetDescription(v string) {
	o.Description = &v
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
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.SpeakingName != nil {
		toSerialize["speakingName"] = o.SpeakingName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
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


