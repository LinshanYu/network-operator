/*
Copyright 2022 linshanyu.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MutilCniNetworkSpec defines the desired state of MutilCniNetwork
type MutilCniNetworkSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of MutilCniNetwork. Edit mutilcninetwork_types.go to remove/update
	Foo string `json:"foo,omitempty"`
	//TODO
	//多个CNI对象，默认第一个cni作为主要网络 返回给k8s
	CNIs []CNI `json:"cnis"`
}

type CNI struct {
	
}

// MutilCniNetworkStatus defines the observed state of MutilCniNetwork
type MutilCniNetworkStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Msg string `json:"msg"`
	Status bool `json:"status"`
	Created bool `json:"created"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MutilCniNetwork is the Schema for the mutilcninetworks API
type MutilCniNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MutilCniNetworkSpec   `json:"spec,omitempty"`
	Status MutilCniNetworkStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MutilCniNetworkList contains a list of MutilCniNetwork
type MutilCniNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MutilCniNetwork `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MutilCniNetwork{}, &MutilCniNetworkList{})
}
