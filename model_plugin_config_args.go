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

// PluginConfigArgs PluginConfigArgs plugin config args
type PluginConfigArgs struct {
	// description
	Description string `json:"Description"`
	// name
	Name string `json:"Name"`
	// settable
	Settable []string `json:"Settable"`
	// value
	Value []string `json:"Value"`
}

// NewPluginConfigArgs instantiates a new PluginConfigArgs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginConfigArgs(description string, name string, settable []string, value []string) *PluginConfigArgs {
	this := PluginConfigArgs{}
	this.Description = description
	this.Name = name
	this.Settable = settable
	this.Value = value
	return &this
}

// NewPluginConfigArgsWithDefaults instantiates a new PluginConfigArgs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginConfigArgsWithDefaults() *PluginConfigArgs {
	this := PluginConfigArgs{}
	return &this
}

// GetDescription returns the Description field value
func (o *PluginConfigArgs) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PluginConfigArgs) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PluginConfigArgs) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *PluginConfigArgs) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PluginConfigArgs) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PluginConfigArgs) SetName(v string) {
	o.Name = v
}

// GetSettable returns the Settable field value
func (o *PluginConfigArgs) GetSettable() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Settable
}

// GetSettableOk returns a tuple with the Settable field value
// and a boolean to check if the value has been set.
func (o *PluginConfigArgs) GetSettableOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Settable, true
}

// SetSettable sets field value
func (o *PluginConfigArgs) SetSettable(v []string) {
	o.Settable = v
}

// GetValue returns the Value field value
func (o *PluginConfigArgs) GetValue() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PluginConfigArgs) GetValueOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *PluginConfigArgs) SetValue(v []string) {
	o.Value = v
}

func (o PluginConfigArgs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Description"] = o.Description
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	if true {
		toSerialize["Settable"] = o.Settable
	}
	if true {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullablePluginConfigArgs struct {
	value *PluginConfigArgs
	isSet bool
}

func (v NullablePluginConfigArgs) Get() *PluginConfigArgs {
	return v.value
}

func (v *NullablePluginConfigArgs) Set(val *PluginConfigArgs) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginConfigArgs) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginConfigArgs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginConfigArgs(val *PluginConfigArgs) *NullablePluginConfigArgs {
	return &NullablePluginConfigArgs{value: val, isSet: true}
}

func (v NullablePluginConfigArgs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginConfigArgs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


