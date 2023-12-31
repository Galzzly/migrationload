/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MigrationStats struct {
	FilesMigratedScan int64 `json:"filesMigratedScan,omitempty"`
	DirsMigratedScan int64 `json:"dirsMigratedScan,omitempty"`
	RawBytesMigratedScan int64 `json:"rawBytesMigratedScan,omitempty"`
	SuccessfulBytesMigratedScan int64 `json:"successfulBytesMigratedScan,omitempty"`
	FilesMigratedClient int64 `json:"filesMigratedClient,omitempty"`
	DirsMigratedClient int64 `json:"dirsMigratedClient,omitempty"`
	RawBytesMigratedClient int64 `json:"rawBytesMigratedClient,omitempty"`
	SuccessfulBytesMigratedClient int64 `json:"successfulBytesMigratedClient,omitempty"`
	FilesMigratedClientScan int64 `json:"filesMigratedClientScan,omitempty"`
	DirsMigratedClientScan int64 `json:"dirsMigratedClientScan,omitempty"`
	RawBytesMigratedClientScan int64 `json:"rawBytesMigratedClientScan,omitempty"`
	SuccessfulBytesMigratedClientScan int64 `json:"successfulBytesMigratedClientScan,omitempty"`
	BytesExcluded int64 `json:"bytesExcluded,omitempty"`
	FilesExcluded int64 `json:"filesExcluded,omitempty"`
	DirsExcluded int64 `json:"dirsExcluded,omitempty"`
	EventsExcluded int64 `json:"eventsExcluded,omitempty"`
	FilesScanned int64 `json:"filesScanned,omitempty"`
	DirsScanned int64 `json:"dirsScanned,omitempty"`
	FilesRescanned int64 `json:"filesRescanned,omitempty"`
	DirsRescanned int64 `json:"dirsRescanned,omitempty"`
	FileCountBySize map[string]int64 `json:"fileCountBySize,omitempty"`
	SuccessfulBytesMigrated int64 `json:"successfulBytesMigrated,omitempty"`
	RawBytesMigrated int64 `json:"rawBytesMigrated,omitempty"`
	FileSystemAction int64 `json:"fileSystemAction,omitempty"`
	FileSystemActionsScan int64 `json:"fileSystemActionsScan,omitempty"`
	DirsMigrated int64 `json:"dirsMigrated,omitempty"`
	FilesMigrated int64 `json:"filesMigrated,omitempty"`
	IncomingTransportBytes int64 `json:"incomingTransportBytes,omitempty"`
	OutgoingTransportBytes int64 `json:"outgoingTransportBytes,omitempty"`
}
