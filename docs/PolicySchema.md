# PolicySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | **map[string]interface{}** | List of channel notification | 
**ClientSource** | **string** | Policy Client source | 
**ClientUuid** | **string** | Policy Client UUID | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Deleted** | **bool** | Policy deleted | 
**Enabled** | **bool** | Policy enabled | 
**Filters** | **map[string]interface{}** | Policy filters | 
**Frequency** | **bool** | Enable alert frequency for the policy | 
**FrequencyMinutes** | **int32** | Quantity of alert interval time | 
**FrequencyOccurrences** | **int32** | Quantity of alert occurrencies | 
**Id** | **string** | Id | 
**Labels** | **map[string]interface{}** | Policy labels | 
**Name** | **string** | Policy name | 
**Severity** | **string** | Policy severity | 
**Type** | **string** | Policy type | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPolicySchema

`func NewPolicySchema(channels map[string]interface{}, clientSource string, clientUuid string, deleted bool, enabled bool, filters map[string]interface{}, frequency bool, frequencyMinutes int32, frequencyOccurrences int32, id string, labels map[string]interface{}, name string, severity string, type_ string, ) *PolicySchema`

NewPolicySchema instantiates a new PolicySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicySchemaWithDefaults

`func NewPolicySchemaWithDefaults() *PolicySchema`

NewPolicySchemaWithDefaults instantiates a new PolicySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *PolicySchema) GetChannels() map[string]interface{}`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *PolicySchema) GetChannelsOk() (*map[string]interface{}, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *PolicySchema) SetChannels(v map[string]interface{})`

SetChannels sets Channels field to given value.


### GetClientSource

`func (o *PolicySchema) GetClientSource() string`

GetClientSource returns the ClientSource field if non-nil, zero value otherwise.

### GetClientSourceOk

`func (o *PolicySchema) GetClientSourceOk() (*string, bool)`

GetClientSourceOk returns a tuple with the ClientSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSource

`func (o *PolicySchema) SetClientSource(v string)`

SetClientSource sets ClientSource field to given value.


### GetClientUuid

`func (o *PolicySchema) GetClientUuid() string`

GetClientUuid returns the ClientUuid field if non-nil, zero value otherwise.

### GetClientUuidOk

`func (o *PolicySchema) GetClientUuidOk() (*string, bool)`

GetClientUuidOk returns a tuple with the ClientUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUuid

`func (o *PolicySchema) SetClientUuid(v string)`

SetClientUuid sets ClientUuid field to given value.


### GetCreatedAt

`func (o *PolicySchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PolicySchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PolicySchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PolicySchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeleted

`func (o *PolicySchema) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PolicySchema) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PolicySchema) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetEnabled

`func (o *PolicySchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PolicySchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PolicySchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFilters

`func (o *PolicySchema) GetFilters() map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PolicySchema) GetFiltersOk() (*map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PolicySchema) SetFilters(v map[string]interface{})`

SetFilters sets Filters field to given value.


### GetFrequency

`func (o *PolicySchema) GetFrequency() bool`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *PolicySchema) GetFrequencyOk() (*bool, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *PolicySchema) SetFrequency(v bool)`

SetFrequency sets Frequency field to given value.


### GetFrequencyMinutes

`func (o *PolicySchema) GetFrequencyMinutes() int32`

GetFrequencyMinutes returns the FrequencyMinutes field if non-nil, zero value otherwise.

### GetFrequencyMinutesOk

`func (o *PolicySchema) GetFrequencyMinutesOk() (*int32, bool)`

GetFrequencyMinutesOk returns a tuple with the FrequencyMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyMinutes

`func (o *PolicySchema) SetFrequencyMinutes(v int32)`

SetFrequencyMinutes sets FrequencyMinutes field to given value.


### GetFrequencyOccurrences

`func (o *PolicySchema) GetFrequencyOccurrences() int32`

GetFrequencyOccurrences returns the FrequencyOccurrences field if non-nil, zero value otherwise.

### GetFrequencyOccurrencesOk

`func (o *PolicySchema) GetFrequencyOccurrencesOk() (*int32, bool)`

GetFrequencyOccurrencesOk returns a tuple with the FrequencyOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyOccurrences

`func (o *PolicySchema) SetFrequencyOccurrences(v int32)`

SetFrequencyOccurrences sets FrequencyOccurrences field to given value.


### GetId

`func (o *PolicySchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicySchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicySchema) SetId(v string)`

SetId sets Id field to given value.


### GetLabels

`func (o *PolicySchema) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PolicySchema) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PolicySchema) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.


### GetName

`func (o *PolicySchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicySchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicySchema) SetName(v string)`

SetName sets Name field to given value.


### GetSeverity

`func (o *PolicySchema) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *PolicySchema) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *PolicySchema) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetType

`func (o *PolicySchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicySchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicySchema) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *PolicySchema) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PolicySchema) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PolicySchema) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PolicySchema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


