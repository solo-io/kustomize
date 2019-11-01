// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package main_test

import (
	"testing"

	"sigs.k8s.io/kustomize/api/testutils/kusttest"
)

func TestSomeServiceGeneratorPlugin(t *testing.T) {
	tc := kusttest_test.NewPluginTestEnv(t).Set()
	defer tc.Reset()

	tc.BuildGoPlugin(
		"someteam.example.com", "v1", "SomeServiceGenerator")

	th := kusttest_test.NewKustTestHarnessAllowPlugins(t, "/app")

	m := th.LoadAndRunGenerator(`
apiVersion: someteam.example.com/v1
kind: SomeServiceGenerator
metadata:
  name: my-service
  namespace: test
port: "12345"
`)
	th.AssertActualEqualsExpected(m, `
apiVersion: v1
kind: Service
metadata:
  labels:
    app: dev
  name: my-service
  namespace: test
spec:
  ports:
  - port: 12345
  selector:
    app: dev
`)
}
