# PaginationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int32** |  | 
**PageSize** | **int32** |  | 
**Rows** | **[]map[string]interface{}** |  | 
**TotalRows** | **int32** |  | 

## Methods

### NewPaginationSchema

`func NewPaginationSchema(page int32, pageSize int32, rows []map[string]interface{}, totalRows int32, ) *PaginationSchema`

NewPaginationSchema instantiates a new PaginationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationSchemaWithDefaults

`func NewPaginationSchemaWithDefaults() *PaginationSchema`

NewPaginationSchemaWithDefaults instantiates a new PaginationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PaginationSchema) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PaginationSchema) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PaginationSchema) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *PaginationSchema) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PaginationSchema) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PaginationSchema) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetRows

`func (o *PaginationSchema) GetRows() []map[string]interface{}`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *PaginationSchema) GetRowsOk() (*[]map[string]interface{}, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *PaginationSchema) SetRows(v []map[string]interface{})`

SetRows sets Rows field to given value.


### GetTotalRows

`func (o *PaginationSchema) GetTotalRows() int32`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *PaginationSchema) GetTotalRowsOk() (*int32, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *PaginationSchema) SetTotalRows(v int32)`

SetTotalRows sets TotalRows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


