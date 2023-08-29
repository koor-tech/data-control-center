package ceph

import (
	"fmt"

	"github.com/ceph/go-ceph/rados"
)

type Rados struct {
	conn *rados.Conn
}

func NewRadosConnection() (*Rados, error) {
	conn, _ := rados.NewConn()
	err := conn.ReadDefaultConfigFile()
	if err != nil {
		return nil, fmt.Errorf("error reading the configuration: %w", err)
	}
	err = conn.Connect()
	if err != nil {
		return nil, fmt.Errorf("error during the connection")
	}
	return &Rados{
		conn: conn,
	}, nil
}

func (r *Rados) Cluster() (*rados.ClusterStat, error) {
	stats, err := r.conn.GetClusterStats()
	if err != nil {
		return nil, err
	}
	return &stats, nil
}
