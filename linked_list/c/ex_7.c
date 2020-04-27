/**
 * Give two (singly) linked lists, determine the the two lists intersect.
 * Return the intersecting node. Note that the intersection is defined based on
 * reference, not value. That is, if the kth node of the first linked list is
 * the exact same node (by reference) as the jth node of the second linked list,
 * then they are intersecting.
 */

#include <stdio.h>
#include "linked_list.h"
#include "../../lib/c/uthash/uthash.h"

node *hash_table_intersection(list *list_1, list *list_2);
node *iterative_intersection(list *list_1, list *list_2);
node *get_last_node_and_len(node *head, int *len);
node *advance_list(node *head, int places);

int main() {
    list *list_1 = create_list();
    insert_multiple(list_1, 3, insert_unsorted, 3, 7, 3);

    list *list_2 = create_list();
    insert_multiple(list_2, 4, insert_unsorted, 4, 8, 5, 1);

    list *list_3 = create_list();
    insert_multiple(list_3, 3, insert_unsorted, 2, 1, 2);
    append(list_1, list_3);
    append(list_2, list_3);

    node *intersect = hash_table_intersection(list_1, list_2);
    if (intersect == NULL) {
        printf("Null intersection\n");
    } else {
        printf("hash table intersect is head of list_3? %i\n", intersect == list_3->head);
    }

    intersect = iterative_intersection(list_1, list_2);
    if (intersect == NULL) {
        printf("Null intersection\n");
    } else {
        printf("iterative intersect is head of list_3? %i\n", intersect == list_3->head);
    }
    
}

/* O(A) space complexity, O(A + B) time complexity */
node *hash_table_intersection(list *list_1, list *list_2) {
    if (list_1 == NULL || list_2 == NULL) {
        return NULL;
    }

    node *l1_current = list_1->head;
    node *l2_current = list_2->head;
    if (l1_current == NULL || l2_current == NULL) {
        return NULL;
    }

    node *node_table = NULL;
    node *result = NULL;
    for ( ; l1_current != NULL; l1_current = l1_current->next) {
        HASH_ADD_PTR(node_table, address, l1_current);
    }
    for ( ; l2_current != NULL; l2_current = l2_current->next) {
        HASH_FIND_PTR(node_table, l2_current, result);
        if (result != NULL) {
            return result;
        }
    }
    return NULL;
}

/* O(1) space complexity, O(A + B) time complexity */
node *iterative_intersection(list *list_1, list *list_2) {
    if (list_1 == NULL || list_2 == NULL) {
        return NULL;
    }

    node *l1_current = list_1->head;
    node *l2_current = list_2->head;
    if (l1_current == NULL || l2_current == NULL) {
        return NULL;
    }

    int len_1 = 1;
    int len_2 = 1;
    node *last_1 = get_last_node_and_len(list_1->head, &len_1);
    node *last_2 = get_last_node_and_len(list_2->head, &len_2);

    if (last_1 != last_2) { // Last nodes are different; lists don't intersect
        return NULL;
    }
    
    if (len_1 < len_2) {
        l2_current = advance_list(l2_current, len_2 - len_1);
    } else if (len_2 < len_1) {
        l1_current = advance_list(l1_current, len_1 - len_2);
    }

    while (l1_current != l2_current) {
        l1_current = l1_current->next;
        l2_current = l2_current->next;
    }

    return l1_current;
}

node *get_last_node_and_len(node *head, int *len) {
    if (head == NULL) return NULL;

    node *current = head;
    for ( ; current->next != NULL; current = current->next, ++*len)
        ;

    return current;
}

node *advance_list(node *head, int places) {
    while (places--) {
        head = head->next;
    }
    return head;
}