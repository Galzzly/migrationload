/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"os"
)

type MultipartGcsWithJsonKeyBody struct {
	FsProperties *GcsFilesystemProperties `json:"fsProperties"`
	// JSON KeyFile
	File **os.File `json:"file,omitempty"`
}