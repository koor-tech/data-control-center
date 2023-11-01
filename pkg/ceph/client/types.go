package client

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Permissions struct {
	CephFS            []string `json:"cephfs"`
	ConfigOpt         []string `json:"config-opt"`
	DashboardSettings []string `json:"dashboard-settings"`
	Grafana           []string `json:"grafana"`
	Hosts             []string `json:"hosts"`
	ISCSI             []string `json:"iscsi"`
	Log               []string `json:"log"`
	Manager           []string `json:"manager"`
	Monitor           []string `json:"monitor"`
	NFSGanesha        []string `json:"nfs-ganesha"`
	Osd               []string `json:"osd"`
	Pool              []string `json:"pool"`
	Prometheus        []string `json:"prometheus"`
	RBDImage          []string `json:"rbd-image"`
	RBDMirroring      []string `json:"rbd-mirroring"`
	RGW               []string `json:"rgw"`
	User              []string `json:"user"`
}

type ResponseData struct {
	Token             string      `json:"token"`
	Username          string      `json:"username"`
	Permissions       Permissions `json:"permissions"`
	PwdExpirationDate interface{} `json:"pwdExpirationDate"`
	SSO               bool        `json:"sso"`
	PwdUpdateRequired bool        `json:"pwdUpdateRequired"`
}

// CephApiVersion defines custom type to manage versions of the api
type CephApiVersion string

const (
	CephApiVersionV01 CephApiVersion = "vnd.ceph.api.v0.1+json"
	CephApiVersionV10 CephApiVersion = "vnd.ceph.api.v1.0+json"
	CephApiVersionV20 CephApiVersion = "vnd.ceph.api.v2.0+json"
)

// CephApiEndpoint define each endpoint available on ceph api
type CephApiEndpoint string

const (
	CephApiEndpointAuth       CephApiEndpoint = "/auth"
	CephApiEndpointHealthFull CephApiEndpoint = "/health/full"
	CephApiEndpointBlockImage CephApiEndpoint = "/block/image"
	CephApiEndpointUsers      CephApiEndpoint = "/user"
)
