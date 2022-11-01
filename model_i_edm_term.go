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

// IEdmTerm struct for IEdmTerm
type IEdmTerm struct {
	Type *IEdmTypeReference `json:"type,omitempty"`
	AppliesTo NullableString `json:"appliesTo,omitempty"`
	DefaultValue NullableString `json:"defaultValue,omitempty"`
	SchemaElementKind *EdmSchemaElementKind `json:"schemaElementKind,omitempty"`
	Namespace NullableString `json:"namespace,omitempty"`
	Name NullableString `json:"name,omitempty"`
}

// NewIEdmTerm instantiates a new IEdmTerm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIEdmTerm() *IEdmTerm {
	this := IEdmTerm{}
	return &this
}

// NewIEdmTermWithDefaults instantiates a new IEdmTerm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIEdmTermWithDefaults() *IEdmTerm {
	this := IEdmTerm{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IEdmTerm) GetType() IEdmTypeReference {
	if o == nil || isNil(o.Type) {
		var ret IEdmTypeReference
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmTerm) GetTypeOk() (*IEdmTypeReference, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IEdmTerm) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given IEdmTypeReference and assigns it to the Type field.
func (o *IEdmTerm) SetType(v IEdmTypeReference) {
	o.Type = &v
}

// GetAppliesTo returns the AppliesTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IEdmTerm) GetAppliesTo() string {
	if o == nil || isNil(o.AppliesTo.Get()) {
		var ret string
		return ret
	}
	return *o.AppliesTo.Get()
}

// GetAppliesToOk returns a tuple with the AppliesTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IEdmTerm) GetAppliesToOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AppliesTo.Get(), o.AppliesTo.IsSet()
}

// HasAppliesTo returns a boolean if a field has been set.
func (o *IEdmTerm) HasAppliesTo() bool {
	if o != nil && o.AppliesTo.IsSet() {
		return true
	}

	return false
}

// SetAppliesTo gets a reference to the given NullableString and assigns it to the AppliesTo field.
func (o *IEdmTerm) SetAppliesTo(v string) {
	o.AppliesTo.Set(&v)
}
// SetAppliesToNil sets the value for AppliesTo to be an explicit nil
func (o *IEdmTerm) SetAppliesToNil() {
	o.AppliesTo.Set(nil)
}

// UnsetAppliesTo ensures that no value is present for AppliesTo, not even an explicit nil
func (o *IEdmTerm) UnsetAppliesTo() {
	o.AppliesTo.Unset()
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IEdmTerm) GetDefaultValue() string {
	if o == nil || isNil(o.DefaultValue.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultValue.Get()
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IEdmTerm) GetDefaultValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DefaultValue.Get(), o.DefaultValue.IsSet()
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *IEdmTerm) HasDefaultValue() bool {
	if o != nil && o.DefaultValue.IsSet() {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given NullableString and assigns it to the DefaultValue field.
func (o *IEdmTerm) SetDefaultValue(v string) {
	o.DefaultValue.Set(&v)
}
// SetDefaultValueNil sets the value for DefaultValue to be an explicit nil
func (o *IEdmTerm) SetDefaultValueNil() {
	o.DefaultValue.Set(nil)
}

// UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
func (o *IEdmTerm) UnsetDefaultValue() {
	o.DefaultValue.Unset()
}

// GetSchemaElementKind returns the SchemaElementKind field value if set, zero value otherwise.
func (o *IEdmTerm) GetSchemaElementKind() EdmSchemaElementKind {
	if o == nil || isNil(o.SchemaElementKind) {
		var ret EdmSchemaElementKind
		return ret
	}
	return *o.SchemaElementKind
}

// GetSchemaElementKindOk returns a tuple with the SchemaElementKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmTerm) GetSchemaElementKindOk() (*EdmSchemaElementKind, bool) {
	if o == nil || isNil(o.SchemaElementKind) {
    return nil, false
	}
	return o.SchemaElementKind, true
}

// HasSchemaElementKind returns a boolean if a field has been set.
func (o *IEdmTerm) HasSchemaElementKind() bool {
	if o != nil && !isNil(o.SchemaElementKind) {
		return true
	}

	return false
}

// SetSchemaElementKind gets a reference to the given EdmSchemaElementKind and assigns it to the SchemaElementKind field.
func (o *IEdmTerm) SetSchemaElementKind(v EdmSchemaElementKind) {
	o.SchemaElementKind = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IEdmTerm) GetNamespace() string {
	if o == nil || isNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IEdmTerm) GetNamespaceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *IEdmTerm) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *IEdmTerm) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *IEdmTerm) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *IEdmTerm) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IEdmTerm) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IEdmTerm) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IEdmTerm) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IEdmTerm) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IEdmTerm) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IEdmTerm) UnsetName() {
	o.Name.Unset()
}

func (o IEdmTerm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.AppliesTo.IsSet() {
		toSerialize["appliesTo"] = o.AppliesTo.Get()
	}
	if o.DefaultValue.IsSet() {
		toSerialize["defaultValue"] = o.DefaultValue.Get()
	}
	if !isNil(o.SchemaElementKind) {
		toSerialize["schemaElementKind"] = o.SchemaElementKind
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIEdmTerm struct {
	value *IEdmTerm
	isSet bool
}

func (v NullableIEdmTerm) Get() *IEdmTerm {
	return v.value
}

func (v *NullableIEdmTerm) Set(val *IEdmTerm) {
	v.value = val
	v.isSet = true
}

func (v NullableIEdmTerm) IsSet() bool {
	return v.isSet
}

func (v *NullableIEdmTerm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIEdmTerm(val *IEdmTerm) *NullableIEdmTerm {
	return &NullableIEdmTerm{value: val, isSet: true}
}

func (v NullableIEdmTerm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIEdmTerm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


