# \EventsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEventsParametersDelete**](EventsAPI.md#DeleteEventsParametersDelete) | **Delete** /events/parameters/{app_key} | 
[**GetEventsGet**](EventsAPI.md#GetEventsGet) | **Get** /events/{id} | 
[**GetEventsGetAll**](EventsAPI.md#GetEventsGetAll) | **Get** /events | 
[**GetEventsParametersGet**](EventsAPI.md#GetEventsParametersGet) | **Get** /events/parameters | 
[**PostEventsPost**](EventsAPI.md#PostEventsPost) | **Post** /events | 



## DeleteEventsParametersDelete

> string DeleteEventsParametersDelete(ctx, appKey).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    appKey := "appKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.DeleteEventsParametersDelete(context.Background(), appKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.DeleteEventsParametersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEventsParametersDelete`: string
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.DeleteEventsParametersDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventsParametersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[auth_token](../README.md#auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsGet

> EventSchema GetEventsGet(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.GetEventsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEventsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventsGet`: EventSchema
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEventsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventSchema**](EventSchema.md)

### Authorization

[auth_token](../README.md#auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsGetAll

> PaginationSchema GetEventsGetAll(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.GetEventsGetAll(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEventsGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventsGetAll`: PaginationSchema
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEventsGetAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsGetAllRequest struct via the builder pattern


### Return type

[**PaginationSchema**](PaginationSchema.md)

### Authorization

[auth_token](../README.md#auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsParametersGet

> []EventSchema GetEventsParametersGet(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.GetEventsParametersGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEventsParametersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventsParametersGet`: []EventSchema
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEventsParametersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsParametersGetRequest struct via the builder pattern


### Return type

[**[]EventSchema**](EventSchema.md)

### Authorization

[auth_token](../README.md#auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEventsPost

> EventSchema PostEventsPost(ctx).RequestBody(requestBody).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    requestBody := *openapiclient.NewEventSchema("ClientSource_example", "ClientUuid_example", map[string]interface{}(123), "EventType_example", "Id_example", map[string]interface{}(123), "SchemaVersion_example", "Severity_example", "Status_example") // EventSchema |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.PostEventsPost(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostEventsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostEventsPost`: EventSchema
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.PostEventsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEventsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | [**EventSchema**](EventSchema.md) |  | 

### Return type

[**EventSchema**](EventSchema.md)

### Authorization

[auth_token](../README.md#auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

