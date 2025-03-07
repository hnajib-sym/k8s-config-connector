{# AUTOGENERATED. DO NOT EDIT. #}

{% extends "config-connector/_base.html" %}

{% block page_title %}ComputeAddress{% endblock %}
{% block body %}


<table>
<thead>
<tr>
<th><strong>Property</strong></th>
<th><strong>Value</strong></th>
</tr>
</thead>
<tbody>
<tr>
<td>{{gcp_name_short}} Service Name</td>
<td>Compute Engine</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/compute/docs/">/compute/docs/</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>
<pre>v1.addresses</pre>
<pre>v1.globalAddresses</pre>
</td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Documentation</td>
<td>
<pre><a href="/compute/docs/reference/rest/v1/globalAddresses">/compute/docs/reference/rest/v1/globalAddresses</a></pre>
<pre><a href="/compute/docs/reference/rest/v1/addresses">/compute/docs/reference/rest/v1/addresses</a></pre>
</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>gcpcomputeaddress<br>gcpcomputeaddresses<br>computeaddress</td>
</tr>
<tr>
<td>{{product_name_short}} Service Name</td>
<td>compute.googleapis.com</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>computeaddresses.compute.cnrm.cloud.google.com</td>
</tr>

<tr>
    <td>Can Be Referenced by IAMPolicy/IAMPolicyMember</td>
    <td>No</td>
</tr>


<tr>
<td>{{product_name_short}} Default Average Reconcile Interval In Seconds</td>
<td>600</td>
</tr>
</tbody>
</table>

ComputeAddress can manage both global and regional addresses. To manage a global ComputeAddress use a value of `global` in
the `spec.location` field. To manage a regional ComputeAddress, use a region name in the `spec.location` field.

## Custom Resource Definition Properties


### Annotations
<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>cnrm.cloud.google.com/project-id</code></td>
    </tr>
</tbody>
</table>


### Spec
#### Schema
```yaml
address: string
addressType: string
description: string
ipVersion: string
location: string
networkRef:
  external: string
  name: string
  namespace: string
networkTier: string
prefixLength: integer
purpose: string
resourceID: string
subnetworkRef:
  external: string
  name: string
  namespace: string
```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td>
            <p><code>address</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The static external IP address represented by this resource. Only
IPv4 is supported. An address may only be specified for INTERNAL
address types. The IP address must be inside the specified subnetwork,
if any. Set by the API if undefined.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>addressType</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The type of address to reserve. Default value: "EXTERNAL" Possible values: ["INTERNAL", "EXTERNAL"].{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>description</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. An optional description of this resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>ipVersion</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The IP Version that will be used by this address. The default value is 'IPV4'. Possible values: ["IPV4", "IPV6"]. This field can only be specified for a global address.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>location</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Location represents the geographical location of the ComputeAddress. Specify a region name or "global" for global resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/){% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>networkRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}The network in which to reserve the address. If global, the address
must be within the RFC1918 IP space. The network cannot be deleted
if there are any reserved IP ranges referring to it. This field can
only be used with INTERNAL type with the VPC_PEERING and
IPSEC_INTERCONNECT purposes.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>networkRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `selfLink` field of a `ComputeNetwork` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>networkRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>networkRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>networkTier</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The networking tier used for configuring this address. If this field is not
specified, it is assumed to be PREMIUM. Possible values: ["PREMIUM", "STANDARD"].{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>prefixLength</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Immutable. The prefix length if the resource represents an IP range.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>purpose</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The purpose of this resource, which can be one of the following values:

* GCE_ENDPOINT for addresses that are used by VM instances, alias IP
  ranges, internal load balancers, and similar resources.

* SHARED_LOADBALANCER_VIP for an address that can be used by multiple
  internal load balancers.

* VPC_PEERING for addresses that are reserved for VPC peer networks.

* IPSEC_INTERCONNECT for addresses created from a private IP range
  that are reserved for a VLAN attachment in an IPsec-encrypted Cloud
  Interconnect configuration. These addresses are regional resources.

* PRIVATE_SERVICE_CONNECT for a private network address that is used
to configure Private Service Connect. Only global internal addresses
can use this purpose.

This should only be set when using an Internal address.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>resourceID</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>subnetworkRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}The subnetwork in which to reserve the address. If an IP address is
specified, it must be within the subnetwork's IP range.  This field
can only be used with INTERNAL type with GCE_ENDPOINT/DNS_RESOLVER
purposes.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>subnetworkRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `selfLink` field of a `ComputeSubnetwork` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>subnetworkRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>subnetworkRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>



### Status
#### Schema
```yaml
conditions:
- lastTransitionTime: string
  message: string
  reason: string
  status: string
  type: string
creationTimestamp: string
labelFingerprint: string
observedGeneration: integer
selfLink: string
users:
- string
```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>conditions</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Conditions represent the latest available observation of the resource's current state.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].lastTransitionTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Last time the condition transitioned from one status to another.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].message</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Human-readable message indicating details about last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].reason</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Unique, one-word, CamelCase reason for the condition's last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].status</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Status is the status of the condition. Can be True, False, Unknown.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].type</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Type is the type of the condition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>creationTimestamp</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Creation timestamp in RFC3339 text format.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>labelFingerprint</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The fingerprint used for optimistic locking of this resource.  Used
internally during updates.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedGeneration</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>selfLink</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>users</code></td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}The URLs of the resources that are using this address.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>users[]</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### Global Compute Address
```yaml
# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeAddress
metadata:
  name: computeaddress-sample-global
  labels:
    label-one: "value-one"
spec:
  addressType: INTERNAL
  description: a test global address
  location: global
  ipVersion: IPV4
  purpose: VPC_PEERING
  prefixLength: 16
  networkRef:
    name: computeaddress-dep-global
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeNetwork
metadata:
  name: computeaddress-dep-global
spec:
  routingMode: REGIONAL
  autoCreateSubnetworks: false
```

### Regional Compute Address
```yaml
# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeAddress
metadata:
  name: computeaddress-sample-regional
  labels:
    label-one: "value-one"
spec:
  addressType: INTERNAL
  description: a test regional address
  location: us-central1
  ipVersion: IPV6
  subnetworkRef:
    name: computeaddress-dep-regional
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeNetwork
metadata:
  name: computeaddress-dep-regional
spec:
  routingMode: REGIONAL
  autoCreateSubnetworks: false
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeSubnetwork
metadata:
  name: computeaddress-dep-regional
spec:
  ipCidrRange: 10.2.0.0/16
  region: us-central1
  networkRef:
    name: computeaddress-dep-regional
```


{% endblock %}