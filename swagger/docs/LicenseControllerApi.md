# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchConsumptionLicense**](LicenseControllerApi.md#FetchConsumptionLicense) | **Get** /license/components/consumption | Get Consumption License Component for LiveData Migrator
[**FetchLicense**](LicenseControllerApi.md#FetchLicense) | **Get** /license | Get License for LiveData Migrator
[**FetchLicenseComponents**](LicenseControllerApi.md#FetchLicenseComponents) | **Get** /license/components | Get License Components for LiveData Migrator
[**FetchLicenseStatus**](LicenseControllerApi.md#FetchLicenseStatus) | **Get** /license/status | Get License Status for LiveData Migrator
[**FetchLicenseValid**](LicenseControllerApi.md#FetchLicenseValid) | **Get** /license/valid | Get whether License is valid for LiveData Migrator
[**FetchMeteredLicense**](LicenseControllerApi.md#FetchMeteredLicense) | **Get** /license/components/metered | Get Metered License Component for LiveData Migrator
[**FetchTrialLicense**](LicenseControllerApi.md#FetchTrialLicense) | **Get** /license/components/trial | Get Trial License Component for LiveData Migrator
[**FetchVolumeLicense**](LicenseControllerApi.md#FetchVolumeLicense) | **Get** /license/components/volume | Get Volume License Component for LiveData Migrator
[**UpdateLicense**](LicenseControllerApi.md#UpdateLicense) | **Post** /license | Update License for LiveData Migrator

# **FetchConsumptionLicense**
> ConsumptionLicense FetchConsumptionLicense(ctx, )
Get Consumption License Component for LiveData Migrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConsumptionLicense**](ConsumptionLicense.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchLicense**
> License FetchLicense(ctx, optional)
Get License for LiveData Migrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LicenseControllerApiFetchLicenseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LicenseControllerApiFetchLicenseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fullLicense** | **optional.Bool**| Flag return complete license information | [default to false]

### Return type

[**License**](License.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchLicenseComponents**
> []LicenseComponent FetchLicenseComponents(ctx, )
Get License Components for LiveData Migrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]LicenseComponent**](LicenseComponent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchLicenseStatus**
> string FetchLicenseStatus(ctx, )
Get License Status for LiveData Migrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchLicenseValid**
> string FetchLicenseValid(ctx, )
Get whether License is valid for LiveData Migrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchMeteredLicense**
> MeteredLicense FetchMeteredLicense(ctx, )
Get Metered License Component for LiveData Migrator

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

# **FetchTrialLicense**
> TrialLicense FetchTrialLicense(ctx, )
Get Trial License Component for LiveData Migrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TrialLicense**](TrialLicense.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchVolumeLicense**
> DataLicense FetchVolumeLicense(ctx, )
Get Volume License Component for LiveData Migrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DataLicense**](DataLicense.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLicense**
> License UpdateLicense(ctx, optional)
Update License for LiveData Migrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LicenseControllerApiUpdateLicenseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LicenseControllerApiUpdateLicenseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **optional.Interface of *os.File****optional.**|  | 

### Return type

[**License**](License.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

