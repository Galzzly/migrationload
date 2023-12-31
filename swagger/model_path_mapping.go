/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PathMapping struct {
	Id string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	SourceFileSystem string `json:"sourceFileSystem,omitempty"`
	SourcePath string `json:"sourcePath,omitempty"`
	TargetFileSystem string `json:"targetFileSystem,omitempty"`
	TargetPath string `json:"targetPath,omitempty"`
}
