/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the CreateIPSpaceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateIPSpaceResponse{}

// CreateIPSpaceResponse The response format to create the __IPSpace__ object.
type CreateIPSpaceResponse struct {
	// The created IP Space object.
	Result               *IPSpace `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateIPSpaceResponse CreateIPSpaceResponse

// NewCreateIPSpaceResponse instantiates a new CreateIPSpaceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIPSpaceResponse() *CreateIPSpaceResponse {
	this := CreateIPSpaceResponse{}
	return &this
}

// NewCreateIPSpaceResponseWithDefaults instantiates a new CreateIPSpaceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIPSpaceResponseWithDefaults() *CreateIPSpaceResponse {
	this := CreateIPSpaceResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateIPSpaceResponse) GetResult() IPSpace {
	if o == nil || IsNil(o.Result) {
		var ret IPSpace
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIPSpaceResponse) GetResultOk() (*IPSpace, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateIPSpaceResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given IPSpace and assigns it to the Result field.
func (o *CreateIPSpaceResponse) SetResult(v IPSpace) {
	o.Result = &v
}

func (o CreateIPSpaceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateIPSpaceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateIPSpaceResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateIPSpaceResponse := _CreateIPSpaceResponse{}

	err = json.Unmarshal(data, &varCreateIPSpaceResponse)

	if err != nil {
		return err
	}

	*o = CreateIPSpaceResponse(varCreateIPSpaceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateIPSpaceResponse struct {
	value *CreateIPSpaceResponse
	isSet bool
}

func (v NullableCreateIPSpaceResponse) Get() *CreateIPSpaceResponse {
	return v.value
}

func (v *NullableCreateIPSpaceResponse) Set(val *CreateIPSpaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIPSpaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIPSpaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIPSpaceResponse(val *CreateIPSpaceResponse) *NullableCreateIPSpaceResponse {
	return &NullableCreateIPSpaceResponse{value: val, isSet: true}
}

func (v NullableCreateIPSpaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIPSpaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
