# \ConfigApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListScmsSupportingVersionAutoDetect**](ConfigApi.md#ListScmsSupportingVersionAutoDetect) | **Get** /versions/auto_detect_supported_scms | List all SCM types supporting auto detecting the project Unity version
[**ListUnityVersions**](ConfigApi.md#ListUnityVersions) | **Get** /versions/unity | List all unity versions
[**ListXcodeVersions**](ConfigApi.md#ListXcodeVersions) | **Get** /versions/xcode | List all xcode versions



## ListScmsSupportingVersionAutoDetect

> []string ListScmsSupportingVersionAutoDetect(ctx, )
List all SCM types supporting auto detecting the project Unity version

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[apikey](../README.md#apikey), [filetoken](../README.md#filetoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUnityVersions

> []InlineResponse200 ListUnityVersions(ctx, )
List all unity versions

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]InlineResponse200**](inline_response_200.md)

### Authorization

[apikey](../README.md#apikey), [filetoken](../README.md#filetoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListXcodeVersions

> []InlineResponse2001 ListXcodeVersions(ctx, )
List all xcode versions

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]InlineResponse2001**](inline_response_200_1.md)

### Authorization

[apikey](../README.md#apikey), [filetoken](../README.md#filetoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

