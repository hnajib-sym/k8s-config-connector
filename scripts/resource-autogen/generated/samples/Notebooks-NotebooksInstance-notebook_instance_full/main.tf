/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

```hcl
resource "google_notebooks_instance" "instance" {
  name = "notebooks-instance"
  location = "us-central1-a"
  machine_type = "e2-medium"

  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-cpu"
  }

  instance_owners = ["admin@hashicorptest.com"]
  service_account = "emailAddress:my@service-account.com"

  install_gpu_driver = true
  boot_disk_type = "PD_SSD"
  boot_disk_size_gb = 110

  no_public_ip = true
  no_proxy_access = true

  network = data.google_compute_network.my_network.id
  subnet = data.google_compute_subnetwork.my_subnetwork.id

  labels = {
    k = "val"
  }

  metadata = {
    terraform = "true"
  }
}

data "google_compute_network" "my_network" {
  name = "default"
}

data "google_compute_subnetwork" "my_subnetwork" {
  name   = "default"
  region = "us-central1"
}
```
