/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InotifyDiagnostic struct {
	TimeStamp int64 `json:"timeStamp,omitempty"`
	Type_ string `json:"type"`
	AvgEventsBehind int64 `json:"avgEventsBehind,omitempty"`
	MaxEventsBehind int64 `json:"maxEventsBehind,omitempty"`
	TotalEventsRead int64 `json:"totalEventsRead,omitempty"`
	AvgEventsReadPerCall int64 `json:"avgEventsReadPerCall,omitempty"`
	MaxEventsReadPerCall int64 `json:"maxEventsReadPerCall,omitempty"`
	AvgRpcCallTime int64 `json:"avgRpcCallTime,omitempty"`
	MaxRpcCallTime int64 `json:"maxRpcCallTime,omitempty"`
	EventsBehind int64 `json:"eventsBehind,omitempty"`
	NameNodeRetained int64 `json:"nameNodeRetained,omitempty"`
	NameNodeMaxEventsPerRpc int64 `json:"nameNodeMaxEventsPerRpc,omitempty"`
	NameNodeCheckpointTxs int64 `json:"nameNodeCheckpointTxs,omitempty"`
}