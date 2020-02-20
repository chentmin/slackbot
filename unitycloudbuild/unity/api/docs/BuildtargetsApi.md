# \BuildtargetsApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBuildTargetsForOrg**](BuildtargetsApi.md#GetBuildTargetsForOrg) | **Get** /orgs/{orgid}/buildtargets | List all build targets for an org
[**GetProjectsBuildTargetsStats**](BuildtargetsApi.md#GetProjectsBuildTargetsStats) | **Get** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/stats | Get build target statistics
[**OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidDelete**](BuildtargetsApi.md#OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidDelete) | **Delete** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid} | Delete build target
[**OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidEnvvarsGet**](BuildtargetsApi.md#OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidEnvvarsGet) | **Get** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/envvars | Get environment variables
[**OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidEnvvarsPut**](BuildtargetsApi.md#OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidEnvvarsPut) | **Put** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/envvars | Set environment variables
[**OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidGet**](BuildtargetsApi.md#OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidGet) | **Get** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid} | Get a build target
[**OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidPut**](BuildtargetsApi.md#OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidPut) | **Put** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid} | Update build target details
[**OrgsOrgidProjectsProjectidBuildtargetsGet**](BuildtargetsApi.md#OrgsOrgidProjectsProjectidBuildtargetsGet) | **Get** /orgs/{orgid}/projects/{projectid}/buildtargets | List all build targets for a project
[**OrgsOrgidProjectsProjectidBuildtargetsPost**](BuildtargetsApi.md#OrgsOrgidProjectsProjectidBuildtargetsPost) | **Post** /orgs/{orgid}/projects/{projectid}/buildtargets | Create build target for a project



## GetBuildTargetsForOrg

> []InlineResponse2006 GetBuildTargetsForOrg(ctx, orgid, optional)
List all build targets for an org

Gets all configured build targets for an org, regardless of whether they are enabled. Add \"?include=settings,credentials\" as a query parameter to include the build target settings and credentials with the response. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
 **optional** | ***GetBuildTargetsForOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBuildTargetsForOrgOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **optional.String**| Extra fields to include in the response | 
 **includeLastSuccess** | **optional.Bool**| Include last successful build | [default to false]

### Return type

[**[]InlineResponse2006**](inline_response_200_6.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsBuildTargetsStats

> map[string]interface{} GetProjectsBuildTargetsStats(ctx, orgid, projectid, buildtargetid, optional)
Get build target statistics

Get statistics for the specified build target

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 
 **optional** | ***GetProjectsBuildTargetsStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProjectsBuildTargetsStatsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **buildStatus** | **optional.String**| Query for only builds of a specific status | 
 **cleanBuild** | **optional.Bool**| Query for builds that have either been built clean or using caches | 
 **limit** | **optional.Float32**| Get stats for last limit builds | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidDelete

> string OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidDelete(ctx, orgid, projectid, buildtargetid)
Delete build target

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 

### Return type

**string**

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidEnvvarsGet

> map[string]string OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidEnvvarsGet(ctx, orgid, projectid, buildtargetid)
Get environment variables

Get all configured environment variables for a given build target

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 

### Return type

**map[string]string**

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidEnvvarsPut

> map[string]string OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidEnvvarsPut(ctx, orgid, projectid, buildtargetid, envvars)
Set environment variables

Set all configured environment variables for a given build target

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 
**envvars** | [**map[string]string**](string.md)| Environment variables | 

### Return type

**map[string]string**

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidGet

> map[string]interface{} OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidGet(ctx, orgid, projectid, buildtargetid)
Get a build target

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidPut

> map[string]interface{} OrgsOrgidProjectsProjectidBuildtargetsBuildtargetidPut(ctx, orgid, projectid, buildtargetid, options)
Update build target details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 
**options** | [**InlineObject7**](InlineObject7.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsOrgidProjectsProjectidBuildtargetsGet

> []InlineResponse2006 OrgsOrgidProjectsProjectidBuildtargetsGet(ctx, orgid, projectid, optional)
List all build targets for a project

Gets all configured build targets for a project, regardless of whether they are enabled. Add \"?include=settings,credentials\" as a query parameter to include the build target settings and credentials with the response. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
 **optional** | ***OrgsOrgidProjectsProjectidBuildtargetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrgsOrgidProjectsProjectidBuildtargetsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | **optional.String**| Extra fields to include in the response | 
 **includeLastSuccess** | **optional.Bool**| Include last successful build | [default to false]

### Return type

[**[]InlineResponse2006**](inline_response_200_6.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsOrgidProjectsProjectidBuildtargetsPost

> map[string]interface{} OrgsOrgidProjectsProjectidBuildtargetsPost(ctx, orgid, projectid, options)
Create build target for a project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**options** | [**InlineObject6**](InlineObject6.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

