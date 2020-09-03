/*


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

// Metadata metadata for object
type Metadata struct {
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

// Object defines a single object to be created
type Object struct {
	Kind       string   `json:"kind"`
	APIVersion string   `json:"apiVersion"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Name       string   `json:"name"`
	Spec       string   `json:"spec"`
}

// ObjectTemplateSpec defines the desired state of ObjectTemplate
type ObjectTemplateSpec struct {
	Description string   `json:"description,omitempty"`
	Objects     []Object `json:"objects"`
}

// ObjectTemplateStatus defines the observed state of ObjectTemplate
type ObjectTemplateStatus struct {
	Status string `json:"status"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=objecttemplates,scope=Cluster
// +kubebuilder:printcolumn:name="status",type=string,JSONPath=`.status.status`
// +kubebuilder:subresource:status

// ObjectTemplate is the Schema for the objecttemplates API
type ObjectTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ObjectTemplateSpec   `json:"spec,omitempty"`
	Status ObjectTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectTemplateList contains a list of ObjectTemplate
type ObjectTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ObjectTemplate{}, &ObjectTemplateList{})
}
