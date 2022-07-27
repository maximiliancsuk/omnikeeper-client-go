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

// ODataServiceDocument struct for ODataServiceDocument
type ODataServiceDocument struct {
	TypeAnnotation *ODataTypeAnnotation `json:"typeAnnotation,omitempty"`
	EntitySets []ODataEntitySetInfo `json:"entitySets,omitempty"`
	Singletons []ODataSingletonInfo `json:"singletons,omitempty"`
	FunctionImports []ODataFunctionImportInfo `json:"functionImports,omitempty"`
}

// NewODataServiceDocument instantiates a new ODataServiceDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewODataServiceDocument() *ODataServiceDocument {
	this := ODataServiceDocument{}
	return &this
}

// NewODataServiceDocumentWithDefaults instantiates a new ODataServiceDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewODataServiceDocumentWithDefaults() *ODataServiceDocument {
	this := ODataServiceDocument{}
	return &this
}

// GetTypeAnnotation returns the TypeAnnotation field value if set, zero value otherwise.
func (o *ODataServiceDocument) GetTypeAnnotation() ODataTypeAnnotation {
	if o == nil || o.TypeAnnotation == nil {
		var ret ODataTypeAnnotation
		return ret
	}
	return *o.TypeAnnotation
}

// GetTypeAnnotationOk returns a tuple with the TypeAnnotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ODataServiceDocument) GetTypeAnnotationOk() (*ODataTypeAnnotation, bool) {
	if o == nil || o.TypeAnnotation == nil {
		return nil, false
	}
	return o.TypeAnnotation, true
}

// HasTypeAnnotation returns a boolean if a field has been set.
func (o *ODataServiceDocument) HasTypeAnnotation() bool {
	if o != nil && o.TypeAnnotation != nil {
		return true
	}

	return false
}

// SetTypeAnnotation gets a reference to the given ODataTypeAnnotation and assigns it to the TypeAnnotation field.
func (o *ODataServiceDocument) SetTypeAnnotation(v ODataTypeAnnotation) {
	o.TypeAnnotation = &v
}

// GetEntitySets returns the EntitySets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataServiceDocument) GetEntitySets() []ODataEntitySetInfo {
	if o == nil {
		var ret []ODataEntitySetInfo
		return ret
	}
	return o.EntitySets
}

// GetEntitySetsOk returns a tuple with the EntitySets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataServiceDocument) GetEntitySetsOk() ([]ODataEntitySetInfo, bool) {
	if o == nil || o.EntitySets == nil {
		return nil, false
	}
	return o.EntitySets, true
}

// HasEntitySets returns a boolean if a field has been set.
func (o *ODataServiceDocument) HasEntitySets() bool {
	if o != nil && o.EntitySets != nil {
		return true
	}

	return false
}

// SetEntitySets gets a reference to the given []ODataEntitySetInfo and assigns it to the EntitySets field.
func (o *ODataServiceDocument) SetEntitySets(v []ODataEntitySetInfo) {
	o.EntitySets = v
}

// GetSingletons returns the Singletons field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataServiceDocument) GetSingletons() []ODataSingletonInfo {
	if o == nil {
		var ret []ODataSingletonInfo
		return ret
	}
	return o.Singletons
}

// GetSingletonsOk returns a tuple with the Singletons field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataServiceDocument) GetSingletonsOk() ([]ODataSingletonInfo, bool) {
	if o == nil || o.Singletons == nil {
		return nil, false
	}
	return o.Singletons, true
}

// HasSingletons returns a boolean if a field has been set.
func (o *ODataServiceDocument) HasSingletons() bool {
	if o != nil && o.Singletons != nil {
		return true
	}

	return false
}

// SetSingletons gets a reference to the given []ODataSingletonInfo and assigns it to the Singletons field.
func (o *ODataServiceDocument) SetSingletons(v []ODataSingletonInfo) {
	o.Singletons = v
}

// GetFunctionImports returns the FunctionImports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataServiceDocument) GetFunctionImports() []ODataFunctionImportInfo {
	if o == nil {
		var ret []ODataFunctionImportInfo
		return ret
	}
	return o.FunctionImports
}

// GetFunctionImportsOk returns a tuple with the FunctionImports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataServiceDocument) GetFunctionImportsOk() ([]ODataFunctionImportInfo, bool) {
	if o == nil || o.FunctionImports == nil {
		return nil, false
	}
	return o.FunctionImports, true
}

// HasFunctionImports returns a boolean if a field has been set.
func (o *ODataServiceDocument) HasFunctionImports() bool {
	if o != nil && o.FunctionImports != nil {
		return true
	}

	return false
}

// SetFunctionImports gets a reference to the given []ODataFunctionImportInfo and assigns it to the FunctionImports field.
func (o *ODataServiceDocument) SetFunctionImports(v []ODataFunctionImportInfo) {
	o.FunctionImports = v
}

func (o ODataServiceDocument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TypeAnnotation != nil {
		toSerialize["typeAnnotation"] = o.TypeAnnotation
	}
	if o.EntitySets != nil {
		toSerialize["entitySets"] = o.EntitySets
	}
	if o.Singletons != nil {
		toSerialize["singletons"] = o.Singletons
	}
	if o.FunctionImports != nil {
		toSerialize["functionImports"] = o.FunctionImports
	}
	return json.Marshal(toSerialize)
}

type NullableODataServiceDocument struct {
	value *ODataServiceDocument
	isSet bool
}

func (v NullableODataServiceDocument) Get() *ODataServiceDocument {
	return v.value
}

func (v *NullableODataServiceDocument) Set(val *ODataServiceDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableODataServiceDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableODataServiceDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableODataServiceDocument(val *ODataServiceDocument) *NullableODataServiceDocument {
	return &NullableODataServiceDocument{value: val, isSet: true}
}

func (v NullableODataServiceDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableODataServiceDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

