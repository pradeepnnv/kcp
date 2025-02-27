/*
Copyright 2022 The KCP Authors.

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

package plugin

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"

	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func TestExecute(t *testing.T) {
	streams, stdin, stdout, _ := genericclioptions.NewTestIOStreams()

	opts := NewOptions(streams)
	opts.Prefix = "testing"
	opts.Filename = "-"

	c := NewCRDSnapshot(opts)

	n, err := stdin.WriteString(endpointsYAML)
	require.NoError(t, err)
	require.Equal(t, len(endpointsYAML), n)

	err = c.Execute()
	require.NoError(t, err)

	require.Empty(t, cmp.Diff(expectedYAML, strings.Trim(stdout.String(), "\n")))
}

var endpointsYAML = `

---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.5.0
  creationTimestamp: null
  name: endpoints.core
spec:
  group: ""
  names:
    kind: Endpoints
    listKind: EndpointsList
    plural: endpoints
    singular: endpoints
  scope: Namespaced
  versions:
  - name: v1
    schema:
      openAPIV3Schema:
        description: 'Endpoints is a collection of endpoints that implement the actual service. Example:   Name: "mysvc",   Subsets: [     {       Addresses: [{"ip": "10.10.1.1"}, {"ip": "10.10.2.2"}],       Ports: [{"name": "a", "port": 8675}, {"name": "b", "port": 309}]     },     {       Addresses: [{"ip": "10.10.3.3"}],       Ports: [{"name": "a", "port": 93}, {"name": "b", "port": 76}]     },  ]'
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          subsets:
            description: The set of all endpoints is the union of all subsets. Addresses are placed into subsets according to the IPs they share. A single address with multiple ports, some of which are ready and some of which are not (because they come from different containers) will result in the address being displayed in different subsets for the different ports. No address will appear in both Addresses and NotReadyAddresses in the same subset. Sets of addresses and ports that comprise a service.
            items:
              description: 'EndpointSubset is a group of addresses with a common set of ports. The expanded set of endpoints is the Cartesian product of Addresses x Ports. For example, given:   {     Addresses: [{"ip": "10.10.1.1"}, {"ip": "10.10.2.2"}],     Ports:     [{"name": "a", "port": 8675}, {"name": "b", "port": 309}]   } The resulting set of endpoints can be viewed as:     a: [ 10.10.1.1:8675, 10.10.2.2:8675 ],     b: [ 10.10.1.1:309, 10.10.2.2:309 ]'
              properties:
                addresses:
                  description: IP addresses which offer the related ports that are marked as ready. These endpoints should be considered safe for load balancers and clients to utilize.
                  items:
                    description: EndpointAddress is a tuple that describes single IP address.
                    properties:
                      hostname:
                        description: The Hostname of this endpoint
                        type: string
                      ip:
                        description: 'The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24). IPv6 is also accepted but not fully supported on all platforms. Also, certain kubernetes components, like kube-proxy, are not IPv6 ready. TODO: This should allow hostname or IP, See #4447.'
                        type: string
                      nodeName:
                        description: 'Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.'
                        type: string
                      targetRef:
                        description: Reference to object providing the endpoint.
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          fieldPath:
                            description: 'If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: "spec.containers{name}" (where "name" refers to the name of the container that triggered the event) or if no container name is specified "spec.containers[2]" (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.'
                            type: string
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                          resourceVersion:
                            description: 'Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                            type: string
                          uid:
                            description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                            type: string
                        type: object
                    required:
                    - ip
                    type: object
                  type: array
                notReadyAddresses:
                  description: IP addresses which offer the related ports but are not currently marked as ready because they have not yet finished starting, have recently failed a readiness check, or have recently failed a liveness check.
                  items:
                    description: EndpointAddress is a tuple that describes single IP address.
                    properties:
                      hostname:
                        description: The Hostname of this endpoint
                        type: string
                      ip:
                        description: 'The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24). IPv6 is also accepted but not fully supported on all platforms. Also, certain kubernetes components, like kube-proxy, are not IPv6 ready. TODO: This should allow hostname or IP, See #4447.'
                        type: string
                      nodeName:
                        description: 'Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.'
                        type: string
                      targetRef:
                        description: Reference to object providing the endpoint.
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          fieldPath:
                            description: 'If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: "spec.containers{name}" (where "name" refers to the name of the container that triggered the event) or if no container name is specified "spec.containers[2]" (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.'
                            type: string
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                          resourceVersion:
                            description: 'Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                            type: string
                          uid:
                            description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                            type: string
                        type: object
                    required:
                    - ip
                    type: object
                  type: array
                ports:
                  description: Port numbers available on the related IP addresses.
                  items:
                    description: EndpointPort is a tuple that describes a single port.
                    properties:
                      appProtocol:
                        description: The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol. Field can be enabled with ServiceAppProtocol feature gate.
                        type: string
                      name:
                        description: The name of this port.  This must match the 'name' field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined.
                        type: string
                      port:
                        description: The port number of the endpoint.
                        format: int32
                        type: integer
                      protocol:
                        default: TCP
                        description: The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
                        type: string
                    required:
                    - port
                    type: object
                  type: array
              type: object
            type: array
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`

var expectedYAML = `apiVersion: apis.kcp.dev/v1alpha1
kind: APIResourceSchema
metadata:
  creationTimestamp: null
  name: testing.endpoints.core
spec:
  group: ""
  names:
    kind: Endpoints
    listKind: EndpointsList
    plural: endpoints
    singular: endpoints
  scope: Namespaced
  versions:
  - name: v1
    schema:
      description: 'Endpoints is a collection of endpoints that implement the actual
        service. Example:   Name: "mysvc",   Subsets: [     {       Addresses: [{"ip":
        "10.10.1.1"}, {"ip": "10.10.2.2"}],       Ports: [{"name": "a", "port": 8675},
        {"name": "b", "port": 309}]     },     {       Addresses: [{"ip": "10.10.3.3"}],       Ports:
        [{"name": "a", "port": 93}, {"name": "b", "port": 76}]     },  ]'
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        subsets:
          description: The set of all endpoints is the union of all subsets. Addresses
            are placed into subsets according to the IPs they share. A single address
            with multiple ports, some of which are ready and some of which are not
            (because they come from different containers) will result in the address
            being displayed in different subsets for the different ports. No address
            will appear in both Addresses and NotReadyAddresses in the same subset.
            Sets of addresses and ports that comprise a service.
          items:
            description: 'EndpointSubset is a group of addresses with a common set
              of ports. The expanded set of endpoints is the Cartesian product of
              Addresses x Ports. For example, given:   {     Addresses: [{"ip": "10.10.1.1"},
              {"ip": "10.10.2.2"}],     Ports:     [{"name": "a", "port": 8675}, {"name":
              "b", "port": 309}]   } The resulting set of endpoints can be viewed
              as:     a: [ 10.10.1.1:8675, 10.10.2.2:8675 ],     b: [ 10.10.1.1:309,
              10.10.2.2:309 ]'
            properties:
              addresses:
                description: IP addresses which offer the related ports that are marked
                  as ready. These endpoints should be considered safe for load balancers
                  and clients to utilize.
                items:
                  description: EndpointAddress is a tuple that describes single IP
                    address.
                  properties:
                    hostname:
                      description: The Hostname of this endpoint
                      type: string
                    ip:
                      description: 'The IP of this endpoint. May not be loopback (127.0.0.0/8),
                        link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24).
                        IPv6 is also accepted but not fully supported on all platforms.
                        Also, certain kubernetes components, like kube-proxy, are
                        not IPv6 ready. TODO: This should allow hostname or IP, See
                        #4447.'
                      type: string
                    nodeName:
                      description: 'Optional: Node hosting this endpoint. This can
                        be used to determine endpoints local to a node.'
                      type: string
                    targetRef:
                      description: Reference to object providing the endpoint.
                      properties:
                        apiVersion:
                          description: API version of the referent.
                          type: string
                        fieldPath:
                          description: 'If referring to a piece of an object instead
                            of an entire object, this string should contain a valid
                            JSON/Go field access statement, such as desiredState.manifest.containers[2].
                            For example, if the object reference is to a container
                            within a pod, this would take on a value like: "spec.containers{name}"
                            (where "name" refers to the name of the container that
                            triggered the event) or if no container name is specified
                            "spec.containers[2]" (container with index 2 in this pod).
                            This syntax is chosen only to have some well-defined way
                            of referencing a part of an object. TODO: this design
                            is not final and this field is subject to change in the
                            future.'
                          type: string
                        kind:
                          description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                          type: string
                        name:
                          description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                          type: string
                        namespace:
                          description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                          type: string
                        resourceVersion:
                          description: 'Specific resourceVersion to which this reference
                            is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                          type: string
                        uid:
                          description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                          type: string
                      type: object
                  required:
                  - ip
                  type: object
                type: array
              notReadyAddresses:
                description: IP addresses which offer the related ports but are not
                  currently marked as ready because they have not yet finished starting,
                  have recently failed a readiness check, or have recently failed
                  a liveness check.
                items:
                  description: EndpointAddress is a tuple that describes single IP
                    address.
                  properties:
                    hostname:
                      description: The Hostname of this endpoint
                      type: string
                    ip:
                      description: 'The IP of this endpoint. May not be loopback (127.0.0.0/8),
                        link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24).
                        IPv6 is also accepted but not fully supported on all platforms.
                        Also, certain kubernetes components, like kube-proxy, are
                        not IPv6 ready. TODO: This should allow hostname or IP, See
                        #4447.'
                      type: string
                    nodeName:
                      description: 'Optional: Node hosting this endpoint. This can
                        be used to determine endpoints local to a node.'
                      type: string
                    targetRef:
                      description: Reference to object providing the endpoint.
                      properties:
                        apiVersion:
                          description: API version of the referent.
                          type: string
                        fieldPath:
                          description: 'If referring to a piece of an object instead
                            of an entire object, this string should contain a valid
                            JSON/Go field access statement, such as desiredState.manifest.containers[2].
                            For example, if the object reference is to a container
                            within a pod, this would take on a value like: "spec.containers{name}"
                            (where "name" refers to the name of the container that
                            triggered the event) or if no container name is specified
                            "spec.containers[2]" (container with index 2 in this pod).
                            This syntax is chosen only to have some well-defined way
                            of referencing a part of an object. TODO: this design
                            is not final and this field is subject to change in the
                            future.'
                          type: string
                        kind:
                          description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                          type: string
                        name:
                          description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                          type: string
                        namespace:
                          description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                          type: string
                        resourceVersion:
                          description: 'Specific resourceVersion to which this reference
                            is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                          type: string
                        uid:
                          description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                          type: string
                      type: object
                  required:
                  - ip
                  type: object
                type: array
              ports:
                description: Port numbers available on the related IP addresses.
                items:
                  description: EndpointPort is a tuple that describes a single port.
                  properties:
                    appProtocol:
                      description: The application protocol for this port. This field
                        follows standard Kubernetes label syntax. Un-prefixed names
                        are reserved for IANA standard service names (as per RFC-6335
                        and http://www.iana.org/assignments/service-names). Non-standard
                        protocols should use prefixed names such as mycompany.com/my-custom-protocol.
                        Field can be enabled with ServiceAppProtocol feature gate.
                      type: string
                    name:
                      description: The name of this port.  This must match the 'name'
                        field in the corresponding ServicePort. Must be a DNS_LABEL.
                        Optional only if one port is defined.
                      type: string
                    port:
                      description: The port number of the endpoint.
                      format: int32
                      type: integer
                    protocol:
                      default: TCP
                      description: The IP protocol for this port. Must be UDP, TCP,
                        or SCTP. Default is TCP.
                      type: string
                  required:
                  - port
                  type: object
                type: array
            type: object
          type: array
      type: object
    served: true
    storage: true
    subresources: {}`
