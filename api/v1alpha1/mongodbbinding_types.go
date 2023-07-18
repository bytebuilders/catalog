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

// MongoDBBindingSpec defines the desired state of MongoDBBinding
type MongoDBBindingSpec struct {
	// SourceRef refers to the source app instance.
	SourceRef kmapi.ObjectReference `json:"sourceRef"`
}

//+kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:subresource:status

// MongoDBBinding is the Schema for the mongodbbindings API
type MongoDBBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MongoDBBindingSpec `json:"spec,omitempty"`
	Status BindingStatus      `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MongoDBBindingList contains a list of MongoDBBinding
type MongoDBBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongoDBBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MongoDBBinding{}, &MongoDBBindingList{})
}

var _ BindingInterface = &MongoDBBinding{}

func (in *MongoDBBinding) GetStatus() *BindingStatus {
	return &in.Status
}

func (in *MongoDBBinding) GetConditions() kmapi.Conditions {
	return in.Status.Conditions
}

func (in *MongoDBBinding) SetConditions(conditions kmapi.Conditions) {
	in.Status.Conditions = conditions
}
