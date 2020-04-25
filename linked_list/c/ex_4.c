/**
 * Partition: Write code to partition a linked list around a value x, such that
 * all nodes less than x come before all nodes greater than or equal to x. If
 * x is contained within the list, the values of x only need to be after the
 * elements less than x. The partition element x can appear anywhere in the
 * "right partition"; it does not need to appear between the left and right
 * paritions.
 */

#include <stdio.h>
#include "linked_list.h"

list *stable_partition(list *list_p, int partition);
list *unstable_partition(list *list_p, int partition);

int main() {
    list *test_list = create_list();
    insert_multiple(test_list, 15, insert_unsorted, 3, 6, 5, 7, 8, 2, 5, 0, 1, 3, 4, 7, 4, 2, 9);
    unstable_partition(test_list, 5);
    printf("Returned\n");
    print_list(test_list);
}

list *stable_partition(list *list_p, int partition) {
    node *current = list_p->head;
    if (current == NULL || list_p->size < 2) {
        return NULL;
    }

    node *less_than_head = NULL;
    node *less_than_tail = NULL;
    node *more_or_equal_head = NULL;
    node *more_or_equal_tail = NULL;

    for ( ; current != NULL; current = current->next) {
        if (current->data < partition) {
            if (less_than_head == NULL) {
                less_than_head = current;
                less_than_tail = less_than_head;
            } else {
                less_than_tail->next = current;
                less_than_tail = less_than_tail->next;
            }
        } else {
            if (more_or_equal_head == NULL) {
                more_or_equal_head = current;
                more_or_equal_tail = more_or_equal_head;
            } else {
                more_or_equal_tail->next = current;
                more_or_equal_tail = more_or_equal_tail->next;
            }
        }
    }

    list_p->head = less_than_head;
    less_than_tail->next = more_or_equal_head;
    return list_p;
}

list *unstable_partition(list *list_p, int partition) {
    node *current = list_p->head;
    if (current == NULL || list_p->size < 2) {
        return NULL;
    }

    node *head = current;
    node *tail = current;

    for (node *next; current != NULL; current = next) {
        next = current->next;
        if (current->data < partition) {
            current->next = head;
            head = current;
        } else {
            tail->next = current;
            tail = tail->next;
        }
    }

    tail->next = NULL;
    list_p->head = head;
    return list_p;
}