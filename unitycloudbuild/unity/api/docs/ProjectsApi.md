# \ProjectsApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProject**](ProjectsApi.md#AddProject) | **Post** /orgs/{orgid}/projects | Create project
[**ArchiveProject**](ProjectsApi.md#ArchiveProject) | **Delete** /orgs/{orgid}/projects/{projectid} | Archive project
[**GetAuditLog**](ProjectsApi.md#GetAuditLog) | **Get** /orgs/{orgid}/projects/{projectid}/auditlog | Get audit log
[**GetProject**](ProjectsApi.md#GetProject) | **Get** /orgs/{orgid}/projects/{projectid} | Get project details
[**GetProjectByUpid**](ProjectsApi.md#GetProjectByUpid) | **Get** /projects/{projectupid} | Get project details
[**GetProjectsBillingPlans**](ProjectsApi.md#GetProjectsBillingPlans) | **Get** /orgs/{orgid}/projects/{projectid}/billingplan | Get billing plan
[**GetProjectsSSHKey**](ProjectsApi.md#GetProjectsSSHKey) | **Get** /orgs/{orgid}/projects/{projectid}/sshkey | Get SSH Key
[**GetStats**](ProjectsApi.md#GetStats) | **Get** /orgs/{orgid}/projects/{projectid}/stats | Get project statistics
[**ListProjectsForOrg**](ProjectsApi.md#ListProjectsForOrg) | **Get** /orgs/{orgid}/projects | List all projects (org)
[**ListProjectsForUser**](ProjectsApi.md#ListProjectsForUser) | **Get** /projects | List all projects (user)
[**OrgsOrgidProjectsProjectidEnvvarsGet**](ProjectsApi.md#OrgsOrgidProjectsProjectidEnvvarsGet) | **Get** /orgs/{orgid}/projects/{projectid}/envvars | Get environment variables
[**OrgsOrgidProjectsProjectidEnvvarsPut**](ProjectsApi.md#OrgsOrgidProjectsProjectidEnvvarsPut) | **Put** /orgs/{orgid}/projects/{projectid}/envvars | Set environment variables
[**UpdateProject**](ProjectsApi.md#UpdateProject) | **Put** /orgs/{orgid}/projects/{projectid} | Update project details



## AddProject

> map[string]interface{} AddProject(ctx, orgid, options)
Create project

Create a project for the specified organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**options** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| Options for project create/update | 

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


## ArchiveProject

> string ArchiveProject(ctx, orgid, projectid)
Archive project

This will archive the project in Cloud Build ONLY. Use with caution - this process is not reversible. The projects UPID will be removed from Cloud Build allowing the project to be reconfigured.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 

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


## GetAuditLog

> []InlineResponse2005 GetAuditLog(ctx, orgid, projectid, optional)
Get audit log

Retrieve a list of historical settings changes for this project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
 **optional** | ***GetAuditLogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAuditLogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Float32**| Number of audit log records to retrieve | [default to 25.0]
 **page** | **optional.Float32**| Skip to page number, based on per_page value | [default to 1.0]

### Return type

[**[]InlineResponse2005**](inline_response_200_5.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProject

> map[string]interface{} GetProject(ctx, orgid, projectid, optional)
Get project details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
 **optional** | ***GetProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProjectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | **optional.String**| Extra fields to include in the response | 

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


## GetProjectByUpid

> map[string]interface{} GetProjectByUpid(ctx, projectupid)
Get project details

Gets the same data as /orgs/{orgid}/project/{projectid} but looks up the project by the Unity Project ID. This value is returned in the project's guid field.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectupid** | **string**| Project UPID - Unity global id | 

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


## GetProjectsBillingPlans

> map[string]interface{} GetProjectsBillingPlans(ctx, orgid, projectid)
Get billing plan

Get the billing plan for the specified project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 

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


## GetProjectsSSHKey

> map[string]interface{} GetProjectsSSHKey(ctx, orgid, projectid)
Get SSH Key

Get the ssh public key for the specified project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 

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


## GetStats

> map[string]interface{} GetStats(ctx, orgid, projectid)
Get project statistics

Get statistics for the specified project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 

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


## ListProjectsForOrg

> []map[string]interface{} ListProjectsForOrg(ctx, orgid, optional)
List all projects (org)

List all projects that belong to the specified organization. Add \"?include=settings\" as a query parameter to include the project settings with the response. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
 **optional** | ***ListProjectsForOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListProjectsForOrgOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **optional.String**| Extra fields to include in the response | 

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv, application/json-default, application/json-include-settings

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjectsForUser

> []map[string]interface{} ListProjectsForUser(ctx, optional)
List all projects (user)

List all projects that you have permission to access across all organizations. Add \"?include=settings\" as a query parameter to include the project settings with the response. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListProjectsForUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListProjectsForUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **optional.String**| Extra fields to include in the response | 

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [filetoken](../README.md#filetoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsOrgidProjectsProjectidEnvvarsGet

> map[string]string OrgsOrgidProjectsProjectidEnvvarsGet(ctx, orgid, projectid)
Get environment variables

Get all configured environment variables for a given project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 

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


## OrgsOrgidProjectsProjectidEnvvarsPut

> map[string]string OrgsOrgidProjectsProjectidEnvvarsPut(ctx, orgid, projectid, envvars)
Set environment variables

Set all configured environment variables for a given project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
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


## UpdateProject

> map[string]interface{} UpdateProject(ctx, orgid, projectid, options)
Update project details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**options** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| Options for project create/update | 

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

