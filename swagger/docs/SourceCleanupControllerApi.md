# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSourceCleanupReport**](SourceCleanupControllerApi.md#DeleteSourceCleanupReport) | **Delete** /source-cleanup/reports/{migrationId} | Delete source cleanup report(s) for a given migration ID and report name
[**GetOrphanedCleanupReportsList**](SourceCleanupControllerApi.md#GetOrphanedCleanupReportsList) | **Get** /source-cleanup/orphaned-reports | Get a list of the metadata for all orphaned deletion reports
[**GetSourceCleanupReport**](SourceCleanupControllerApi.md#GetSourceCleanupReport) | **Get** /source-cleanup/reports/{migrationId} | Get a source cleanup report for a given migration ID and report name
[**GetSourceCleanupReportsList**](SourceCleanupControllerApi.md#GetSourceCleanupReportsList) | **Get** /source-cleanup/{migrationId}/reports | Get the metadata of the available source cleanup reports for the given Migration
[**PauseDelayedDeletionCacheReader**](SourceCleanupControllerApi.md#PauseDelayedDeletionCacheReader) | **Post** /source-cleanup/pauseDelayedDeletionCacheReader | Pause the Delayed Deletion Cache Reader from deleting files
[**ResumeDelayedDeletionCacheReader**](SourceCleanupControllerApi.md#ResumeDelayedDeletionCacheReader) | **Post** /source-cleanup/resumeDelayedDeletionCacheReader | Resume the Delayed Deletion Cache Reader

# **DeleteSourceCleanupReport**
> DeleteSourceCleanupReport(ctx, migrationId, reportName)
Delete source cleanup report(s) for a given migration ID and report name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the Migration object | 
  **reportName** | **string**| filename of the report to be returned | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrphanedCleanupReportsList**
> []SourceCleanupReportFileMetadata GetOrphanedCleanupReportsList(ctx, )
Get a list of the metadata for all orphaned deletion reports

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]SourceCleanupReportFileMetadata**](SourceCleanupReportFileMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourceCleanupReport**
> *os.File GetSourceCleanupReport(ctx, migrationId, reportName)
Get a source cleanup report for a given migration ID and report name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the Migration object | 
  **reportName** | **string**| filename of the report to be returned | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourceCleanupReportsList**
> []SourceCleanupReportFileMetadata GetSourceCleanupReportsList(ctx, migrationId)
Get the metadata of the available source cleanup reports for the given Migration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the Migration object | 

### Return type

[**[]SourceCleanupReportFileMetadata**](SourceCleanupReportFileMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PauseDelayedDeletionCacheReader**
> PauseDelayedDeletionCacheReader(ctx, )
Pause the Delayed Deletion Cache Reader from deleting files

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResumeDelayedDeletionCacheReader**
> ResumeDelayedDeletionCacheReader(ctx, )
Resume the Delayed Deletion Cache Reader

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

