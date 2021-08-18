/*
Landscape omnikeeper REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package okclient

import (
	"encoding/json"
	"fmt"
)

// RelationState the model 'RelationState'
type RelationState string

// List of RelationState
const (
	RELATIONSTATE_NEW RelationState = "New"
	RELATIONSTATE_REMOVED RelationState = "Removed"
	RELATIONSTATE_RENEWED RelationState = "Renewed"
)

var allowedRelationStateEnumValues = []RelationState{
	"New",
	"Removed",
	"Renewed",
}

func (v *RelationState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RelationState(value)
	for _, existing := range allowedRelationStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RelationState", value)
}

// NewRelationStateFromValue returns a pointer to a valid RelationState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRelationStateFromValue(v string) (*RelationState, error) {
	ev := RelationState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RelationState: valid values are %v", v, allowedRelationStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RelationState) IsValid() bool {
	for _, existing := range allowedRelationStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RelationState value
func (v RelationState) Ptr() *RelationState {
	return &v
}

type NullableRelationState struct {
	value *RelationState
	isSet bool
}

func (v NullableRelationState) Get() *RelationState {
	return v.value
}

func (v *NullableRelationState) Set(val *RelationState) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationState) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationState(val *RelationState) *NullableRelationState {
	return &NullableRelationState{value: val, isSet: true}
}

func (v NullableRelationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

