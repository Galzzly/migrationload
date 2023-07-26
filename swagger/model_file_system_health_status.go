/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type FileSystemHealthStatus struct {
	FsId string `json:"fsId,omitempty"`
	Healthy bool `json:"healthy,omitempty"`
	TimeStampEpochMillis int64 `json:"timeStampEpochMillis,omitempty"`
	Cause string `json:"cause,omitempty"`
	HealthReason string `json:"healthReason,omitempty"`
}
