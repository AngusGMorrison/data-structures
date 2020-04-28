/**
 * Basic hash table error handling.
 */

#include <stdio.h>
#include <stdlib.h>
#include "error.h"

void error(char *msg, int code) {
    fprintf(stderr, "%s\n", msg);
    exit(code);
} 