// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/example-inc/memcached-operator/pkg/apis/cache/v1alpha1.Memcached":       schema_pkg_apis_cache_v1alpha1_Memcached(ref),
		"github.com/example-inc/memcached-operator/pkg/apis/cache/v1alpha1.MemcachedSpec":   schema_pkg_apis_cache_v1alpha1_MemcachedSpec(ref),
		"github.com/example-inc/memcached-operator/pkg/apis/cache/v1alpha1.MemcachedStatus": schema_pkg_apis_cache_v1alpha1_MemcachedStatus(ref),
	}
}

func schema_pkg_apis_cache_v1alpha1_Memcached(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Memcached is the Schema for the memcacheds API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/example-inc/memcached-operator/pkg/apis/cache/v1alpha1.MemcachedSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/example-inc/memcached-operator/pkg/apis/cache/v1alpha1.MemcachedStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/example-inc/memcached-operator/pkg/apis/cache/v1alpha1.MemcachedSpec", "github.com/example-inc/memcached-operator/pkg/apis/cache/v1alpha1.MemcachedStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_cache_v1alpha1_MemcachedSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MemcachedSpec defines the desired state of Memcached",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"size": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html Size is the size of the memcached deployment",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"size"},
			},
		},
	}
}

func schema_pkg_apis_cache_v1alpha1_MemcachedStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MemcachedStatus defines the observed state of Memcached",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nodes": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html Nodes are the names of the memcached pods",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"nodes"},
			},
		},
	}
}
