/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RegexExclusion struct {
	Id string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	FsRestriction bool `json:"fsRestriction,omitempty"`
	SourceFsExclusion bool `json:"sourceFsExclusion,omitempty"`
	GlobalExclusion bool `json:"globalExclusion,omitempty"`
	DefaultExclusion bool `json:"defaultExclusion,omitempty"`
	Type_ string `json:"type"`
	PatternType *RegexPatternType `json:"patternType,omitempty"`
	Regex string `json:"regex,omitempty"`
}
