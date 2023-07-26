# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackup**](BackupsControllerApi.md#CreateBackup) | **Post** /backups | Generate a Backup.
[**GetBackupByName**](BackupsControllerApi.md#GetBackupByName) | **Get** /backups/{backupName} | 
[**GetBackupConfig**](BackupsControllerApi.md#GetBackupConfig) | **Get** /backups/config | backup config entity
[**GetBackupSchedule**](BackupsControllerApi.md#GetBackupSchedule) | **Get** /backups/config/schedule | Returns current schedule configuration
[**GetBackupsList**](BackupsControllerApi.md#GetBackupsList) | **Get** /backups | list of available LM2 backups
[**RestoreBackup**](BackupsControllerApi.md#RestoreBackup) | **Post** /backups/restore/{backupName} | Restore application state by backup by provided backup ID
[**UpdateSchedule**](BackupsControllerApi.md#UpdateSchedule) | **Put** /backups/config/schedule | Update current schedule.

# **CreateBackup**
> BackupEntry CreateBackup(ctx, )
Generate a Backup.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BackupEntry**](BackupEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupByName**
> BackupEntry GetBackupByName(ctx, backupName)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupName** | **string**| provided backup file name | 

### Return type

[**BackupEntry**](BackupEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupConfig**
> BackupConfigInfo GetBackupConfig(ctx, )
backup config entity

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BackupConfigInfo**](BackupConfigInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupSchedule**
> BackupSchedule GetBackupSchedule(ctx, )
Returns current schedule configuration

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BackupSchedule**](BackupSchedule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupsList**
> []BackupEntry GetBackupsList(ctx, optional)
list of available LM2 backups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BackupsControllerApiGetBackupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupsControllerApiGetBackupsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Size limit of returned backups list. | 

### Return type

[**[]BackupEntry**](BackupEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreBackup**
> BackupEntry RestoreBackup(ctx, backupName)
Restore application state by backup by provided backup ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupName** | **string**| Unique identifier of backup entry | 

### Return type

[**BackupEntry**](BackupEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSchedule**
> BackupSchedule UpdateSchedule(ctx, body)
Update current schedule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BackupSchedule**](BackupSchedule.md)| Backup schedule entity for update | 

### Return type

[**BackupSchedule**](BackupSchedule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

