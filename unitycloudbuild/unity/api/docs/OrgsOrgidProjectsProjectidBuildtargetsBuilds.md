# OrgsOrgidProjectsProjectidBuildtargetsBuilds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Build** | **float32** |  | [optional] 
**Buildtargetid** | **string** | unique id auto-generated from the build target name | [optional] 
**BuildTargetName** | **string** |  | [optional] 
**BuildGUID** | **string** | unique GUID identifying this build | [optional] 
**BuildStatus** | **string** |  | [optional] 
**CleanBuild** | **bool** | if the build was built without using data cached from previous builds | [optional] 
**FailureDetails** | [**[]OrgsOrgidProjectsProjectidBuildtargetsFailureDetails**](_orgs_orgid_projects_projectid_buildtargets_failureDetails.md) | list of failure details for this build attempt, when available | [optional] 
**CanceledBy** | **string** |  | [optional] 
**Platform** | **string** |  | [optional] 
**WorkspaceSize** | **float32** | size of workspace in bytes | [optional] 
**Created** | **string** | when the build was created | [optional] 
**Finished** | **string** | when the build completely finished | [optional] 
**CheckoutStartTime** | **string** | when the build starting checking out code | [optional] 
**CheckoutTimeInSeconds** | **float32** | amount of time spent checking out code | [optional] 
**BuildStartTime** | **string** | when the build started compiling | [optional] 
**BuildTimeInSeconds** | **float32** | amount of time spend compiling | [optional] 
**PublishStartTime** | **string** | when the build started saving build artifacts | [optional] 
**PublishTimeInSeconds** | **float32** | amount of time spent saving build artifacts | [optional] 
**TotalTimeInSeconds** | **float32** | total time for the build | [optional] 
**LastBuiltRevision** | **string** | source control commit id for the build | [optional] 
**Changeset** | [**[]map[string]interface{}**](map[string]interface{}.md) | a list of source control changes between this and the last build | [optional] 
**Favorited** | **bool** | if the build is marked as do not delete or not | [optional] 
**Label** | **string** | description given when a build is favorited | [optional] 
**Deleted** | **bool** | if the build is deleted or not | [optional] 
**Headless** | [**map[string]interface{}**](.md) | if the build was built to run in linux headless mode | [optional] 
**CredentialsOutdated** | **bool** | if a newer credential has been attached to this buildtarget and the build can be re-signed | [optional] 
**DeletedBy** | **string** | email address of the user who deleted this attempt | [optional] 
**QueuedReason** | **string** | reason the build is currently waiting | [optional] 
**CooldownDate** | **string** | time until this build will be reconsidered for building | [optional] 
**ScmBranch** | **string** | scm branch to be built | [optional] 
**UnityVersion** | **string** | &#39;latest&#39; or a unity dot version with underscores (ex. &#39;4_6_5&#39;) | [optional] 
**XcodeVersion** | **string** | &#39;latest&#39; or a supported xcode version (ex. &#39;xcode7&#39;) | [optional] 
**AuditChanges** | **float32** |  | [optional] 
**ProjectVersion** | [**OrgsOrgidProjectsProjectidBuildtargetsProjectVersion**](_orgs_orgid_projects_projectid_buildtargets_projectVersion.md) |  | [optional] 
**ProjectName** | **string** |  | [optional] 
**ProjectId** | **string** |  | [optional] 
**OrgId** | **string** |  | [optional] 
**OrgFk** | **string** |  | [optional] 
**Filetoken** | **string** |  | [optional] 
**Links** | [**map[string]OrgsOrgidProjectsProjectidBuildtargetsLinks**](_orgs_orgid_projects_projectid_buildtargets_links.md) |  | [optional] 
**BuildReport** | [**OrgsOrgidProjectsProjectidBuildtargetsBuildReport**](_orgs_orgid_projects_projectid_buildtargets_buildReport.md) |  | [optional] 
**TestResults** | [**OrgsOrgidProjectsProjectidBuildtargetsTestResults**](_orgs_orgid_projects_projectid_buildtargets_testResults.md) |  | [optional] 
**Error** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


