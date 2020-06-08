/*
 * Instaclustr Cassandra Sidecar
 *
 * REST API for Cassandra Sidecar from Instaclustr
 *
 * API version: 2.0.0
 * Contact: support@instaclustr.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cassandra_sidecar

// bandwidth used during uploads 
type DataRate struct {
	// quantified value of bandwidth, an integer 
	Value int32 `json:"value,omitempty"`
	// unit of 'data bandwidth' 
	Unit string `json:"unit"`
}