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

// BulkCIAttributeLayerScopeDTO struct for BulkCIAttributeLayerScopeDTO
type BulkCIAttributeLayerScopeDTO struct {
	NamePrefix string `json:"namePrefix"`
	LayerID string `json:"layerID"`
	Fragments []FragmentDTO `json:"fragments"`
}

// NewBulkCIAttributeLayerScopeDTO instantiates a new BulkCIAttributeLayerScopeDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkCIAttributeLayerScopeDTO(namePrefix string, layerID string, fragments []FragmentDTO) *BulkCIAttributeLayerScopeDTO {
	this := BulkCIAttributeLayerScopeDTO{}
	this.NamePrefix = namePrefix
	this.LayerID = layerID
	this.Fragments = fragments
	return &this
}

// NewBulkCIAttributeLayerScopeDTOWithDefaults instantiates a new BulkCIAttributeLayerScopeDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkCIAttributeLayerScopeDTOWithDefaults() *BulkCIAttributeLayerScopeDTO {
	this := BulkCIAttributeLayerScopeDTO{}
	return &this
}

// GetNamePrefix returns the NamePrefix field value
func (o *BulkCIAttributeLayerScopeDTO) GetNamePrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NamePrefix
}

// GetNamePrefixOk returns a tuple with the NamePrefix field value
// and a boolean to check if the value has been set.
func (o *BulkCIAttributeLayerScopeDTO) GetNamePrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NamePrefix, true
}

// SetNamePrefix sets field value
func (o *BulkCIAttributeLayerScopeDTO) SetNamePrefix(v string) {
	o.NamePrefix = v
}

// GetLayerID returns the LayerID field value
func (o *BulkCIAttributeLayerScopeDTO) GetLayerID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LayerID
}

// GetLayerIDOk returns a tuple with the LayerID field value
// and a boolean to check if the value has been set.
func (o *BulkCIAttributeLayerScopeDTO) GetLayerIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LayerID, true
}

// SetLayerID sets field value
func (o *BulkCIAttributeLayerScopeDTO) SetLayerID(v string) {
	o.LayerID = v
}

// GetFragments returns the Fragments field value
func (o *BulkCIAttributeLayerScopeDTO) GetFragments() []FragmentDTO {
	if o == nil {
		var ret []FragmentDTO
		return ret
	}

	return o.Fragments
}

// GetFragmentsOk returns a tuple with the Fragments field value
// and a boolean to check if the value has been set.
func (o *BulkCIAttributeLayerScopeDTO) GetFragmentsOk() ([]FragmentDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fragments, true
}

// SetFragments sets field value
func (o *BulkCIAttributeLayerScopeDTO) SetFragments(v []FragmentDTO) {
	o.Fragments = v
}

func (o BulkCIAttributeLayerScopeDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["namePrefix"] = o.NamePrefix
	}
	if true {
		toSerialize["layerID"] = o.LayerID
	}
	if true {
		toSerialize["fragments"] = o.Fragments
	}
	return json.Marshal(toSerialize)
}

type NullableBulkCIAttributeLayerScopeDTO struct {
	value *BulkCIAttributeLayerScopeDTO
	isSet bool
}

func (v NullableBulkCIAttributeLayerScopeDTO) Get() *BulkCIAttributeLayerScopeDTO {
	return v.value
}

func (v *NullableBulkCIAttributeLayerScopeDTO) Set(val *BulkCIAttributeLayerScopeDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkCIAttributeLayerScopeDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkCIAttributeLayerScopeDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkCIAttributeLayerScopeDTO(val *BulkCIAttributeLayerScopeDTO) *NullableBulkCIAttributeLayerScopeDTO {
	return &NullableBulkCIAttributeLayerScopeDTO{value: val, isSet: true}
}

func (v NullableBulkCIAttributeLayerScopeDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkCIAttributeLayerScopeDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


