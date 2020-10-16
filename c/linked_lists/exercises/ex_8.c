/**
 * Loop Detection: Given a circular linked list, implement an algorithm that
 * returns the node at the beginning of the loop.
 */

#include <stdio.h>
#include "linked_list.h"

node *find_loop_start(list *lst);

int main() {
    list *test_list = create_list();
    insert_multiple(test_list, 3, insert_unsorted, 5, 3, 9);
    list *part_2 = create_list();
    insert_multiple(part_2, 4, insert_unsorted, 1, 2, 4, 8);
    list *part_3 = create_list();
    insert_unsorted(part_3, 0);
    // Create a loop
    append(test_list, part_2);
    append(part_2, part_3);
    append(part_3, part_2);
    
    node *loop_start = find_loop_start(test_list);
    printf("Found correct node? %i\n", loop_start->data == 1);
}

node *find_loop_start(list *lst) {
    if (lst == NULL || lst->size < 2) {
        return NULL;
    }

    node *slow, *fast;
    slow = fast = lst->head;

    // Find collision point in loop
    do {
        slow = slow->next;
        fast = fast->next->next;
    } while (slow != fast);

    /* Collision point is equidistant from start of list and start of loop.
       By moving one pointer forward from the start of the loop, and another
       forward from the collision point, they will meet at the node to be
       returned. */
    slow = lst->head;
    while (slow != fast) {
        slow = slow->next;
        fast = fast->next;
    }

    return slow;
}