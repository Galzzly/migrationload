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
	"time"
)

type FailedPath struct {
	Path string `json:"path,omitempty"`
	RenamePath string `json:"renamePath,omitempty"`
	FailureMessage string `json:"failureMessage,omitempty"`
	TimeStamp int64 `json:"timeStamp,omitempty"`
	Action *ActionType `json:"action,omitempty"`
	DateFailed time.Time `json:"dateFailed,omitempty"`
}
