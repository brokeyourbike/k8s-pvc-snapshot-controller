/*
Copyright 2026.

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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BackupSpec defines the desired state of Backup
type BackupSpec struct {
	// SourcePVC is the name of the PersistentVolumeClaim you want to backup
	// +kubebuilder:validation:Required
	SourcePVC string `json:"sourcePVC"`

	// GCSBucket is the name of the Google Cloud Storage bucket
	// +kubebuilder:validation:Required
	GCSBucket string `json:"gcsBucket"`

	// GCPProject is the Project ID where the bucket resides
	// +kubebuilder:validation:Required
	GCPProject string `json:"gcpProject"`
}

// BackupStatus defines the observed state of Backup.
type BackupStatus struct {
	// Phase is a high-level summary (Pending, Running, Completed, Failed)
	Phase string `json:"phase,omitempty"`

	// JobName stores the name of the worker Job doing the actual upload
	JobName string `json:"jobName,omitempty"`

	// Conditions represent the granular state of the Backup resource.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Backup is the Schema for the backups API
type Backup struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of Backup
	// +required
	Spec BackupSpec `json:"spec"`

	// status defines the observed state of Backup
	// +optional
	Status BackupStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// BackupList contains a list of Backup
type BackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []Backup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Backup{}, &BackupList{})
}
