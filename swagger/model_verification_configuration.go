/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type VerificationConfiguration struct {
	MigrationId string `json:"migrationId,omitempty"`
	IgnoreAfterTimestamp int64 `json:"ignoreAfterTimestamp,omitempty"`
	OriginalPaths []string `json:"originalPaths,omitempty"`
	VerificationDepth int32 `json:"verificationDepth,omitempty"`
}
