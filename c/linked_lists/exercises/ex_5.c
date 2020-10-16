/**
 * Sum Lists: You have two numbers represented by a linked list, where each
 * node contains a single digit. The digits are stored in reverse order, such
 * that the digit with the highest power of 10 is at the head of the list. Write
 * a function that adds the two numbers and returns the sum as a linked list.
 * 
 * FOLLOW UP: Supose the digits are stored in forward order. Repeat the above
 * problem.
 */

#include <stdio.h>
#include <stdlib.h>
#include "linked_list.h"

typedef struct carry_node {
    node *nd;
    int carry;
} carry_node;

list *sum_reversed_lists(list *list_1, list *list_2);
int sum_reversed_list(list *lst);
list *sum_forward_lists(list *list_1, list *list_2);
int sum_forward_list(list *lst);
node *recursive_sum_reversed(node *head_1, node *head_2, int carry);
node *recursive_sum_forward(list *list_1, list *list_2);
void pad_list(list *lst, int padding);
carry_node *recursive_sum_forward_helper(node *head_1, node *head_2);
carry_node *create_carry_node();

int main() {
    list *list_1 = create_list();
    insert_multiple(list_1, 3, insert_unsorted, 1, 2, 3);
    list *list_2 = create_list();
    insert_multiple(list_2, 3, insert_unsorted, 1, 2, 3);

    list *reverse_sum = sum_reversed_lists(list_1, list_2);
    print_list(reverse_sum);
    printf("\n");

    list *forward_sum = sum_forward_lists(list_1, list_2);
    print_list(forward_sum);
    printf("\n");

    list *recursive_reverse_sum = create_list();
    recursive_reverse_sum->head =
        recursive_sum_reversed(list_1->head, list_2->head, 0);
    print_list(recursive_reverse_sum);
    printf("\n");

    list *recursive_forward_sum = create_list();
    recursive_forward_sum->head =
        recursive_sum_forward(list_1, list_2);
    print_list(recursive_forward_sum);
    printf("\n");
}

/* On the order of n, but suboptimal, requiring more iterations than a
   recursive approach and risking integer overflow for large lists/lists with
   large numbers. */
list *sum_reversed_lists(list *list_1, list *list_2) {
    int sum = sum_reversed_list(list_1) + sum_reversed_list(list_2);

    node *head = NULL;
    node *current = NULL;
    do { // Use do-while loop to ensure output list will contain at least 0
        if (head == NULL) {
            current = create_node(sum % 10);
            head = current;
        } else {
            current->next = create_node(sum % 10);
            current = current->next;
        }
    } while (sum /= 10);

    list *sum_list = create_list();
    sum_list->head = head;
    return sum_list;
}

int sum_reversed_list(list *lst) {
    node *current = lst->head;
    if (current == NULL) {
        return 0;
    }

    int sum = 0;
    int exp = 1;
    while (current != NULL) {
        sum += current->data * exp;
        exp *= 10;
        current = current->next;
    }

    return sum;
}

list *sum_forward_lists(list *list_1, list* list_2) {
    int sum = sum_forward_list(list_1) + sum_forward_list(list_2);
    node *previous = NULL;
    node *current = NULL;

    do {
        current = create_node(sum % 10);
        current->next = previous;
        previous = current;
    } while (sum /= 10);

    list *sum_list = create_list();
    sum_list->head = previous;
    return sum_list;
}

int sum_forward_list(list *lst) {
    node *current = lst->head;
    if (current == NULL) {
        return 0;
    }

    int total = 0;
    while (current != NULL) {
        total = total * 10 + current->data;
        current = current->next;
    }

    return total;
}

node *recursive_sum_reversed(node *head_1, node *head_2, int carry) {
    if (head_1 == NULL && head_2 == NULL && carry == 0) {
        return NULL;
    }

    int value = carry;
    if (head_1 != NULL) {
        value += head_1->data;
    }

    if (head_2 != NULL) {
        value += head_2->data;
    }

    node *result = create_node(value % 10); // 1s column of number

    // Recurse
    if (head_1 != NULL || head_2 != NULL) {
        result->next = recursive_sum_reversed(
            head_1 == NULL ? NULL : head_1->next,
            head_2 == NULL ? NULL : head_2->next,
            value >= 10 ? 1 : 0
        );
    }

    return result;
}

node *recursive_sum_forward(list *list_1, list *list_2) {
    int size_1 = list_1->size;
    int size_2 = list_2->size;
    // Pad the shorter list with 0s
    if (size_1 < size_2) {
        pad_list(list_1, size_2 - size_1);
    } else {
        pad_list(list_2, size_1 - size_2);
    }

    // Add lists
    carry_node *sum = recursive_sum_forward_helper(list_1->head, list_2->head);

    // Insert any remaining carry value at the front of the list
    if (sum->carry == 0) {
        return sum->nd;
    } else {
        node *result = create_node(sum->carry);
        result->next = sum->nd;
        return result;
    }
}

void pad_list(list *lst, int padding) {
    node *next = lst->head;
    while (padding--) {
        node *pad_node = create_node(0);
        pad_node->next = next;
        next = pad_node;
    };
    lst->head = next;
}

carry_node *recursive_sum_forward_helper(node *head_1, node *head_2) {
    if (head_1 == NULL && head_2 == NULL) {
        carry_node *sum = create_carry_node();
        return sum;
    }

    // Add smaller digits recursively
    carry_node *sum = recursive_sum_forward_helper(head_1->next, head_2->next);
    int value = sum->carry + head_1->data + head_2->data;

    // Insert sum of current digits
    node *full_result = create_node(value % 10);
    full_result->next = sum->nd;

    // Return sum so far, and the carry value
    sum->nd = full_result;
    sum->carry = value / 10;
    return sum;
}

carry_node *create_carry_node() {
    carry_node *carry = malloc(sizeof(carry_node));
    if (carry == NULL) {
        error("create_carry_node: nd is null\n", NULL_NODE);
    }

    carry->nd = NULL;
    return carry;
}