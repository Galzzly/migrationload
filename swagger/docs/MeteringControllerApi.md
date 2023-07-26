# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchMeteringLicense**](MeteringControllerApi.md#FetchMeteringLicense) | **Get** /metering/license | Get Metered License for LiveData Migrator
[**FetchMeteringStats**](MeteringControllerApi.md#FetchMeteringStats) | **Get** /metering/stats | Get metering statistics for LiveData Migrator
[**GetMeteringStatsAndExtendLicense**](MeteringControllerApi.md#GetMeteringStatsAndExtendLicense) | **Put** /metering | Get metering statistics for LiveData Migrator and extend license 

# **FetchMeteringLicense**
> MeteredLicense FetchMeteringLicense(ctx, )
Get Metered License for LiveData Migrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MeteredLicense**](MeteredLicense.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchMeteringStats**
> Metering FetchMeteringStats(ctx, optional)
Get metering statistics for LiveData Migrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeteringControllerApiFetchMeteringStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeteringControllerApiFetchMeteringStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **useMinutes** | **optional.Bool**| Flag return Metering buckets in minutes | 

### Return type

[**Metering**](Metering.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMeteringStatsAndExtendLicense**
> Metering GetMeteringStatsAndExtendLicense(ctx, optional)
Get metering statistics for LiveData Migrator and extend license 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeteringControllerApiGetMeteringStatsAndExtendLicenseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeteringControllerApiGetMeteringStatsAndExtendLicenseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **useMinutes** | **optional.Bool**| Flag return Metering buckets in minutes | 
 **overrideExtensionTime** | **optional.Int64**| Override extension period. | [default to -1]

### Return type

[**Metering**](Metering.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

