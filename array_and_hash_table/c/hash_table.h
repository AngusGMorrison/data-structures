#define MAX_BUCKET_IDX 5 // Max index of pre-generated prime table size
#define MAX_DENSITY 5 // Max avg. entries per bucket

/* Bucket sizes are primes close to powers of two for more even distribution */
int BUCKETS[] = { 8 + 3, 16 + 3, 32 + 5, 64 + 3, 128 + 3, 256 + 27 };

enum value_types { INT, DBL, STR };

typedef struct hash_node {
    char *key;
    int value;
    struct hash_node *next;
} hash_node;

typedef struct hash_table {
    int buckets_idx; // The current index of the BUCKETS array
    int buckets; // The number of buckets the table currently has
    int size; // The number of entries currently in the table
    struct hash_node **table;
} hash_table;

hash_table *new_hash_table();
hash_table *hash_put(hash_table *hsh, char *key, int value);
int max_density_reached(hash_table *hsh);
void rehash(hash_table *hsh);
hash_node *rehash_node(hash_node *current, hash_table *hsh, hash_node **new_table);
hash_node *hash_get(hash_table *hsh, char *key);
int hash(char *key, int buckets);
hash_node *new_hash_node(char *key, int value);