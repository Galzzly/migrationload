/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ScannerContentMatrix struct {
	ByteCount int64 `json:"byteCount,omitempty"`
	FileCount int64 `json:"fileCount,omitempty"`
	DirectoryCount int64 `json:"directoryCount,omitempty"`
}