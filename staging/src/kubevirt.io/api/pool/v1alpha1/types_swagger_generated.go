// Code generated by swagger-doc. DO NOT EDIT.

package v1alpha1

func (VirtualMachinePool) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachinePool resource contains a VirtualMachine configuration\nthat can be used to replicate multiple VirtualMachine resources.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true\n+genclient",
	}
}

func (VirtualMachineTemplateSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "+k8s:openapi-gen=true",
		"metadata": "+kubebuilder:pruning:PreserveUnknownFields\n+nullable",
		"spec":     "VirtualMachineSpec contains the VirtualMachine specification.",
	}
}

func (VirtualMachinePoolCondition) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                   "+k8s:openapi-gen=true",
		"lastProbeTime":      "+nullable",
		"lastTransitionTime": "+nullable",
	}
}

func (VirtualMachinePoolStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "+k8s:openapi-gen=true",
		"conditions":    "+listType=atomic",
		"labelSelector": "Canonical form of the label selector for HPA which consumes it through the scale subresource.",
	}
}

func (VirtualMachinePoolSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                       "+k8s:openapi-gen=true",
		"replicas":               "Number of desired pods. This is a pointer to distinguish between explicit\nzero and not specified. Defaults to 1.\n+optional",
		"selector":               "Label selector for pods. Existing Poolss whose pods are\nselected by this will be the ones affected by this deployment.",
		"virtualMachineTemplate": "Template describes the VM that will be created.",
		"paused":                 "Indicates that the pool is paused.\n+optional",
		"nameGeneration":         "Options for the name generation in a pool.\n+optional",
		"maxUnavailable":         "(Defaults to 100%) Integer or string pointer, that when set represents either a percentage or number of VMs in a pool that can be unavailable (ready condition false) at a time during automated update.\n+optional",
		"scaleInStrategy":        "ScaleInStrategy specifies how the VMPool controller manages scaling in VMs within a VMPool\n+optional",
	}
}

func (VirtualMachinePoolNameGeneration) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "+k8s:openapi-gen=true",
	}
}

func (VirtualMachinePoolList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachinePoolList is a list of VirtualMachinePool resources.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true",
	}
}

func (VirtualMachinePoolScaleInStrategy) SwaggerDoc() map[string]string {
	return map[string]string{
		"":          "VirtualMachinePoolScaleInStrategy specifies how the VMPool controller manages scaling in VMs within a VMPool\n+k8s:openapi-gen=true",
		"proactive": "Proactive scale-in by forcing VMs to shutdown during scale-in (Default)\n+optional",
	}
}

func (VirtualMachinePoolProactiveScaleInStrategy) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                "VirtualMachinePoolProactiveScaleInStrategy represents proactive scale-in strategy\n+k8s:openapi-gen=true",
		"selectionPolicy": "SelectionPolicy defines the priority in which VM instances are selected for proactive scale-in\nDefaults to \"Random\" base policy when no SelectionPolicy is configured\n+optional",
	}
}

func (VirtualMachinePoolSelectionPolicy) SwaggerDoc() map[string]string {
	return map[string]string{
		"":           "VirtualMachinePoolSelectionPolicy defines the priority in which VM instances are selected for scale-in\n+k8s:openapi-gen=true",
		"basePolicy": "BasePolicy is a catch-all policy [Random|DescendingOrder]\n+optional\n+kubebuilder:validation:Enum=Random;DescendingOrder",
	}
}
