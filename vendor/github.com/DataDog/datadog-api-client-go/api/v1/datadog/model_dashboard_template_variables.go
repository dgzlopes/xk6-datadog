/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// DashboardTemplateVariables Template variable.
type DashboardTemplateVariables struct {
	// The default value for the template variable on dashboard load.
	Default NullableString `json:"default,omitempty"`
	// The name of the variable.
	Name string `json:"name"`
	// The tag prefix associated with the variable. Only tags with this prefix appear in the variable drop-down.
	Prefix NullableString `json:"prefix,omitempty"`
}

// NewDashboardTemplateVariables instantiates a new DashboardTemplateVariables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardTemplateVariables(name string) *DashboardTemplateVariables {
	this := DashboardTemplateVariables{}
	this.Name = name
	return &this
}

// NewDashboardTemplateVariablesWithDefaults instantiates a new DashboardTemplateVariables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardTemplateVariablesWithDefaults() *DashboardTemplateVariables {
	this := DashboardTemplateVariables{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardTemplateVariables) GetDefault() string {
	if o == nil || o.Default.Get() == nil {
		var ret string
		return ret
	}
	return *o.Default.Get()
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DashboardTemplateVariables) GetDefaultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Default.Get(), o.Default.IsSet()
}

// HasDefault returns a boolean if a field has been set.
func (o *DashboardTemplateVariables) HasDefault() bool {
	if o != nil && o.Default.IsSet() {
		return true
	}

	return false
}

// SetDefault gets a reference to the given NullableString and assigns it to the Default field.
func (o *DashboardTemplateVariables) SetDefault(v string) {
	o.Default.Set(&v)
}

// SetDefaultNil sets the value for Default to be an explicit nil
func (o *DashboardTemplateVariables) SetDefaultNil() {
	o.Default.Set(nil)
}

// UnsetDefault ensures that no value is present for Default, not even an explicit nil
func (o *DashboardTemplateVariables) UnsetDefault() {
	o.Default.Unset()
}

// GetName returns the Name field value
func (o *DashboardTemplateVariables) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DashboardTemplateVariables) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DashboardTemplateVariables) SetName(v string) {
	o.Name = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardTemplateVariables) GetPrefix() string {
	if o == nil || o.Prefix.Get() == nil {
		var ret string
		return ret
	}
	return *o.Prefix.Get()
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DashboardTemplateVariables) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefix.Get(), o.Prefix.IsSet()
}

// HasPrefix returns a boolean if a field has been set.
func (o *DashboardTemplateVariables) HasPrefix() bool {
	if o != nil && o.Prefix.IsSet() {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given NullableString and assigns it to the Prefix field.
func (o *DashboardTemplateVariables) SetPrefix(v string) {
	o.Prefix.Set(&v)
}

// SetPrefixNil sets the value for Prefix to be an explicit nil
func (o *DashboardTemplateVariables) SetPrefixNil() {
	o.Prefix.Set(nil)
}

// UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
func (o *DashboardTemplateVariables) UnsetPrefix() {
	o.Prefix.Unset()
}

func (o DashboardTemplateVariables) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Default.IsSet() {
		toSerialize["default"] = o.Default.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Prefix.IsSet() {
		toSerialize["prefix"] = o.Prefix.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDashboardTemplateVariables struct {
	value *DashboardTemplateVariables
	isSet bool
}

func (v NullableDashboardTemplateVariables) Get() *DashboardTemplateVariables {
	return v.value
}

func (v *NullableDashboardTemplateVariables) Set(val *DashboardTemplateVariables) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardTemplateVariables) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardTemplateVariables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardTemplateVariables(val *DashboardTemplateVariables) *NullableDashboardTemplateVariables {
	return &NullableDashboardTemplateVariables{value: val, isSet: true}
}

func (v NullableDashboardTemplateVariables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardTemplateVariables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}