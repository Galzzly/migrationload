/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DataAgentStatsHistory struct {
	AgentName string `json:"agentName,omitempty"`
	DataAgentStatsRecordStore *DataAgentStatsRecordStore `json:"dataAgentStatsRecordStore,omitempty"`
	TotalAgentStats *DataAgentStats `json:"totalAgentStats,omitempty"`
}
