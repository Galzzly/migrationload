/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RequeuingActionsStats struct {
	RequeueCount int64 `json:"requeueCount,omitempty"`
	RequeueCountByMigration map[string]int64 `json:"requeueCountByMigration,omitempty"`
}
