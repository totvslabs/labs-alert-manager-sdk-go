/*
alertmanager

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package labs_alert_manager_client

import (
	"encoding/json"
	"time"
)

// checks if the PolicySchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicySchema{}

// PolicySchema struct for PolicySchema
type PolicySchema struct {
	// List of channel notification
	Channels map[string]interface{} `json:"channels"`
	// Policy Client source
	ClientSource string `json:"client_source"`
	// Policy Client UUID
	ClientUuid string `json:"client_uuid"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Policy deleted
	Deleted bool `json:"deleted"`
	// Policy enabled
	Enabled bool `json:"enabled"`
	// Policy filters
	Filters map[string]interface{} `json:"filters"`
	// Enable alert frequency for the policy
	Frequency bool `json:"frequency"`
	// Quantity of alert interval time
	FrequencyMinutes int32 `json:"frequency_minutes"`
	// Quantity of alert occurrencies
	FrequencyOccurrences int32 `json:"frequency_occurrences"`
	// Id
	Id string `json:"id"`
	// Policy labels
	Labels map[string]interface{} `json:"labels"`
	// Policy name
	Name string `json:"name"`
	// Policy severity
	Severity string `json:"severity"`
	// Policy type
	Type string `json:"type"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewPolicySchema instantiates a new PolicySchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicySchema(channels map[string]interface{}, clientSource string, clientUuid string, deleted bool, enabled bool, filters map[string]interface{}, frequency bool, frequencyMinutes int32, frequencyOccurrences int32, id string, labels map[string]interface{}, name string, severity string, type_ string) *PolicySchema {
	this := PolicySchema{}
	this.Channels = channels
	this.ClientSource = clientSource
	this.ClientUuid = clientUuid
	this.Deleted = deleted
	this.Enabled = enabled
	this.Filters = filters
	this.Frequency = frequency
	this.FrequencyMinutes = frequencyMinutes
	this.FrequencyOccurrences = frequencyOccurrences
	this.Id = id
	this.Labels = labels
	this.Name = name
	this.Severity = severity
	this.Type = type_
	return &this
}

// NewPolicySchemaWithDefaults instantiates a new PolicySchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicySchemaWithDefaults() *PolicySchema {
	this := PolicySchema{}
	return &this
}

// GetChannels returns the Channels field value
func (o *PolicySchema) GetChannels() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value
// and a boolean to check if the value has been set.
func (o *PolicySchema) GetChannelsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Channels, true
}

// SetChannels sets field value
func (o *PolicySchema) SetChannels(v map[string]interface{}) {
	o.Channels = v
}

// GetClientSource returns the ClientSource field value
func (o *PolicySchema) GetClientSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSource
}

// GetClientSourceOk returns a tuple with the ClientSource field value
// and a boolean to check if the value has been set.
func (o *PolicySchema) GetClientSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSource, true
}

// SetClientSource sets field value
func (o *PolicySchema) SetClientSource(v string) {
	o.ClientSource = v
}

// GetClientUuid returns the ClientUuid field value
func (o *PolicySchema) GetClientUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientUuid
}

// GetClientUuidOk returns a tuple with the ClientUuid field value
// and a boolean to check if the value has been set.
func (o *PolicySchema) GetClientUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientUuid, true
}

// SetClientUuid sets field value
func (o *PolicySchema) SetClientUuid(v string) {
	o.ClientUuid = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PolicySchema) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PolicySchema) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PolicySchema) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeleted returns the Deleted field value
func (o *PolicySchema) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *PolicySchema) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *PolicySchema) SetDeleted(v bool) {
	o.Deleted = v
}

// GetEnabled returns the Enabled field value
func (o *PolicySchema) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PolicySchema) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PolicySchema) SetEnabled(v bool) {
	o.Enabled = v
}

// GetFilters returns the Filters field value
func (o *PolicySchema) GetFilters() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *PolicySchema) GetFiltersOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Filters, true
}

// SetFilters sets field value
func (o *PolicySchema) SetFilters(v map[string]interface{}) {
	o.Filters = v
}

// GetFrequency returns the Frequency field value
func (o *PolicySchema) GetFrequency() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *PolicySchema) GetFrequencyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *PolicySchema) SetFrequency(v bool) {
	o.Frequency = v
}

// GetFrequencyMinutes returns the FrequencyMinutes field value
func (o *PolicySchema) GetFrequencyMinutes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FrequencyMinutes
}

// GetFrequencyMinutesOk returns a tuple with the FrequencyMinutes field value
// and a boolean to check if the value has been set.
func (o *PolicySchema) GetFrequencyMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FrequencyMinutes, true
}

// SetFrequencyMinutes sets field value
func (o *PolicySchema) SetFrequencyMinutes(v int32) {
	o.FrequencyMinutes = v
}

// GetFrequencyOccurrences returns the FrequencyOccurrences field value
func (o *PolicySchema) GetFrequencyOccurrences() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FrequencyOccurrences
}

// GetFrequencyOccurrencesOk returns a tuple with the FrequencyOccurrences field value
// and a boolean to check if the value has been set.
func (o *PolicySchema) GetFrequencyOccurrencesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FrequencyOccurrences, true
}

// SetFrequencyOccurrences sets field value
func (o *PolicySchema) SetFrequencyOccurrences(v int32) {
	o.FrequencyOccurrences = v
}

// GetId returns the Id field value
func (o *PolicySchema) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PolicySchema) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PolicySchema) SetId(v string) {
	o.Id = v
}

// GetLabels returns the Labels field value
func (o *PolicySchema) GetLabels() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *PolicySchema) GetLabelsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *PolicySchema) SetLabels(v map[string]interface{}) {
	o.Labels = v
}

// GetName returns the Name field value
func (o *PolicySchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PolicySchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PolicySchema) SetName(v string) {
	o.Name = v
}

// GetSeverity returns the Severity field value
func (o *PolicySchema) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *PolicySchema) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *PolicySchema) SetSeverity(v string) {
	o.Severity = v
}

// GetType returns the Type field value
func (o *PolicySchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PolicySchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PolicySchema) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PolicySchema) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySchema) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PolicySchema) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PolicySchema) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o PolicySchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicySchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channels"] = o.Channels
	toSerialize["client_source"] = o.ClientSource
	toSerialize["client_uuid"] = o.ClientUuid
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	toSerialize["deleted"] = o.Deleted
	toSerialize["enabled"] = o.Enabled
	toSerialize["filters"] = o.Filters
	toSerialize["frequency"] = o.Frequency
	toSerialize["frequency_minutes"] = o.FrequencyMinutes
	toSerialize["frequency_occurrences"] = o.FrequencyOccurrences
	toSerialize["id"] = o.Id
	toSerialize["labels"] = o.Labels
	toSerialize["name"] = o.Name
	toSerialize["severity"] = o.Severity
	toSerialize["type"] = o.Type
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullablePolicySchema struct {
	value *PolicySchema
	isSet bool
}

func (v NullablePolicySchema) Get() *PolicySchema {
	return v.value
}

func (v *NullablePolicySchema) Set(val *PolicySchema) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicySchema) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicySchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicySchema(val *PolicySchema) *NullablePolicySchema {
	return &NullablePolicySchema{value: val, isSet: true}
}

func (v NullablePolicySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicySchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


