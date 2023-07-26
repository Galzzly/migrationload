/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type BandwidthPolicy struct {
	Id string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	LimitBytesPerSecond int64 `json:"limitBytesPerSecond,omitempty"`
}
