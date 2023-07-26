# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDateExclusion**](ExclusionsControllerApi.md#AddDateExclusion) | **Put** /exclusions/date/{id} | Add a new date-based exclusion
[**AddFileSizeExclusion**](ExclusionsControllerApi.md#AddFileSizeExclusion) | **Put** /exclusions/fileSize/{id} | Add a new file size exclusion
[**AddRegexExclusion**](ExclusionsControllerApi.md#AddRegexExclusion) | **Put** /exclusions/regex/{id} | Add a new regex exclusion
[**GetExclusion**](ExclusionsControllerApi.md#GetExclusion) | **Get** /exclusions/{exclusionId} | Get exclusion for given exclusion ID
[**GetFsRestrictions**](ExclusionsControllerApi.md#GetFsRestrictions) | **Get** /exclusions/target/{fsType} | See the restrictions applied automatically for the given target filesystem
[**GetSourceFsRestrictions**](ExclusionsControllerApi.md#GetSourceFsRestrictions) | **Get** /exclusions/source/{fsType} | See the restrictions applied automatically for the given source filesystem
[**ListExclusions**](ExclusionsControllerApi.md#ListExclusions) | **Get** /exclusions | Get all exclusions
[**ListGlobalExclusions**](ExclusionsControllerApi.md#ListGlobalExclusions) | **Get** /exclusions/global | Get all global exclusions
[**ListUserDefinedExclusions**](ExclusionsControllerApi.md#ListUserDefinedExclusions) | **Get** /exclusions/userDefined | Get all user defined exclusions
[**RemoveExclusion**](ExclusionsControllerApi.md#RemoveExclusion) | **Delete** /exclusions/{exclusionId} | Delete exclusion for given ID

# **AddDateExclusion**
> Exclusion AddDateExclusion(ctx, id, description, beforeIsoDateTime)
Add a new date-based exclusion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of the exclusion to create | 
  **description** | **string**| Description of the exclusion to create | 
  **beforeIsoDateTime** | **time.Time**| If a file&#x27;s modification time is earlier than this, it will be excluded from migration. Expects the ISO format e.g. \&quot;2020-02-03T10:15:05-05:00\&quot;. | 

### Return type

[**Exclusion**](Exclusion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddFileSizeExclusion**
> Exclusion AddFileSizeExclusion(ctx, id, description, value, unit)
Add a new file size exclusion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of the exclusion to create | 
  **description** | **string**| Description of the exclusion to create | 
  **value** | **int64**| File size, if greater, file will be excluded from migration | 
  **unit** | [**FileSizeUnit**](.md)| Unit of file size | 

### Return type

[**Exclusion**](Exclusion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRegexExclusion**
> Exclusion AddRegexExclusion(ctx, id, description, regex, optional)
Add a new regex exclusion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of the exclusion to create | 
  **description** | **string**| Description of the exclusion to create | 
  **regex** | **string**| Regex string, if a file&#x27;s name matches, it will be excluded from migration | 
 **optional** | ***ExclusionsControllerApiAddRegexExclusionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExclusionsControllerApiAddRegexExclusionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patternType** | **optional.String**| Load configuration from files that are accessible from server | [default to JAVA_PCRE]

### Return type

[**Exclusion**](Exclusion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExclusion**
> Exclusion GetExclusion(ctx, exclusionId)
Get exclusion for given exclusion ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **exclusionId** | **string**|  | 

### Return type

[**Exclusion**](Exclusion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFsRestrictions**
> []Exclusion GetFsRestrictions(ctx, fsType)
See the restrictions applied automatically for the given target filesystem

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fsType** | **string**| Filesystem type | 

### Return type

[**[]Exclusion**](Exclusion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourceFsRestrictions**
> []Exclusion GetSourceFsRestrictions(ctx, fsType)
See the restrictions applied automatically for the given source filesystem

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fsType** | **string**| Filesystem type | 

### Return type

[**[]Exclusion**](Exclusion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListExclusions**
> []Exclusion ListExclusions(ctx, )
Get all exclusions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Exclusion**](Exclusion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGlobalExclusions**
> []Exclusion ListGlobalExclusions(ctx, )
Get all global exclusions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Exclusion**](Exclusion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserDefinedExclusions**
> []Exclusion ListUserDefinedExclusions(ctx, )
Get all user defined exclusions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Exclusion**](Exclusion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveExclusion**
> RemoveExclusion(ctx, exclusionId)
Delete exclusion for given ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **exclusionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

