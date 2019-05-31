package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NSXError is a top-level type
type NSXError struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`


	// +optional

	Status NSXErrorStatus `json:"status,omitempty"`

	// This is where you can define
	// your own custom spec
	Spec NSXErrorSpec `json:"spec,omitempty"`
  }


  // custom spec 
  type NSXErrorSpec struct {
	  Message []string `json:"message,omitempty"`
	  ErrorObjectID string `json:"error-object-id,omitempty"`
	  ErrorObjectType string `json:"error-object-type,omitempty"`
	  ErrorObjectName string `json:"error-object-name,omitempty"`
	  ErrorObjectNamespace string `json:"error-object-ns,omitempty"`
  }

  // custom status
  type NSXErrorStatus struct {
	  Name string
  }

  // +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

  // no client needed for list as it's been created in above
  type NSXErrorList struct {
	  metav1.TypeMeta `json:",inline"`
	  // +optional
	  metav1.ListMeta `son:"metadata,omitempty"`
	  Items []NSXError `json:"items"`
  }
