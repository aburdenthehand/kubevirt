/*
Copyright The KubeVirt Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	scheme "kubevirt.io/client-go/prometheusoperator/scheme"
)

// ServiceMonitorsGetter has a method to return a ServiceMonitorInterface.
// A group's client should implement this interface.
type ServiceMonitorsGetter interface {
	ServiceMonitors(namespace string) ServiceMonitorInterface
}

// ServiceMonitorInterface has methods to work with ServiceMonitor resources.
type ServiceMonitorInterface interface {
	Create(ctx context.Context, serviceMonitor *v1.ServiceMonitor, opts metav1.CreateOptions) (*v1.ServiceMonitor, error)
	Update(ctx context.Context, serviceMonitor *v1.ServiceMonitor, opts metav1.UpdateOptions) (*v1.ServiceMonitor, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ServiceMonitor, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ServiceMonitorList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ServiceMonitor, err error)
	ServiceMonitorExpansion
}

// serviceMonitors implements ServiceMonitorInterface
type serviceMonitors struct {
	*gentype.ClientWithList[*v1.ServiceMonitor, *v1.ServiceMonitorList]
}

// newServiceMonitors returns a ServiceMonitors
func newServiceMonitors(c *MonitoringV1Client, namespace string) *serviceMonitors {
	return &serviceMonitors{
		gentype.NewClientWithList[*v1.ServiceMonitor, *v1.ServiceMonitorList](
			"servicemonitors",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.ServiceMonitor { return &v1.ServiceMonitor{} },
			func() *v1.ServiceMonitorList { return &v1.ServiceMonitorList{} }),
	}
}