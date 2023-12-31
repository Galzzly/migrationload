/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Pressure struct {
	Type_ string `json:"type,omitempty"`
	Sample string `json:"sample,omitempty"`
	TenSecond float64 `json:"tenSecond,omitempty"`
	ThirtySecond float64 `json:"thirtySecond,omitempty"`
	SixtySecond float64 `json:"sixtySecond,omitempty"`
	ThreeHundredSecond float64 `json:"threeHundredSecond,omitempty"`
}
