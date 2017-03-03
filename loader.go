package cqlstress

import (
	"time"

	"github.com/gocql/gocql"
)

type source struct {
	Session *gocql.Session
	Schema  Schema
}

const (
	StaticDistribution      = 0
	NormalDistribution      = 1
	ExponentialDistribution = 2
	
	// Specific to DateTime
	LinearDistribution = 3
)

type DistributionType struct {
}

type StressConfig struct {
	CqlConfig gocql.ClusterConfig

	WarmUp   bool
	Duration time.Duration

	CasSchema   Schema
	ReadWeight  int
	WriteWeight int
}

func (sc StressConfig) NewSource(cfg gocql.ClusterConfig, warmup bool,
	duration time.Duration, schema Schema) (Source Source, err Error) {

}

type distribution struct {
	function func() interface{}
}

type Schema struct {
	Fields            []Field
	PrimaryKey        []string
	ResolvedTableName string
}

type Field struct {
	Name         string
	Type         interface{}
	Distribution DistributionType
}

func NewField(name string, type i

// cell is used to leverage custom unmarshalling in the gocql driver so that we don't unmarshall
// anything unnecessarily.
type cell struct {
	data []byte
	info info.gocql
}

func (sc Schema) NewRow() {
	for k, v := range FieldMappings {
		fv := sc.FieldDistributions[k]

	}
}
