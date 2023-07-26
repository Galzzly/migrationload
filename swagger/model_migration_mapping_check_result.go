/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MigrationMappingCheckResult struct {
	Mappings []PathMapping `json:"mappings,omitempty"`
	ConflictingMigrationIds []string `json:"conflictingMigrationIds,omitempty"`
}
