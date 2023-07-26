/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DataTransferTiming struct {
	// Time to send and retrieve request/response to/from DTA on migrating file
	FileTransferNetworkOverheadNs int64 `json:"FileTransferNetworkOverheadNs,omitempty"`
	// Total time of a file migration on DTA side, the time measured from the start of grpc-method execution to transfer the file and to the time when FINISHED report will be sent to the lm2 side 
	TotalFileTransferTimeNs int64 `json:"TotalFileTransferTimeNs,omitempty"`
	// Time of reporting of PREPARING state of the file migration on data agent side in nanoseconds
	FileTransferPreparingReportTimeNs int64 `json:"FileTransferPreparingReportTimeNs,omitempty"`
	// Time to map source fs from a data transfer request to a real internal one in nanoseconds
	SrcFsMapTimeNs int64 `json:"SrcFsMapTimeNs,omitempty"`
	// Time to map target fs from a data transfer request to a real internal one in nanoseconds
	TgtFsMapTimeNs int64 `json:"TgtFsMapTimeNs,omitempty"`
	// Time to initialize a source filesystem in nanoseconds
	SrcFsInitTimeNs int64 `json:"SrcFsInitTimeNs,omitempty"`
	// Time to initialize a target filesystem in nanoseconds
	TgtFsInitTimeNs int64 `json:"TgtFsInitTimeNs,omitempty"`
	// Time to get file status info from the source filesystem in nanoseconds
	FileStatusRetrievingTimeNs int64 `json:"FileStatusRetrievingTimeNs,omitempty"`
	// Time to copy filer from the source to the target on data-agent side in nanoseconds
	FileTransferTimeNs int64 `json:"FileTransferTimeNs,omitempty"`
	// Time of opening of the source fs input stream to read the data that will be migrated
	SrcInStreamOpenTimeNs int64 `json:"SrcInStreamOpenTimeNs,omitempty"`
	// Time of opening of the target fs output stream where the data will be migrated to
	TgtOutStreamOpenTimeNs int64 `json:"TgtOutStreamOpenTimeNs,omitempty"`
	// Time of closing of the source fs input stream after migrating the data or on happening some exceptions on the file transferring 
	SrcInStreamCloseTimeNs int64 `json:"SrcInStreamCloseTimeNs,omitempty"`
	// Time of closing of the target fs output stream after migrating the data
	TgtOutStreamCloseTimeNs int64 `json:"TgtOutStreamCloseTimeNs,omitempty"`
	// Total time that was spent to read bytes from source fs input stream
	SrcInStreamBytesReadTimeNs int64 `json:"SrcInStreamBytesReadTimeNs,omitempty"`
	// Total time that was spent to write bytes to target fs output stream
	TgtOutStreamBytesWriteTimeNs int64 `json:"TgtOutStreamBytesWriteTimeNs,omitempty"`
}