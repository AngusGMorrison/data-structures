/**
 * Simple demonstration of hash table methods.
 */ 

#include <stdio.h>
#include "hash_table.h"

int main() {
    hash_table *test_table = new_hash_table();

    char buf[3];
    for (int i = 0; i < 56; i++) { // 56th item triggers rehash
        sprintf(buf, "%i", i); // Convert integer to string
        hash_put(test_table, buf, i);
    }

    hash_node *sample_1 = hash_get(test_table, "55");
    hash_node *sample_2 = hash_get(test_table, "14");
    printf("Found new entry %i and rehashed entry %i\n", sample_1->value, sample_2->value);

    hash_put(test_table, "14", 13);
    hash_node *sample_3 = hash_get(test_table, "14");
    printf("hash_put overwrites existing key 14 with value: %i\n", sample_3->value);
}