# ClientAppSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Contract** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | **string** |  | 
**Token** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**BillingId** | Pointer to **string** |  | [optional] 

## Methods

### NewClientAppSchema

`func NewClientAppSchema(id string, ) *ClientAppSchema`

NewClientAppSchema instantiates a new ClientAppSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAppSchemaWithDefaults

`func NewClientAppSchemaWithDefaults() *ClientAppSchema`

NewClientAppSchemaWithDefaults instantiates a new ClientAppSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAt

`func (o *ClientAppSchema) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ClientAppSchema) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ClientAppSchema) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ClientAppSchema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *ClientAppSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientAppSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientAppSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientAppSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeleted

`func (o *ClientAppSchema) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ClientAppSchema) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ClientAppSchema) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ClientAppSchema) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetContract

`func (o *ClientAppSchema) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *ClientAppSchema) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *ClientAppSchema) SetContract(v string)`

SetContract sets Contract field to given value.

### HasContract

`func (o *ClientAppSchema) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetLabels

`func (o *ClientAppSchema) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClientAppSchema) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClientAppSchema) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ClientAppSchema) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ClientAppSchema) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ClientAppSchema) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetId

`func (o *ClientAppSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientAppSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientAppSchema) SetId(v string)`

SetId sets Id field to given value.


### GetToken

`func (o *ClientAppSchema) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ClientAppSchema) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ClientAppSchema) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ClientAppSchema) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ClientAppSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClientAppSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClientAppSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ClientAppSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *ClientAppSchema) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ClientAppSchema) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ClientAppSchema) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ClientAppSchema) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetBillingId

`func (o *ClientAppSchema) GetBillingId() string`

GetBillingId returns the BillingId field if non-nil, zero value otherwise.

### GetBillingIdOk

`func (o *ClientAppSchema) GetBillingIdOk() (*string, bool)`

GetBillingIdOk returns a tuple with the BillingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingId

`func (o *ClientAppSchema) SetBillingId(v string)`

SetBillingId sets BillingId field to given value.

### HasBillingId

`func (o *ClientAppSchema) HasBillingId() bool`

HasBillingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


