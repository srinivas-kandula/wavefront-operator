// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/wavefronthq/wavefront-operator/pkg/apis/wavefront/v1alpha1.WavefrontCollector":       schema_pkg_apis_wavefront_v1alpha1_WavefrontCollector(ref),
		"github.com/wavefronthq/wavefront-operator/pkg/apis/wavefront/v1alpha1.WavefrontCollectorSpec":   schema_pkg_apis_wavefront_v1alpha1_WavefrontCollectorSpec(ref),
		"github.com/wavefronthq/wavefront-operator/pkg/apis/wavefront/v1alpha1.WavefrontCollectorStatus": schema_pkg_apis_wavefront_v1alpha1_WavefrontCollectorStatus(ref),
		"github.com/wavefronthq/wavefront-operator/pkg/apis/wavefront/v1alpha1.WavefrontProxy":           schema_pkg_apis_wavefront_v1alpha1_WavefrontProxy(ref),
		"github.com/wavefronthq/wavefront-operator/pkg/apis/wavefront/v1alpha1.WavefrontProxySpec":       schema_pkg_apis_wavefront_v1alpha1_WavefrontProxySpec(ref),
		"github.com/wavefronthq/wavefront-operator/pkg/apis/wavefront/v1alpha1.WavefrontProxyStatus":     schema_pkg_apis_wavefront_v1alpha1_WavefrontProxyStatus(ref),
	}
}

func schema_pkg_apis_wavefront_v1alpha1_WavefrontCollector(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WavefrontCollector is the Schema for the wavefrontcollectors API",
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
							Ref: ref("github.com/wavefronthq/wavefront-operator/pkg/apis/wavefront/v1alpha1.WavefrontCollectorSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wavefronthq/wavefront-operator/pkg/apis/wavefront/v1alpha1.WavefrontCollectorStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/wavefronthq/wavefront-operator/pkg/apis/wavefront/v1alpha1.WavefrontCollectorSpec", "github.com/wavefronthq/wavefront-operator/pkg/apis/wavefront/v1alpha1.WavefrontCollectorStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_wavefront_v1alpha1_WavefrontCollectorSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WavefrontCollectorSpec defines the desired state of WavefrontCollector",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Defaults to wavefronthq/wavefront-kubernetes-collector:latest",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"daemon": {
						SchemaProps: spec.SchemaProps{
							Description: "Whether to deploy the collector as a daemonset. False will roll out as a deployment.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"enableDebug": {
						SchemaProps: spec.SchemaProps{
							Description: "Whether to enable debug logging and profiling",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"env": {
						SchemaProps: spec.SchemaProps{
							Description: "List of environment variables to set for the Collector containers.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.EnvVar"),
									},
								},
							},
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "Compute resources required by the Collector containers.",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"tolerations": {
						SchemaProps: spec.SchemaProps{
							Description: "Tolerations for the collector pods",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Toleration"),
									},
								},
							},
						},
					},
					"configName": {
						SchemaProps: spec.SchemaProps{
							Description: "The name of the config map providing the configuration for the collector instance. If empty, a default name of \"collectorName-config\" is assumed.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"enableAutoUpgrade": {
						SchemaProps: spec.SchemaProps{
							Description: "If set to true, Collector pods will be upgraded automatically in case new minor upgrade version is available. For pinning Collector to a specific version, you will need to set this option to false. We support only minor version Auto Upgrades.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"openshift": {
						SchemaProps: spec.SchemaProps{
							Description: "Set to true when running collector in Openshift platform.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"useOpenshiftDefaultConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "If set to true, Collector will use default config bundled in the image else it will use the config from ConfigName",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.EnvVar", "k8s.io/api/core/v1.ResourceRequirements", "k8s.io/api/core/v1.Toleration"},
	}
}

func schema_pkg_apis_wavefront_v1alpha1_WavefrontCollectorStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WavefrontCollectorStatus defines the observed state of WavefrontCollector",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"updatedTimestamp": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_wavefront_v1alpha1_WavefrontProxy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WavefrontProxy is the Schema for the wavefrontproxies API",
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
							Ref: ref("github.com/wavefronthq/wavefront-operator/pkg/apis/wavefront/v1alpha1.WavefrontProxySpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wavefronthq/wavefront-operator/pkg/apis/wavefront/v1alpha1.WavefrontProxyStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/wavefronthq/wavefront-operator/pkg/apis/wavefront/v1alpha1.WavefrontProxySpec", "github.com/wavefronthq/wavefront-operator/pkg/apis/wavefront/v1alpha1.WavefrontProxyStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_wavefront_v1alpha1_WavefrontProxySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WavefrontProxySpec defines the desired state of WavefrontProxy",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "The WavefrontProxy image to use. Defaults to wavefronthq/proxy:latest",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"url": {
						SchemaProps: spec.SchemaProps{
							Description: "Wavefront URL (cluster).",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"token": {
						SchemaProps: spec.SchemaProps{
							Description: "Wavefront API Token.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"size": {
						SchemaProps: spec.SchemaProps{
							Description: "The no. of replicas for Wavefront Proxy. Defaults to 1",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"metricPort": {
						SchemaProps: spec.SchemaProps{
							Description: "The port number the proxy will listen on for metrics in Wavefront data format. This is usually port 2878 by default.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"tracePort": {
						SchemaProps: spec.SchemaProps{
							Description: "The port to listen on for Wavefront trace formatted data. Defaults to none. This is usually 30000",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"jaegerPort": {
						SchemaProps: spec.SchemaProps{
							Description: "The port to listen on for Jaeger Thrift formatted data. Defaults to none. This is usually 30001",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"zipkinPort": {
						SchemaProps: spec.SchemaProps{
							Description: "The port to listen on for Zipkin formatted data. Defaults to none. This is usually 9411",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"traceSamplingRate": {
						SchemaProps: spec.SchemaProps{
							Description: "Sampling rate to apply to tracing spans sent to the proxy. This rate is applied to all data formats the proxy is listening on. Value should be between 0.0 and 1.0.  Default is 1.0",
							Type:        []string{"number"},
							Format:      "double",
						},
					},
					"traceSamplingDuration": {
						SchemaProps: spec.SchemaProps{
							Description: "When this is set to a value greater than 0, spans that are greater than or equal to this value will be sampled.",
							Type:        []string{"number"},
							Format:      "double",
						},
					},
					"histogramDistPort": {
						SchemaProps: spec.SchemaProps{
							Description: "The port to listen on for Wavefront histogram distribution formatted data. This is usually 40000",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"preprocessor": {
						SchemaProps: spec.SchemaProps{
							Description: "The name of the config map providing the preprocessor rules for the Wavefront proxy.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"advanced": {
						SchemaProps: spec.SchemaProps{
							Description: "The name of the config map providing the advanced configurations for the Wavefront proxy.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"additionalPorts": {
						SchemaProps: spec.SchemaProps{
							Description: "The comma separated list of ports that need to be opened on Proxy Pod and Services. Needs to be explicitly specified when using \"Advanced\" configuration.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"enableAutoUpgrade": {
						SchemaProps: spec.SchemaProps{
							Description: "If set to true, Proxy pods will be upgraded automatically in case new minor upgrade version is available. For pinning Proxy to a specific version, you will need to set this option to false. We support only minor version Auto Upgrades.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"openshift": {
						SchemaProps: spec.SchemaProps{
							Description: "Set to true when running proxy in Openshift platform.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"storageClaimName": {
						SchemaProps: spec.SchemaProps{
							Description: "The name of the storage claim to be used for creating proxy buffers directory. This is applicable only in an Openshift environment.\"",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"url", "token"},
			},
		},
	}
}

func schema_pkg_apis_wavefront_v1alpha1_WavefrontProxyStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WavefrontProxyStatus defines the observed state of WavefrontProxy",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"updatedTimestamp": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}
