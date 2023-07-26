/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type FileStatus struct {
	Path string `json:"path,omitempty"`
	IsDirectory bool `json:"isDirectory,omitempty"`
	Length int64 `json:"length,omitempty"`
	ModificationTime int64 `json:"modificationTime,omitempty"`
}