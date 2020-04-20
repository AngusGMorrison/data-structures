#ifndef LINKED_LIST_H
#define LINKED_LIST_H

typedef struct list {
    struct node *head;
    int size;
} list;

typedef struct node {
    struct node *next;
    int data;
} node;

list *create_list();
node *insert(list *list_p, int value);
node *create_node(int value);
node *find_node(list *list_p, int value);
node *remove_node(list *list_p, int value);
void reverse(list *list_p);
void print_list(list *list_p);
void error(char *msg, int code, ...);

#endif