/*
Copyright 2021 The Kubernetes Authors.

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

// Package config implements clusterctl config functionality.
package config

import _ "embed"

// ClusterctlAPIManifest contains the clusterctl manifests in raw bytes format.
//
//go:embed manifest/clusterctl-api.yaml
var ClusterctlAPIManifest []byte
