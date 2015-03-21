// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sources

import (
	"time"

	"github.com/GoogleCloudPlatform/heapster/sources/api"
)

type Source interface {
	// Fetches containers or pod information from all the nodes in the cluster.
	// start, end: Represents the time range for stats
	// resolution: Represents the intervals at which samples are collected.
	// Returns:
	// AggregateData: A composite object that contains node, pod and container stats.
	GetInfo(start, end time.Time, resolution time.Duration) (api.AggregateData, error)
	// Returns debug information for the source.
	DebugInfo() string
}

func NewSource(pollDuration time.Duration) (Source, error) {
	if len(*argMaster) > 0 {
		return newKubeSource(pollDuration)
	} else if *argCoreOSMode {
		return newCoreOSCadvisorSource(pollDuration)
	} else {
		return newExternalCadvisorSource(pollDuration)
	}
}
