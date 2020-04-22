/**
 * Remove Dups: Write code to remove duplicates from an unsorted linked list.
 * FOLLOW UP: How would you solve this problems if a temporary buffer is not
 * allowed?
 */

#include <stdio.h>
#include "linked_list.h"
#include "../../lib/c/uthash/uthash.h"

void remove_dups_buffered(list *list_p);
void remove_dups_unbuffered(list *list_p);

int main() {
    list *test_list = create_list();
    insert(test_list, 1);
    insert(test_list, 1);
    insert(test_list, 1);
    insert(test_list, 4);
    insert(test_list, 5);
    insert(test_list, 7);
    insert(test_list, 7);
    insert(test_list, 2);
    insert(test_list, 3);
    insert(test_list, 3);
    insert(test_list, 3);

    remove_dups_unbuffered(test_list);
    print_list(test_list);
}

// O(n) time complexity and O(n) space complexity
void remove_dups_buffered(list *list_p) {
    node *node_table = NULL;
    node *found = NULL;

    node *current = list_p->head;
    // Head can't be a duplicate, so add it to the hash table
    HASH_ADD_INT(node_table, data, current);

    while (current->next != NULL) {
        HASH_FIND_INT(node_table, &(current->next->data), found);
        if (found == NULL) {
            HASH_ADD_INT(node_table, data, current->next);
            current = current->next;
        } else {
            node *dup = current->next;
            current->next = dup->next;
            free(dup);
        }
    }
}

// O(n^2) time complexity and O(1) space complexity
void remove_dups_unbuffered(list *list_p) {
    node *current = list_p->head;
    
    while (current != NULL) {
        node *runner = current;
        while (runner->next != NULL) {
            if (current->data == runner->next->data) {
                // Duplicate
                node *dup = runner->next;
                runner->next = dup->next;
                free(dup);
            } else {
                runner = runner->next;
            }
        }
        current = current->next;
    }
}