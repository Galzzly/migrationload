/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MigrationInfo struct {
	InternalId string `json:"internalId,omitempty"`
	PathRetryCount int64 `json:"pathRetryCount,omitempty"`
	Id string `json:"id,omitempty"`
	Path string `json:"path,omitempty"`
	Progress *MigrationProgress `json:"progress,omitempty"`
}
