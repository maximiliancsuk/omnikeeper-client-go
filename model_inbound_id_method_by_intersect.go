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

// InboundIDMethodByIntersect struct for InboundIDMethodByIntersect
type InboundIDMethodByIntersect struct {
	Inner []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect `json:"inner,omitempty"`
}

// NewInboundIDMethodByIntersect instantiates a new InboundIDMethodByIntersect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundIDMethodByIntersect(type_ NullableString) *InboundIDMethodByIntersect {
	this := InboundIDMethodByIntersect{}
	this.Type = type_
	return &this
}

// NewInboundIDMethodByIntersectWithDefaults instantiates a new InboundIDMethodByIntersect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundIDMethodByIntersectWithDefaults() *InboundIDMethodByIntersect {
	this := InboundIDMethodByIntersect{}
	return &this
}

// GetInner returns the Inner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InboundIDMethodByIntersect) GetInner() []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect {
	if o == nil {
		var ret []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect
		return ret
	}
	return o.Inner
}

// GetInnerOk returns a tuple with the Inner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InboundIDMethodByIntersect) GetInnerOk() ([]OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect, bool) {
	if o == nil || o.Inner == nil {
		return nil, false
	}
	return o.Inner, true
}

// HasInner returns a boolean if a field has been set.
func (o *InboundIDMethodByIntersect) HasInner() bool {
	if o != nil && o.Inner != nil {
		return true
	}

	return false
}

// SetInner gets a reference to the given []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect and assigns it to the Inner field.
func (o *InboundIDMethodByIntersect) SetInner(v []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect) {
	o.Inner = v
}

func (o InboundIDMethodByIntersect) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Inner != nil {
		toSerialize["inner"] = o.Inner
	}
	return json.Marshal(toSerialize)
}

type NullableInboundIDMethodByIntersect struct {
	value *InboundIDMethodByIntersect
	isSet bool
}

func (v NullableInboundIDMethodByIntersect) Get() *InboundIDMethodByIntersect {
	return v.value
}

func (v *NullableInboundIDMethodByIntersect) Set(val *InboundIDMethodByIntersect) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundIDMethodByIntersect) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundIDMethodByIntersect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundIDMethodByIntersect(val *InboundIDMethodByIntersect) *NullableInboundIDMethodByIntersect {
	return &NullableInboundIDMethodByIntersect{value: val, isSet: true}
}

func (v NullableInboundIDMethodByIntersect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundIDMethodByIntersect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


