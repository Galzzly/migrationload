/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DataAgentStatsSummary struct {
	LastWeekTotalBytes int64 `json:"lastWeekTotalBytes,omitempty"`
	LastDayTotalBytes int64 `json:"lastDayTotalBytes,omitempty"`
	AgentName string `json:"agentName,omitempty"`
}
