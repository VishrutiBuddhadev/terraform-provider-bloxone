/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsconfig

import (
	"encoding/json"
)

// checks if the DTCPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DTCPolicy{}

// DTCPolicy The __DTC Policy__ object.
type DTCPolicy struct {
	// __DTC Policy__ display name.
	Name *string `json:"name,omitempty"`
	// The resource identifier.
	PolicyId             *string `json:"policy_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DTCPolicy DTCPolicy

// NewDTCPolicy instantiates a new DTCPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDTCPolicy() *DTCPolicy {
	this := DTCPolicy{}
	return &this
}

// NewDTCPolicyWithDefaults instantiates a new DTCPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDTCPolicyWithDefaults() *DTCPolicy {
	this := DTCPolicy{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DTCPolicy) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCPolicy) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DTCPolicy) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DTCPolicy) SetName(v string) {
	o.Name = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *DTCPolicy) GetPolicyId() string {
	if o == nil || IsNil(o.PolicyId) {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCPolicy) GetPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyId) {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *DTCPolicy) HasPolicyId() bool {
	if o != nil && !IsNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *DTCPolicy) SetPolicyId(v string) {
	o.PolicyId = &v
}

func (o DTCPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DTCPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PolicyId) {
		toSerialize["policy_id"] = o.PolicyId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DTCPolicy) UnmarshalJSON(data []byte) (err error) {
	varDTCPolicy := _DTCPolicy{}

	err = json.Unmarshal(data, &varDTCPolicy)

	if err != nil {
		return err
	}

	*o = DTCPolicy(varDTCPolicy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "policy_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDTCPolicy struct {
	value *DTCPolicy
	isSet bool
}

func (v NullableDTCPolicy) Get() *DTCPolicy {
	return v.value
}

func (v *NullableDTCPolicy) Set(val *DTCPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableDTCPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableDTCPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDTCPolicy(val *DTCPolicy) *NullableDTCPolicy {
	return &NullableDTCPolicy{value: val, isSet: true}
}

func (v NullableDTCPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDTCPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}