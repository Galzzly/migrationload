/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RecurringMigrationTracker struct {
	Iterations []RecurringMigrationIteration `json:"iterations,omitempty"`
	RecurringCount int32 `json:"recurringCount,omitempty"`
	ScansInLastTwentyFourHours int32 `json:"scansInLastTwentyFourHours,omitempty"`
	ScansInLastSevenDays int32 `json:"scansInLastSevenDays,omitempty"`
	AvgDurationInSeconds int32 `json:"avgDurationInSeconds,omitempty"`
}
