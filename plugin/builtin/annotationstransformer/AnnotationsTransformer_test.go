// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package main_test

import (
	"testing"

	"sigs.k8s.io/kustomize/api/testutils/kusttest"
)

func TestAnnotationsTransformer(t *testing.T) {
	tc := kusttest_test.NewPluginTestEnv(t).Set()
	defer tc.Reset()

	tc.BuildGoPlugin(
		"builtin", "", "AnnotationsTransformer")

	th := kusttest_test.NewKustTestHarnessAllowPlugins(t, "/app")

	rm := th.LoadAndRunTransformer(`
apiVersion: builtin
kind: AnnotationsTransformer
metadata:
  name: notImportantHere
annotations:
  app: myApp
  greeting/morning: a string with blanks
fieldSpecs:
  - path: metadata/annotations
    create: true
`, `
apiVersion: v1
kind: Service
metadata:
  name: myService
spec:
  ports:
  - port: 7002
`)

	th.AssertActualEqualsExpected(rm, `
apiVersion: v1
kind: Service
metadata:
  annotations:
    app: myApp
    greeting/morning: a string with blanks
  name: myService
spec:
  ports:
  - port: 7002
`)
}
