# DataTransferTiming

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileTransferNetworkOverheadNs** | **int64** | Time to send and retrieve request/response to/from DTA on migrating file | [optional] [default to null]
**TotalFileTransferTimeNs** | **int64** | Total time of a file migration on DTA side, the time measured from the start of grpc-method execution to transfer the file and to the time when FINISHED report will be sent to the lm2 side  | [optional] [default to null]
**FileTransferPreparingReportTimeNs** | **int64** | Time of reporting of PREPARING state of the file migration on data agent side in nanoseconds | [optional] [default to null]
**SrcFsMapTimeNs** | **int64** | Time to map source fs from a data transfer request to a real internal one in nanoseconds | [optional] [default to null]
**TgtFsMapTimeNs** | **int64** | Time to map target fs from a data transfer request to a real internal one in nanoseconds | [optional] [default to null]
**SrcFsInitTimeNs** | **int64** | Time to initialize a source filesystem in nanoseconds | [optional] [default to null]
**TgtFsInitTimeNs** | **int64** | Time to initialize a target filesystem in nanoseconds | [optional] [default to null]
**FileStatusRetrievingTimeNs** | **int64** | Time to get file status info from the source filesystem in nanoseconds | [optional] [default to null]
**FileTransferTimeNs** | **int64** | Time to copy filer from the source to the target on data-agent side in nanoseconds | [optional] [default to null]
**SrcInStreamOpenTimeNs** | **int64** | Time of opening of the source fs input stream to read the data that will be migrated | [optional] [default to null]
**TgtOutStreamOpenTimeNs** | **int64** | Time of opening of the target fs output stream where the data will be migrated to | [optional] [default to null]
**SrcInStreamCloseTimeNs** | **int64** | Time of closing of the source fs input stream after migrating the data or on happening some exceptions on the file transferring  | [optional] [default to null]
**TgtOutStreamCloseTimeNs** | **int64** | Time of closing of the target fs output stream after migrating the data | [optional] [default to null]
**SrcInStreamBytesReadTimeNs** | **int64** | Total time that was spent to read bytes from source fs input stream | [optional] [default to null]
**TgtOutStreamBytesWriteTimeNs** | **int64** | Total time that was spent to write bytes to target fs output stream | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

