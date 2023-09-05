package ceph

import (
	"encoding/json"
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
	Pools      []Pool     `json:"pools"`
	PGInfo     PGInfo     `json:"pg_info"`
	DF         DF         `json:"df"`
	ClientPerf ClientPerf `json:"client_perf"`
}
