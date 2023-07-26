# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelVerification**](VerificationControllerApi.md#CancelVerification) | **Post** /verifications/{verificationId}/cancel | Cancel verification with given ID
[**DeleteVerification**](VerificationControllerApi.md#DeleteVerification) | **Delete** /verifications/{verificationId} | Delete verification with given ID
[**FetchAllFilteredVerifications**](VerificationControllerApi.md#FetchAllFilteredVerifications) | **Get** /verifications/filtered | Get all Verifications filtered by migrationId and/or verification status
[**FetchAllVerifications**](VerificationControllerApi.md#FetchAllVerifications) | **Get** /verifications | Get all Verifications
[**GetVerification**](VerificationControllerApi.md#GetVerification) | **Get** /verifications/{verificationId} | Get verification for given ID
[**GetVerificationReport**](VerificationControllerApi.md#GetVerificationReport) | **Get** /verifications/reports/{verificationId} | Get verification report for a given ID
[**GetVerificationReportsList**](VerificationControllerApi.md#GetVerificationReportsList) | **Get** /verifications/{verificationId}/reports | Get the metadata for the available reports for the given verification
[**TriggerVerification**](VerificationControllerApi.md#TriggerVerification) | **Put** /verifications | Trigger a new Verification

# **CancelVerification**
> CancelVerification(ctx, verificationId)
Cancel verification with given ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **verificationId** | **string**| Unique identifier of the verification Object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVerification**
> DeleteVerification(ctx, verificationId)
Delete verification with given ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **verificationId** | **string**| Unique identifier of the verification Object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchAllFilteredVerifications**
> []Verification FetchAllFilteredVerifications(ctx, optional)
Get all Verifications filtered by migrationId and/or verification status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VerificationControllerApiFetchAllFilteredVerificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VerificationControllerApiFetchAllFilteredVerificationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **migrationIds** | [**optional.Interface of []string**](string.md)| If this parameter is supplied only verifications for the given migrationIds will be returned | [default to []]
 **states** | [**optional.Interface of []string**](string.md)| If this parameter is supplied, only verifications in the specified states will be returned | [default to []]

### Return type

[**[]Verification**](Verification.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchAllVerifications**
> []Verification FetchAllVerifications(ctx, optional)
Get all Verifications

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VerificationControllerApiFetchAllVerificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VerificationControllerApiFetchAllVerificationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **migrationId** | **optional.String**| If this parameter is supplied only verifications for the given migrationId will be returned | 

### Return type

[**[]Verification**](Verification.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVerification**
> Verification GetVerification(ctx, verificationId)
Get verification for given ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **verificationId** | **string**| Unique identifier of the verification Object | 

### Return type

[**Verification**](Verification.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVerificationReport**
> *os.File GetVerificationReport(ctx, verificationId, optional)
Get verification report for a given ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **verificationId** | **string**| Unique identifier of the verification whose report is to be returned | 
 **optional** | ***VerificationControllerApiGetVerificationReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VerificationControllerApiGetVerificationReportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportId** | **optional.String**| Identifier of the report to be returned | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-tar, application/json, application/gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVerificationReportsList**
> []VerificationReportFileMetadata GetVerificationReportsList(ctx, verificationId)
Get the metadata for the available reports for the given verification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **verificationId** | **string**| Unique identifier of the verification Object | 

### Return type

[**[]VerificationReportFileMetadata**](VerificationReportFileMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerVerification**
> Verification TriggerVerification(ctx, optional)
Trigger a new Verification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VerificationControllerApiTriggerVerificationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VerificationControllerApiTriggerVerificationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of VerificationConfiguration**](VerificationConfiguration.md)|  | 

### Return type

[**Verification**](Verification.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

