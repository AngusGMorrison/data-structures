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

list *partition(list *list_p, int partition);

int main() {
    list *test_list = create_list();
    insert_multiple(test_list, 15, insert_unsorted, 3, 6, 5, 7, 8, 2, 5, 0, 1, 3, 4, 7, 4, 2, 9);
    partition(test_list, 5);
    print_list(test_list);
}

list *partition(list *list_p, int partition) {
    if (list_p->size < 2) return list_p;

    node *less, *more;
    less = more = list_p->head;

    while(1) {
        while (less != NULL && less->data >= partition) {
            less = less->next;
            printf("less->data = %i\n", less->data);
        }

        while (more != NULL && more->data <= partition) {
            more = more->next;
            printf("more->data = %i \n", more->data);
        }

        if (more == NULL || less == NULL) break;

        printf("Swapping less = %i and more = %i\n", less->data, more->data);
        int temp = less->data;
        less->data = more->data;
        more->data = temp;
    }

    return list_p;
}