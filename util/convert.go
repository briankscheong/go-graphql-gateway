package util

import (
	"github.com/briankscheong/go-graphql-gateway/graph/model"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ConvertPod(p v1.Pod) *model.Pod {
	return &model.Pod{
		Metadata: ConvertMetadata(p.ObjectMeta),
		Status: &model.PodStatus{
			Phase:     strPtr(string(p.Status.Phase)),
			HostIP:    strPtr(p.Status.HostIP),
			PodIP:     strPtr(p.Status.PodIP),
			StartTime: strPtr(p.Status.StartTime.String()),
			Conditions: func() []*model.PodCondition {
				var pcs []*model.PodCondition
				for _, c := range p.Status.Conditions {
					pcs = append(pcs, &model.PodCondition{
						Type:               strPtr(string(c.Type)),
						Status:             strPtr(string(c.Status)),
						LastProbeTime:      strPtr(c.LastProbeTime.String()),
						LastTransitionTime: strPtr(c.LastTransitionTime.String()),
					})
				}
				return pcs
			}(),
			ContainerStatuses: func() []*model.ContainerStatus {
				var cs []*model.ContainerStatus
				for _, c := range p.Status.ContainerStatuses {
					cs = append(cs, &model.ContainerStatus{
						Name:         strPtr(c.Name),
						Ready:        boolPtr(c.Ready),
						RestartCount: intPtr(c.RestartCount),
						Image:        strPtr(c.Image),
					})
				}
				return cs
			}(),
		},
		Spec: &model.PodSpec{
			NodeName:           strPtr(p.Spec.NodeName),
			ServiceAccountName: strPtr(p.Spec.ServiceAccountName),
			Containers: func() []*model.ContainerSpec {
				var containers []*model.ContainerSpec
				for _, c := range p.Spec.Containers {
					containers = append(containers, &model.ContainerSpec{
						Name:    strPtr(c.Name),
						Image:   strPtr(c.Image),
						Command: toStringPtrSlice(c.Command),
						Args:    toStringPtrSlice(c.Args),
					})
				}
				return containers
			}(),
		},
	}
}

func ConvertDeployment(d appsv1.Deployment) *model.Deployment {
	return &model.Deployment{
		Metadata: ConvertMetadata(d.ObjectMeta),
		Spec: &model.DeploymentSpec{
			Replicas: d.Spec.Replicas,
			Selector: &model.LabelSelector{
				MatchLabels: mapToJSONStringPtr(d.Spec.Selector.MatchLabels),
			},
			Template: &model.PodTemplateSpec{
				Metadata: ConvertMetadata(d.Spec.Template.ObjectMeta),
				Spec:     ConvertPodSpec(d.Spec.Template.Spec),
			},
		},
		Status: &model.DeploymentStatus{
			AvailableReplicas: intPtr(d.Status.AvailableReplicas),
			ReadyReplicas:     intPtr(d.Status.ReadyReplicas),
			UpdatedReplicas:   intPtr(d.Status.UpdatedReplicas),
		},
	}
}

func ConvertService(s v1.Service) *model.Service {
	var ports []*model.ServicePort
	for _, p := range s.Spec.Ports {
		ports = append(ports, &model.ServicePort{
			Name:       strPtr(p.Name),
			Protocol:   strPtr(string(p.Protocol)),
			Port:       intPtr(p.Port),
			TargetPort: strPtr(p.TargetPort.String()),
		})
	}

	var ingress []*model.LoadBalancerIngress
	if s.Status.LoadBalancer.Ingress != nil {
		for _, i := range s.Status.LoadBalancer.Ingress {
			ingress = append(ingress, &model.LoadBalancerIngress{
				IP:       strPtr(i.IP),
				Hostname: strPtr(i.Hostname),
			})
		}
	}

	return &model.Service{
		Metadata: ConvertMetadata(s.ObjectMeta),
		Spec: &model.ServiceSpec{
			Type:     strPtr(string(s.Spec.Type)),
			Selector: mapToJSONStringPtr(s.Spec.Selector),
			Ports:    ports,
		},
		Status: &model.ServiceStatus{
			LoadBalancer: &model.LoadBalancerStatus{
				Ingress: ingress,
			},
		},
	}
}

func ConvertMetadata(m metav1.ObjectMeta) *model.Metadata {
	return &model.Metadata{
		Name:              m.Name,
		Namespace:         strPtr(m.Namespace),
		Labels:            mapToJSONStringPtr(m.Labels),
		Annotations:       mapToJSONStringPtr(m.Annotations),
		CreationTimestamp: strPtr(m.CreationTimestamp.String()),
		UID:               strPtr(string(m.UID)),
	}
}

func ConvertPodSpec(spec v1.PodSpec) *model.PodSpec {
	var containers []*model.ContainerSpec
	for _, c := range spec.Containers {
		containers = append(containers, &model.ContainerSpec{
			Name:    strPtr(c.Name),
			Image:   strPtr(c.Image),
			Command: toStringPtrSlice(c.Command),
			Args:    toStringPtrSlice(c.Args),
		})
	}

	return &model.PodSpec{
		NodeName:           strPtr(spec.NodeName),
		ServiceAccountName: strPtr(spec.ServiceAccountName),
		Containers:         containers,
	}
}
