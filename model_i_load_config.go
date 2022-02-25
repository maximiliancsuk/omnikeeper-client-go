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

// ILoadConfig struct for ILoadConfig
type ILoadConfig struct {
	SearchLayerIDs []string `json:"searchLayerIDs,omitempty"`
	WriteLayerID NullableString `json:"writeLayerID,omitempty"`
}

// NewILoadConfig instantiates a new ILoadConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewILoadConfig() *ILoadConfig {
	this := ILoadConfig{}
	return &this
}

// NewILoadConfigWithDefaults instantiates a new ILoadConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewILoadConfigWithDefaults() *ILoadConfig {
	this := ILoadConfig{}
	return &this
}

// GetSearchLayerIDs returns the SearchLayerIDs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ILoadConfig) GetSearchLayerIDs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SearchLayerIDs
}

// GetSearchLayerIDsOk returns a tuple with the SearchLayerIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ILoadConfig) GetSearchLayerIDsOk() ([]string, bool) {
	if o == nil || o.SearchLayerIDs == nil {
		return nil, false
	}
	return o.SearchLayerIDs, true
}

// HasSearchLayerIDs returns a boolean if a field has been set.
func (o *ILoadConfig) HasSearchLayerIDs() bool {
	if o != nil && o.SearchLayerIDs != nil {
		return true
	}

	return false
}

// SetSearchLayerIDs gets a reference to the given []string and assigns it to the SearchLayerIDs field.
func (o *ILoadConfig) SetSearchLayerIDs(v []string) {
	o.SearchLayerIDs = v
}

// GetWriteLayerID returns the WriteLayerID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ILoadConfig) GetWriteLayerID() string {
	if o == nil || o.WriteLayerID.Get() == nil {
		var ret string
		return ret
	}
	return *o.WriteLayerID.Get()
}

// GetWriteLayerIDOk returns a tuple with the WriteLayerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ILoadConfig) GetWriteLayerIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WriteLayerID.Get(), o.WriteLayerID.IsSet()
}

// HasWriteLayerID returns a boolean if a field has been set.
func (o *ILoadConfig) HasWriteLayerID() bool {
	if o != nil && o.WriteLayerID.IsSet() {
		return true
	}

	return false
}

// SetWriteLayerID gets a reference to the given NullableString and assigns it to the WriteLayerID field.
func (o *ILoadConfig) SetWriteLayerID(v string) {
	o.WriteLayerID.Set(&v)
}
// SetWriteLayerIDNil sets the value for WriteLayerID to be an explicit nil
func (o *ILoadConfig) SetWriteLayerIDNil() {
	o.WriteLayerID.Set(nil)
}

// UnsetWriteLayerID ensures that no value is present for WriteLayerID, not even an explicit nil
func (o *ILoadConfig) UnsetWriteLayerID() {
	o.WriteLayerID.Unset()
}

func (o ILoadConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SearchLayerIDs != nil {
		toSerialize["searchLayerIDs"] = o.SearchLayerIDs
	}
	if o.WriteLayerID.IsSet() {
		toSerialize["writeLayerID"] = o.WriteLayerID.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableILoadConfig struct {
	value *ILoadConfig
	isSet bool
}

func (v NullableILoadConfig) Get() *ILoadConfig {
	return v.value
}

func (v *NullableILoadConfig) Set(val *ILoadConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableILoadConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableILoadConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableILoadConfig(val *ILoadConfig) *NullableILoadConfig {
	return &NullableILoadConfig{value: val, isSet: true}
}

func (v NullableILoadConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableILoadConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


