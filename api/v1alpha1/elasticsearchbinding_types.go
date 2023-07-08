/*
Copyright 2023.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
)

// ElasticsearchBindingSpec defines the desired state of ElasticsearchBinding
type ElasticsearchBindingSpec struct {
	// SourceRef refers to the source app instance.
	SourceRef kmapi.ObjectReference `json:"sourceRef"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ElasticsearchBinding is the Schema for the elasticsearchbindings API
type ElasticsearchBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElasticsearchBindingSpec `json:"spec,omitempty"`
	Status AppStatus                `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ElasticsearchBindingList contains a list of ElasticsearchBinding
type ElasticsearchBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticsearchBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ElasticsearchBinding{}, &ElasticsearchBindingList{})
}
