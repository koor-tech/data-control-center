package ceph

type Health struct {
	Status string        `json:"status"`
	Checks []interface{} `json:"checks"`
	Mutes  []interface{} `json:"mutes"`
}

type Mons []struct {
	Name string `json:"name"`
}

type Monmap struct {
	Fsid string `json:"fsid"` //change to uuid
	Mons Mons   `json:"mons"`
}

type MonStatus struct {
	Name   string `json:"name"`
	Monmap Monmap `json:"monmap"`
}

type Standbys struct {
	Name string `json:"name"`
}

type MgrMap struct {
	ActiveName string     `json:"active_name"`
	Standbys   []Standbys `json:"standbys"`
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
}

type HealthStatus struct {
	Health    Health    `json:"health"`
	MonStatus MonStatus `json:"mon_status"`
	MgrMap    MgrMap    `json:"mgr_map"`
	FsMap     FsMap     `json:"fs_map"`
	OsdMap    OsdMap    `json:"osd_map"`
}
