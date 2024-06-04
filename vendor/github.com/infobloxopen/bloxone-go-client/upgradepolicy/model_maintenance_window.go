/*
Schedule Software/Config Updates

Infoblox by default does automatic software updates when they become available. Updates are applied to all on-prem hosts, physical or virtual. However, you can override and schedule the software updates. You can also defer the updates to a later date and time. You can configure up to a total of 50 deferrals (scheduled and deferred software updates), which means you have the flexibility to create up to 50 update groups across different on-prem hosts by mapping with appropriate tags. Tags are be used to associate deferrals (scheduled or deferred) with a specific or group of onprem-hosts. Apart from software update deferrals, config update deferrals also can be configured using these overrides.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package upgradepolicy

import (
	"encoding/json"
)

// checks if the MaintenanceWindow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaintenanceWindow{}

// MaintenanceWindow struct for MaintenanceWindow
type MaintenanceWindow struct {
	DeferredWindow       *DeferredWindow                   `json:"deferred_window,omitempty"`
	Id                   *string                           `json:"id,omitempty"`
	ScheduledWindow      *ScheduledWindow                  `json:"scheduled_window,omitempty"`
	Tags                 map[string]map[string]interface{} `json:"tags,omitempty"`
	WindowType           *string                           `json:"window_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MaintenanceWindow MaintenanceWindow

// NewMaintenanceWindow instantiates a new MaintenanceWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaintenanceWindow() *MaintenanceWindow {
	this := MaintenanceWindow{}
	return &this
}

// NewMaintenanceWindowWithDefaults instantiates a new MaintenanceWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaintenanceWindowWithDefaults() *MaintenanceWindow {
	this := MaintenanceWindow{}
	return &this
}

// GetDeferredWindow returns the DeferredWindow field value if set, zero value otherwise.
func (o *MaintenanceWindow) GetDeferredWindow() DeferredWindow {
	if o == nil || IsNil(o.DeferredWindow) {
		var ret DeferredWindow
		return ret
	}
	return *o.DeferredWindow
}

// GetDeferredWindowOk returns a tuple with the DeferredWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetDeferredWindowOk() (*DeferredWindow, bool) {
	if o == nil || IsNil(o.DeferredWindow) {
		return nil, false
	}
	return o.DeferredWindow, true
}

// HasDeferredWindow returns a boolean if a field has been set.
func (o *MaintenanceWindow) HasDeferredWindow() bool {
	if o != nil && !IsNil(o.DeferredWindow) {
		return true
	}

	return false
}

// SetDeferredWindow gets a reference to the given DeferredWindow and assigns it to the DeferredWindow field.
func (o *MaintenanceWindow) SetDeferredWindow(v DeferredWindow) {
	o.DeferredWindow = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MaintenanceWindow) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MaintenanceWindow) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MaintenanceWindow) SetId(v string) {
	o.Id = &v
}

// GetScheduledWindow returns the ScheduledWindow field value if set, zero value otherwise.
func (o *MaintenanceWindow) GetScheduledWindow() ScheduledWindow {
	if o == nil || IsNil(o.ScheduledWindow) {
		var ret ScheduledWindow
		return ret
	}
	return *o.ScheduledWindow
}

// GetScheduledWindowOk returns a tuple with the ScheduledWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetScheduledWindowOk() (*ScheduledWindow, bool) {
	if o == nil || IsNil(o.ScheduledWindow) {
		return nil, false
	}
	return o.ScheduledWindow, true
}

// HasScheduledWindow returns a boolean if a field has been set.
func (o *MaintenanceWindow) HasScheduledWindow() bool {
	if o != nil && !IsNil(o.ScheduledWindow) {
		return true
	}

	return false
}

// SetScheduledWindow gets a reference to the given ScheduledWindow and assigns it to the ScheduledWindow field.
func (o *MaintenanceWindow) SetScheduledWindow(v ScheduledWindow) {
	o.ScheduledWindow = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MaintenanceWindow) GetTags() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetTagsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MaintenanceWindow) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]map[string]interface{} and assigns it to the Tags field.
func (o *MaintenanceWindow) SetTags(v map[string]map[string]interface{}) {
	o.Tags = v
}

// GetWindowType returns the WindowType field value if set, zero value otherwise.
func (o *MaintenanceWindow) GetWindowType() string {
	if o == nil || IsNil(o.WindowType) {
		var ret string
		return ret
	}
	return *o.WindowType
}

// GetWindowTypeOk returns a tuple with the WindowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetWindowTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WindowType) {
		return nil, false
	}
	return o.WindowType, true
}

// HasWindowType returns a boolean if a field has been set.
func (o *MaintenanceWindow) HasWindowType() bool {
	if o != nil && !IsNil(o.WindowType) {
		return true
	}

	return false
}

// SetWindowType gets a reference to the given string and assigns it to the WindowType field.
func (o *MaintenanceWindow) SetWindowType(v string) {
	o.WindowType = &v
}

func (o MaintenanceWindow) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaintenanceWindow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeferredWindow) {
		toSerialize["deferred_window"] = o.DeferredWindow
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ScheduledWindow) {
		toSerialize["scheduled_window"] = o.ScheduledWindow
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.WindowType) {
		toSerialize["window_type"] = o.WindowType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MaintenanceWindow) UnmarshalJSON(data []byte) (err error) {
	varMaintenanceWindow := _MaintenanceWindow{}

	err = json.Unmarshal(data, &varMaintenanceWindow)

	if err != nil {
		return err
	}

	*o = MaintenanceWindow(varMaintenanceWindow)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "deferred_window")
		delete(additionalProperties, "id")
		delete(additionalProperties, "scheduled_window")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "window_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMaintenanceWindow struct {
	value *MaintenanceWindow
	isSet bool
}

func (v NullableMaintenanceWindow) Get() *MaintenanceWindow {
	return v.value
}

func (v *NullableMaintenanceWindow) Set(val *MaintenanceWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableMaintenanceWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableMaintenanceWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaintenanceWindow(val *MaintenanceWindow) *NullableMaintenanceWindow {
	return &NullableMaintenanceWindow{value: val, isSet: true}
}

func (v NullableMaintenanceWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaintenanceWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}