/*
Copyright 2023 The ImagineKube Authors.

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

package helmapplication

import (
	compbasemetrics "k8s.io/component-base/metrics"

	"imaginekube.com/imaginekube/pkg/utils/metrics"
)

var (
	appOperationTotal = compbasemetrics.NewCounterVec(
		&compbasemetrics.CounterOpts{
			Subsystem:      "ks_cm",
			Name:           "helm_application_operation_total",
			Help:           "Counter of app creation and deletion",
			StabilityLevel: compbasemetrics.ALPHA,
		},
		[]string{"verb", "name", "appstore"},
	)

	metricsList = []compbasemetrics.Registerable{
		appOperationTotal,
	}
)

func registerMetrics() {
	for _, m := range metricsList {
		metrics.MustRegister(m)
	}
}
