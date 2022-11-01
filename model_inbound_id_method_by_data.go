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

// InboundIDMethodByData struct for InboundIDMethodByData
type InboundIDMethodByData struct {
	Attributes []string `json:"attributes,omitempty"`
}

// NewInboundIDMethodByData instantiates a new InboundIDMethodByData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundIDMethodByData(type_ string) *InboundIDMethodByData {
	this := InboundIDMethodByData{}
	this.Type = type_
	return &this
}

// NewInboundIDMethodByDataWithDefaults instantiates a new InboundIDMethodByData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundIDMethodByDataWithDefaults() *InboundIDMethodByData {
	this := InboundIDMethodByData{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *InboundIDMethodByData) GetAttributes() []string {
	if o == nil || isNil(o.Attributes) {
		var ret []string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundIDMethodByData) GetAttributesOk() ([]string, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *InboundIDMethodByData) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []string and assigns it to the Attributes field.
func (o *InboundIDMethodByData) SetAttributes(v []string) {
	o.Attributes = v
}

func (o InboundIDMethodByData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableInboundIDMethodByData struct {
	value *InboundIDMethodByData
	isSet bool
}

func (v NullableInboundIDMethodByData) Get() *InboundIDMethodByData {
	return v.value
}

func (v *NullableInboundIDMethodByData) Set(val *InboundIDMethodByData) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundIDMethodByData) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundIDMethodByData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundIDMethodByData(val *InboundIDMethodByData) *NullableInboundIDMethodByData {
	return &NullableInboundIDMethodByData{value: val, isSet: true}
}

func (v NullableInboundIDMethodByData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundIDMethodByData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


