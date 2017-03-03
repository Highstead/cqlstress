package cqlstress_test

import (
	"testing"
	"time"

	"github.com/gocql/gocql"
)

func stressConfig() gocql.ClusterConfig {
	cfg := gocql.NewCluster("127.0.0.1")
	cfg.Port = 9042
	cfg.Keyspace = "test"
	cfg.ProtoVersion = 4
	cfg.Consistency = gocql.One
	cfg.Timeout, _ = time.ParseDuration("5s")
	cfg.Compressor = nil

	duration := time.Second * 5
	scfg := StressConfig{
		WarmUp:      false,
		Duration:    duration,
		cqlCfg:      cfg,
		ReadWeight:  0,
		WriteWeight: 1,
	}

	return scfg
}

func testTimeSchema() {

}

func TestSchemaWrite(t *testing.T) {

}
