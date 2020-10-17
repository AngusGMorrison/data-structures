package hashtable

import (
	"math"
)

const (
	loadFactorNum = 13
	loadFactorDen = 2
	nbuckets      = 8
)

type hmap struct {
	nitems  int
	buckets []*entry
}

type entry struct {
	key  string
	val  string
	next *entry
}

func New() *hmap {
	return &hmap{buckets: make([]*entry, nbuckets)}
}

func (h *hmap) Put(key string, val string) {
	if h.shouldRehash() {
		h.rehash()
	}

	idx := h.hashToIndex(key)
	item := h.buckets[idx]
	// Handle list head
	if item == nil {
		h.buckets[idx] = &entry{key: key, val: val}
		h.nitems++
		return
	}
	// Iterate through list, updating a matching entry or appending a new
	// entry to the tail.
	for ; ; item = item.next {
		if item.key == key {
			item.val = val
			return
		}

		if item.next == nil {
			item.next = &entry{key: key, val: val}
			h.nitems++
			return
		}
	}

}

func (h *hmap) shouldRehash() bool {
	return ((h.nitems + 1) / len(h.buckets)) >= (loadFactorNum / loadFactorDen)
}

func (h *hmap) rehash() {
	temp := &hmap{buckets: make([]*entry, len(h.buckets)<<1)}
	for _, bucket := range h.buckets {
		for cur := bucket; cur != nil; cur = cur.next {
			temp.Put(cur.key, cur.val)
		}
	}
	h.buckets = temp.buckets
}

func (h *hmap) hashToIndex(key string) uint {
	keyhash := hash(key)
	return uint(keyhash) % uint(len(h.buckets))
}

func hash(key string) uint64 {
	var hashTotal uint64
	keylen := float64(len(key))
	for i, r := range key {
		hashTotal += uint64(r) * uint64(math.Pow(31, keylen-(float64(i)+1)))
	}
	return hashTotal
}

func (h *hmap) Get(key string) (string, bool) {
	idx := h.hashToIndex(key)
	for cur := h.buckets[idx]; cur != nil; cur = cur.next {
		if cur.key == key {
			return cur.val, true
		}
	}
	return "", false
}

func (h *hmap) Delete(key string) {
	idx := h.hashToIndex(key)
	var prev *entry
	for cur := h.buckets[idx]; cur != nil; prev, cur = cur, cur.next {
		if cur.key == key {
			if prev == nil {
				// Head of list
				h.buckets[idx] = cur.next
			} else {
				prev.next = cur.next
			}

			h.nitems--
			return
		}

	}
}
