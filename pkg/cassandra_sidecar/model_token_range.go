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

type TokenRange struct {
	// starting token of a range 
	Start string `json:"start"`
	// ending token of a range 
	End string `json:"end"`
}
