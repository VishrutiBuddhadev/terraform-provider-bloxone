/*
BloxOne FW API

BloxOne Threat Defense Cloud is an extension of the BloxOne Suite that provides visibility into infected and compromised off-premises devices, roaming users, remote sites, and branch offices. You can subscribe to BloxOne Cloud and use its functionality to mitigate and control malware as well as provide unprecedented insight into your network security posture and enable timely action. BloxOne Cloud also offers unified policy management, reporting, and threat analytics across the entire spectrum. Using automated and high-quality threat intelligence feeds and unique behavioral analytics, it automatically stops device communications with C&Cs/botnets and prevents DNS based data exfiltration.  The mission-critical DNS infrastructure can become a vulnerable component in your network when it is inadequately protected by traditional security solutions and consequently used as an attack surface. Compromised DNS services can result in catastrophic network and system failures. To fully protect your network in today’s cyber security threat environment, Infoblox sets a new DNS security standard by offering scalable, enterprise-grade, and integrated protection for your DNS infrastructure.  Through the Infoblox Cloud Services Portal, you can view the status of your subscription and threat intelligence feeds, manage your network scope and roaming end users, and learn more about threats on your networks through the Infoblox Threat Lookup tool and predefined reports.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fw

import (
	"encoding/json"
)

// checks if the AppApprovalsUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppApprovalsUpdateRequest{}

// AppApprovalsUpdateRequest struct for AppApprovalsUpdateRequest
type AppApprovalsUpdateRequest struct {
	// Field Mask.
	Fields               *ProtobufFieldMask          `json:"fields,omitempty"`
	InsertedApprovals    []AppApproval               `json:"inserted_approvals,omitempty"`
	RemovedApprovals     []AppApprovalRemovalRequest `json:"removed_approvals,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppApprovalsUpdateRequest AppApprovalsUpdateRequest

// NewAppApprovalsUpdateRequest instantiates a new AppApprovalsUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppApprovalsUpdateRequest() *AppApprovalsUpdateRequest {
	this := AppApprovalsUpdateRequest{}
	return &this
}

// NewAppApprovalsUpdateRequestWithDefaults instantiates a new AppApprovalsUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppApprovalsUpdateRequestWithDefaults() *AppApprovalsUpdateRequest {
	this := AppApprovalsUpdateRequest{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *AppApprovalsUpdateRequest) GetFields() ProtobufFieldMask {
	if o == nil || IsNil(o.Fields) {
		var ret ProtobufFieldMask
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppApprovalsUpdateRequest) GetFieldsOk() (*ProtobufFieldMask, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *AppApprovalsUpdateRequest) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given ProtobufFieldMask and assigns it to the Fields field.
func (o *AppApprovalsUpdateRequest) SetFields(v ProtobufFieldMask) {
	o.Fields = &v
}

// GetInsertedApprovals returns the InsertedApprovals field value if set, zero value otherwise.
func (o *AppApprovalsUpdateRequest) GetInsertedApprovals() []AppApproval {
	if o == nil || IsNil(o.InsertedApprovals) {
		var ret []AppApproval
		return ret
	}
	return o.InsertedApprovals
}

// GetInsertedApprovalsOk returns a tuple with the InsertedApprovals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppApprovalsUpdateRequest) GetInsertedApprovalsOk() ([]AppApproval, bool) {
	if o == nil || IsNil(o.InsertedApprovals) {
		return nil, false
	}
	return o.InsertedApprovals, true
}

// HasInsertedApprovals returns a boolean if a field has been set.
func (o *AppApprovalsUpdateRequest) HasInsertedApprovals() bool {
	if o != nil && !IsNil(o.InsertedApprovals) {
		return true
	}

	return false
}

// SetInsertedApprovals gets a reference to the given []AppApproval and assigns it to the InsertedApprovals field.
func (o *AppApprovalsUpdateRequest) SetInsertedApprovals(v []AppApproval) {
	o.InsertedApprovals = v
}

// GetRemovedApprovals returns the RemovedApprovals field value if set, zero value otherwise.
func (o *AppApprovalsUpdateRequest) GetRemovedApprovals() []AppApprovalRemovalRequest {
	if o == nil || IsNil(o.RemovedApprovals) {
		var ret []AppApprovalRemovalRequest
		return ret
	}
	return o.RemovedApprovals
}

// GetRemovedApprovalsOk returns a tuple with the RemovedApprovals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppApprovalsUpdateRequest) GetRemovedApprovalsOk() ([]AppApprovalRemovalRequest, bool) {
	if o == nil || IsNil(o.RemovedApprovals) {
		return nil, false
	}
	return o.RemovedApprovals, true
}

// HasRemovedApprovals returns a boolean if a field has been set.
func (o *AppApprovalsUpdateRequest) HasRemovedApprovals() bool {
	if o != nil && !IsNil(o.RemovedApprovals) {
		return true
	}

	return false
}

// SetRemovedApprovals gets a reference to the given []AppApprovalRemovalRequest and assigns it to the RemovedApprovals field.
func (o *AppApprovalsUpdateRequest) SetRemovedApprovals(v []AppApprovalRemovalRequest) {
	o.RemovedApprovals = v
}

func (o AppApprovalsUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppApprovalsUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.InsertedApprovals) {
		toSerialize["inserted_approvals"] = o.InsertedApprovals
	}
	if !IsNil(o.RemovedApprovals) {
		toSerialize["removed_approvals"] = o.RemovedApprovals
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppApprovalsUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	varAppApprovalsUpdateRequest := _AppApprovalsUpdateRequest{}

	err = json.Unmarshal(data, &varAppApprovalsUpdateRequest)

	if err != nil {
		return err
	}

	*o = AppApprovalsUpdateRequest(varAppApprovalsUpdateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fields")
		delete(additionalProperties, "inserted_approvals")
		delete(additionalProperties, "removed_approvals")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppApprovalsUpdateRequest struct {
	value *AppApprovalsUpdateRequest
	isSet bool
}

func (v NullableAppApprovalsUpdateRequest) Get() *AppApprovalsUpdateRequest {
	return v.value
}

func (v *NullableAppApprovalsUpdateRequest) Set(val *AppApprovalsUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppApprovalsUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppApprovalsUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppApprovalsUpdateRequest(val *AppApprovalsUpdateRequest) *NullableAppApprovalsUpdateRequest {
	return &NullableAppApprovalsUpdateRequest{value: val, isSet: true}
}

func (v NullableAppApprovalsUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppApprovalsUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
