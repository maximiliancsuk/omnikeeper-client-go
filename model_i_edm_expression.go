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

// IEdmExpression struct for IEdmExpression
type IEdmExpression struct {
	ExpressionKind *EdmExpressionKind `json:"expressionKind,omitempty"`
}

// NewIEdmExpression instantiates a new IEdmExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIEdmExpression() *IEdmExpression {
	this := IEdmExpression{}
	return &this
}

// NewIEdmExpressionWithDefaults instantiates a new IEdmExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIEdmExpressionWithDefaults() *IEdmExpression {
	this := IEdmExpression{}
	return &this
}

// GetExpressionKind returns the ExpressionKind field value if set, zero value otherwise.
func (o *IEdmExpression) GetExpressionKind() EdmExpressionKind {
	if o == nil || isNil(o.ExpressionKind) {
		var ret EdmExpressionKind
		return ret
	}
	return *o.ExpressionKind
}

// GetExpressionKindOk returns a tuple with the ExpressionKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmExpression) GetExpressionKindOk() (*EdmExpressionKind, bool) {
	if o == nil || isNil(o.ExpressionKind) {
    return nil, false
	}
	return o.ExpressionKind, true
}

// HasExpressionKind returns a boolean if a field has been set.
func (o *IEdmExpression) HasExpressionKind() bool {
	if o != nil && !isNil(o.ExpressionKind) {
		return true
	}

	return false
}

// SetExpressionKind gets a reference to the given EdmExpressionKind and assigns it to the ExpressionKind field.
func (o *IEdmExpression) SetExpressionKind(v EdmExpressionKind) {
	o.ExpressionKind = &v
}

func (o IEdmExpression) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExpressionKind) {
		toSerialize["expressionKind"] = o.ExpressionKind
	}
	return json.Marshal(toSerialize)
}

type NullableIEdmExpression struct {
	value *IEdmExpression
	isSet bool
}

func (v NullableIEdmExpression) Get() *IEdmExpression {
	return v.value
}

func (v *NullableIEdmExpression) Set(val *IEdmExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableIEdmExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableIEdmExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIEdmExpression(val *IEdmExpression) *NullableIEdmExpression {
	return &NullableIEdmExpression{value: val, isSet: true}
}

func (v NullableIEdmExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIEdmExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


