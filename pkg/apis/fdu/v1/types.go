package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FDU struct {
	meta_v1.TypeMeta `json:",inline"`

	meta_v1.ObjectMeta `json:"metadata,omitempty`

	Spec FDUSpec `json:"spec"`
}

type FDUSpec struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Version            string `json:"version"`
	FDUVersion         string `json:"fdu_version"`
	Hypervisor         string `json:"hypervisor"`
	HypervisorSpecific string `json:"hypervisor_specific,omitempty`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FDUList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`

	Items []FDU `json:"items"`
}
