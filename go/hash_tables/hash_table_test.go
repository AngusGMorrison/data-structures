package hashtable

import (
	"reflect"
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	h := New()
	manual := &hmap{buckets: make([]*entry, nbuckets), nitems: 0}
	if !reflect.DeepEqual(h, manual) {
		t.Errorf("New() did not return correct correct hashmap, got %+v", h)
	}
}

func newOneBucketHMap() *hmap {
	return &hmap{buckets: make([]*entry, 1)}
}

type testCase struct {
	key, want, got string
	ok             bool
}

func assertNItems(gotItems, wantItems int, t *testing.T) {
	if gotItems != wantItems {
		t.Errorf("got %d items, want %d", gotItems, wantItems)
	}
}

func assertGot(tc *testCase, t *testing.T) {
	if !tc.ok {
		t.Errorf("h.Get(%q) returned false, want true", tc.key)
	}
	if tc.got != tc.want {
		t.Errorf("got %q for key %q, want %q", tc.got, tc.key, tc.want)
	}
}

func TestPutGet(t *testing.T) {
	key := "key"
	val := "val"

	t.Run("inserts and retrieves from the head of a chain", func(t *testing.T) {
		h := newOneBucketHMap()
		h.Put(key, val)
		assertNItems(h.nitems, 1, t)
		got, ok := h.Get(key)
		tc := &testCase{key, val, got, ok}
		assertGot(tc, t)
	})

	t.Run("inserts and retrieves from the tail of a chain", func(t *testing.T) {
		h := newOneBucketHMap()
		h.Put("head", "don't want")
		h.Put(key, val)
		assertNItems(h.nitems, 2, t)
		got, ok := h.Get(key)
		tc := &testCase{key, val, got, ok}
		assertGot(tc, t)
	})

	t.Run("inserts and retrieves from the middle of a chain", func(t *testing.T) {
		h := newOneBucketHMap()
		h.Put("head", "don't want")
		h.Put(key, val)
		h.Put("tail", "don't want")
		assertNItems(h.nitems, 3, t)
		got, ok := h.Get(key)
		tc := &testCase{key, val, got, ok}
		assertGot(tc, t)
	})

	t.Run("correctly inserts and retrieves when there are multiple buckets", func(t *testing.T) {
		h := New()
		h.Put(key, val)
		assertNItems(h.nitems, 1, t)
		got, ok := h.Get(key)
		tc := &testCase{key, val, got, ok}
		assertGot(tc, t)
	})

	t.Run("rehashes when loadFactor is exceeded", func(t *testing.T) {
		h := newOneBucketHMap()
		loadFactor := loadFactorNum / loadFactorDen
		for i := 0; i < loadFactor+1; i++ {
			str := strconv.Itoa(i)
			h.Put(str, str)
		}

		wantNBuckets := 2
		if gotNBuckets := len(h.buckets); gotNBuckets != wantNBuckets {
			t.Errorf("wanted hmap to rehash with %d buckets; got %d bucket(s)",
				wantNBuckets, gotNBuckets)
		}
	})

	t.Run("returns (\"\", false) when entry is not found", func(t *testing.T) {
		h := New()
		got, ok := h.Get(key)
		if ok {
			t.Errorf("got ok == true, want false")
		}
		if got != "" {
			t.Errorf("got val %q, want \"\"", got)
		}
	})
}
