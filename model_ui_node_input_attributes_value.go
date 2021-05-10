/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.6.0-alpha.11
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// UiNodeInputAttributesValue - struct for UiNodeInputAttributesValue
type UiNodeInputAttributesValue struct {
	Bool *bool
	Float32 *float32
	String *string
}

// boolAsUiNodeInputAttributesValue is a convenience function that returns bool wrapped in UiNodeInputAttributesValue
func BoolAsUiNodeInputAttributesValue(v *bool) UiNodeInputAttributesValue {
	return UiNodeInputAttributesValue{
		Bool: v,
	}
}

// float32AsUiNodeInputAttributesValue is a convenience function that returns float32 wrapped in UiNodeInputAttributesValue
func Float32AsUiNodeInputAttributesValue(v *float32) UiNodeInputAttributesValue {
	return UiNodeInputAttributesValue{
		Float32: v,
	}
}

// stringAsUiNodeInputAttributesValue is a convenience function that returns string wrapped in UiNodeInputAttributesValue
func StringAsUiNodeInputAttributesValue(v *string) UiNodeInputAttributesValue {
	return UiNodeInputAttributesValue{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *UiNodeInputAttributesValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			match++
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into Float32
	err = newStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			match++
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Bool = nil
		dst.Float32 = nil
		dst.String = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(UiNodeInputAttributesValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(UiNodeInputAttributesValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UiNodeInputAttributesValue) MarshalJSON() ([]byte, error) {
	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UiNodeInputAttributesValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.Float32 != nil {
		return obj.Float32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUiNodeInputAttributesValue struct {
	value *UiNodeInputAttributesValue
	isSet bool
}

func (v NullableUiNodeInputAttributesValue) Get() *UiNodeInputAttributesValue {
	return v.value
}

func (v *NullableUiNodeInputAttributesValue) Set(val *UiNodeInputAttributesValue) {
	v.value = val
	v.isSet = true
}

func (v NullableUiNodeInputAttributesValue) IsSet() bool {
	return v.isSet
}

func (v *NullableUiNodeInputAttributesValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiNodeInputAttributesValue(val *UiNodeInputAttributesValue) *NullableUiNodeInputAttributesValue {
	return &NullableUiNodeInputAttributesValue{value: val, isSet: true}
}

func (v NullableUiNodeInputAttributesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiNodeInputAttributesValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


