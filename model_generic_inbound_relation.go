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

// GenericInboundRelation struct for GenericInboundRelation
type GenericInboundRelation struct {
	From *string `json:"from,omitempty"`
	Predicate *string `json:"predicate,omitempty"`
	To *string `json:"to,omitempty"`
}

// NewGenericInboundRelation instantiates a new GenericInboundRelation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericInboundRelation() *GenericInboundRelation {
	this := GenericInboundRelation{}
	return &this
}

// NewGenericInboundRelationWithDefaults instantiates a new GenericInboundRelation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericInboundRelationWithDefaults() *GenericInboundRelation {
	this := GenericInboundRelation{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *GenericInboundRelation) GetFrom() string {
	if o == nil || isNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericInboundRelation) GetFromOk() (*string, bool) {
	if o == nil || isNil(o.From) {
    return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *GenericInboundRelation) HasFrom() bool {
	if o != nil && !isNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *GenericInboundRelation) SetFrom(v string) {
	o.From = &v
}

// GetPredicate returns the Predicate field value if set, zero value otherwise.
func (o *GenericInboundRelation) GetPredicate() string {
	if o == nil || isNil(o.Predicate) {
		var ret string
		return ret
	}
	return *o.Predicate
}

// GetPredicateOk returns a tuple with the Predicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericInboundRelation) GetPredicateOk() (*string, bool) {
	if o == nil || isNil(o.Predicate) {
    return nil, false
	}
	return o.Predicate, true
}

// HasPredicate returns a boolean if a field has been set.
func (o *GenericInboundRelation) HasPredicate() bool {
	if o != nil && !isNil(o.Predicate) {
		return true
	}

	return false
}

// SetPredicate gets a reference to the given string and assigns it to the Predicate field.
func (o *GenericInboundRelation) SetPredicate(v string) {
	o.Predicate = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *GenericInboundRelation) GetTo() string {
	if o == nil || isNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericInboundRelation) GetToOk() (*string, bool) {
	if o == nil || isNil(o.To) {
    return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *GenericInboundRelation) HasTo() bool {
	if o != nil && !isNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *GenericInboundRelation) SetTo(v string) {
	o.To = &v
}

func (o GenericInboundRelation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !isNil(o.Predicate) {
		toSerialize["predicate"] = o.Predicate
	}
	if !isNil(o.To) {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableGenericInboundRelation struct {
	value *GenericInboundRelation
	isSet bool
}

func (v NullableGenericInboundRelation) Get() *GenericInboundRelation {
	return v.value
}

func (v *NullableGenericInboundRelation) Set(val *GenericInboundRelation) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericInboundRelation) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericInboundRelation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericInboundRelation(val *GenericInboundRelation) *NullableGenericInboundRelation {
	return &NullableGenericInboundRelation{value: val, isSet: true}
}

func (v NullableGenericInboundRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericInboundRelation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


