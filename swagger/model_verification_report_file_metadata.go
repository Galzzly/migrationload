/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type VerificationReportFileMetadata struct {
	Descriptor *VerificationReportFileDescriptor `json:"descriptor,omitempty"`
	Size int64 `json:"size,omitempty"`
	Created int64 `json:"created,omitempty"`
	Checksum int64 `json:"checksum,omitempty"`
}