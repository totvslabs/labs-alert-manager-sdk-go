# PolicyChannelSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientUuid** | **string** | Client uuid. This is the id defined by client app | 
**Config** | **map[string]interface{}** | Policy Channel config | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Deleted** | **bool** | Policy Channel deleted flag | 
**Enabled** | **bool** | Policy Channel enabled flag | 
**Id** | **string** | Id | 
**LastNotification** | Pointer to **NullableTime** |  | [optional] 
**Name** | **string** | Policy Channel name | 
**Type** | **string** | Policy Channel type | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPolicyChannelSchema

`func NewPolicyChannelSchema(clientUuid string, config map[string]interface{}, deleted bool, enabled bool, id string, name string, type_ string, ) *PolicyChannelSchema`

NewPolicyChannelSchema instantiates a new PolicyChannelSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyChannelSchemaWithDefaults

`func NewPolicyChannelSchemaWithDefaults() *PolicyChannelSchema`

NewPolicyChannelSchemaWithDefaults instantiates a new PolicyChannelSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientUuid

`func (o *PolicyChannelSchema) GetClientUuid() string`

GetClientUuid returns the ClientUuid field if non-nil, zero value otherwise.

### GetClientUuidOk

`func (o *PolicyChannelSchema) GetClientUuidOk() (*string, bool)`

GetClientUuidOk returns a tuple with the ClientUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUuid

`func (o *PolicyChannelSchema) SetClientUuid(v string)`

SetClientUuid sets ClientUuid field to given value.


### GetConfig

`func (o *PolicyChannelSchema) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PolicyChannelSchema) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PolicyChannelSchema) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetCreatedAt

`func (o *PolicyChannelSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PolicyChannelSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PolicyChannelSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PolicyChannelSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeleted

`func (o *PolicyChannelSchema) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PolicyChannelSchema) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PolicyChannelSchema) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetEnabled

`func (o *PolicyChannelSchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PolicyChannelSchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PolicyChannelSchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *PolicyChannelSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyChannelSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyChannelSchema) SetId(v string)`

SetId sets Id field to given value.


### GetLastNotification

`func (o *PolicyChannelSchema) GetLastNotification() time.Time`

GetLastNotification returns the LastNotification field if non-nil, zero value otherwise.

### GetLastNotificationOk

`func (o *PolicyChannelSchema) GetLastNotificationOk() (*time.Time, bool)`

GetLastNotificationOk returns a tuple with the LastNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNotification

`func (o *PolicyChannelSchema) SetLastNotification(v time.Time)`

SetLastNotification sets LastNotification field to given value.

### HasLastNotification

`func (o *PolicyChannelSchema) HasLastNotification() bool`

HasLastNotification returns a boolean if a field has been set.

### SetLastNotificationNil

`func (o *PolicyChannelSchema) SetLastNotificationNil(b bool)`

 SetLastNotificationNil sets the value for LastNotification to be an explicit nil

### UnsetLastNotification
`func (o *PolicyChannelSchema) UnsetLastNotification()`

UnsetLastNotification ensures that no value is present for LastNotification, not even an explicit nil
### GetName

`func (o *PolicyChannelSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyChannelSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyChannelSchema) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *PolicyChannelSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicyChannelSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicyChannelSchema) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *PolicyChannelSchema) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PolicyChannelSchema) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PolicyChannelSchema) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PolicyChannelSchema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


