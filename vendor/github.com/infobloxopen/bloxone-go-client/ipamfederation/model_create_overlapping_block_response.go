/*
IPAM Federation API

The DDI IPAM Federation application enables a SaaS administrator to manage multiple IPAM systems from one central control point CSP.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipamfederation

import (
	"encoding/json"
)

// checks if the CreateOverlappingBlockResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOverlappingBlockResponse{}

// CreateOverlappingBlockResponse The response format to create the __OverlappingBlock__ object.
type CreateOverlappingBlockResponse struct {
	// The created OverlappingBlock object.
	Result               *OverlappingBlock `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateOverlappingBlockResponse CreateOverlappingBlockResponse

// NewCreateOverlappingBlockResponse instantiates a new CreateOverlappingBlockResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOverlappingBlockResponse() *CreateOverlappingBlockResponse {
	this := CreateOverlappingBlockResponse{}
	return &this
}

// NewCreateOverlappingBlockResponseWithDefaults instantiates a new CreateOverlappingBlockResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOverlappingBlockResponseWithDefaults() *CreateOverlappingBlockResponse {
	this := CreateOverlappingBlockResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateOverlappingBlockResponse) GetResult() OverlappingBlock {
	if o == nil || IsNil(o.Result) {
		var ret OverlappingBlock
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOverlappingBlockResponse) GetResultOk() (*OverlappingBlock, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateOverlappingBlockResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given OverlappingBlock and assigns it to the Result field.
func (o *CreateOverlappingBlockResponse) SetResult(v OverlappingBlock) {
	o.Result = &v
}

func (o CreateOverlappingBlockResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOverlappingBlockResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateOverlappingBlockResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateOverlappingBlockResponse := _CreateOverlappingBlockResponse{}

	err = json.Unmarshal(data, &varCreateOverlappingBlockResponse)

	if err != nil {
		return err
	}

	*o = CreateOverlappingBlockResponse(varCreateOverlappingBlockResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateOverlappingBlockResponse struct {
	value *CreateOverlappingBlockResponse
	isSet bool
}

func (v NullableCreateOverlappingBlockResponse) Get() *CreateOverlappingBlockResponse {
	return v.value
}

func (v *NullableCreateOverlappingBlockResponse) Set(val *CreateOverlappingBlockResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOverlappingBlockResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOverlappingBlockResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOverlappingBlockResponse(val *CreateOverlappingBlockResponse) *NullableCreateOverlappingBlockResponse {
	return &NullableCreateOverlappingBlockResponse{value: val, isSet: true}
}

func (v NullableCreateOverlappingBlockResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOverlappingBlockResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}