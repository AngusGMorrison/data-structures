/**
 * Delete Middle Node: Implement an algorithm to delete a node in the middle
 * (i.e., any node but the first and last node, not necessarily the exact
 * middle) of a singly linked list, given only access to that node.
 */

#include <stdio.h>
#include "linked_list.h"

void delete_middle_node(node *nd);

int main() {
    list *test_list = create_list();
    insert_multiple(test_list, 5, 1, 2, 3, 4, 5);
    node *middle = find_node(test_list, 3);
    delete_middle_node(middle);
    print_list(test_list);
}

void delete_middle_node(node *nd) {
    if (nd == NULL) {
        return;
    }

    node *next = nd->next;
    nd->data = next->data;
    nd->next = next->next;
    free(next);
} 