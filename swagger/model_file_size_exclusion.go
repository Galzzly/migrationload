/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type FileSizeExclusion struct {
	Id string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	FsRestriction bool `json:"fsRestriction,omitempty"`
	SourceFsExclusion bool `json:"sourceFsExclusion,omitempty"`
	GlobalExclusion bool `json:"globalExclusion,omitempty"`
	DefaultExclusion bool `json:"defaultExclusion,omitempty"`
	Type_ string `json:"type"`
	MaxBytes int64 `json:"maxBytes,omitempty"`
	DisplayUnit *FileSizeUnit `json:"displayUnit,omitempty"`
	MaxBytesFormatted string `json:"maxBytesFormatted,omitempty"`
}
