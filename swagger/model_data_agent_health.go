/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DataAgentHealth struct {
	LastStatusUpdateTime int64 `json:"lastStatusUpdateTime,omitempty"`
	LastHealthMessage string `json:"lastHealthMessage,omitempty"`
	Status *DataAgentStatus `json:"status,omitempty"`
}
