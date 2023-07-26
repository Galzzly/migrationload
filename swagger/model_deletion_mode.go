/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DeletionMode string

// List of DeletionMode
const (
	NO_DELETION_DeletionMode DeletionMode = "NO_DELETION"
	IMMEDIATE_DELETION_DeletionMode DeletionMode = "IMMEDIATE_DELETION"
	DELAYED_DELETION_DeletionMode DeletionMode = "DELAYED_DELETION"
)