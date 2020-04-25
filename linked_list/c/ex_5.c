/**
 * Sum Lists: You have two numbers represented by a linked list, where each
 * node contains a single digit. The digits are stored in reverse order, such
 * that the digit with the highest power of 10 is at the head of the list. Write
 * a function that adds the two numbers and returns the sum as a linked list.
 * 
 * FOLLOW UP: Supose the digits are stored in forward order. Repeat the above
 * problem.
 */

#include "linked_list.h"

list *sum_reversed_lists(list *list_1, list *list_2);
int sum_reversed_list(list *lst);
list *sum_forward_lists(list *list_1, list *list_2);
int sum_forward_list(list *lst);

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
}

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