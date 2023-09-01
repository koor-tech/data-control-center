package ceph

import "errors"

var (
	ErrorUnableToAuthenticate   = errors.New("unable to authenticate against the ceph mgr api")
	ErrorUnableToConnectWithApi = errors.New("unable to connect with the ceph mgr api")
)
