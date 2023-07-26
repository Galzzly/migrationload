# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationBandwidthPolicy**](BandwidthPolicyControllerApi.md#GetApplicationBandwidthPolicy) | **Get** /bandwidthPolicies/application | Get the application&#x27;s bandwidth policy
[**ResetApplicationBandwidthPolicy**](BandwidthPolicyControllerApi.md#ResetApplicationBandwidthPolicy) | **Delete** /bandwidthPolicies/application | Restore default, unlimited application bandwidth policy
[**UpdateApplicationBandwidthPolicy**](BandwidthPolicyControllerApi.md#UpdateApplicationBandwidthPolicy) | **Post** /bandwidthPolicies/application | Set the application&#x27;s bandwidth policy. A limit of -1 will mean no bandwidth limit is imposed.

# **GetApplicationBandwidthPolicy**
> BandwidthPolicy GetApplicationBandwidthPolicy(ctx, )
Get the application's bandwidth policy

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BandwidthPolicy**](BandwidthPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetApplicationBandwidthPolicy**
> BandwidthPolicy ResetApplicationBandwidthPolicy(ctx, )
Restore default, unlimited application bandwidth policy

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BandwidthPolicy**](BandwidthPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApplicationBandwidthPolicy**
> BandwidthPolicy UpdateApplicationBandwidthPolicy(ctx, limit, unit)
Set the application's bandwidth policy. A limit of -1 will mean no bandwidth limit is imposed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **limit** | **int64**| Overall application bandwidth limit to apply. | 
  **unit** | [**FileSizeUnit**](.md)| Unit of bandwidth - bytes per second | 

### Return type

[**BandwidthPolicy**](BandwidthPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

