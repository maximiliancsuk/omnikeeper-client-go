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

// checks if the IEdmEntityContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IEdmEntityContainer{}

// IEdmEntityContainer struct for IEdmEntityContainer
type IEdmEntityContainer struct {
	Elements []IEdmEntityContainerElement `json:"elements,omitempty"`
	SchemaElementKind *EdmSchemaElementKind `json:"schemaElementKind,omitempty"`
	Namespace NullableString `json:"namespace,omitempty"`
	Name NullableString `json:"name,omitempty"`
}

// NewIEdmEntityContainer instantiates a new IEdmEntityContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIEdmEntityContainer() *IEdmEntityContainer {
	this := IEdmEntityContainer{}
	return &this
}

// NewIEdmEntityContainerWithDefaults instantiates a new IEdmEntityContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIEdmEntityContainerWithDefaults() *IEdmEntityContainer {
	this := IEdmEntityContainer{}
	return &this
}

// GetElements returns the Elements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IEdmEntityContainer) GetElements() []IEdmEntityContainerElement {
	if o == nil {
		var ret []IEdmEntityContainerElement
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IEdmEntityContainer) GetElementsOk() ([]IEdmEntityContainerElement, bool) {
	if o == nil || isNil(o.Elements) {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *IEdmEntityContainer) HasElements() bool {
	if o != nil && isNil(o.Elements) {
		return true
	}

	return false
}

// SetElements gets a reference to the given []IEdmEntityContainerElement and assigns it to the Elements field.
func (o *IEdmEntityContainer) SetElements(v []IEdmEntityContainerElement) {
	o.Elements = v
}

// GetSchemaElementKind returns the SchemaElementKind field value if set, zero value otherwise.
func (o *IEdmEntityContainer) GetSchemaElementKind() EdmSchemaElementKind {
	if o == nil || isNil(o.SchemaElementKind) {
		var ret EdmSchemaElementKind
		return ret
	}
	return *o.SchemaElementKind
}

// GetSchemaElementKindOk returns a tuple with the SchemaElementKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmEntityContainer) GetSchemaElementKindOk() (*EdmSchemaElementKind, bool) {
	if o == nil || isNil(o.SchemaElementKind) {
		return nil, false
	}
	return o.SchemaElementKind, true
}

// HasSchemaElementKind returns a boolean if a field has been set.
func (o *IEdmEntityContainer) HasSchemaElementKind() bool {
	if o != nil && !isNil(o.SchemaElementKind) {
		return true
	}

	return false
}

// SetSchemaElementKind gets a reference to the given EdmSchemaElementKind and assigns it to the SchemaElementKind field.
func (o *IEdmEntityContainer) SetSchemaElementKind(v EdmSchemaElementKind) {
	o.SchemaElementKind = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IEdmEntityContainer) GetNamespace() string {
	if o == nil || isNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IEdmEntityContainer) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *IEdmEntityContainer) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *IEdmEntityContainer) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *IEdmEntityContainer) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *IEdmEntityContainer) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IEdmEntityContainer) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IEdmEntityContainer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IEdmEntityContainer) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IEdmEntityContainer) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IEdmEntityContainer) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IEdmEntityContainer) UnsetName() {
	o.Name.Unset()
}

func (o IEdmEntityContainer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IEdmEntityContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Elements != nil {
		toSerialize["elements"] = o.Elements
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
	return toSerialize, nil
}

type NullableIEdmEntityContainer struct {
	value *IEdmEntityContainer
	isSet bool
}

func (v NullableIEdmEntityContainer) Get() *IEdmEntityContainer {
	return v.value
}

func (v *NullableIEdmEntityContainer) Set(val *IEdmEntityContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableIEdmEntityContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableIEdmEntityContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIEdmEntityContainer(val *IEdmEntityContainer) *NullableIEdmEntityContainer {
	return &NullableIEdmEntityContainer{value: val, isSet: true}
}

func (v NullableIEdmEntityContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIEdmEntityContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


