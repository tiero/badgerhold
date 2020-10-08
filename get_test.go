// Copyright 2019 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package badgerhold_test

import (
	"testing"
	"time"

	"github.com/tiero/badgerhold"
)

func TestGet(t *testing.T) {
	testWrap(t, func(store *badgerhold.Store, t *testing.T) {
		key := "testKey"
		data := &ItemTest{
			Name:    "Test Name",
			Created: time.Now(),
		}
		err := store.Insert(key, data)
		if err != nil {
			t.Fatalf("Error creating data for get test: %s", err)
		}

		result := &ItemTest{}

		err = store.Get(key, result)
		if err != nil {
			t.Fatalf("Error getting data from badgerhold: %s", err)
		}

		if !data.equal(result) {
			t.Fatalf("Got %v wanted %v.", result, data)
		}
	})
}
