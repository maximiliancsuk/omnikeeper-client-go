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

// IEdmEntityContainerElement struct for IEdmEntityContainerElement
type IEdmEntityContainerElement struct {
	ContainerElementKind *EdmContainerElementKind `json:"containerElementKind,omitempty"`
	Container *IEdmEntityContainer `json:"container,omitempty"`
	Name NullableString `json:"name,omitempty"`
}

// NewIEdmEntityContainerElement instantiates a new IEdmEntityContainerElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIEdmEntityContainerElement() *IEdmEntityContainerElement {
	this := IEdmEntityContainerElement{}
	return &this
}

// NewIEdmEntityContainerElementWithDefaults instantiates a new IEdmEntityContainerElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIEdmEntityContainerElementWithDefaults() *IEdmEntityContainerElement {
	this := IEdmEntityContainerElement{}
	return &this
}

// GetContainerElementKind returns the ContainerElementKind field value if set, zero value otherwise.
func (o *IEdmEntityContainerElement) GetContainerElementKind() EdmContainerElementKind {
	if o == nil || o.ContainerElementKind == nil {
		var ret EdmContainerElementKind
		return ret
	}
	return *o.ContainerElementKind
}

// GetContainerElementKindOk returns a tuple with the ContainerElementKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmEntityContainerElement) GetContainerElementKindOk() (*EdmContainerElementKind, bool) {
	if o == nil || o.ContainerElementKind == nil {
		return nil, false
	}
	return o.ContainerElementKind, true
}

// HasContainerElementKind returns a boolean if a field has been set.
func (o *IEdmEntityContainerElement) HasContainerElementKind() bool {
	if o != nil && o.ContainerElementKind != nil {
		return true
	}

	return false
}

// SetContainerElementKind gets a reference to the given EdmContainerElementKind and assigns it to the ContainerElementKind field.
func (o *IEdmEntityContainerElement) SetContainerElementKind(v EdmContainerElementKind) {
	o.ContainerElementKind = &v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *IEdmEntityContainerElement) GetContainer() IEdmEntityContainer {
	if o == nil || o.Container == nil {
		var ret IEdmEntityContainer
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmEntityContainerElement) GetContainerOk() (*IEdmEntityContainer, bool) {
	if o == nil || o.Container == nil {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *IEdmEntityContainerElement) HasContainer() bool {
	if o != nil && o.Container != nil {
		return true
	}

	return false
}

// SetContainer gets a reference to the given IEdmEntityContainer and assigns it to the Container field.
func (o *IEdmEntityContainerElement) SetContainer(v IEdmEntityContainer) {
	o.Container = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IEdmEntityContainerElement) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IEdmEntityContainerElement) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IEdmEntityContainerElement) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IEdmEntityContainerElement) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IEdmEntityContainerElement) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IEdmEntityContainerElement) UnsetName() {
	o.Name.Unset()
}

func (o IEdmEntityContainerElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContainerElementKind != nil {
		toSerialize["containerElementKind"] = o.ContainerElementKind
	}
	if o.Container != nil {
		toSerialize["container"] = o.Container
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIEdmEntityContainerElement struct {
	value *IEdmEntityContainerElement
	isSet bool
}

func (v NullableIEdmEntityContainerElement) Get() *IEdmEntityContainerElement {
	return v.value
}

func (v *NullableIEdmEntityContainerElement) Set(val *IEdmEntityContainerElement) {
	v.value = val
	v.isSet = true
}

func (v NullableIEdmEntityContainerElement) IsSet() bool {
	return v.isSet
}

func (v *NullableIEdmEntityContainerElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIEdmEntityContainerElement(val *IEdmEntityContainerElement) *NullableIEdmEntityContainerElement {
	return &NullableIEdmEntityContainerElement{value: val, isSet: true}
}

func (v NullableIEdmEntityContainerElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIEdmEntityContainerElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


