/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Metering struct {
	HistoryBuckets []MeterBucket `json:"historyBuckets,omitempty"`
	TotalCurrentTransactionCnt int64 `json:"totalCurrentTransactionCnt,omitempty"`
	TotalCurrentKbCnt int64 `json:"totalCurrentKbCnt,omitempty"`
	StaticKbCnt int64 `json:"staticKbCnt,omitempty"`
	LiveKbCnt int64 `json:"liveKbCnt,omitempty"`
}
