// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package inventory_test

import (
	"testing"

	. "sigs.k8s.io/kustomize/api/inventory"
	"sigs.k8s.io/kustomize/api/resid"
)

func makeRefs() (Refs, Refs) {
	a := resid.FromString("G1_V1_K1|ns1|nm1")
	b := resid.FromString("G2_V2_K2|ns2|nm2")
	c := resid.FromString("G3_V3_K3|ns3|nm3")
	current := NewRefs()
	current[a] = []resid.ResId{b, c}
	current[b] = []resid.ResId{}
	current[c] = []resid.ResId{}
	newRefs := NewRefs()
	newRefs[a] = []resid.ResId{b}
	newRefs[b] = []resid.ResId{}
	return current, newRefs
}

func TestInventory(t *testing.T) {
	inventory := NewInventory()
	curref, _ := makeRefs()

	inventory.UpdateCurrent(curref)
	if len(inventory.Current) != 3 {
		t.Fatalf("not getting the correct inventory %v", inventory)
	}
	curref, newref := makeRefs()
	inventory.UpdateCurrent(curref)
	if len(inventory.Current) != 3 {
		t.Fatalf("not getting the corrent inventory %v", inventory)
	}
	if len(inventory.Previous) != 3 {
		t.Fatalf("not getting the corrent inventory %v", inventory)
	}

	items := inventory.Prune()
	if len(items) != 0 {
		t.Fatalf("not getting the corrent items %v", items)
	}
	if len(inventory.Previous) != 0 {
		t.Fatalf("not getting the corrent inventory %v", inventory)
	}

	inventory.UpdateCurrent(newref)
	items = inventory.Prune()
	if len(items) != 1 {
		t.Fatalf("not getting the corrent items %v", items)
	}
	if len(inventory.Previous) != 0 {
		t.Fatalf("not getting the corrent inventory %v", inventory.Previous)
	}
}
