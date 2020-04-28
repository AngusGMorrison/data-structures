int BUCKETS[] = { 8 + 3, 16 + 3, 32 + 5, 64 + 3, 128 + 3, 256 + 27 };

enum value_types { INT, DBL, STR };

typedef struct hash_node {
    char *key;
    union {
        int i_val;
        double d_val;
        char *s_val;
    } value;
    int value_type; // Keep track of what type of value is currently stored
    struct hash_node *next;
} hash_node;

typedef struct hash_table {
    int buckets_idx; // The current index of the BUCKETS array
    int buckets; // The number of buckets the table currently has
    int size; // The number of entries currently in the table
    struct hash_node *(*table)[];
} hash_table;

hash_table *new_hash_table();
hash_table *insert(char *key, )