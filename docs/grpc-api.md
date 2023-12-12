# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/resources/ceph/v1/ceph.proto](#api_resources_ceph_v1_ceph-proto)
    - [User](#api-resources-ceph-v1-User)
  
- [api/resources/ceph/v1/resources.proto](#api_resources_ceph_v1_resources-proto)
    - [Resource](#api-resources-ceph-v1-Resource)
    - [Resources](#api-resources-ceph-v1-Resources)
  
- [api/resources/koor/v1/koor.proto](#api_resources_koor_v1_koor-proto)
    - [ClusterResources](#api-resources-koor-v1-ClusterResources)
    - [DetailedProductVersions](#api-resources-koor-v1-DetailedProductVersions)
    - [DetailedVersion](#api-resources-koor-v1-DetailedVersion)
    - [KoorCluster](#api-resources-koor-v1-KoorCluster)
    - [KoorClusterSpec](#api-resources-koor-v1-KoorClusterSpec)
    - [KoorClusterStatus](#api-resources-koor-v1-KoorClusterStatus)
    - [OSDScrubbingSchedule](#api-resources-koor-v1-OSDScrubbingSchedule)
    - [ProductVersions](#api-resources-koor-v1-ProductVersions)
    - [UpgradeOptions](#api-resources-koor-v1-UpgradeOptions)
  
    - [UpgradeMode](#api-resources-koor-v1-UpgradeMode)
  
- [api/resources/stats/v1/stats.proto](#api_resources_stats_v1_stats-proto)
    - [ClusterRadar](#api-resources-stats-v1-ClusterRadar)
    - [ClusterStats](#api-resources-stats-v1-ClusterStats)
    - [Crash](#api-resources-stats-v1-Crash)
    - [Data](#api-resources-stats-v1-Data)
    - [IOPS](#api-resources-stats-v1-IOPS)
    - [MdsService](#api-resources-stats-v1-MdsService)
    - [MgrService](#api-resources-stats-v1-MgrService)
    - [MonService](#api-resources-stats-v1-MonService)
    - [NodeInfo](#api-resources-stats-v1-NodeInfo)
    - [Objects](#api-resources-stats-v1-Objects)
    - [OsdService](#api-resources-stats-v1-OsdService)
    - [PGs](#api-resources-stats-v1-PGs)
    - [Pools](#api-resources-stats-v1-Pools)
    - [ResourceInfo](#api-resources-stats-v1-ResourceInfo)
    - [RgwService](#api-resources-stats-v1-RgwService)
    - [Services](#api-resources-stats-v1-Services)
    - [Usage](#api-resources-stats-v1-Usage)
  
    - [ClusterHealth](#api-resources-stats-v1-ClusterHealth)
    - [ReliabilityScore](#api-resources-stats-v1-ReliabilityScore)
    - [ResourceStatus](#api-resources-stats-v1-ResourceStatus)
  
- [api/resources/timestamp/v1/timestamp.proto](#api_resources_timestamp_v1_timestamp-proto)
    - [Timestamp](#api-resources-timestamp-v1-Timestamp)
  
- [api/services/auth/v1/auth.proto](#api_services_auth_v1_auth-proto)
    - [CheckTokenRequest](#api-services-auth-v1-CheckTokenRequest)
    - [CheckTokenResponse](#api-services-auth-v1-CheckTokenResponse)
    - [LoginRequest](#api-services-auth-v1-LoginRequest)
    - [LoginResponse](#api-services-auth-v1-LoginResponse)
    - [LogoutRequest](#api-services-auth-v1-LogoutRequest)
    - [LogoutResponse](#api-services-auth-v1-LogoutResponse)
  
    - [AuthService](#api-services-auth-v1-AuthService)
  
- [api/services/ceph/v1/ceph.proto](#api_services_ceph_v1_ceph-proto)
    - [CreateCephUserRequest](#api-services-ceph-v1-CreateCephUserRequest)
    - [CreateCephUserResponse](#api-services-ceph-v1-CreateCephUserResponse)
    - [DeleteCephUserRequest](#api-services-ceph-v1-DeleteCephUserRequest)
    - [DeleteCephUserResponse](#api-services-ceph-v1-DeleteCephUserResponse)
    - [GetCephUsersRequest](#api-services-ceph-v1-GetCephUsersRequest)
    - [GetCephUsersResponse](#api-services-ceph-v1-GetCephUsersResponse)
  
    - [CephService](#api-services-ceph-v1-CephService)
  
- [api/services/cluster/v1/cluster.proto](#api_services_cluster_v1_cluster-proto)
    - [CancelNetworkTestRequest](#api-services-cluster-v1-CancelNetworkTestRequest)
    - [CancelNetworkTestResponse](#api-services-cluster-v1-CancelNetworkTestResponse)
    - [GetKoorClusterRequest](#api-services-cluster-v1-GetKoorClusterRequest)
    - [GetKoorClusterResponse](#api-services-cluster-v1-GetKoorClusterResponse)
    - [GetNetworkTestResultsRequest](#api-services-cluster-v1-GetNetworkTestResultsRequest)
    - [GetNetworkTestResultsResponse](#api-services-cluster-v1-GetNetworkTestResultsResponse)
    - [GetNetworkTestStatusRequest](#api-services-cluster-v1-GetNetworkTestStatusRequest)
    - [GetNetworkTestStatusResponse](#api-services-cluster-v1-GetNetworkTestStatusResponse)
    - [GetResourcesRequest](#api-services-cluster-v1-GetResourcesRequest)
    - [GetResourcesResponse](#api-services-cluster-v1-GetResourcesResponse)
    - [GetTroubleshootReportRequest](#api-services-cluster-v1-GetTroubleshootReportRequest)
    - [GetTroubleshootReportResponse](#api-services-cluster-v1-GetTroubleshootReportResponse)
    - [SaveResourcesRequest](#api-services-cluster-v1-SaveResourcesRequest)
    - [SaveResourcesResponse](#api-services-cluster-v1-SaveResourcesResponse)
    - [SetScrubbingScheduleRequest](#api-services-cluster-v1-SetScrubbingScheduleRequest)
    - [SetScrubbingScheduleResponse](#api-services-cluster-v1-SetScrubbingScheduleResponse)
    - [StartNetworkTestRequest](#api-services-cluster-v1-StartNetworkTestRequest)
    - [StartNetworkTestResponse](#api-services-cluster-v1-StartNetworkTestResponse)
  
    - [NetworkTestOutputFormat](#api-services-cluster-v1-NetworkTestOutputFormat)
  
    - [ClusterService](#api-services-cluster-v1-ClusterService)
  
- [api/services/stats/v1/stats.proto](#api_services_stats_v1_stats-proto)
    - [GetClusterNodesRequest](#api-services-stats-v1-GetClusterNodesRequest)
    - [GetClusterNodesResponse](#api-services-stats-v1-GetClusterNodesResponse)
    - [GetClusterRadarRequest](#api-services-stats-v1-GetClusterRadarRequest)
    - [GetClusterRadarResponse](#api-services-stats-v1-GetClusterRadarResponse)
    - [GetClusterResourcesRequest](#api-services-stats-v1-GetClusterResourcesRequest)
    - [GetClusterResourcesResponse](#api-services-stats-v1-GetClusterResourcesResponse)
    - [GetClusterStatsRequest](#api-services-stats-v1-GetClusterStatsRequest)
    - [GetClusterStatsResponse](#api-services-stats-v1-GetClusterStatsResponse)
  
    - [StatsService](#api-services-stats-v1-StatsService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_resources_ceph_v1_ceph-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/resources/ceph/v1/ceph.proto



<a name="api-resources-ceph-v1-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Username | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Email | [string](#string) |  |  |
| Password | [string](#string) |  |  |
| Enabled | [bool](#bool) |  |  |
| Roles | [string](#string) | repeated |  |





 

 

 

 



<a name="api_resources_ceph_v1_resources-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/resources/ceph/v1/resources.proto



<a name="api-resources-ceph-v1-Resource"></a>

### Resource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| content | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| kind | [string](#string) |  |  |
| object | [string](#string) |  |  |






<a name="api-resources-ceph-v1-Resources"></a>

### Resources



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resources | [Resource](#api-resources-ceph-v1-Resource) | repeated |  |





 

 

 

 



<a name="api_resources_koor_v1_koor-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/resources/koor/v1/koor.proto



<a name="api-resources-koor-v1-ClusterResources"></a>

### ClusterResources



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodes | [string](#string) |  | The number of nodes in the cluster |
| storage | [string](#string) |  | Ephemeral Storage available |
| cpu | [string](#string) |  | CPU cores available |
| memory | [string](#string) |  | Memory available |






<a name="api-resources-koor-v1-DetailedProductVersions"></a>

### DetailedProductVersions
Represents a map of products to detailed versions, which include images or helm charts.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| koor_operator | [DetailedVersion](#api-resources-koor-v1-DetailedVersion) |  | The detailed Koor Operator version. |
| ksd | [DetailedVersion](#api-resources-koor-v1-DetailedVersion) |  | The detailed Koor Storage Distribution version. |
| ceph | [DetailedVersion](#api-resources-koor-v1-DetailedVersion) |  | The detailed Ceph version. |






<a name="api-resources-koor-v1-DetailedVersion"></a>

### DetailedVersion
Defines a detailed version of a product, which includes a container image or a helm chart.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  | The product version, must be a valid semver. |
| image_uri | [string](#string) |  | The URI of the container image. |
| image_hash | [string](#string) |  | The hash of the container image. |
| helm_repository | [string](#string) |  | The URI of the helm repository. |
| helm_chart | [string](#string) |  | The name of the helm chart in the repository. |






<a name="api-resources-koor-v1-KoorCluster"></a>

### KoorCluster



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| spec | [KoorClusterSpec](#api-resources-koor-v1-KoorClusterSpec) |  |  |
| status | [KoorClusterStatus](#api-resources-koor-v1-KoorClusterStatus) |  |  |






<a name="api-resources-koor-v1-KoorClusterSpec"></a>

### KoorClusterSpec
Represents the state of KoorCluster


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| use_all_devices | [google.protobuf.BoolValue](#google-protobuf-BoolValue) |  | Use all devices on nodes |
| monitoring_enabled | [google.protobuf.BoolValue](#google-protobuf-BoolValue) |  | Enable monitoring. Requires Prometheus to be pre-installed. |
| dashboard_enabled | [google.protobuf.BoolValue](#google-protobuf-BoolValue) |  | Enable the ceph dashboard for viewing cluster status |
| toolbox_enabled | [google.protobuf.BoolValue](#google-protobuf-BoolValue) |  | Installs a debugging toolbox deployment |
| upgrade_options | [UpgradeOptions](#api-resources-koor-v1-UpgradeOptions) |  | Specifies the upgrade options for new ceph versions |
| ksd_release_name | [string](#string) |  | The name to use for KSD helm release. |
| ksd_cluster_release_name | [string](#string) |  | The name to use for KSD cluster helm release. |
| osd_scrubbing_schedule | [OSDScrubbingSchedule](#api-resources-koor-v1-OSDScrubbingSchedule) |  | OSD scrubbing schedule config |






<a name="api-resources-koor-v1-KoorClusterStatus"></a>

### KoorClusterStatus
Represents the status of the KoorCluster CRD


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total_resources | [ClusterResources](#api-resources-koor-v1-ClusterResources) |  | The total resources available in the cluster nodes |
| meets_minimum_resources | [bool](#bool) |  | Does the cluster meet the minimum recommended resources |
| current_versions | [ProductVersions](#api-resources-koor-v1-ProductVersions) |  | The current versions of rook and ceph |
| latest_versions | [DetailedProductVersions](#api-resources-koor-v1-DetailedProductVersions) |  | The latest versions of rook and ceph |






<a name="api-resources-koor-v1-OSDScrubbingSchedule"></a>

### OSDScrubbingSchedule
Osd Scrubbing schedule config


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apply_schedule | [bool](#bool) |  |  |
| max_scrub_ops | [int64](#int64) | optional |  |
| begin_hour | [int64](#int64) | optional |  |
| end_hour | [int64](#int64) | optional |  |
| begin_week_day | [int64](#int64) | optional |  |
| end_week_day | [int64](#int64) | optional |  |
| min_scrub_interval | [string](#string) | optional |  |
| max_scrub_interval | [string](#string) | optional |  |
| deep_scrub_interval | [string](#string) | optional |  |
| scrub_sleep_seconds | [string](#string) | optional |  |






<a name="api-resources-koor-v1-ProductVersions"></a>

### ProductVersions
Represents a map of products to version strings.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kube | [string](#string) |  | Kubernetes version, must be a valid semver. |
| koor_operator | [string](#string) |  | Koor Operator version, must be a valid semver. |
| ksd | [string](#string) |  | Koor Storage Distribution version, must be a valid semver. |
| ceph | [string](#string) |  | Ceph version, must be a valid semver. |






<a name="api-resources-koor-v1-UpgradeOptions"></a>

### UpgradeOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mode | [UpgradeMode](#api-resources-koor-v1-UpgradeMode) |  | Upgrade mode |
| endpoint | [string](#string) |  | The api endpoint used to find the ceph latest version |
| schedule | [string](#string) |  | The schedule to check for new versions. Uses CRON format as specified by https://github.com/robfig/cron/tree/v3. Defaults to everyday at midnight in the local timezone. To change the timezone, prefix the schedule with CRON_TZ=&lt;Timezone&gt;. For example: &#34;CRON_TZ=UTC 0 0 * * *&#34; is midnight UTC. |





 


<a name="api-resources-koor-v1-UpgradeMode"></a>

### UpgradeMode
The mode of the upgrade

| Name | Number | Description |
| ---- | ------ | ----------- |
| UPGRADE_MODE_UNSPECIFIED | 0 |  |
| UPGRADE_MODE_DISABLED | 1 | Disable upgrades |
| UPGRADE_MODE_NOTIFY | 2 | Notify about new upgrades but do not apply them |
| UPGRADE_MODE_UPGRADE | 3 | Notify about new upgrades and apply them |


 

 

 



<a name="api_resources_stats_v1_stats-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/resources/stats/v1/stats.proto



<a name="api-resources-stats-v1-ClusterRadar"></a>

### ClusterRadar



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster_health | [float](#float) |  |  |
| nodes_health | [float](#float) |  |  |
| capacity_available | [float](#float) |  |  |
| stability | [float](#float) |  |  |
| reliability | [float](#float) |  |  |






<a name="api-resources-stats-v1-ClusterStats"></a>

### ClusterStats



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [ClusterHealth](#api-resources-stats-v1-ClusterHealth) |  |  |
| crashes | [Crash](#api-resources-stats-v1-Crash) | repeated |  |
| services | [Services](#api-resources-stats-v1-Services) |  |  |
| data | [Data](#api-resources-stats-v1-Data) |  |  |
| test | [string](#string) |  |  |
| iops | [IOPS](#api-resources-stats-v1-IOPS) |  |  |






<a name="api-resources-stats-v1-Crash"></a>

### Crash



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| description | [string](#string) |  |  |






<a name="api-resources-stats-v1-Data"></a>

### Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| volumes | [int32](#int32) |  |  |
| pools | [Pools](#api-resources-stats-v1-Pools) |  |  |
| objects | [Objects](#api-resources-stats-v1-Objects) |  |  |
| usage | [Usage](#api-resources-stats-v1-Usage) |  |  |






<a name="api-resources-stats-v1-IOPS"></a>

### IOPS



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| client_read | [int64](#int64) |  |  |
| client_write | [int64](#int64) |  |  |
| client_read_ops | [int64](#int64) |  |  |
| client_write_ops | [int64](#int64) |  |  |






<a name="api-resources-stats-v1-MdsService"></a>

### MdsService



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| daemons_up | [int32](#int32) |  |  |
| hot_standby_count | [int32](#int32) |  |  |






<a name="api-resources-stats-v1-MgrService"></a>

### MgrService



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| active | [string](#string) |  |  |
| standbys | [string](#string) | repeated |  |
| updated_since | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="api-resources-stats-v1-MonService"></a>

### MonService



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| daemon_count | [int32](#int32) |  |  |
| quorum | [string](#string) | repeated |  |
| created_since | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_since | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="api-resources-stats-v1-NodeInfo"></a>

### NodeInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| status | [ResourceStatus](#api-resources-stats-v1-ResourceStatus) |  |  |
| roles | [string](#string) | repeated |  |
| internal_ip | [string](#string) | optional |  |
| external_ip | [string](#string) | optional |  |
| age | [google.protobuf.Timestamp](#google-protobuf-Timestamp) | optional |  |






<a name="api-resources-stats-v1-Objects"></a>

### Objects



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_count | [int32](#int32) |  |  |
| size | [int64](#int64) |  |  |






<a name="api-resources-stats-v1-OsdService"></a>

### OsdService



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| osd_count | [int32](#int32) |  |  |
| osd_up | [int32](#int32) |  |  |
| osd_in | [int32](#int32) |  |  |
| osd_up_updated_since | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| osd_in_updated_since | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="api-resources-stats-v1-PGs"></a>

### PGs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| active_clean | [int32](#int32) |  |  |






<a name="api-resources-stats-v1-Pools"></a>

### Pools



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pools | [int32](#int32) |  |  |
| pgs | [PGs](#api-resources-stats-v1-PGs) |  |  |






<a name="api-resources-stats-v1-ResourceInfo"></a>

### ResourceInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apiversion | [string](#string) |  |  |
| kind | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| name | [string](#string) |  |  |
| status | [ResourceStatus](#api-resources-stats-v1-ResourceStatus) |  |  |
| replicas | [int32](#int32) |  |  |
| reliability | [ReliabilityScore](#api-resources-stats-v1-ReliabilityScore) |  |  |
| version | [string](#string) | optional |  |






<a name="api-resources-stats-v1-RgwService"></a>

### RgwService



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| active_daemon | [int32](#int32) |  |  |
| host_count | [int32](#int32) |  |  |
| zone_count | [int32](#int32) |  |  |






<a name="api-resources-stats-v1-Services"></a>

### Services



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mon | [MonService](#api-resources-stats-v1-MonService) |  |  |
| mgr | [MgrService](#api-resources-stats-v1-MgrService) |  |  |
| mds | [MdsService](#api-resources-stats-v1-MdsService) |  |  |
| osd | [OsdService](#api-resources-stats-v1-OsdService) |  |  |
| rgw | [RgwService](#api-resources-stats-v1-RgwService) |  |  |






<a name="api-resources-stats-v1-Usage"></a>

### Usage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| used | [int64](#int64) |  |  |
| available | [int64](#int64) |  |  |
| total | [int64](#int64) |  |  |





 


<a name="api-resources-stats-v1-ClusterHealth"></a>

### ClusterHealth


| Name | Number | Description |
| ---- | ------ | ----------- |
| CLUSTER_HEALTH_UNSPECIFIED | 0 |  |
| CLUSTER_HEALTH_OFFLINE | 1 |  |
| CLUSTER_HEALTH_OK | 2 |  |
| CLUSTER_HEALTH_WARN | 3 |  |
| CLUSTER_HEALTH_ERR | 4 |  |



<a name="api-resources-stats-v1-ReliabilityScore"></a>

### ReliabilityScore


| Name | Number | Description |
| ---- | ------ | ----------- |
| RELIABILITY_SCORE_UNSPECIFIED | 0 |  |
| RELIABILITY_SCORE_UNKNOWN | 1 |  |
| RELIABILITY_SCORE_NONE | 2 |  |
| RELIABILITY_SCORE_DEGRADED | 3 |  |
| RELIABILITY_SCORE_OK | 4 |  |



<a name="api-resources-stats-v1-ResourceStatus"></a>

### ResourceStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| RESOURCE_STATUS_UNSPECIFIED | 0 |  |
| RESOURCE_STATUS_UNKNOWN | 1 |  |
| RESOURCE_STATUS_READY | 2 |  |
| RESOURCE_STATUS_PROGRESSING | 3 |  |
| RESOURCE_STATUS_NOT_READY | 4 |  |


 

 

 



<a name="api_resources_timestamp_v1_timestamp-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/resources/timestamp/v1/timestamp.proto



<a name="api-resources-timestamp-v1-Timestamp"></a>

### Timestamp
Timestamp for storage messages.  We&#39;ve defined a new local type wrapper
of google.protobuf.Timestamp so we can implement sql.Scanner and sql.Valuer
interfaces.  See:
https://golang.org/pkg/database/sql/#Scanner
https://golang.org/pkg/database/sql/driver/#Valuer


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 

 

 

 



<a name="api_services_auth_v1_auth-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/services/auth/v1/auth.proto



<a name="api-services-auth-v1-CheckTokenRequest"></a>

### CheckTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="api-services-auth-v1-CheckTokenResponse"></a>

### CheckTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="api-services-auth-v1-LoginRequest"></a>

### LoginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="api-services-auth-v1-LoginResponse"></a>

### LoginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| expires | [api.resources.timestamp.v1.Timestamp](#api-resources-timestamp-v1-Timestamp) |  |  |
| account_id | [string](#string) |  |  |






<a name="api-services-auth-v1-LogoutRequest"></a>

### LogoutRequest







<a name="api-services-auth-v1-LogoutResponse"></a>

### LogoutResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |





 

 

 


<a name="api-services-auth-v1-AuthService"></a>

### AuthService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Login | [LoginRequest](#api-services-auth-v1-LoginRequest) | [LoginResponse](#api-services-auth-v1-LoginResponse) |  |
| Logout | [LogoutRequest](#api-services-auth-v1-LogoutRequest) | [LogoutResponse](#api-services-auth-v1-LogoutResponse) |  |
| CheckToken | [CheckTokenRequest](#api-services-auth-v1-CheckTokenRequest) | [CheckTokenResponse](#api-services-auth-v1-CheckTokenResponse) |  |

 



<a name="api_services_ceph_v1_ceph-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/services/ceph/v1/ceph.proto



<a name="api-services-ceph-v1-CreateCephUserRequest"></a>

### CreateCephUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ceph_user | [api.resources.ceph.v1.User](#api-resources-ceph-v1-User) |  |  |






<a name="api-services-ceph-v1-CreateCephUserResponse"></a>

### CreateCephUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ceph_user | [api.resources.ceph.v1.User](#api-resources-ceph-v1-User) |  |  |






<a name="api-services-ceph-v1-DeleteCephUserRequest"></a>

### DeleteCephUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |






<a name="api-services-ceph-v1-DeleteCephUserResponse"></a>

### DeleteCephUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [string](#string) |  |  |






<a name="api-services-ceph-v1-GetCephUsersRequest"></a>

### GetCephUsersRequest







<a name="api-services-ceph-v1-GetCephUsersResponse"></a>

### GetCephUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ceph_users | [api.resources.ceph.v1.User](#api-resources-ceph-v1-User) | repeated |  |





 

 

 


<a name="api-services-ceph-v1-CephService"></a>

### CephService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetCephUsers | [GetCephUsersRequest](#api-services-ceph-v1-GetCephUsersRequest) | [GetCephUsersResponse](#api-services-ceph-v1-GetCephUsersResponse) |  |
| CreateCephUser | [CreateCephUserRequest](#api-services-ceph-v1-CreateCephUserRequest) | [CreateCephUserResponse](#api-services-ceph-v1-CreateCephUserResponse) |  |
| DeleteCephUser | [DeleteCephUserRequest](#api-services-ceph-v1-DeleteCephUserRequest) | [DeleteCephUserResponse](#api-services-ceph-v1-DeleteCephUserResponse) |  |

 



<a name="api_services_cluster_v1_cluster-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/services/cluster/v1/cluster.proto



<a name="api-services-cluster-v1-CancelNetworkTestRequest"></a>

### CancelNetworkTestRequest







<a name="api-services-cluster-v1-CancelNetworkTestResponse"></a>

### CancelNetworkTestResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| not_running | [bool](#bool) |  |  |






<a name="api-services-cluster-v1-GetKoorClusterRequest"></a>

### GetKoorClusterRequest







<a name="api-services-cluster-v1-GetKoorClusterResponse"></a>

### GetKoorClusterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| koor_cluster | [api.resources.koor.v1.KoorCluster](#api-resources-koor-v1-KoorCluster) |  |  |






<a name="api-services-cluster-v1-GetNetworkTestResultsRequest"></a>

### GetNetworkTestResultsRequest







<a name="api-services-cluster-v1-GetNetworkTestResultsResponse"></a>

### GetNetworkTestResultsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| file_name | [string](#string) |  |  |
| file_type | [string](#string) |  |  |
| file_contents | [bytes](#bytes) |  |  |






<a name="api-services-cluster-v1-GetNetworkTestStatusRequest"></a>

### GetNetworkTestStatusRequest







<a name="api-services-cluster-v1-GetNetworkTestStatusResponse"></a>

### GetNetworkTestStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| running | [bool](#bool) |  |  |
| logs | [string](#string) |  |  |






<a name="api-services-cluster-v1-GetResourcesRequest"></a>

### GetResourcesRequest







<a name="api-services-cluster-v1-GetResourcesResponse"></a>

### GetResourcesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resources | [api.resources.ceph.v1.Resources](#api-resources-ceph-v1-Resources) |  |  |






<a name="api-services-cluster-v1-GetTroubleshootReportRequest"></a>

### GetTroubleshootReportRequest







<a name="api-services-cluster-v1-GetTroubleshootReportResponse"></a>

### GetTroubleshootReportResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| report | [string](#string) |  |  |






<a name="api-services-cluster-v1-SaveResourcesRequest"></a>

### SaveResourcesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resource | [api.resources.ceph.v1.Resource](#api-resources-ceph-v1-Resource) |  |  |






<a name="api-services-cluster-v1-SaveResourcesResponse"></a>

### SaveResourcesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resource | [api.resources.ceph.v1.Resource](#api-resources-ceph-v1-Resource) |  |  |






<a name="api-services-cluster-v1-SetScrubbingScheduleRequest"></a>

### SetScrubbingScheduleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| osd_scrubbing_schedule | [api.resources.koor.v1.OSDScrubbingSchedule](#api-resources-koor-v1-OSDScrubbingSchedule) |  |  |






<a name="api-services-cluster-v1-SetScrubbingScheduleResponse"></a>

### SetScrubbingScheduleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| osd_scrubbing_schedule | [api.resources.koor.v1.OSDScrubbingSchedule](#api-resources-koor-v1-OSDScrubbingSchedule) |  |  |






<a name="api-services-cluster-v1-StartNetworkTestRequest"></a>

### StartNetworkTestRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host_network | [bool](#bool) |  |  |
| output_format | [NetworkTestOutputFormat](#api-services-cluster-v1-NetworkTestOutputFormat) |  |  |






<a name="api-services-cluster-v1-StartNetworkTestResponse"></a>

### StartNetworkTestResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| already_running | [bool](#bool) |  |  |





 


<a name="api-services-cluster-v1-NetworkTestOutputFormat"></a>

### NetworkTestOutputFormat


| Name | Number | Description |
| ---- | ------ | ----------- |
| NETWORK_TEST_OUTPUT_FORMAT_UNSPECIFIED | 0 |  |
| NETWORK_TEST_OUTPUT_FORMAT_CSV | 1 |  |
| NETWORK_TEST_OUTPUT_FORMAT_EXCELIZE | 2 |  |


 

 


<a name="api-services-cluster-v1-ClusterService"></a>

### ClusterService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetKoorCluster | [GetKoorClusterRequest](#api-services-cluster-v1-GetKoorClusterRequest) | [GetKoorClusterResponse](#api-services-cluster-v1-GetKoorClusterResponse) |  |
| GetTroubleshootReport | [GetTroubleshootReportRequest](#api-services-cluster-v1-GetTroubleshootReportRequest) | [GetTroubleshootReportResponse](#api-services-cluster-v1-GetTroubleshootReportResponse) |  |
| GetNetworkTestStatus | [GetNetworkTestStatusRequest](#api-services-cluster-v1-GetNetworkTestStatusRequest) | [GetNetworkTestStatusResponse](#api-services-cluster-v1-GetNetworkTestStatusResponse) |  |
| StartNetworkTest | [StartNetworkTestRequest](#api-services-cluster-v1-StartNetworkTestRequest) | [StartNetworkTestResponse](#api-services-cluster-v1-StartNetworkTestResponse) |  |
| CancelNetworkTest | [CancelNetworkTestRequest](#api-services-cluster-v1-CancelNetworkTestRequest) | [CancelNetworkTestResponse](#api-services-cluster-v1-CancelNetworkTestResponse) |  |
| GetNetworkTestResults | [GetNetworkTestResultsRequest](#api-services-cluster-v1-GetNetworkTestResultsRequest) | [GetNetworkTestResultsResponse](#api-services-cluster-v1-GetNetworkTestResultsResponse) |  |
| SetScrubbingSchedule | [SetScrubbingScheduleRequest](#api-services-cluster-v1-SetScrubbingScheduleRequest) | [SetScrubbingScheduleResponse](#api-services-cluster-v1-SetScrubbingScheduleResponse) |  |
| GetResources | [GetResourcesRequest](#api-services-cluster-v1-GetResourcesRequest) | [GetResourcesResponse](#api-services-cluster-v1-GetResourcesResponse) |  |
| SaveResources | [SaveResourcesRequest](#api-services-cluster-v1-SaveResourcesRequest) | [SaveResourcesResponse](#api-services-cluster-v1-SaveResourcesResponse) |  |

 



<a name="api_services_stats_v1_stats-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/services/stats/v1/stats.proto



<a name="api-services-stats-v1-GetClusterNodesRequest"></a>

### GetClusterNodesRequest







<a name="api-services-stats-v1-GetClusterNodesResponse"></a>

### GetClusterNodesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodes | [api.resources.stats.v1.NodeInfo](#api-resources-stats-v1-NodeInfo) | repeated |  |






<a name="api-services-stats-v1-GetClusterRadarRequest"></a>

### GetClusterRadarRequest







<a name="api-services-stats-v1-GetClusterRadarResponse"></a>

### GetClusterRadarResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| radar | [api.resources.stats.v1.ClusterRadar](#api-resources-stats-v1-ClusterRadar) |  |  |






<a name="api-services-stats-v1-GetClusterResourcesRequest"></a>

### GetClusterResourcesRequest







<a name="api-services-stats-v1-GetClusterResourcesResponse"></a>

### GetClusterResourcesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resources | [api.resources.stats.v1.ResourceInfo](#api-resources-stats-v1-ResourceInfo) | repeated |  |
| deployments | [api.resources.stats.v1.ResourceInfo](#api-resources-stats-v1-ResourceInfo) | repeated |  |






<a name="api-services-stats-v1-GetClusterStatsRequest"></a>

### GetClusterStatsRequest







<a name="api-services-stats-v1-GetClusterStatsResponse"></a>

### GetClusterStatsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stats | [api.resources.stats.v1.ClusterStats](#api-resources-stats-v1-ClusterStats) |  |  |





 

 

 


<a name="api-services-stats-v1-StatsService"></a>

### StatsService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetClusterStats | [GetClusterStatsRequest](#api-services-stats-v1-GetClusterStatsRequest) | [GetClusterStatsResponse](#api-services-stats-v1-GetClusterStatsResponse) |  |
| GetClusterResources | [GetClusterResourcesRequest](#api-services-stats-v1-GetClusterResourcesRequest) | [GetClusterResourcesResponse](#api-services-stats-v1-GetClusterResourcesResponse) |  |
| GetClusterNodes | [GetClusterNodesRequest](#api-services-stats-v1-GetClusterNodesRequest) | [GetClusterNodesResponse](#api-services-stats-v1-GetClusterNodesResponse) |  |
| GetClusterRadar | [GetClusterRadarRequest](#api-services-stats-v1-GetClusterRadarRequest) | [GetClusterRadarResponse](#api-services-stats-v1-GetClusterRadarResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

