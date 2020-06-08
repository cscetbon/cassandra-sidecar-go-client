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

type TruncateOperationRequest struct {
	Type_ string `json:"type"`
	// keyspace to truncate
	Keyspace string `json:"keyspace"`
	// table to truncate
	Table string `json:"table"`
}