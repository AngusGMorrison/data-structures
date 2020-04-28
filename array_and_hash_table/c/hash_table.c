/**
 * Separate chaining hash table implementation similar to the Ruby
 * implementation. Table sizes are primes near progressive powers of 2, as prime
 * table sizes have been found to result in more even distributions of entries.
 * Rehash to the next table size occurs when the average density of the table
 * exceeds 5 entries per bin.
 */

#include <stdio.h>
#include <stdlib.h>
#include "error.h"
#include "hash_table.h"

int main() {
    hash_table *new = new_hash_table();
}

hash_table *new_hash_table() {
    hash_table *new = malloc(sizeof(hash_table));
    if (new == NULL) {
        error("new_hash_table: couldn't allocate new", NULL_PTR);
    }

    new->buckets_idx = 0;
    new->buckets = BUCKETS[0];
    new->size = 0;
    new->table = malloc(new->buckets * sizeof(hash_node *));
    if (new->table == NULL) {
        error("new_hash_table: couldn't allocate table", NULL_PTR);
    }

    return new;
}