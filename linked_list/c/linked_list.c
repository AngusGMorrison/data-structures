/**
 * C implementation of a linked list of integers maintained in sorted order.
 */

#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include "linked_list.h"

enum {
    NULL_LIST = 1,
    NULL_NODE
};

/* Return a newly created list with size 0. */
list *create_list() {
    list *list_p = malloc(sizeof(list));
    if (list_p == NULL) {
        error("main: couldn't allocate list\n", NULL_LIST);
    }
    list_p->size = 0;
    return list_p;
}

/* Iterate through the list until the correct position for value is found, then
   insert a new node to hold it. Returns the new node. */
node *insert(list *list_p, int value) {
    if (list_p == NULL) {
        error("insert: list is NULL\n", NULL_LIST);
    }

    node *new_p = create_node(value);
    node *current = list_p->head;
    if (current == NULL || current->data > value) {
        // Create node as head
        new_p->next = list_p->head;
        list_p->head = new_p;
        list_p->size++;
        return new_p;
    }

    // Find correct position in list
    while (current->next != NULL && current->next->data < value) {
        current = current->next;
    }
    // Correct position found; insert node
    new_p->next = current->next;
    current->next = new_p;
    list_p->size++;

    return new_p;
}

list *insert_multiple(list *list_p, int count, ...) {
    if (count < 1) {    // No data to add
        return list_p;
    }

    va_list args;
    int data;
    va_start(args, count);

    while (count--) {
        data = va_arg(args, int);
        insert(list_p, data);
    }
    va_end(args);

    return list_p;
}

/* Return a newly created node. */
node *create_node(int value) {
    node *new_p = malloc(sizeof(node));
    if (new_p == NULL) {
        error("insert: couldn't allocate node\n", NULL_NODE);
    }
    new_p->next = NULL;
    new_p->data = value;
    return new_p;
}

/* Iterate through the list and return the first matching node, or NULL if a
   matching node is not found. */
node *find_node(list *list_p, int value) {
    if (list_p == NULL) {
        error("find_node: list is NULL\n", NULL_LIST);
    }

    node *current = list_p->head;
    while (current != NULL && current->data != value) {
        current = current->next;
    }
    return current;
}

/* Removes the first node with value from the list and returns it. */
node *remove_node(list *list_p, int value) {
    if (list_p == NULL) {
        error("remove_node: list is NULL\n", NULL_LIST);
    }

    node *current = list_p->head;
    if (current == NULL) {
        return NULL;
    }

    if (current->data == value) {
        // Remove the head of the list
        list_p->head = current->next;
        list_p->size--;
        return current;
    }

    for ( ; current->next != NULL; current = current->next) {
        if (current->next->data == value) {
            node *match = current->next;
            current->next = match->next;
            list_p->size--;
            return match;
        }
    }

    return NULL;
}

void reverse(list *list_p) {
    if (list_p == NULL) {
        error("reverse: list is NULL\n", NULL_LIST);
    }

    node *next;
    node *current = list_p->head;
    node *previous = NULL;
    while (current != NULL) {
        next = current->next;
        current->next = previous;
        previous = current;
        current = next;
    }

    list_p->head = previous;
}

/* Print each node in the list. */
void print_list(list *list_p) {
    for (node *current = list_p->head; current != NULL; current = current->next) {
        printf("%i\n", current->data);
    }
}

/* Handle errors with variable args. */
void error(char *msg, int code, ...) {
    va_list vargs;
    va_start(vargs, code);
    vfprintf(stderr, msg, vargs);
    fprintf(stderr, "\n");
    va_end(vargs);
    exit(code);
}