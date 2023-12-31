/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ThreadDumpDiagnostic struct {
	TimeStamp int64 `json:"timeStamp,omitempty"`
	Type_ string `json:"type"`
	Threads []ThreadInfo `json:"threads,omitempty"`
	TimeToCollect int64 `json:"timeToCollect,omitempty"`
	ThreadStateSummary *ThreadStates `json:"threadStateSummary,omitempty"`
}
