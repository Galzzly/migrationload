/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ThreadStates struct {
	NEW_THREADS int32 `json:"NEW_THREADS,omitempty"`
	RUNNABLE_THREADS int32 `json:"RUNNABLE_THREADS,omitempty"`
	BLOCKED_THREADS int32 `json:"BLOCKED_THREADS,omitempty"`
	WAITING_THREADS int32 `json:"WAITING_THREADS,omitempty"`
	TIMED_WAITING_THREADS int32 `json:"TIMED_WAITING_THREADS,omitempty"`
	TERMINATED_THREADS int32 `json:"TERMINATED_THREADS,omitempty"`
	TOTAL_THREADS int32 `json:"TOTAL_THREADS,omitempty"`
}
