package cqlstress

import "github.com/gocql/gocql"

type Source struct {
	Session *gocql.Session
	Schema  Schema
}

const (
	NormalDistribution      = 0
	ExponentialDistribution = 1
)

type DistributionType struct {
}

type distribution struct {
	function func() interface{}
}

func (s Source) NewSource(sess gocql.Session, schema Schema) {
}

func (s Source) Write() {

}

func (s Source) Read() {

}

type Schema struct {
	FieldMappings      map[string]interface{}
	FieldDistributions map[string]distribution
	PartitionKey       string
	PrimaryKey         []string
	ResolvedTableName  string
}

func (sc Schema) NewRow() {
	for k, v := range FieldMappings {
		fv := sc.FieldDistributions[k]

	}
}
