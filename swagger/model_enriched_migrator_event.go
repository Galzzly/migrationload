/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EnrichedMigratorEvent struct {
	EventType string `json:"eventType,omitempty"`
	SrcPath string `json:"srcPath,omitempty"`
	DestPath string `json:"destPath,omitempty"`
	FileSize int64 `json:"fileSize,omitempty"`
	MigrationId string `json:"migrationId,omitempty"`
	TargetFsType string `json:"targetFsType,omitempty"`
	Excluded bool `json:"excluded,omitempty"`
	ExcludedSource bool `json:"excludedSource,omitempty"`
	ExcludedTarget bool `json:"excludedTarget,omitempty"`
}