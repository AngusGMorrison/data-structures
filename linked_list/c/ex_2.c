/**
 * Return Kth to Last: Implement an algorithm to find the kth to last element
 * of a singly linked list.
 */ 

#include <stdio.h>
#include <stdlib.h>
#include "linked_list.h"

node *kth_to_last_iterative(list *list_p, int k);
node *kth_to_last_recursive(node *head, int k, int *count);
node *kth_to_last_iterative_optimal(list *list_p, int k);

int main() {
    list *test_list = create_list();
    insert(test_list, 1);
    insert(test_list, 2);
    insert(test_list, 3);
    insert(test_list, 4);
    insert(test_list, 5);

    node *kth_last = kth_to_last_iterative(test_list, 2);
    printf("%i\n", kth_last->data);

    int count = 0;
    kth_last = kth_to_last_recursive(test_list->head, 2, &count);
    printf("%i\n", kth_last->data);

    kth_last = kth_to_last_iterative_optimal(test_list, 2);
    printf("%i\n", kth_last->data);
}

/* Suboptimal solution for large lists, O(2n) time, O(1) space.
   Count the nodes in the list, then move to the correct position on a second
   iteration. */
node *kth_to_last_iterative(list *list_p, int k) {
    node *current = list_p->head;
    node *runner = current;
    int count;

    for (count = 0; runner != NULL; runner = runner->next, count++)
        ;

    for (int i = 0; i < (count - k); i++, current = current->next)
        ;

    return current;
}

/* Recursive solution, O(n) time, O(n) space.
   When k recursive function calls have returned, return the node that was
   originally passed in to that function. */
node *kth_to_last_recursive(node *head, int k, int *count) {
    if (head == NULL) {
        return NULL;
    }

    node *nd = kth_to_last_recursive(head->next, k, count);
    (*count)++;
    if (*count == k) {
        return head;
    }
    return nd;
}

/* Optimal solution, O(n) time and O(1) space.
   Separate runners by k elements, then move both until the first reaches the
   end. */
node *kth_to_last_iterative_optimal(list *list_p, int k) {
    node *current = list_p->head;
    node *runner = current;

    for (int i = 0; i < k && runner != NULL; i++) {
        runner = runner->next;
    }

    if (runner == NULL) {
        fprintf(stderr, "List has fewer than %i elements\n", k);
        exit(1);
    }

    while (runner != NULL) {
        runner = runner->next;
        current = current->next;
    }

    return current;
}