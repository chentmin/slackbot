# \BuildsApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteBuildArtifacts**](BuildsApi.md#BatchDeleteBuildArtifacts) | **Post** /orgs/{orgid}/projects/{projectid}/artifacts/delete | Delete artifacts for a batch of builds
[**CancelAllBuilds**](BuildsApi.md#CancelAllBuilds) | **Delete** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/builds | Cancel all builds
[**CancelBuild**](BuildsApi.md#CancelBuild) | **Delete** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/builds/{number} | Cancel build
[**CancelBuildsForOrg**](BuildsApi.md#CancelBuildsForOrg) | **Delete** /orgs/{orgid}/builds | Cancel builds for org
[**CreateShare**](BuildsApi.md#CreateShare) | **Post** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/builds/{number}/share | Create a new link to share a project
[**DeleteAllBuildArtifacts**](BuildsApi.md#DeleteAllBuildArtifacts) | **Delete** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/builds/artifacts | Delete all artifacts associated with all non-favorited builds for a specified buildtargetid (_all is allowed).
[**DeleteBuildArtifacts**](BuildsApi.md#DeleteBuildArtifacts) | **Delete** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/builds/{number}/artifacts | Delete all artifacts associated with a specific build
[**GetBuild**](BuildsApi.md#GetBuild) | **Get** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/builds/{number} | Build Status
[**GetBuildLog**](BuildsApi.md#GetBuildLog) | **Get** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/builds/{number}/log | Get build log
[**GetBuildSteps**](BuildsApi.md#GetBuildSteps) | **Get** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/builds/{number}/steps | Get the build steps for a given build
[**GetBuilds**](BuildsApi.md#GetBuilds) | **Get** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/builds | List all builds
[**GetBuildsForOrg**](BuildsApi.md#GetBuildsForOrg) | **Get** /orgs/{orgid}/builds | List all builds for org
[**GetProjectsBuildTargetsAuditLog**](BuildsApi.md#GetProjectsBuildTargetsAuditLog) | **Get** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/auditlog | Get audit log
[**GetProjectsBuildTargetsBuildsAuditLog**](BuildsApi.md#GetProjectsBuildTargetsBuildsAuditLog) | **Get** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/builds/{number}/auditlog | Get audit log
[**GetShare**](BuildsApi.md#GetShare) | **Get** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/builds/{number}/share | Get the share link
[**ResignBuildArtifact**](BuildsApi.md#ResignBuildArtifact) | **Post** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/builds/{number}/resign | Re-sign a build artifact
[**RevokeShare**](BuildsApi.md#RevokeShare) | **Delete** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/builds/{number}/share | Revoke a shared link
[**StartBuilds**](BuildsApi.md#StartBuilds) | **Post** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/builds | Create new build
[**UpdateBuild**](BuildsApi.md#UpdateBuild) | **Put** /orgs/{orgid}/projects/{projectid}/buildtargets/{buildtargetid}/builds/{number} | Update build information



## BatchDeleteBuildArtifacts

> string BatchDeleteBuildArtifacts(ctx, orgid, projectid, options)
Delete artifacts for a batch of builds

Delete all artifacts associated with the builds identified by the provided build target ids and build numbers. Builds marked as do not delete or that are currently building will be ignored. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**options** | [**InlineObject8**](InlineObject8.md)|  | 

### Return type

**string**

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelAllBuilds

> string CancelAllBuilds(ctx, orgid, projectid, buildtargetid)
Cancel all builds

Cancel all builds in progress for this build target (or all targets, if '_all' is specified as the buildtargetid). Canceling an already finished build will do nothing and respond successfully. 

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


## CancelBuild

> string CancelBuild(ctx, orgid, projectid, buildtargetid, number)
Cancel build

Cancel a build in progress. Canceling an already finished build will do nothing and respond successfully. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 
**number** | **string**| Build number or in some cases _all | 

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


## CancelBuildsForOrg

> string CancelBuildsForOrg(ctx, orgid)
Cancel builds for org

Cancel all in progress builds for an organization. Canceling an already finished build will do nothing and respond successfully. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 

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


## CreateShare

> map[string]interface{} CreateShare(ctx, orgid, projectid, buildtargetid, number)
Create a new link to share a project

Create a new short link to share a project. If this is called when a share already exists, that share will be revoked and a new one created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 
**number** | **string**| Build number or in some cases _all | 

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


## DeleteAllBuildArtifacts

> string DeleteAllBuildArtifacts(ctx, orgid, projectid, buildtargetid)
Delete all artifacts associated with all non-favorited builds for a specified buildtargetid (_all is allowed).

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


## DeleteBuildArtifacts

> string DeleteBuildArtifacts(ctx, orgid, projectid, buildtargetid, number)
Delete all artifacts associated with a specific build

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 
**number** | **string**| Build number or in some cases _all | 

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


## GetBuild

> map[string]interface{} GetBuild(ctx, orgid, projectid, buildtargetid, number, optional)
Build Status

Retrieve information for a specific build. A Build resource contains information related to a build attempt for a build target, including the build number, changeset, build times, and other pertinent data. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 
**number** | **string**| Build number or in some cases _all | 
 **optional** | ***GetBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBuildOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **include** | **optional.String**| Extra fields to include in the response | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [filetoken](../README.md#filetoken), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv, application/json-accepted

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuildLog

> GetBuildLog(ctx, orgid, projectid, buildtargetid, number, optional)
Get build log

Retrieve the plain text log for a specifc build.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 
**number** | **string**| Build number or in some cases _all | 
 **optional** | ***GetBuildLogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBuildLogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **offsetlines** | **optional.Float32**| Stream log from the given line number | [default to 1.0]
 **linenumbers** | **optional.Bool**| Include log line numbers in the text output | [default to false]
 **lastLineNumber** | **optional.Float32**| The last line number seen, numbering will continue from here | [default to 0.0]
 **compact** | **optional.Bool**| Return the compact log, showing only errors and warnings | [default to false]
 **withHtml** | **optional.Bool**| Surround important lines (errors, warnings) with SPAN and CSS markup  | [default to false]

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain-full, text/plain-compact, text/html-compact, application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuildSteps

> []InlineResponse2009 GetBuildSteps(ctx, orgid, projectid, buildtargetid, number)
Get the build steps for a given build

Retrieves all build steps for a build, this replaces the old method where we would manually download the build report artifacts and allows us to add more functionality into build steps. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 
**number** | **string**| Build number or in some cases _all | 

### Return type

[**[]InlineResponse2009**](inline_response_200_9.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuilds

> []OrgsOrgidProjectsProjectidBuildtargetsBuilds GetBuilds(ctx, orgid, projectid, buildtargetid, optional)
List all builds

List all running and finished builds, sorted by build number (optionally paginating the results). Use '_all' as the buildtargetid to get all configured build targets. The response includes a Content-Range header that identifies the range of results returned and the total number of results matching the given query parameters. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 
 **optional** | ***GetBuildsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBuildsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **include** | **optional.String**| Extra fields to include in the response | 
 **perPage** | **optional.Float32**| Number of audit log records to retrieve | [default to 25.0]
 **page** | **optional.Float32**| Skip to page number, based on per_page value | [default to 1.0]
 **buildStatus** | **optional.String**| Query for only builds of a specific status | 
 **platform** | **optional.String**| Query for only builds of a specific platform | 
 **showDeleted** | **optional.Bool**| Query for builds that have been deleted | [default to false]
 **onlyFavorites** | **optional.Bool**| Query for builds that have been favorited | [default to false]
 **cleanBuild** | **optional.Bool**| Query for builds that have either been built clean or using caches | 

### Return type

[**[]OrgsOrgidProjectsProjectidBuildtargetsBuilds**](_orgs_orgid_projects_projectid_buildtargets_builds.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuildsForOrg

> []OrgsOrgidProjectsProjectidBuildtargetsBuilds GetBuildsForOrg(ctx, orgid, optional)
List all builds for org

List all running and finished builds, sorted by build number (optionally paginating the results). The response includes a Content-Range header that identifies the range of results returned and the total number of results matching the given query parameters. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
 **optional** | ***GetBuildsForOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBuildsForOrgOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **optional.String**| Extra fields to include in the response | 
 **perPage** | **optional.Float32**| Number of audit log records to retrieve | [default to 25.0]
 **page** | **optional.Float32**| Skip to page number, based on per_page value | [default to 1.0]
 **buildStatus** | **optional.String**| Query for only builds of a specific status | 
 **platform** | **optional.String**| Query for only builds of a specific platform | 
 **showDeleted** | **optional.Bool**| Query for builds that have been deleted | [default to false]
 **onlyFavorites** | **optional.Bool**| Query for builds that have been favorited | [default to false]
 **cleanBuild** | **optional.Bool**| Query for builds that have either been built clean or using caches | 

### Return type

[**[]OrgsOrgidProjectsProjectidBuildtargetsBuilds**](_orgs_orgid_projects_projectid_buildtargets_builds.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsBuildTargetsAuditLog

> []InlineResponse2005 GetProjectsBuildTargetsAuditLog(ctx, orgid, projectid, buildtargetid, optional)
Get audit log

Retrieve a list of historical settings changes for this build target.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 
 **optional** | ***GetProjectsBuildTargetsAuditLogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProjectsBuildTargetsAuditLogOpts struct


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


## GetProjectsBuildTargetsBuildsAuditLog

> []InlineResponse2005 GetProjectsBuildTargetsBuildsAuditLog(ctx, orgid, projectid, buildtargetid, number, optional)
Get audit log

Retrieve a list of settings changes between the last and current build.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 
**number** | **string**| Build number or in some cases _all | 
 **optional** | ***GetProjectsBuildTargetsBuildsAuditLogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProjectsBuildTargetsBuildsAuditLogOpts struct


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


## GetShare

> map[string]interface{} GetShare(ctx, orgid, projectid, buildtargetid, number)
Get the share link

Gets a share link if it exists

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 
**number** | **string**| Build number or in some cases _all | 

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


## ResignBuildArtifact

> []OrgsOrgidProjectsProjectidBuildtargetsBuilds ResignBuildArtifact(ctx, orgid, projectid, buildtargetid, number)
Re-sign a build artifact

Re-sign a build artifact using the most recent credentials associated with the buildtarget. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 
**number** | **string**| Build number or in some cases _all | 

### Return type

[**[]OrgsOrgidProjectsProjectidBuildtargetsBuilds**](_orgs_orgid_projects_projectid_buildtargets_builds.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv, application/json-accepted, application/json-already-building

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeShare

> string RevokeShare(ctx, orgid, projectid, buildtargetid, number)
Revoke a shared link

Revoke a shared link, both {buildtargetid} and {number} may use _all to revoke all share links for a given buildtarget or entire project.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 
**number** | **string**| Build number or in some cases _all | 

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


## StartBuilds

> []OrgsOrgidProjectsProjectidBuildtargetsBuilds StartBuilds(ctx, orgid, projectid, buildtargetid, optional)
Create new build

Start the build process for this build target (or all targets, if '_all' is specified as the buildtargetid), if there is not one currently in process.  If a build is currently in process that information will be related in the 'error' field. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 
 **optional** | ***StartBuildsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StartBuildsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **options** | [**optional.Interface of InlineObject9**](InlineObject9.md)|  | 

### Return type

[**[]OrgsOrgidProjectsProjectidBuildtargetsBuilds**](_orgs_orgid_projects_projectid_buildtargets_builds.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html, text/csv, application/json-accepted, application/json-already-building

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBuild

> map[string]interface{} UpdateBuild(ctx, orgid, projectid, buildtargetid, number, options)
Update build information

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**buildtargetid** | **string**| unique id auto-generated from the build target name | 
**number** | **string**| Build number or in some cases _all | 
**options** | [**InlineObject10**](InlineObject10.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html, text/csv, application/json-accepted

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

