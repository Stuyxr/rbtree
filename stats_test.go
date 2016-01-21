// Copyright 2015, Hu Keping. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rbtree

import (
	"testing"
)

func TestDeleteReturnValue(t *testing.T) {
	rbt := New()

	rbt.Insert(String("go"))
	rbt.Insert(String("lang"))

	if rbt.Len() != 2 {
		t.Errorf("tree.Len() = %d, expect %d", rbt.Len(), 2)
	}

	// go should be in the rbtree
	deletedGo := rbt.Delete(String("go"))
	if deletedGo != String("go") {
		t.Errorf("expect %v, got %v", "go", deletedGo)
	}

	// C should not be in the rbtree
	deletedC := rbt.Delete(String("C"))
	if deletedC != nil {
		t.Errorf("expect %v, got %v", nil, deletedC)
	}
}

func TestMin(t *testing.T) {
	rbt := New()

	m := 0
	n := 1000
	for m < n {
		rbt.Insert(Int(m))
		m++
	}
	if rbt.Len() != uint(n) {
		t.Errorf("tree.Len() = %d, expect %d", rbt.Len(), n)
	}

	for m >= 0 {
		rbt.Delete(rbt.Min())
		m--
	}

	if rbt.Len() != 0 {
		t.Errorf("tree.Len() = %d, expect %d", rbt.Len(), 0)
	}
}