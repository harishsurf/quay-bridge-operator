package logging

import (
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var Log = logf.Log.WithName("quay-openshift-registry-operator")
