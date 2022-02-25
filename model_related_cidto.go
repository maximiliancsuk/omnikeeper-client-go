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

// RelatedCIDTO struct for RelatedCIDTO
type RelatedCIDTO struct {
	FromCIID string `json:"fromCIID"`
	ToCIID string `json:"toCIID"`
	PredicateID string `json:"predicateID"`
}

// NewRelatedCIDTO instantiates a new RelatedCIDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelatedCIDTO(fromCIID string, toCIID string, predicateID string) *RelatedCIDTO {
	this := RelatedCIDTO{}
	this.FromCIID = fromCIID
	this.ToCIID = toCIID
	this.PredicateID = predicateID
	return &this
}

// NewRelatedCIDTOWithDefaults instantiates a new RelatedCIDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelatedCIDTOWithDefaults() *RelatedCIDTO {
	this := RelatedCIDTO{}
	return &this
}

// GetFromCIID returns the FromCIID field value
func (o *RelatedCIDTO) GetFromCIID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromCIID
}

// GetFromCIIDOk returns a tuple with the FromCIID field value
// and a boolean to check if the value has been set.
func (o *RelatedCIDTO) GetFromCIIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromCIID, true
}

// SetFromCIID sets field value
func (o *RelatedCIDTO) SetFromCIID(v string) {
	o.FromCIID = v
}

// GetToCIID returns the ToCIID field value
func (o *RelatedCIDTO) GetToCIID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToCIID
}

// GetToCIIDOk returns a tuple with the ToCIID field value
// and a boolean to check if the value has been set.
func (o *RelatedCIDTO) GetToCIIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToCIID, true
}

// SetToCIID sets field value
func (o *RelatedCIDTO) SetToCIID(v string) {
	o.ToCIID = v
}

// GetPredicateID returns the PredicateID field value
func (o *RelatedCIDTO) GetPredicateID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PredicateID
}

// GetPredicateIDOk returns a tuple with the PredicateID field value
// and a boolean to check if the value has been set.
func (o *RelatedCIDTO) GetPredicateIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PredicateID, true
}

// SetPredicateID sets field value
func (o *RelatedCIDTO) SetPredicateID(v string) {
	o.PredicateID = v
}

func (o RelatedCIDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fromCIID"] = o.FromCIID
	}
	if true {
		toSerialize["toCIID"] = o.ToCIID
	}
	if true {
		toSerialize["predicateID"] = o.PredicateID
	}
	return json.Marshal(toSerialize)
}

type NullableRelatedCIDTO struct {
	value *RelatedCIDTO
	isSet bool
}

func (v NullableRelatedCIDTO) Get() *RelatedCIDTO {
	return v.value
}

func (v *NullableRelatedCIDTO) Set(val *RelatedCIDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableRelatedCIDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableRelatedCIDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelatedCIDTO(val *RelatedCIDTO) *NullableRelatedCIDTO {
	return &NullableRelatedCIDTO{value: val, isSet: true}
}

func (v NullableRelatedCIDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelatedCIDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


