package ceph

import (
	"log"

	"github.com/ceph/go-ceph/rados"
)

type Rados struct {
	conn *rados.Conn
}

func NewRadosConnection() *Rados {
	conn, _ := rados.NewConn()
	err := conn.ReadDefaultConfigFile()
	if err != nil {
		log.Println("error reading the configuration")
		log.Fatal(err)
	}
	err = conn.Connect()
	if err != nil {
		log.Println("error during the connection")
		log.Fatal(err)
	}
	return &Rados{
		conn: conn,
	}
}

func (r *Rados) Cluster() (*rados.ClusterStat, error) {
	stats, err := r.conn.GetClusterStats()
	if err != nil {
		return nil, err
	}
	return &stats, nil
}
