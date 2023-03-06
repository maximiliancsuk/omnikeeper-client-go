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

// checks if the GraphQLQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphQLQuery{}

// GraphQLQuery struct for GraphQLQuery
type GraphQLQuery struct {
	OperationName NullableString `json:"operationName,omitempty"`
	Query NullableString `json:"query,omitempty"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// NewGraphQLQuery instantiates a new GraphQLQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphQLQuery() *GraphQLQuery {
	this := GraphQLQuery{}
	return &this
}

// NewGraphQLQueryWithDefaults instantiates a new GraphQLQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphQLQueryWithDefaults() *GraphQLQuery {
	this := GraphQLQuery{}
	return &this
}

// GetOperationName returns the OperationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphQLQuery) GetOperationName() string {
	if o == nil || IsNil(o.OperationName.Get()) {
		var ret string
		return ret
	}
	return *o.OperationName.Get()
}

// GetOperationNameOk returns a tuple with the OperationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphQLQuery) GetOperationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperationName.Get(), o.OperationName.IsSet()
}

// HasOperationName returns a boolean if a field has been set.
func (o *GraphQLQuery) HasOperationName() bool {
	if o != nil && o.OperationName.IsSet() {
		return true
	}

	return false
}

// SetOperationName gets a reference to the given NullableString and assigns it to the OperationName field.
func (o *GraphQLQuery) SetOperationName(v string) {
	o.OperationName.Set(&v)
}
// SetOperationNameNil sets the value for OperationName to be an explicit nil
func (o *GraphQLQuery) SetOperationNameNil() {
	o.OperationName.Set(nil)
}

// UnsetOperationName ensures that no value is present for OperationName, not even an explicit nil
func (o *GraphQLQuery) UnsetOperationName() {
	o.OperationName.Unset()
}

// GetQuery returns the Query field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphQLQuery) GetQuery() string {
	if o == nil || IsNil(o.Query.Get()) {
		var ret string
		return ret
	}
	return *o.Query.Get()
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphQLQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Query.Get(), o.Query.IsSet()
}

// HasQuery returns a boolean if a field has been set.
func (o *GraphQLQuery) HasQuery() bool {
	if o != nil && o.Query.IsSet() {
		return true
	}

	return false
}

// SetQuery gets a reference to the given NullableString and assigns it to the Query field.
func (o *GraphQLQuery) SetQuery(v string) {
	o.Query.Set(&v)
}
// SetQueryNil sets the value for Query to be an explicit nil
func (o *GraphQLQuery) SetQueryNil() {
	o.Query.Set(nil)
}

// UnsetQuery ensures that no value is present for Query, not even an explicit nil
func (o *GraphQLQuery) UnsetQuery() {
	o.Query.Unset()
}

// GetVariables returns the Variables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphQLQuery) GetVariables() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphQLQuery) GetVariablesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Variables) {
		return map[string]interface{}{}, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *GraphQLQuery) HasVariables() bool {
	if o != nil && IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]interface{} and assigns it to the Variables field.
func (o *GraphQLQuery) SetVariables(v map[string]interface{}) {
	o.Variables = v
}

func (o GraphQLQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphQLQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OperationName.IsSet() {
		toSerialize["operationName"] = o.OperationName.Get()
	}
	if o.Query.IsSet() {
		toSerialize["query"] = o.Query.Get()
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	return toSerialize, nil
}

type NullableGraphQLQuery struct {
	value *GraphQLQuery
	isSet bool
}

func (v NullableGraphQLQuery) Get() *GraphQLQuery {
	return v.value
}

func (v *NullableGraphQLQuery) Set(val *GraphQLQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphQLQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphQLQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphQLQuery(val *GraphQLQuery) *NullableGraphQLQuery {
	return &NullableGraphQLQuery{value: val, isSet: true}
}

func (v NullableGraphQLQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphQLQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


