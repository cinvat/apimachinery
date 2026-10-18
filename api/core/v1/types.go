// Package v1 holds the shared API contract spoken across the cinvat kloud:
// the envelope every resource API serves, the metadata common to every
// object, and the discovery document that lets a client turn a manifest's
// kind into the collection path it lives at.
//
// It is a standalone module (github.com/cinvat/apimachinery) on purpose,
// mirroring k8s.io/apimachinery: the aggregation origin (kloudapi) and every
// downstream apiService (arsan, alvand, iam, …) import it, so nobody depends
// on anybody and the contract cannot drift. It stores no state and makes no
// decisions; it only names the shapes on the wire.
//
// General envelope and discovery types live here. Domain types — Node,
// VirtualMachine, Network, Port — never do: they belong to the service that
// owns the records. See docs/business-logics.md in kloudapi.
package v1

// TypeMeta identifies the API type of an object on the wire.
type TypeMeta struct {
	APIVersion string `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
	Kind       string `json:"kind,omitempty" yaml:"kind,omitempty"`
}

// ObjectMeta is the metadata common to every object: the human-readable
// address a client uses, never the server-side id (which lives in status).
type ObjectMeta struct {
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}

// APIResource mirrors one Kubernetes APIResource entry: enough for a client
// to turn a manifest's kind into the collection path it lives at.
type APIResource struct {
	Name         string   `json:"name" yaml:"name"`
	SingularName string   `json:"singularName,omitempty" yaml:"singularName,omitempty"`
	Namespaced   bool     `json:"namespaced" yaml:"namespaced"`
	Kind         string   `json:"kind" yaml:"kind"`
	Verbs        []string `json:"verbs,omitempty" yaml:"verbs,omitempty"`
}

// APIResourceList is the groupVersion discovery document an apiService serves
// at its groupVersion path (mount prefix joined with version).
type APIResourceList struct {
	GroupVersion string        `json:"groupVersion" yaml:"groupVersion"`
	Resources    []APIResource `json:"resources" yaml:"resources"`
}
