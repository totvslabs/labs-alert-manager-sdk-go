# NotificationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientUuid** | **string** | Notification Client UUID | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Events** | Pointer to **[]string** | List of notification events | [optional] 
**FirstEvent** | Pointer to **NullableTime** |  | [optional] 
**Id** | **string** | Id | 
**Retries** | Pointer to **NullableInt32** | Quantity notification retries | [optional] 
**Status** | **string** | Notification status | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewNotificationSchema

`func NewNotificationSchema(clientUuid string, id string, status string, ) *NotificationSchema`

NewNotificationSchema instantiates a new NotificationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSchemaWithDefaults

`func NewNotificationSchemaWithDefaults() *NotificationSchema`

NewNotificationSchemaWithDefaults instantiates a new NotificationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientUuid

`func (o *NotificationSchema) GetClientUuid() string`

GetClientUuid returns the ClientUuid field if non-nil, zero value otherwise.

### GetClientUuidOk

`func (o *NotificationSchema) GetClientUuidOk() (*string, bool)`

GetClientUuidOk returns a tuple with the ClientUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUuid

`func (o *NotificationSchema) SetClientUuid(v string)`

SetClientUuid sets ClientUuid field to given value.


### GetCreatedAt

`func (o *NotificationSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEvents

`func (o *NotificationSchema) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *NotificationSchema) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *NotificationSchema) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *NotificationSchema) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *NotificationSchema) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *NotificationSchema) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetFirstEvent

`func (o *NotificationSchema) GetFirstEvent() time.Time`

GetFirstEvent returns the FirstEvent field if non-nil, zero value otherwise.

### GetFirstEventOk

`func (o *NotificationSchema) GetFirstEventOk() (*time.Time, bool)`

GetFirstEventOk returns a tuple with the FirstEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstEvent

`func (o *NotificationSchema) SetFirstEvent(v time.Time)`

SetFirstEvent sets FirstEvent field to given value.

### HasFirstEvent

`func (o *NotificationSchema) HasFirstEvent() bool`

HasFirstEvent returns a boolean if a field has been set.

### SetFirstEventNil

`func (o *NotificationSchema) SetFirstEventNil(b bool)`

 SetFirstEventNil sets the value for FirstEvent to be an explicit nil

### UnsetFirstEvent
`func (o *NotificationSchema) UnsetFirstEvent()`

UnsetFirstEvent ensures that no value is present for FirstEvent, not even an explicit nil
### GetId

`func (o *NotificationSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationSchema) SetId(v string)`

SetId sets Id field to given value.


### GetRetries

`func (o *NotificationSchema) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *NotificationSchema) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *NotificationSchema) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *NotificationSchema) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *NotificationSchema) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *NotificationSchema) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetStatus

`func (o *NotificationSchema) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotificationSchema) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotificationSchema) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *NotificationSchema) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationSchema) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationSchema) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationSchema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


