/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type WorkloadIdentityPoolObservation struct {

	// an identifier for the resource with format projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource name of the pool as
	// projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The state of the pool.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type WorkloadIdentityPoolParameters struct {

	// A description of the pool. Cannot exceed 256 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
	// existing tokens to access resources. If the pool is re-enabled, existing tokens grant
	// access again.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// A display name for the pool. Cannot exceed 32 characters.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// WorkloadIdentityPoolSpec defines the desired state of WorkloadIdentityPool
type WorkloadIdentityPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkloadIdentityPoolParameters `json:"forProvider"`
}

// WorkloadIdentityPoolStatus defines the observed state of WorkloadIdentityPool.
type WorkloadIdentityPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkloadIdentityPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WorkloadIdentityPool is the Schema for the WorkloadIdentityPools API. Represents a collection of external workload identities.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type WorkloadIdentityPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkloadIdentityPoolSpec   `json:"spec"`
	Status            WorkloadIdentityPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkloadIdentityPoolList contains a list of WorkloadIdentityPools
type WorkloadIdentityPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkloadIdentityPool `json:"items"`
}

// Repository type metadata.
var (
	WorkloadIdentityPool_Kind             = "WorkloadIdentityPool"
	WorkloadIdentityPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WorkloadIdentityPool_Kind}.String()
	WorkloadIdentityPool_KindAPIVersion   = WorkloadIdentityPool_Kind + "." + CRDGroupVersion.String()
	WorkloadIdentityPool_GroupVersionKind = CRDGroupVersion.WithKind(WorkloadIdentityPool_Kind)
)

func init() {
	SchemeBuilder.Register(&WorkloadIdentityPool{}, &WorkloadIdentityPoolList{})
}