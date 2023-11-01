package ceph

import (
	"encoding/json"
	"errors"
	statsv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/stats/v1"
	"time"
)

type Summary struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
}

type Detail struct {
	Message string `json:"message"`
}

type Check struct {
	Severity string   `json:"severity"`
	Summary  Summary  `json:"summary"`
	Details  []Detail `json:"detail"`
	Muted    bool     `json:"muted"`
	Type     string   `json:"type"`
}
type Health struct {
	Status string        `json:"status"`
	Checks []Check       `json:"checks"`
	Mutes  []interface{} `json:"mutes"`
}

type Mons []struct {
	Name string `json:"name"`
}

type Monmap struct {
	Fsid string `json:"fsid"` //change to uuid
	Mons Mons   `json:"mons"`

	Modified time.Time `json:"modified"`
	Created  time.Time `json:"created"`
}

type MonStatus struct {
	Name   string `json:"name"`
	Monmap Monmap `json:"monmap"`
}

type Standbys struct {
	Name string `json:"name"`
}

// cephTime is used to parse the time
type cephTime struct {
	time.Time
}

func (mt *cephTime) UnmarshalJSON(data []byte) error {
	var timeStr string
	const layout = "2006-01-02T15:04:05-0700"
	if err := json.Unmarshal(data, &timeStr); err != nil {
		return err
	}

	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return err
	}

	mt.Time = parsedTime
	return nil
}

type MgrMap struct {
	ActiveName string     `json:"active_name"`
	Standbys   []Standbys `json:"standbys"`

	ActiveChange cephTime `json:"active_change"`
}

type MdsMap struct {
	Up                 interface{} `json:"up"`
	Failed             interface{} `json:"failed"`
	Damaged            interface{} `json:"damaged"`
	Stopped            interface{} `json:"stopped"`
	StandbyCountWanted int         `json:"standby_count_wanted"`
}

type Filesystems struct {
	Id     int    `json:"id"`
	MdsMap MdsMap `json:"mdsmap"`
}

type FsMap struct {
	Epoch       int           `json:"epoch"`
	Filesystems []Filesystems `json:"filesystems"`
}

type Osds struct {
	Osd int `json:"osd"`
	Up  int `json:"up"`
	In  int `json:"in"`
}

type OsdMap struct {
	Osds []Osds `json:"osds"`

	LastUpChange cephTime `json:"last_up_change"`
	LastInChange cephTime `json:"last_in_change"`
}

type PgStatus struct {
	ActiveClean int `json:"active+clean"`
}

type Pool struct {
	Pool     int      `json:"pool"`
	PgStatus PgStatus `json:"pg_status"`
}

type ObjectStats struct {
	NumObjects int `json:"num_objects"`
}

type PGInfo struct {
	ObjectStats ObjectStats `json:"object_stats"`
}

type Stats struct {
	TotalBytes         int64   `json:"total_bytes"`
	TotalAvailBytes    int64   `json:"total_avail_bytes"`
	TotalUsedBytes     int64   `json:"total_used_bytes"`
	TotalUsedRawBytes  int     `json:"total_used_raw_bytes"`
	TotalUsedRawRatio  float64 `json:"total_used_raw_ratio"`
	NumOsds            int     `json:"num_osds"`
	NumPerPoolOsds     int     `json:"num_per_pool_osds"`
	NumPerPoolOmapOsds int     `json:"num_per_pool_omap_osds"`
}

type DFPoolStats struct {
	Stored             int     `json:"stored"`
	StoredData         int     `json:"stored_data"`
	StoredOmap         int     `json:"stored_omap"`
	Objects            int     `json:"objects"`
	KbUsed             int     `json:"kb_used"`
	BytesUsed          int     `json:"bytes_used"`
	DataBytesUsed      int     `json:"data_bytes_used"`
	OmapBytesUsed      int     `json:"omap_bytes_used"`
	PercentUsed        float64 `json:"percent_used"`
	MaxAvail           int64   `json:"max_avail"`
	QuotaObjects       int     `json:"quota_objects"`
	QuotaBytes         int     `json:"quota_bytes"`
	Dirty              int     `json:"dirty"`
	Rd                 int     `json:"rd"`
	RdBytes            int     `json:"rd_bytes"`
	Wr                 int     `json:"wr"`
	WrBytes            int     `json:"wr_bytes"`
	CompressBytesUsed  int     `json:"compress_bytes_used"`
	CompressUnderBytes int     `json:"compress_under_bytes"`
	StoredRaw          int     `json:"stored_raw"`
	AvailRaw           int64   `json:"avail_raw"`
}

type DFPools struct {
	Name        string      `json:"name"`
	Id          int         `json:"id"`
	DFPoolStats DFPoolStats `json:"stats"`
}

type DF struct {
	Stats Stats     `json:"stats"`
	Pools []DFPools `json:"pools"`
}

type ClientPerf struct {
	ReadBytesSec          int `json:"read_bytes_sec"`
	ReadOpPerSec          int `json:"read_op_per_sec"`
	WriteBytesSec         int `json:"write_bytes_sec"`
	WriteOpPerSec         int `json:"write_op_per_sec"`
	RecoveringBytesPerSec int `json:"recovering_bytes_per_sec"`
}

type HealthStatus struct {
	Health     Health     `json:"health"`
	MonStatus  MonStatus  `json:"mon_status"`
	MgrMap     MgrMap     `json:"mgr_map"`
	FsMap      FsMap      `json:"fs_map"`
	OsdMap     OsdMap     `json:"osd_map"`
	Rgw        int        `json:"rgw"`
	Hosts      int        `json:"hosts"`
	Pools      []Pool     `json:"pools"`
	PGInfo     PGInfo     `json:"pg_info"`
	DF         DF         `json:"df"`
	ClientPerf ClientPerf `json:"client_perf"`
}

func (hs *HealthStatus) ClusterHealth() float32 {
	switch hs.Health.Status {
	case "HEALTH_OK":
		return 100
	case "HEALTH_WARN":
		return 50
	case "HEALTH_ERR":
		fallthrough
	default:
		return 0
	}
}

func (hs *HealthStatus) ClusterHealthStatus() statsv1.ClusterHealth {
	clusterHealthStatus := statsv1.ClusterHealth_CLUSTER_HEALTH_OFFLINE
	switch hs.Health.Status {
	case "HEALTH_OK":
		clusterHealthStatus = statsv1.ClusterHealth_CLUSTER_HEALTH_OK
	case "HEALTH_WARN":
		clusterHealthStatus = statsv1.ClusterHealth_CLUSTER_HEALTH_WARN
	case "HEALTH_ERR":
		fallthrough
	default:
		clusterHealthStatus = statsv1.ClusterHealth_CLUSTER_HEALTH_ERR
	}

	return clusterHealthStatus
}

func (hs *HealthStatus) Crashes() []*statsv1.Crash {
	var crashes []*statsv1.Crash

	for _, check := range hs.Health.Checks {
		crashes = append(crashes, &statsv1.Crash{Description: check.Summary.Message})
	}
	return crashes
}

func (hs *HealthStatus) ObjectSize() int64 {
	size := 0
	for _, pool := range hs.DF.Pools {
		size += pool.DFPoolStats.Stored
	}

	return int64(size)
}

type BlockImage struct {
	BlockImageValue []BlockImageValue `json:"value"`
	PoolName        string            `json:"pool_name"`
}

type BlockImageValue struct {
	Size            int           `json:"size"`
	ObjSize         int           `json:"obj_size"`
	NumObjs         int           `json:"num_objs"`
	Order           int           `json:"order"`
	BlockNamePrefix string        `json:"block_name_prefix"`
	MirrorMode      string        `json:"mirror_mode"`
	Name            string        `json:"name"`
	UniqueId        string        `json:"unique_id"`
	Id              string        `json:"id"`
	ImageFormat     int           `json:"image_format"`
	PoolName        string        `json:"pool_name"`
	Namespace       string        `json:"namespace"`
	Features        int           `json:"features"`
	FeaturesName    []string      `json:"features_name"`
	Timestamp       time.Time     `json:"timestamp"`
	StripeCount     int           `json:"stripe_count"`
	StripeUnit      int           `json:"stripe_unit"`
	DataPool        interface{}   `json:"data_pool"`
	Parent          interface{}   `json:"parent"`
	Snapshots       []interface{} `json:"snapshots"`
	TotalDiskUsage  interface{}   `json:"total_disk_usage"`
	DiskUsage       interface{}   `json:"disk_usage"`
	Configuration   []interface{} `json:"configuration"`
}

type User struct {
	Username          string   `json:"username"`
	Name              string   `json:"name"`
	Email             string   `json:"email"`
	Roles             []string `json:"roles"`
	LastUpdate        int      `json:"lastUpdate"`
	Enabled           bool     `json:"enabled"`
	PwdExpirationDate float64  `json:"pwdExpirationDate"`
	PwdUpdateRequired bool     `json:"pwdUpdateRequired"`
}

type UserCreate struct {
	Username          string   `json:"username"`
	Name              string   `json:"name"`
	Email             string   `json:"email"`
	Password          string   `json:"password"`
	Roles             []string `json:"roles"`
	Enabled           bool     `json:"enabled"`
	PwdExpirationDate float64  `json:"pwdExpirationDate"`
	PwdUpdateRequired bool     `json:"pwdUpdateRequired"`
}

type ErrorCephApi struct {
	Detail    string `json:"detail"`
	Code      string `json:"code"`
	Component string `json:"component"`
}

func (e *ErrorCephApi) Error() error {
	return errors.New(e.Detail)
}
