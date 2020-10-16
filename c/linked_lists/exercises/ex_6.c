/**
 * Palindrome: Implement a function to check if a linked list is a palindrome
 */

#include <stdio.h>
#include <stdlib.h>
#include "linked_list.h"

int is_palindrome(list *lst);
int is_palindrome_recursive(node *front_nd, int length);

int main() {
    list *odd_list = create_list();
    insert_multiple(odd_list, 5, insert_unsorted, 1, 2, 3, 2, 1);
    printf("odd_list is palindrome: %i\n", is_palindrome(odd_list));

    list *even_list = create_list();
    insert_multiple(even_list, 6, insert_unsorted, 1, 2, 3, 3, 2, 1);
    printf("even_list is palindrome: %i\n", is_palindrome(even_list));

    list *random_list = create_list();
    insert_multiple(random_list, 6, insert_unsorted, 3, 8, 1, 0, 3, 5);
    printf("random_list is palindrome: %i\n", is_palindrome(random_list));

    list *empty_list = create_list();
    printf("empty_list is palindrome: %i\n", is_palindrome(empty_list));

    int result = is_palindrome_recursive(odd_list->head, odd_list->size);
    printf("odd_list is palindrome recursively: %i\n", result);

    result = is_palindrome_recursive(even_list->head, even_list->size);
    printf("even_list is palindrome recursively: %i\n", result);

    result = is_palindrome_recursive(random_list->head, random_list->size);
    printf("random_list is palindrome recursively: %i\n", result);

    result = is_palindrome_recursive(empty_list->head, empty_list->size);
    printf("empty_list is palindrome recursively: %i\n", result);
}

/* A stack-based approach can be used where the length of the list is unknown.
   O(n) space and time complexity. */
int is_palindrome(list *lst) {
    void push(node *nd);
    node *pop(), *current, *runner;

    if (lst == NULL) {
        error("is_palindrome: lst is NULL\n", NULL_LIST);
    }

    if ((current = runner = lst->head) == NULL) {
        return 0;
    }
    // Find the end of the list
    while (runner != NULL && runner->next != NULL) {
        push(current);
        current = current->next;
        runner = runner->next->next;
    }

    if (runner != NULL) { // Odd number in list; skip middle item
        current = current->next;
    }

    while (current != NULL) {
        if (pop()->data != current->data) {
            return 0;
        }
        current = current->next;
    }

    return 1;
}

#define MAX_STACK 100
static node *stack[MAX_STACK];
static int stack_pos = 0;

void push(node *nd) {
    if (stack_pos < MAX_STACK) {
        stack[stack_pos++] = nd;
    } else {
        error("push: stack overflow\n", STACK_OVERFLOW);
    }
}

node *pop() {
    if (stack_pos <= 0) {
        error("pop: stack empty\n", STACK_EMPTY);
    }
    return stack[--stack_pos];
}

/* A recursive approach can also be used, ideally if length is know in advance.
   O(n) space and time complexity. */
int is_palindrome_recursive(node *front_nd, int length) {
    static node *back_nd;

    // When the middle of the node has been reached, or the list is empty...
    if (front_nd == NULL) {
        return 0;
    } else if (length == 0) {
        back_nd = front_nd;
        return 1;
    } else if (length == 1) {
        back_nd = front_nd->next;
        return 1;
    }
    
    int result = is_palindrome_recursive(front_nd->next, length - 2);
    // If child calls are not a palindrome, pass a failure back up
    if (result == 0 || front_nd->data != back_nd->data) {
        return 0;
    }
    // Move the back node along and pass true up the chain for the current pair
    back_nd = back_nd->next;
    return result;
}
