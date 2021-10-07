// Code generated from [TypeName]HashSet_test.tt by T4 template. DO NOT EDIT.

//---------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool. Changes to this
//     file may cause incorrect behavior and will be lost
//     if the code is regenerated.
//
//     Generated on 2021 October 07 20:52:09 UTC
// </auto-generated>
//---------------------------------------------------------

//******************************************************************************************************
//  HashSet_test.go - Gbtc
//
//  Copyright © 2021, Grid Protection Alliance.	 All Rights Reserved.
//
//  Licensed to the Grid Protection Alliance (GPA) under one or more contributor license agreements. See
//  the NOTICE file distributed with this work for additional information regarding copyright ownership.
//  The GPA licenses this file to you under the MIT License (MIT), the "License"; you may not use this
//  file except in compliance with the License. You may obtain a copy of the License at:
//
//      http://opensource.org/licenses/MIT
//
//  Unless agreed to in writing, the subject software distributed under the License is distributed on an
//  "AS-IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. Refer to the
//  License for the specific language governing permissions and limitations.
//
//  Code Modification History:
//  ----------------------------------------------------------------------------------------------------
//  09/10/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package guid

import (
	"testing"
)

func random() Guid {
	return New()
}

var empty = Guid{}

func TestNewHashSet(t *testing.T) {
	set := NewHashSet([]Guid{random(), random(), random(), random()})

	if len(set) != 4 {
		t.Fatalf("NewHashSet: len != 4")
	}
}

func TestAdd(t *testing.T) {
	set := NewHashSet([]Guid{})

	set.Add(random())
	set.Add(random())

	if set.Add(set.Keys()[0]) {
		t.Fatalf("Add: Inserted duplicated %s", set.Keys()[0])
	}

	if len(set) != 2 {
		t.Fatalf("Add: len != 2")
	}
}

func TestRemove(t *testing.T) {
	set := NewHashSet([]Guid{})

	set.Add(random())
	set.Add(random())

	item := random()
	set.Add(item)

	if !set.Remove(item) {
		t.Fatalf("Remove: Failed to remove %s", item)
	}

	if len(set) != 2 {
		t.Fatalf("Remove: len != 2")
	}
}

func TestRemoveWhere(t *testing.T) {
	set := NewHashSet([]Guid{})

	set.Add(random())
	set.Add(random())
	set.Add(empty)
	set.Add(empty) // Will not add duplicate

	if count := set.RemoveWhere(func(item Guid) bool { return item == empty }); count != 1 {
		t.Fatalf("RemoveWhere: Failed to remove")
	}

	if len(set) != 2 {
		t.Fatalf("RemoveWhere: len != 2")
	}
}

func TestIsEmpty(t *testing.T) {
	set := NewHashSet([]Guid{})

	if !set.IsEmpty() {
		t.Fatalf("IsEmpty: Set not empty")
	}
}

func TestClear(t *testing.T) {
	set := NewHashSet([]Guid{random(), random(), random()})

	set.Clear()

	if !set.IsEmpty() {
		t.Fatalf("Clear: Set not empty")
	}
}

func TestContains(t *testing.T) {
	set := NewHashSet([]Guid{})

	set.Add(random())
	set.Add(random())

	item := random()
	set.Add(item)

	if !set.Contains(item) {
		t.Fatalf("Contains: Failed to find %s", item)
	}
}

func TestKeys(t *testing.T) {
	set := NewHashSet([]Guid{})

	set.Add(random())
	set.Add(random())

	item := random()
	set.Add(item)

	keys := set.Keys()
	found := false

	for _, v := range keys {
		if v == item {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("Keys: Failed to find key %s", item)
	}
}

func TestExceptWithSet(t *testing.T) {
	set1 := NewHashSet([]Guid{random(), random(), random(), random()})
	set2 := NewHashSet(set1.Keys())

	set1.Add(random())
	set2.Add(random())

	set1.ExceptWithSet(set2)

	if len(set1) != 1 {
		t.Fatalf("ExceptWith: len != 1")
	}
}

func TestSymmetricExceptWithSet(t *testing.T) {
	set1 := NewHashSet([]Guid{random(), random(), random(), random()})
	set2 := NewHashSet(set1.Keys())

	set3 := HashSet{}
	set3.SymmetricExceptWithSet(set1)

	if !set1.SetEqualsSet(set3) {
		t.Fatalf("SymmetricExceptWith: sets not equal")
	}

	set1.Add(random())
	set2.Add(random())

	set1.SymmetricExceptWithSet(set2)

	if len(set1) != 2 {
		t.Fatalf("SymmetricExceptWith: len != 2")
	}
}

func TestIntersectWithSet(t *testing.T) {
	set1 := NewHashSet([]Guid{random(), random(), random(), random()})
	set2 := NewHashSet(set1.Keys())
	set3 := NewHashSet(set1.Keys())

	set1.Add(random())
	set2.Add(random())

	set1.IntersectWithSet(set2)

	if !set1.SetEqualsSet(set3) {
		t.Fatalf("IntersectWith: Sets not equal")
	}

	set4 := HashSet{}
	set4.IntersectWith(set1.Keys())

	if len(set4) != 0 {
		t.Fatalf("IntersectWith: empty set intersect caused change")
	}

	set1.IntersectWith(set4.Keys())

	if len(set1) != 0 {
		t.Fatalf("IntersectWith: insertsect with empty set should be zero")
	}
}

func TestUnionWithSet(t *testing.T) {
	set1 := NewHashSet([]Guid{random(), random(), random(), random()})
	set2 := NewHashSet(set1.Keys())
	set3 := NewHashSet(set1.Keys())

	item1 := random()
	set1.Add(item1)

	item2 := random()
	set2.Add(item2)

	set1.UnionWithSet(set2)

	if !set1.Contains(item1) {
		t.Fatalf("UnionWith: Missing  item1")
	}

	if !set1.Contains(item2) {
		t.Fatalf("UnionWith: Missing  item2")
	}

	set3.UnionWithSet(set1)

	if !set1.SetEqualsSet(set3) {
		t.Fatalf("UnionWith: Sets not equal")
	}
}

func TestSetEqualsSet(t *testing.T) {
	set1 := NewHashSet([]Guid{random(), random(), random(), random()})
	set2 := NewHashSet(set1.Keys())

	if !set1.SetEqualsSet(set2) {
		t.Fatalf("SetEquals: Sets not equal")
	}

	keys := set1.Keys()
	first := keys[0]

	if !set1.Remove(first) || !set2.Remove(first) {
		t.Fatalf("SetEquals: Failed to remove")
	}

	if !set1.SetEqualsSet(set2) {
		t.Fatalf("SetEquals: Sets not equal")
	}

	first = set2.Keys()[0]

	if !set2.Remove(first) {
		t.Fatalf("SetEquals: Failed to remove first key")
	}

	set2.Add(random())

	if len(set1) != len(set2) {
		t.Fatalf("SetEquals: Set lengths are unequal")
	}

	if set1.SetEquals(set2.Keys()) {
		t.Fatalf("SetEquals: Same length unequal sets are equal")
	}

	set2.Add(random())

	if set1.SetEquals(set2.Keys()) {
		t.Fatalf("SetEquals: Different length unequal sets are equal")
	}
}

func TestOverlapsSet(t *testing.T) {
	set1 := NewHashSet([]Guid{random(), random(), random(), random()})
	set2 := NewHashSet(set1.Keys())

	if !set1.OverlapsSet(set2) {
		t.Fatalf("Overlaps: Sets do not overlap")
	}

	if !set1.Remove(set1.Keys()[0]) {
		t.Fatalf("Overlaps: Failed to remove")
	}

	if !set1.OverlapsSet(set2) {
		t.Fatalf("Overlaps: Sets do not overlap")
	}

	set2.Add(random())

	if !set1.OverlapsSet(set2) {
		t.Fatalf("Overlaps: Sets do not overlap")
	}

	if !set2.OverlapsSet(set1) {
		t.Fatalf("Overlaps: Sets do not overlap")
	}

	set3 := HashSet{}

	if set3.Overlaps(set1.Keys()) {
		t.Fatalf("Overlaps: Sets overlap")
	}

	set3.Add(random())
	set3.Add(random())

	if set3.Overlaps(set1.Keys()) {
		t.Fatalf("Overlaps: Sets overlap")
	}
}

func TestIsSubsetOfSet(t *testing.T) {
	set1 := NewHashSet([]Guid{random(), random(), random(), random()})
	set2 := NewHashSet(set1.Keys())

	if !set1.IsSubsetOfSet(set2) {
		t.Fatalf("IsSubsetOf: Set is not subset")
	}

	if !set1.Remove(set1.Keys()[0]) {
		t.Fatalf("IsSubsetOf: Failed to remove")
	}

	if !set1.IsSubsetOfSet(set2) {
		t.Fatalf("IsSubsetOf: Set is not subset")
	}

	set2.Add(random())

	if !set1.IsSubsetOfSet(set2) {
		t.Fatalf("IsSubsetOf: Set is not subset")
	}

	if set2.IsSubsetOfSet(set1) {
		t.Fatalf("IsSubsetOf: Set is not expected to be subset")
	}

	set3 := HashSet{}

	if !set3.IsSubsetOf(set1.Keys()) {
		t.Fatalf("IsSubsetOf: Empty set is not a subset")
	}

	set3.Add(random())
	set3.Add(random())

	if set3.IsSubsetOf(set1.Keys()) {
		t.Fatalf("IsSubsetOf: Unequal set is a subset")
	}
}

func TestIsProperSubsetOfSet(t *testing.T) {
	set1 := NewHashSet([]Guid{random(), random(), random(), random()})
	set2 := NewHashSet(set1.Keys())

	if set1.IsProperSubsetOfSet(set2) {
		t.Fatalf("IsProperSubsetOf: Set is not expected to be proper subset")
	}

	if !set1.Remove(set1.Keys()[0]) {
		t.Fatalf("IsProperSubsetOf: Failed to remove")
	}

	if !set1.IsProperSubsetOfSet(set2) {
		t.Fatalf("IsProperSubsetOf: Set is not proper subset")
	}

	set2.Add(random())

	if !set1.IsProperSubsetOfSet(set2) {
		t.Fatalf("IsProperSubsetOf: Set is not proper subset")
	}

	if set2.IsProperSubsetOfSet(set1) {
		t.Fatalf("IsProperSubsetOf: Set is not expected to be proper subset")
	}

	set3 := HashSet{}

	if !set3.IsProperSubsetOf(set1.Keys()) {
		t.Fatalf("IsSubsetOf: Empty proper set is not a subset")
	}

	set3.Add(random())
	set3.Add(random())

	if set3.IsProperSubsetOf(set1.Keys()) {
		t.Fatalf("IsSubsetOf: Unequal set is a proper subset")
	}
}

func TestIsSupersetOfSet(t *testing.T) {
	set1 := NewHashSet([]Guid{random(), random(), random(), random()})
	set2 := NewHashSet(set1.Keys())

	if !set2.IsSupersetOfSet(set1) {
		t.Fatalf("IsSupersetOf: Set is not superset")
	}

	if !set1.Remove(set1.Keys()[0]) {
		t.Fatalf("IsSupersetOf: Failed to remove")
	}

	if !set2.IsSupersetOfSet(set1) {
		t.Fatalf("IsSupersetOf: Set is not superset")
	}

	set2.Add(random())

	if !set2.IsSupersetOfSet(set1) {
		t.Fatalf("IsSupersetOf: Set is not superset")
	}

	if set1.IsSupersetOfSet(set2) {
		t.Fatalf("IsSupersetOf: Set is not expected to be superset")
	}

	set3 := HashSet{}

	if !set1.IsSupersetOf(set3.Keys()) {
		t.Fatalf("IsSupersetOf: Empty set is not a superset")
	}

	set3.Add(random())
	set3.Add(random())

	if set1.IsSupersetOf(set3.Keys()) {
		t.Fatalf("IsSubsetOf: Unequal set is a superset")
	}
}

func TestIsProperSupersetOfSet(t *testing.T) {
	set1 := NewHashSet([]Guid{random(), random(), random(), random()})
	set2 := NewHashSet(set1.Keys())

	if set2.IsProperSupersetOfSet(set1) {
		t.Fatalf("IsProperSupersetOf: Set is not expected to be proper superset")
	}

	if !set1.Remove(set1.Keys()[0]) {
		t.Fatalf("IsProperSupersetOf: Failed to remove")
	}

	if !set2.IsProperSupersetOfSet(set1) {
		t.Fatalf("IsProperSupersetOf: Set is not proper superset")
	}

	set2.Add(random())

	if !set2.IsProperSupersetOfSet(set1) {
		t.Fatalf("IsProperSupersetOf: Set is not proper superset")
	}

	if set1.IsProperSupersetOfSet(set2) {
		t.Fatalf("IsProperSupersetOf: Set is not expected to be proper superset")
	}

	set3 := HashSet{}

	if !set1.IsProperSupersetOf(set3.Keys()) {
		t.Fatalf("IsSupersetOf: Empty set is not a proper superset")
	}

	if !set3.IsProperSupersetOf(set1.Keys()) {
		t.Fatalf("IsSupersetOf: Empty set is not a proper superset")
	}

	set3.Add(random())
	set3.Add(random())

	if set1.IsProperSupersetOf(set3.Keys()) {
		t.Fatalf("IsSubsetOf: Unequal set is a proper superset")
	}
}
