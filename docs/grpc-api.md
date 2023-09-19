# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/resources/stats/stats.proto](#api_resources_stats_stats-proto)
    - [ClusterRadar](#resources-stats-ClusterRadar)
    - [ClusterResources](#resources-stats-ClusterResources)
    - [ClusterStats](#resources-stats-ClusterStats)
    - [Crash](#resources-stats-Crash)
    - [Data](#resources-stats-Data)
    - [DetailedProductVersions](#resources-stats-DetailedProductVersions)
    - [DetailedVersion](#resources-stats-DetailedVersion)
    - [IOPS](#resources-stats-IOPS)
    - [KoorCluster](#resources-stats-KoorCluster)
    - [KoorClusterSpec](#resources-stats-KoorClusterSpec)
    - [KoorClusterStatus](#resources-stats-KoorClusterStatus)
    - [MdsService](#resources-stats-MdsService)
    - [MgrService](#resources-stats-MgrService)
    - [MonService](#resources-stats-MonService)
    - [NodeInfo](#resources-stats-NodeInfo)
    - [Objects](#resources-stats-Objects)
    - [OsdService](#resources-stats-OsdService)
    - [PGs](#resources-stats-PGs)
    - [Pools](#resources-stats-Pools)
    - [ProductVersions](#resources-stats-ProductVersions)
    - [ResourceInfo](#resources-stats-ResourceInfo)
    - [RgwService](#resources-stats-RgwService)
    - [Services](#resources-stats-Services)
    - [UpgradeOptions](#resources-stats-UpgradeOptions)
    - [Usage](#resources-stats-Usage)
  
    - [ReliabilityScore](#resources-stats-ReliabilityScore)
    - [ResourceStatus](#resources-stats-ResourceStatus)
    - [UpgradeMode](#resources-stats-UpgradeMode)
  
- [api/resources/timestamp/timestamp.proto](#api_resources_timestamp_timestamp-proto)
    - [Timestamp](#resources-timestamp-Timestamp)
  
- [api/services/auth/auth.proto](#api_services_auth_auth-proto)
    - [CheckTokenRequest](#services-auth-CheckTokenRequest)
    - [CheckTokenResponse](#services-auth-CheckTokenResponse)
    - [LoginRequest](#services-auth-LoginRequest)
    - [LoginResponse](#services-auth-LoginResponse)
    - [LogoutRequest](#services-auth-LogoutRequest)
    - [LogoutResponse](#services-auth-LogoutResponse)
  
    - [AuthService](#services-auth-AuthService)
  
- [api/services/stats/stats.proto](#api_services_stats_stats-proto)
    - [GetClusterNodesRequest](#stats-GetClusterNodesRequest)
    - [GetClusterNodesResponse](#stats-GetClusterNodesResponse)
    - [GetClusterRadarRequest](#stats-GetClusterRadarRequest)
    - [GetClusterRadarResponse](#stats-GetClusterRadarResponse)
    - [GetClusterResourcesRequest](#stats-GetClusterResourcesRequest)
    - [GetClusterResourcesResponse](#stats-GetClusterResourcesResponse)
    - [GetClusterStatsRequest](#stats-GetClusterStatsRequest)
    - [GetClusterStatsResponse](#stats-GetClusterStatsResponse)
    - [GetKoorClusterRequest](#stats-GetKoorClusterRequest)
    - [GetKoorClusterResponse](#stats-GetKoorClusterResponse)
  
    - [StatsService](#stats-StatsService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_resources_stats_stats-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/resources/stats/stats.proto



<a name="resources-stats-ClusterRadar"></a>

### ClusterRadar



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster_health | [float](#float) |  |  |
| nodes_health | [float](#float) |  |  |
| capacity_available | [float](#float) |  |  |
| stability | [float](#float) |  |  |
| reliability | [float](#float) |  |  |






<a name="resources-stats-ClusterResources"></a>

### ClusterResources



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodes | [string](#string) |  | The number of nodes in the cluster |
| storage | [string](#string) |  | Ephemeral Storage available |
| cpu | [string](#string) |  | CPU cores available |
| memory | [string](#string) |  | Memory available |






<a name="resources-stats-ClusterStats"></a>

### ClusterStats



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [string](#string) |  |  |
| crashes | [Crash](#resources-stats-Crash) | repeated |  |
| services | [Services](#resources-stats-Services) |  |  |
| data | [Data](#resources-stats-Data) |  |  |
| test | [string](#string) |  |  |
| iops | [IOPS](#resources-stats-IOPS) |  |  |






<a name="resources-stats-Crash"></a>

### Crash



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| description | [string](#string) |  |  |






<a name="resources-stats-Data"></a>

### Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| volumes | [int32](#int32) |  |  |
| pools | [Pools](#resources-stats-Pools) |  |  |
| objects | [Objects](#resources-stats-Objects) |  |  |
| usage | [Usage](#resources-stats-Usage) |  |  |






<a name="resources-stats-DetailedProductVersions"></a>

### DetailedProductVersions
Represents a map of products to detailed versions, which include images or helm charts.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| koor_operator | [DetailedVersion](#resources-stats-DetailedVersion) |  | The detailed Koor Operator version. |
| ksd | [DetailedVersion](#resources-stats-DetailedVersion) |  | The detailed Koor Storage Distribution version. |
| ceph | [DetailedVersion](#resources-stats-DetailedVersion) |  | The detailed Ceph version. |






<a name="resources-stats-DetailedVersion"></a>

### DetailedVersion
Defines a detailed version of a product, which includes a container image or a helm chart.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  | The product version, must be a valid semver. |
| image_uri | [string](#string) |  | The URI of the container image. |
| image_hash | [string](#string) |  | The hash of the container image. |
| helm_repository | [string](#string) |  | The URI of the helm repository. |
| helm_chart | [string](#string) |  | The name of the helm chart in the repository. |






<a name="resources-stats-IOPS"></a>

### IOPS



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| client_read | [string](#string) |  |  |
| client_write | [string](#string) |  |  |
| client_read_ops | [string](#string) |  |  |
| client_write_ops | [string](#string) |  |  |






<a name="resources-stats-KoorCluster"></a>

### KoorCluster



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| spec | [KoorClusterSpec](#resources-stats-KoorClusterSpec) |  |  |
| status | [KoorClusterStatus](#resources-stats-KoorClusterStatus) |  |  |






<a name="resources-stats-KoorClusterSpec"></a>

### KoorClusterSpec
Represents the state of KoorCluster


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| use_all_devices | [bool](#bool) |  | Use all devices on nodes |
| monitoring_enabled | [bool](#bool) |  | Enable monitoring. Requires Prometheus to be pre-installed. |
| dashboard_enabled | [bool](#bool) |  | Enable the ceph dashboard for viewing cluster status |
| toolbox_enabled | [bool](#bool) |  | Installs a debugging toolbox deployment |
| upgrade_options | [UpgradeOptions](#resources-stats-UpgradeOptions) |  | Specifies the upgrade options for new ceph versions |
| ksd_release_name | [string](#string) |  | The name to use for KSD helm release. |
| ksd_cluster_release_name | [string](#string) |  | The name to use for KSD cluster helm release. |






<a name="resources-stats-KoorClusterStatus"></a>

### KoorClusterStatus
Represents the status of the KoorCluster CRD


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total_resources | [ClusterResources](#resources-stats-ClusterResources) |  | The total resources available in the cluster nodes |
| meets_minimum_resources | [bool](#bool) |  | Does the cluster meet the minimum recommended resources |
| current_versions | [ProductVersions](#resources-stats-ProductVersions) |  | The current versions of rook and ceph |
| latest_versions | [DetailedProductVersions](#resources-stats-DetailedProductVersions) |  | The latest versions of rook and ceph |






<a name="resources-stats-MdsService"></a>

### MdsService



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| daemons_up | [int32](#int32) |  |  |
| hot_standby_count | [int32](#int32) |  |  |






<a name="resources-stats-MgrService"></a>

### MgrService



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| active | [string](#string) |  |  |
| standbys | [string](#string) | repeated |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="resources-stats-MonService"></a>

### MonService



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| daemon_count | [int32](#int32) |  |  |
| quorum | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="resources-stats-NodeInfo"></a>

### NodeInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| status | [ResourceStatus](#resources-stats-ResourceStatus) |  |  |
| roles | [string](#string) | repeated |  |
| internal_ip | [string](#string) |  |  |
| external_ip | [string](#string) |  |  |
| age | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="resources-stats-Objects"></a>

### Objects



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_count | [int32](#int32) |  |  |
| size | [string](#string) |  |  |






<a name="resources-stats-OsdService"></a>

### OsdService



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| osd_count | [int32](#int32) |  |  |
| osd_up | [int32](#int32) |  |  |
| osd_in | [int32](#int32) |  |  |
| osd_up_updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| osd_in_updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="resources-stats-PGs"></a>

### PGs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| active_clean | [int32](#int32) |  |  |






<a name="resources-stats-Pools"></a>

### Pools



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pools | [int32](#int32) |  |  |
| pgs | [PGs](#resources-stats-PGs) |  |  |






<a name="resources-stats-ProductVersions"></a>

### ProductVersions
Represents a map of products to version strings.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kube | [string](#string) |  | Kubernetes version, must be a valid semver. |
| koor_operator | [string](#string) |  | Koor Operator version, must be a valid semver. |
| ksd | [string](#string) |  | Koor Storage Distribution version, must be a valid semver. |
| ceph | [string](#string) |  | Ceph version, must be a valid semver. |






<a name="resources-stats-ResourceInfo"></a>

### ResourceInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apiversion | [string](#string) |  |  |
| kind | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| name | [string](#string) |  |  |
| status | [ResourceStatus](#resources-stats-ResourceStatus) |  |  |
| replicas | [int32](#int32) |  |  |
| reliability | [ReliabilityScore](#resources-stats-ReliabilityScore) |  |  |






<a name="resources-stats-RgwService"></a>

### RgwService



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| active_daemon | [int32](#int32) |  |  |
| host_count | [int32](#int32) |  |  |
| zone_count | [int32](#int32) |  |  |






<a name="resources-stats-Services"></a>

### Services



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mon | [MonService](#resources-stats-MonService) |  |  |
| mgr | [MgrService](#resources-stats-MgrService) |  |  |
| mds | [MdsService](#resources-stats-MdsService) |  |  |
| osd | [OsdService](#resources-stats-OsdService) |  |  |
| rgw | [RgwService](#resources-stats-RgwService) |  |  |






<a name="resources-stats-UpgradeOptions"></a>

### UpgradeOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mode | [UpgradeMode](#resources-stats-UpgradeMode) |  | Upgrade mode |
| endpoint | [string](#string) |  | The api endpoint used to find the ceph latest version |
| schedule | [string](#string) |  | The schedule to check for new versions. Uses CRON format as specified by https://github.com/robfig/cron/tree/v3. Defaults to everyday at midnight in the local timezone. To change the timezone, prefix the schedule with CRON_TZ=&lt;Timezone&gt;. For example: &#34;CRON_TZ=UTC 0 0 * * *&#34; is midnight UTC. |






<a name="resources-stats-Usage"></a>

### Usage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| used | [int64](#int64) |  |  |
| available | [int64](#int64) |  |  |
| total | [int64](#int64) |  |  |





 


<a name="resources-stats-ReliabilityScore"></a>

### ReliabilityScore


| Name | Number | Description |
| ---- | ------ | ----------- |
| RELIABILITY_UNKNOWN | 0 |  |
| RELIABILITY_NONE | 1 |  |
| RELIABILITY_DEGRADED | 2 |  |
| RELIABILITY_OK | 3 |  |



<a name="resources-stats-ResourceStatus"></a>

### ResourceStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| RESOURCE_UNKNOWN | 0 |  |
| RESOURCE_READY | 1 |  |
| RESOURCE_PROGRESSING | 2 |  |
| RESOURCE_NOT_READY | 3 |  |



<a name="resources-stats-UpgradeMode"></a>

### UpgradeMode
The mode of the upgrade

| Name | Number | Description |
| ---- | ------ | ----------- |
| UPGRADE_MODE_UNSPECIFIED | 0 |  |
| UPGRADE_MODE_DISABLED | 1 | Disable upgrades |
| UPGRADE_MODE_NOTIFY | 2 | Notify about new upgrades but do not apply them |
| UPGRADE_MODE_UPGRADE | 3 | Notify about new upgrades and apply them |


 

 

 



<a name="api_resources_timestamp_timestamp-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/resources/timestamp/timestamp.proto



<a name="resources-timestamp-Timestamp"></a>

### Timestamp
Timestamp for storage messages.  We&#39;ve defined a new local type wrapper
of google.protobuf.Timestamp so we can implement sql.Scanner and sql.Valuer
interfaces.  See:
https://golang.org/pkg/database/sql/#Scanner
https://golang.org/pkg/database/sql/driver/#Valuer


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 

 

 

 



<a name="api_services_auth_auth-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/services/auth/auth.proto



<a name="services-auth-CheckTokenRequest"></a>

### CheckTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="services-auth-CheckTokenResponse"></a>

### CheckTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="services-auth-LoginRequest"></a>

### LoginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="services-auth-LoginResponse"></a>

### LoginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| expires | [resources.timestamp.Timestamp](#resources-timestamp-Timestamp) |  |  |
| account_id | [string](#string) |  |  |






<a name="services-auth-LogoutRequest"></a>

### LogoutRequest







<a name="services-auth-LogoutResponse"></a>

### LogoutResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |





 

 

 


<a name="services-auth-AuthService"></a>

### AuthService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Login | [LoginRequest](#services-auth-LoginRequest) | [LoginResponse](#services-auth-LoginResponse) |  |
| Logout | [LogoutRequest](#services-auth-LogoutRequest) | [LogoutResponse](#services-auth-LogoutResponse) |  |
| CheckToken | [CheckTokenRequest](#services-auth-CheckTokenRequest) | [CheckTokenResponse](#services-auth-CheckTokenResponse) |  |

 



<a name="api_services_stats_stats-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/services/stats/stats.proto



<a name="stats-GetClusterNodesRequest"></a>

### GetClusterNodesRequest







<a name="stats-GetClusterNodesResponse"></a>

### GetClusterNodesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodes | [resources.stats.NodeInfo](#resources-stats-NodeInfo) | repeated |  |






<a name="stats-GetClusterRadarRequest"></a>

### GetClusterRadarRequest







<a name="stats-GetClusterRadarResponse"></a>

### GetClusterRadarResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| radar | [resources.stats.ClusterRadar](#resources-stats-ClusterRadar) |  |  |






<a name="stats-GetClusterResourcesRequest"></a>

### GetClusterResourcesRequest







<a name="stats-GetClusterResourcesResponse"></a>

### GetClusterResourcesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resources | [resources.stats.ResourceInfo](#resources-stats-ResourceInfo) | repeated |  |
| deployments | [resources.stats.ResourceInfo](#resources-stats-ResourceInfo) | repeated |  |






<a name="stats-GetClusterStatsRequest"></a>

### GetClusterStatsRequest







<a name="stats-GetClusterStatsResponse"></a>

### GetClusterStatsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stats | [resources.stats.ClusterStats](#resources-stats-ClusterStats) |  |  |






<a name="stats-GetKoorClusterRequest"></a>

### GetKoorClusterRequest







<a name="stats-GetKoorClusterResponse"></a>

### GetKoorClusterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| koor_cluster | [resources.stats.KoorCluster](#resources-stats-KoorCluster) |  |  |





 

 

 


<a name="stats-StatsService"></a>

### StatsService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetClusterStats | [GetClusterStatsRequest](#stats-GetClusterStatsRequest) | [GetClusterStatsResponse](#stats-GetClusterStatsResponse) |  |
| GetClusterResources | [GetClusterResourcesRequest](#stats-GetClusterResourcesRequest) | [GetClusterResourcesResponse](#stats-GetClusterResourcesResponse) |  |
| GetClusterNodes | [GetClusterNodesRequest](#stats-GetClusterNodesRequest) | [GetClusterNodesResponse](#stats-GetClusterNodesResponse) |  |
| GetClusterRadar | [GetClusterRadarRequest](#stats-GetClusterRadarRequest) | [GetClusterRadarResponse](#stats-GetClusterRadarResponse) |  |
| GetKoorCluster | [GetKoorClusterRequest](#stats-GetKoorClusterRequest) | [GetKoorClusterResponse](#stats-GetKoorClusterResponse) |  |

 



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

