# S3aFilesystemProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** | Access key | [optional] [default to null]
**SecretKey** | **string** | Secret key | [optional] [default to null]
**CredentialsProvider** | **string** | Credentials Provider | [optional] [default to null]
**AwsProfile** | **string** | aws profile | [optional] [default to null]
**AwsCredentialsConfigFile** | **string** | path to credentials file with profiles | [optional] [default to null]
**BucketName** | **string** | Bucket Name | [optional] [default to null]
**Endpoint** | **string** | s3 endpoint | [optional] [default to null]
**AdditionalProperties** | **map[string]string** | Additional filesystem properties | [optional] [default to null]
**SqsQueue** | **string** | SQS queue used when using s3 as a source | [optional] [default to null]
**S3type** | **string** | Indicates an s3a compatibility filesystem type | [optional] [default to null]
**BootstrapServers** | **string** | Kafka server, host:port | [optional] [default to null]
**Topic** | **string** | Kafka&#x27;s topic where s3 object change notifications are provided | [optional] [default to null]
**SqsEndpoint** | **string** | sqs endpoint | [optional] [default to null]
**CleanUpSqsQueue** | **bool** | ldm will cleanup sqs queue after when the filesystem is deleted | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

