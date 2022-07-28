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

// GenericInboundData struct for GenericInboundData
type GenericInboundData struct {
	Cis []GenericInboundCI `json:"cis,omitempty"`
	Relations []GenericInboundRelation `json:"relations,omitempty"`
}

// NewGenericInboundData instantiates a new GenericInboundData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericInboundData() *GenericInboundData {
	this := GenericInboundData{}
	return &this
}

// NewGenericInboundDataWithDefaults instantiates a new GenericInboundData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericInboundDataWithDefaults() *GenericInboundData {
	this := GenericInboundData{}
	return &this
}

// GetCis returns the Cis field value if set, zero value otherwise.
func (o *GenericInboundData) GetCis() []GenericInboundCI {
	if o == nil || o.Cis == nil {
		var ret []GenericInboundCI
		return ret
	}
	return o.Cis
}

// GetCisOk returns a tuple with the Cis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericInboundData) GetCisOk() ([]GenericInboundCI, bool) {
	if o == nil || o.Cis == nil {
		return nil, false
	}
	return o.Cis, true
}

// HasCis returns a boolean if a field has been set.
func (o *GenericInboundData) HasCis() bool {
	if o != nil && o.Cis != nil {
		return true
	}

	return false
}

// SetCis gets a reference to the given []GenericInboundCI and assigns it to the Cis field.
func (o *GenericInboundData) SetCis(v []GenericInboundCI) {
	o.Cis = v
}

// GetRelations returns the Relations field value if set, zero value otherwise.
func (o *GenericInboundData) GetRelations() []GenericInboundRelation {
	if o == nil || o.Relations == nil {
		var ret []GenericInboundRelation
		return ret
	}
	return o.Relations
}

// GetRelationsOk returns a tuple with the Relations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericInboundData) GetRelationsOk() ([]GenericInboundRelation, bool) {
	if o == nil || o.Relations == nil {
		return nil, false
	}
	return o.Relations, true
}

// HasRelations returns a boolean if a field has been set.
func (o *GenericInboundData) HasRelations() bool {
	if o != nil && o.Relations != nil {
		return true
	}

	return false
}

// SetRelations gets a reference to the given []GenericInboundRelation and assigns it to the Relations field.
func (o *GenericInboundData) SetRelations(v []GenericInboundRelation) {
	o.Relations = v
}

func (o GenericInboundData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cis != nil {
		toSerialize["cis"] = o.Cis
	}
	if o.Relations != nil {
		toSerialize["relations"] = o.Relations
	}
	return json.Marshal(toSerialize)
}

type NullableGenericInboundData struct {
	value *GenericInboundData
	isSet bool
}

func (v NullableGenericInboundData) Get() *GenericInboundData {
	return v.value
}

func (v *NullableGenericInboundData) Set(val *GenericInboundData) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericInboundData) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericInboundData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericInboundData(val *GenericInboundData) *NullableGenericInboundData {
	return &NullableGenericInboundData{value: val, isSet: true}
}

func (v NullableGenericInboundData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericInboundData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


