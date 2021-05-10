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
)

// RecoveryAddress struct for RecoveryAddress
type RecoveryAddress struct {
	Id string `json:"id"`
	Value string `json:"value"`
	Via string `json:"via"`
}

// NewRecoveryAddress instantiates a new RecoveryAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryAddress(id string, value string, via string) *RecoveryAddress {
	this := RecoveryAddress{}
	this.Id = id
	this.Value = value
	this.Via = via
	return &this
}

// NewRecoveryAddressWithDefaults instantiates a new RecoveryAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryAddressWithDefaults() *RecoveryAddress {
	this := RecoveryAddress{}
	return &this
}

// GetId returns the Id field value
func (o *RecoveryAddress) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RecoveryAddress) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RecoveryAddress) SetId(v string) {
	o.Id = v
}

// GetValue returns the Value field value
func (o *RecoveryAddress) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RecoveryAddress) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *RecoveryAddress) SetValue(v string) {
	o.Value = v
}

// GetVia returns the Via field value
func (o *RecoveryAddress) GetVia() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Via
}

// GetViaOk returns a tuple with the Via field value
// and a boolean to check if the value has been set.
func (o *RecoveryAddress) GetViaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Via, true
}

// SetVia sets field value
func (o *RecoveryAddress) SetVia(v string) {
	o.Via = v
}

func (o RecoveryAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["via"] = o.Via
	}
	return json.Marshal(toSerialize)
}

type NullableRecoveryAddress struct {
	value *RecoveryAddress
	isSet bool
}

func (v NullableRecoveryAddress) Get() *RecoveryAddress {
	return v.value
}

func (v *NullableRecoveryAddress) Set(val *RecoveryAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryAddress(val *RecoveryAddress) *NullableRecoveryAddress {
	return &NullableRecoveryAddress{value: val, isSet: true}
}

func (v NullableRecoveryAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


