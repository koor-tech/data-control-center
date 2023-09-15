# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/resources/stats/stats.proto](#api_resources_stats_stats-proto)
    - [ClusterRadar](#resources-stats-ClusterRadar)
    - [ClusterStats](#resources-stats-ClusterStats)
    - [Crash](#resources-stats-Crash)
    - [Data](#resources-stats-Data)
    - [Io](#resources-stats-Io)
    - [MdsService](#resources-stats-MdsService)
    - [MgrService](#resources-stats-MgrService)
    - [MonService](#resources-stats-MonService)
    - [NodeInfo](#resources-stats-NodeInfo)
    - [Objects](#resources-stats-Objects)
    - [OsdService](#resources-stats-OsdService)
    - [Pgs](#resources-stats-Pgs)
    - [Pools](#resources-stats-Pools)
    - [ResourceInfo](#resources-stats-ResourceInfo)
    - [RgwService](#resources-stats-RgwService)
    - [Services](#resources-stats-Services)
    - [Usage](#resources-stats-Usage)
  
    - [ResourceStatus](#resources-stats-ResourceStatus)
  
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
| capacity_used | [float](#float) |  |  |
| stability | [float](#float) |  |  |
| reliability | [float](#float) |  |  |






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
| io | [Io](#resources-stats-Io) |  |  |






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






<a name="resources-stats-Io"></a>

### Io



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| client_read | [string](#string) |  |  |
| client_read_ops | [string](#string) |  |  |
| client_write_ops | [string](#string) |  |  |






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






<a name="resources-stats-Pgs"></a>

### Pgs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| active_clean | [int32](#int32) |  |  |






<a name="resources-stats-Pools"></a>

### Pools



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pools | [int32](#int32) |  |  |
| pgs | [Pgs](#resources-stats-Pgs) |  |  |






<a name="resources-stats-ResourceInfo"></a>

### ResourceInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apiversion | [string](#string) |  |  |
| kind | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| name | [string](#string) |  |  |
| status | [ResourceStatus](#resources-stats-ResourceStatus) |  |  |






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






<a name="resources-stats-Usage"></a>

### Usage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| used | [int64](#int64) |  |  |
| available | [int64](#int64) |  |  |
| total | [int64](#int64) |  |  |





 


<a name="resources-stats-ResourceStatus"></a>

### ResourceStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| RESOURCE_UNKNOWN | 0 |  |
| RESOURCE_READY | 1 |  |
| RESOURCE_NOT_READY | 2 |  |


 

 

 



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





 

 

 


<a name="stats-StatsService"></a>

### StatsService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetClusterStats | [GetClusterStatsRequest](#stats-GetClusterStatsRequest) | [GetClusterStatsResponse](#stats-GetClusterStatsResponse) |  |
| GetClusterResources | [GetClusterResourcesRequest](#stats-GetClusterResourcesRequest) | [GetClusterResourcesResponse](#stats-GetClusterResourcesResponse) |  |
| GetClusterNodes | [GetClusterNodesRequest](#stats-GetClusterNodesRequest) | [GetClusterNodesResponse](#stats-GetClusterNodesResponse) |  |
| GetClusterRadar | [GetClusterRadarRequest](#stats-GetClusterRadarRequest) | [GetClusterRadarResponse](#stats-GetClusterRadarResponse) |  |

 



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

