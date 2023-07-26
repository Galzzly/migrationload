/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ScannerContentSummary struct {
	ProgressSummary *ScannerProgressMatrix `json:"progressSummary,omitempty"`
	ContentSummary *ScannerContentMatrix `json:"contentSummary,omitempty"`
}