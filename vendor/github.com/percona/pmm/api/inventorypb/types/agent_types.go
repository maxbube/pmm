package types

import "fmt"

// this list should be in sync with inventorypb/agents.pb.go
const (
	AgentTypePMMAgent                       = "PMM_AGENT"
	AgentTypeNodeExporter                   = "NODE_EXPORTER"
	AgentTypeMySQLdExporter                 = "MYSQLD_EXPORTER"
	AgentTypeMongoDBExporter                = "MONGODB_EXPORTER"
	AgentTypePostgresExporter               = "POSTGRES_EXPORTER"
	AgentTypeProxySQLExporter               = "PROXYSQL_EXPORTER"
	AgentTypeQANMySQLPerfSchemaAgent        = "QAN_MYSQL_PERFSCHEMA_AGENT"
	AgentTypeQANMySQLSlowlogAgent           = "QAN_MYSQL_SLOWLOG_AGENT"
	AgentTypeQANMongoDBProfilerAgent        = "QAN_MONGODB_PROFILER_AGENT"
	AgentTypeQANPostgreSQLPgStatementsAgent = "QAN_POSTGRESQL_PGSTATEMENTS_AGENT"
	AgentTypeRDSExporter                    = "RDS_EXPORTER"
	AgentTypeExternalExporter               = "EXTERNAL_EXPORTER"
)

var agentTypeNames = map[string]string{
	// no invalid
	AgentTypePMMAgent:                       "pmm_agent",
	AgentTypeNodeExporter:                   "node_exporter",
	AgentTypeMySQLdExporter:                 "mysqld_exporter",
	AgentTypeMongoDBExporter:                "mongodb_exporter",
	AgentTypePostgresExporter:               "postgres_exporter",
	AgentTypeProxySQLExporter:               "proxysql_exporter",
	AgentTypeQANMySQLPerfSchemaAgent:        "mysql_perfschema_agent",
	AgentTypeQANMySQLSlowlogAgent:           "mysql_slowlog_agent",
	AgentTypeQANMongoDBProfilerAgent:        "mongodb_profiler_agent",
	AgentTypeQANPostgreSQLPgStatementsAgent: "postgresql_pgstatements_agent",
	AgentTypeRDSExporter:                    "rds_exporter",
	AgentTypeExternalExporter:               "external-exporter",
}

// AgentTypeName returns human friendly agent type to be used in reports
func AgentTypeName(t string) string {
	res := agentTypeNames[t]
	if res == "" {
		panic(fmt.Sprintf("no nice string for Agent Type %s", t))
	}

	return res
}
