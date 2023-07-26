/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MigrationState string

// List of MigrationState
const (
	NONSCHEDULED_MigrationState MigrationState = "NONSCHEDULED"
	SCHEDULED_MigrationState MigrationState = "SCHEDULED"
	RUNNING_MigrationState MigrationState = "RUNNING"
	LIVE_MigrationState MigrationState = "LIVE"
	PAUSED_MigrationState MigrationState = "PAUSED"
	STOPPED_MigrationState MigrationState = "STOPPED"
	COMPLETED_MigrationState MigrationState = "COMPLETED"
	RECURRENCE_SCHEDULED_MigrationState MigrationState = "RECURRENCE_SCHEDULED"
	WAITING_FOR_NEXT_RECURRENCE_MigrationState MigrationState = "WAITING_FOR_NEXT_RECURRENCE"
	STOPPING_MigrationState MigrationState = "STOPPING"
)