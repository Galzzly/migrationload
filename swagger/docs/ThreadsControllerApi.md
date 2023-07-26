# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchThreadDump**](ThreadsControllerApi.md#FetchThreadDump) | **Get** /threads | Get Thread Dump
[**FetchThreadStates**](ThreadsControllerApi.md#FetchThreadStates) | **Get** /threads/states | Get Thread State Summary

# **FetchThreadDump**
> string FetchThreadDump(ctx, )
Get Thread Dump

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchThreadStates**
> ThreadStates FetchThreadStates(ctx, )
Get Thread State Summary

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ThreadStates**](ThreadStates.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

