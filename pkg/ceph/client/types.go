package client

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Headers struct {
	Accept      string `json:"Accept"`
	ContentType string `json:"Content-Type"`
}

var headers = Headers{
	Accept:      "application/vnd.ceph.api.v1.0+json",
	ContentType: "application/json",
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
