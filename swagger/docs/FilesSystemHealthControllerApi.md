# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHealthStatus**](FilesSystemHealthControllerApi.md#GetHealthStatus) | **Get** /fs/health/{fsId} | Get health status for given ID
[**ListHealthStatuses**](FilesSystemHealthControllerApi.md#ListHealthStatuses) | **Get** /fs/health | List health statues for all filesystems

# **GetHealthStatus**
> FileSystemHealthStatus GetHealthStatus(ctx, fsId)
Get health status for given ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fsId** | **string**| Unique identifier of the source filesystem | 

### Return type

[**FileSystemHealthStatus**](FileSystemHealthStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHealthStatuses**
> []FileSystemHealthStatus ListHealthStatuses(ctx, )
List health statues for all filesystems

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]FileSystemHealthStatus**](FileSystemHealthStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

