# \SharesApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetShareMetadata**](SharesApi.md#GetShareMetadata) | **Get** /shares/{shareid} | Get details on shared build including download link



## GetShareMetadata

> map[string]interface{} GetShareMetadata(ctx, shareid, optional)
Get details on shared build including download link

This is an endpoint accessible without an api key that provides information about a specific build including download links. A shareid is generated by POSTing to a <a href=\"#!/builds/createShare\">build's share endpoint</a>.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareid** | **string**|  | 
 **optional** | ***GetShareMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetShareMetadataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **optional.String**| Extra fields to include in the response | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [filetoken](../README.md#filetoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
