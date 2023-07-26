/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MigrationsSummary struct {
	OverallCount int32 `json:"overallCount,omitempty"`
	Running *MigrationInfoList `json:"running,omitempty"`
	Ready *MigrationInfoList `json:"ready,omitempty"`
	Stopped *MigrationInfoList `json:"stopped,omitempty"`
	Live *MigrationInfoList `json:"live,omitempty"`
	MigrationStats []MigrationRecord `json:"migrationStats,omitempty"`
	Completed *MigrationInfoList `json:"completed,omitempty"`
	WaitingForNextRecurrence *MigrationInfoList `json:"waitingForNextRecurrence,omitempty"`
	RecurrenceScheduled *MigrationInfoList `json:"recurrenceScheduled,omitempty"`
}