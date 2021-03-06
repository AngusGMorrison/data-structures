#ifndef LINKED_LIST_H
#define LINKED_LIST_H

#include "../../lib/c/uthash/uthash.h"

enum {
    NULL_LIST = 1,
    NULL_NODE,
    STACK_OVERFLOW,
    STACK_EMPTY
};

typedef struct list {
    struct node *head;
    int size;
} list;

typedef struct node {
    struct node *address; // For storage in hash tables
    struct node *next;
    int data;
    UT_hash_handle hh;
} node;

list *create_list();
node *insert_sorted(list *list_p, int value);
node *insert_unsorted(list *list_p, int value);
list *insert_multiple(list *list_p, int count, node *(*insert)(list *, int), ...);
node *create_node(int value);
node *find_node(list *list_p, int value);
node *remove_node(list *list_p, int value);
list *append(list *list_p, list *to_append);
void reverse(list *list_p);
void print_list(list *list_p);
void error(char *msg, int code, ...);

#endif