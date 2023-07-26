# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAllCurrentMigrationStats**](StatisticsControllerApi.md#FetchAllCurrentMigrationStats) | **Get** /stats/current | Get current Stats for all migrations
[**FetchCurrentMigrationStats**](StatisticsControllerApi.md#FetchCurrentMigrationStats) | **Get** /stats/current/{migrationId} | Get current Stats for given migration ID
[**FetchDiagnostics**](StatisticsControllerApi.md#FetchDiagnostics) | **Get** /stats/filetrackers | Get File Trackers
[**FetchDiagnosticsActive**](StatisticsControllerApi.md#FetchDiagnosticsActive) | **Get** /stats/filetrackers/active | Get File Trackers for active transfers
[**FetchDiagnosticsActiveId**](StatisticsControllerApi.md#FetchDiagnosticsActiveId) | **Get** /stats/filetrackers/active/{progressId} | Get File Tracker for an active transfer
[**FetchDiagnosticsActiveIdBuckets**](StatisticsControllerApi.md#FetchDiagnosticsActiveIdBuckets) | **Get** /stats/filetrackers/active/{progressId}/buckets | Get File Tracker Bytes per Second Buckets
[**FetchDiagnosticsComplete**](StatisticsControllerApi.md#FetchDiagnosticsComplete) | **Get** /stats/filetrackers/complete | Get File Trackers for Completed Transfers.
[**FetchFileSizeDistribution**](StatisticsControllerApi.md#FetchFileSizeDistribution) | **Get** /stats/filetrackers/fileSizes | Get Transferred File Size Distribution.
[**FetchFileSizeDistributionActive**](StatisticsControllerApi.md#FetchFileSizeDistributionActive) | **Get** /stats/filetrackers/active/fileSizes | Get Active Transfer File Size Distribution.
[**FetchFileTransferRatesDistribution**](StatisticsControllerApi.md#FetchFileTransferRatesDistribution) | **Get** /stats/filetrackers/fileTransferRates | Get File Transfer Rates Distribution.
[**FetchFileTransferRatesDistributionActive**](StatisticsControllerApi.md#FetchFileTransferRatesDistributionActive) | **Get** /stats/filetrackers/active/fileTransferRates | Get Active File Transfer Rates Distribution.
[**FetchMigrationInfo**](StatisticsControllerApi.md#FetchMigrationInfo) | **Get** /stats/migrationInfo/{migrationId} | Get stats information for the given migration ID
[**FetchMigrationStats**](StatisticsControllerApi.md#FetchMigrationStats) | **Get** /stats/{migrationId} | Get Stats over time for given migration ID
[**FetchMigrationSummary**](StatisticsControllerApi.md#FetchMigrationSummary) | **Get** /stats/migrationSummary | Get migrations summary
[**FetchThroughputInLastSecs**](StatisticsControllerApi.md#FetchThroughputInLastSecs) | **Get** /stats/throughputInLastSecs/{nSecs} | Get throughput stats for all migrations during last N seconds
[**FetchThroughputSummary**](StatisticsControllerApi.md#FetchThroughputSummary) | **Get** /stats/throughputSummary | Get throughput summary
[**FetchTotalRollingStats**](StatisticsControllerApi.md#FetchTotalRollingStats) | **Get** /stats | Get rolling stats over time for all migrations
[**FetchTotalStats**](StatisticsControllerApi.md#FetchTotalStats) | **Get** /stats/cumulative | Get cumulative total stats for all migrations
[**LastExcludedEvents**](StatisticsControllerApi.md#LastExcludedEvents) | **Get** /stats/excludedEvents | Get most recent 1000 excluded events.

# **FetchAllCurrentMigrationStats**
> []MigrationRecord FetchAllCurrentMigrationStats(ctx, )
Get current Stats for all migrations

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]MigrationRecord**](MigrationRecord.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchCurrentMigrationStats**
> MigrationRecord FetchCurrentMigrationStats(ctx, migrationId)
Get current Stats for given migration ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 

### Return type

[**MigrationRecord**](MigrationRecord.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchDiagnostics**
> []FileTracker FetchDiagnostics(ctx, )
Get File Trackers

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]FileTracker**](FileTracker.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchDiagnosticsActive**
> []FileTracker FetchDiagnosticsActive(ctx, )
Get File Trackers for active transfers

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]FileTracker**](FileTracker.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchDiagnosticsActiveId**
> FileTracker FetchDiagnosticsActiveId(ctx, progressId)
Get File Tracker for an active transfer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **progressId** | **int64**| Identifier for File Tracker | 

### Return type

[**FileTracker**](FileTracker.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchDiagnosticsActiveIdBuckets**
> []FileTrackerBucket FetchDiagnosticsActiveIdBuckets(ctx, progressId)
Get File Tracker Bytes per Second Buckets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **progressId** | **int64**| Identifier for File Tracker | 

### Return type

[**[]FileTrackerBucket**](FileTrackerBucket.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchDiagnosticsComplete**
> []FileTracker FetchDiagnosticsComplete(ctx, )
Get File Trackers for Completed Transfers.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]FileTracker**](FileTracker.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchFileSizeDistribution**
> FileTrackerPercentiles FetchFileSizeDistribution(ctx, optional)
Get Transferred File Size Distribution.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsControllerApiFetchFileSizeDistributionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsControllerApiFetchFileSizeDistributionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **optional.Int32**| Percentile range | [default to 10]
 **migrationId** | **optional.String**| Migration Id | 

### Return type

[**FileTrackerPercentiles**](FileTrackerPercentiles.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchFileSizeDistributionActive**
> FileTrackerPercentiles FetchFileSizeDistributionActive(ctx, optional)
Get Active Transfer File Size Distribution.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsControllerApiFetchFileSizeDistributionActiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsControllerApiFetchFileSizeDistributionActiveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **optional.Int32**| Percentile range | [default to 10]
 **migrationId** | **optional.String**| Migration Id | 

### Return type

[**FileTrackerPercentiles**](FileTrackerPercentiles.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchFileTransferRatesDistribution**
> FileTrackerPercentiles FetchFileTransferRatesDistribution(ctx, optional)
Get File Transfer Rates Distribution.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsControllerApiFetchFileTransferRatesDistributionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsControllerApiFetchFileTransferRatesDistributionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **optional.Int32**| Percentile range | [default to 10]
 **migrationId** | **optional.String**| Migration Id | 

### Return type

[**FileTrackerPercentiles**](FileTrackerPercentiles.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchFileTransferRatesDistributionActive**
> FileTrackerPercentiles FetchFileTransferRatesDistributionActive(ctx, optional)
Get Active File Transfer Rates Distribution.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StatisticsControllerApiFetchFileTransferRatesDistributionActiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsControllerApiFetchFileTransferRatesDistributionActiveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **optional.Int32**| Percentile range | [default to 10]
 **migrationId** | **optional.String**| Migration Id | 

### Return type

[**FileTrackerPercentiles**](FileTrackerPercentiles.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchMigrationInfo**
> MigrationInfo FetchMigrationInfo(ctx, migrationId)
Get stats information for the given migration ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 

### Return type

[**MigrationInfo**](MigrationInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchMigrationStats**
> []MigrationRecord FetchMigrationStats(ctx, migrationId)
Get Stats over time for given migration ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 

### Return type

[**[]MigrationRecord**](MigrationRecord.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchMigrationSummary**
> MigrationsSummary FetchMigrationSummary(ctx, )
Get migrations summary

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MigrationsSummary**](MigrationsSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchThroughputInLastSecs**
> ThroughputBucket FetchThroughputInLastSecs(ctx, nSecs)
Get throughput stats for all migrations during last N seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nSecs** | **int32**| Number of seconds to check | 

### Return type

[**ThroughputBucket**](ThroughputBucket.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchThroughputSummary**
> ThroughputSummary FetchThroughputSummary(ctx, )
Get throughput summary

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ThroughputSummary**](ThroughputSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchTotalRollingStats**
> []AllMigrationsRollingStats FetchTotalRollingStats(ctx, )
Get rolling stats over time for all migrations

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AllMigrationsRollingStats**](AllMigrationsRollingStats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchTotalStats**
> MigrationStats FetchTotalStats(ctx, )
Get cumulative total stats for all migrations

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MigrationStats**](MigrationStats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LastExcludedEvents**
> []EnrichedMigratorEvent LastExcludedEvents(ctx, )
Get most recent 1000 excluded events.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]EnrichedMigratorEvent**](EnrichedMigratorEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

