/**
 * Separate chaining hash table implementation similar to the Ruby
 * implementation. Table sizes are primes near progressive powers of 2, as prime
 * table sizes have been found to result in more even distributions of entries.
 * Rehash to the next table size occurs when the average density of the table
 * exceeds 5 entries per bin.
 */

#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "error.h"
#include "hash_table.h"

hash_table *new_hash_table() {
    hash_table *new = malloc(sizeof(hash_table));
    if (new == NULL) {
        error("new_hash_table: couldn't allocate new", NULL_PTR);
    }

    new->buckets_idx = 0;
    new->buckets = BUCKETS[0];
    new->size = 0;
    new->table = calloc(new->buckets, sizeof(hash_node *));
    if (new->table == NULL) {
        error("new_hash_table: couldn't allocate table", NULL_PTR);
    }

    return new;
}

hash_table *hash_put(hash_table *hsh, char *key, int value) {
    if  (hsh == NULL || hsh->table == NULL) {
        return NULL;
    }

    if (max_density_reached(hsh)) {
        rehash(hsh);
    }

    hash_node *node;
    if ((node = hash_get(hsh, key)) == NULL) {
        node = new_hash_node(key, value);
        int key_hash = hash(key, hsh->buckets);
        node->next = hsh->table[key_hash];
        hsh->table[key_hash] = node;
        hsh->size++;
    } else {
        node->value = value;
    }

    return  hsh;
}

int max_density_reached(hash_table *hsh) {
    return (float) hsh->size / hsh->buckets >= MAX_DENSITY;
}

void rehash(hash_table *hsh) {
    if (hsh->buckets_idx == MAX_BUCKET_IDX) {
        error("rehash: max hash size reached", MAX_HASH);
    }

    hsh->buckets = BUCKETS[++hsh->buckets_idx];
    hash_node **new_table = calloc(hsh->buckets, sizeof(hash_node *));

    while (*hsh->table++) {
        hash_node *current = *hsh->table;
        while (current != NULL) {
            current = rehash_node(current, hsh, new_table);
        }
    }

    free(hsh->table);
    hsh->table = new_table;
}

hash_node *rehash_node(hash_node *current, hash_table *hsh, hash_node **new_table) {
    hash_node *next = current->next;
    int key_hash = hash(current->key, hsh->buckets);
    current->next = new_table[key_hash];
    new_table[key_hash] = current;
    return next;
}

hash_node *hash_get(hash_table *hsh, char *key) {
    int key_hash = hash(key, hsh->buckets);
    hash_node *node = hsh->table[key_hash];

    for ( ; node != NULL && strcmp(node->key, key) != 0; node = node->next)
        ;

    return node;
}

int hash(char *key, int buckets) {
    int len = strlen(key);
    long hash_total = 0;
    for (int i = 1; *key; i++) {
        hash_total += *key++ * pow(31, len - i);
    }
    return (int) hash_total % buckets;
}

hash_node *new_hash_node(char *key, int value) {
    if (key == NULL) {
        return NULL;
    }

    hash_node *new_node = malloc(sizeof(hash_node));
    if (new_node == NULL) {
        error("new_hash_node: couldn't allocate node", NULL_PTR);
    }

    new_node->key = strdup(key);
    if (new_node->key == NULL) {
        error("new_hash_node: couldn't duplicate key", NULL_PTR);
    }
    new_node->value = value;
    new_node->next = NULL;
    return new_node;
}