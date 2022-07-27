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
	From NullableString `json:"from,omitempty"`
	Predicate NullableString `json:"predicate,omitempty"`
	To NullableString `json:"to,omitempty"`
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

// GetFrom returns the From field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenericInboundRelation) GetFrom() string {
	if o == nil || o.From.Get() == nil {
		var ret string
		return ret
	}
	return *o.From.Get()
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenericInboundRelation) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.From.Get(), o.From.IsSet()
}

// HasFrom returns a boolean if a field has been set.
func (o *GenericInboundRelation) HasFrom() bool {
	if o != nil && o.From.IsSet() {
		return true
	}

	return false
}

// SetFrom gets a reference to the given NullableString and assigns it to the From field.
func (o *GenericInboundRelation) SetFrom(v string) {
	o.From.Set(&v)
}
// SetFromNil sets the value for From to be an explicit nil
func (o *GenericInboundRelation) SetFromNil() {
	o.From.Set(nil)
}

// UnsetFrom ensures that no value is present for From, not even an explicit nil
func (o *GenericInboundRelation) UnsetFrom() {
	o.From.Unset()
}

// GetPredicate returns the Predicate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenericInboundRelation) GetPredicate() string {
	if o == nil || o.Predicate.Get() == nil {
		var ret string
		return ret
	}
	return *o.Predicate.Get()
}

// GetPredicateOk returns a tuple with the Predicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenericInboundRelation) GetPredicateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Predicate.Get(), o.Predicate.IsSet()
}

// HasPredicate returns a boolean if a field has been set.
func (o *GenericInboundRelation) HasPredicate() bool {
	if o != nil && o.Predicate.IsSet() {
		return true
	}

	return false
}

// SetPredicate gets a reference to the given NullableString and assigns it to the Predicate field.
func (o *GenericInboundRelation) SetPredicate(v string) {
	o.Predicate.Set(&v)
}
// SetPredicateNil sets the value for Predicate to be an explicit nil
func (o *GenericInboundRelation) SetPredicateNil() {
	o.Predicate.Set(nil)
}

// UnsetPredicate ensures that no value is present for Predicate, not even an explicit nil
func (o *GenericInboundRelation) UnsetPredicate() {
	o.Predicate.Unset()
}

// GetTo returns the To field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenericInboundRelation) GetTo() string {
	if o == nil || o.To.Get() == nil {
		var ret string
		return ret
	}
	return *o.To.Get()
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenericInboundRelation) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.To.Get(), o.To.IsSet()
}

// HasTo returns a boolean if a field has been set.
func (o *GenericInboundRelation) HasTo() bool {
	if o != nil && o.To.IsSet() {
		return true
	}

	return false
}

// SetTo gets a reference to the given NullableString and assigns it to the To field.
func (o *GenericInboundRelation) SetTo(v string) {
	o.To.Set(&v)
}
// SetToNil sets the value for To to be an explicit nil
func (o *GenericInboundRelation) SetToNil() {
	o.To.Set(nil)
}

// UnsetTo ensures that no value is present for To, not even an explicit nil
func (o *GenericInboundRelation) UnsetTo() {
	o.To.Unset()
}

func (o GenericInboundRelation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.From.IsSet() {
		toSerialize["from"] = o.From.Get()
	}
	if o.Predicate.IsSet() {
		toSerialize["predicate"] = o.Predicate.Get()
	}
	if o.To.IsSet() {
		toSerialize["to"] = o.To.Get()
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


