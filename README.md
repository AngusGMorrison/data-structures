# Hash Tables

## Separate Chaining
Entries are stored in linked lists. Whenever a collision occurrs, the new entry is added to the list. On retrieval, if more than one entry exists at a location, the list is searched sequentially until the matching key is found.

In the best case, when entries are evenly distributed, separate chaining provides O(1) access. In the worst case, access takes O(n).

Requires additional memory to maintain linked list nodes as opposed to storing data directly in the table. Data locality is also poor, as nodes may be stored in very different locations in memory.

## Open Addressing
Entries are stored in the table itself. Whenever a collision occurs, the insertion algorithm continues searching for a free spot by probing the table at intervals.

### Linear Probing
The table is probed at fixed intervals (typically +1) until a free slot is found.

Requires less memory than separate chaining, but suffers from clustering, where consecutive elements form groups and additional time is required to find the next free slot.

Has the best cache locality of all addressing techniques.

### Quadratic Probing
Reduces clustering by probing at quadratic intervals: hash(x), hash(x) + 1 * 1, hash(x) + 2 * 2, ...

Lies between linear probing and double hashing in terms of cache locality and clustering.

### Double Hashing
Uses a second hashing alogrithm and searches for slot i * hash2(x) in the ith rotation if a free slot is not found:
- hash(x) + 1 * hash2(x) % s
- hash(x) + 2 * hash2(x) % s
- hash(x) + 3 * hash3(x) % s
- ...

Has poor cache performance but no clustering. Requires more computation time as two hash functions must be run.