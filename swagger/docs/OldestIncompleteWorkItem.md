# OldestIncompleteWorkItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MigrationId** | **string** | Migration id | [optional] [default to null]
**ItemType** | **string** | Type of item | [optional] [default to null]
**Operation** | [***ActionType**](ActionType.md) |  | [optional] [default to null]
**Path** | **string** | Path of item as provided by the source | [optional] [default to null]
**NewPath** | **string** | Destination path (if rename) | [optional] [default to null]
**SourceEventCreationTimestamp** | **int64** | Timestamp of item provided by source filesystem | [optional] [default to null]
**SourceEventCreationDate** | [**time.Time**](time.Time.md) | Human readable date of item provided by source filesystem | [optional] [default to null]
**LdmActionCreationTimeStamp** | **int64** | Timestamp of action being created in LDM | [optional] [default to null]
**LdmActionCreationDate** | [**time.Time**](time.Time.md) | Human readable date of action being created in LDM | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

