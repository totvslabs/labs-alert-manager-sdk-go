# EventSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSource** | **string** | Event souce | 
**ClientUuid** | **string** | Client uuid. This is the id defined by client app | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Data** | **map[string]interface{}** | Event data | 
**EventType** | **string** | Event type | 
**Id** | **string** | Id | 
**Labels** | **map[string]interface{}** | Event labels | 
**SchemaVersion** | **string** | Event schema version. Can be used by client app to know how to parse the event | 
**Severity** | **string** | Event severity | 
**Status** | **string** | Event status, Received, Stored | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEventSchema

`func NewEventSchema(clientSource string, clientUuid string, data map[string]interface{}, eventType string, id string, labels map[string]interface{}, schemaVersion string, severity string, status string, ) *EventSchema`

NewEventSchema instantiates a new EventSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSchemaWithDefaults

`func NewEventSchemaWithDefaults() *EventSchema`

NewEventSchemaWithDefaults instantiates a new EventSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientSource

`func (o *EventSchema) GetClientSource() string`

GetClientSource returns the ClientSource field if non-nil, zero value otherwise.

### GetClientSourceOk

`func (o *EventSchema) GetClientSourceOk() (*string, bool)`

GetClientSourceOk returns a tuple with the ClientSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSource

`func (o *EventSchema) SetClientSource(v string)`

SetClientSource sets ClientSource field to given value.


### GetClientUuid

`func (o *EventSchema) GetClientUuid() string`

GetClientUuid returns the ClientUuid field if non-nil, zero value otherwise.

### GetClientUuidOk

`func (o *EventSchema) GetClientUuidOk() (*string, bool)`

GetClientUuidOk returns a tuple with the ClientUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUuid

`func (o *EventSchema) SetClientUuid(v string)`

SetClientUuid sets ClientUuid field to given value.


### GetCreatedAt

`func (o *EventSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EventSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetData

`func (o *EventSchema) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventSchema) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventSchema) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetEventType

`func (o *EventSchema) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventSchema) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventSchema) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetId

`func (o *EventSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventSchema) SetId(v string)`

SetId sets Id field to given value.


### GetLabels

`func (o *EventSchema) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EventSchema) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EventSchema) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.


### GetSchemaVersion

`func (o *EventSchema) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *EventSchema) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *EventSchema) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetSeverity

`func (o *EventSchema) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *EventSchema) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *EventSchema) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetStatus

`func (o *EventSchema) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventSchema) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventSchema) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *EventSchema) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EventSchema) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EventSchema) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EventSchema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


