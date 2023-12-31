/*
 * LiveData Migrator
 *
 * LiveData Migrator API Documentation
 *
 * API version: 3.3.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// ADLS2 file system properties, with service principal authentication
type Adls2ServicePrincipalFilesystemProperties struct {
	// Storage account name
	AccountName string `json:"accountName,omitempty"`
	// OAuth Client ID
	OauthClientId string `json:"oauthClientId,omitempty"`
	// OAuth Client Secret
	OauthClientSecret string `json:"oauthClientSecret,omitempty"`
	// OAuth Client Endpoint
	OauthClientEndpoint string `json:"oauthClientEndpoint,omitempty"`
	// Container name
	ContainerName string `json:"containerName,omitempty"`
	// Adls2 endpoint for current account
	Endpoint string `json:"endpoint,omitempty"`
	// Insecure connection using abfs scheme
	Insecure bool `json:"insecure,omitempty"`
	// Additional filesystem properties
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty"`
}
