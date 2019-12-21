/*
 * In The Name of God
 * =======================================
 * [] File Name : s2.c
 *
 * [] Creation Date : 17-12-2019
 *
 * [] Created By : Parham Alvani <parham.alvani@gmail.com>
 * =======================================
 */
/*
 * Copyright (c)  2019 Parham Alvani.
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main() {
  char *mem = NULL;
  int size = 0;

  char input[200];
  do {
    if (fgets(input, 200, stdin) == NULL) {
      return 0;
    }

    mem = realloc(mem, sizeof(char) * (size + strlen(input)));
    strcpy(mem + size, input);
    size += strlen(input);
  } while (strcmp(input, "\n"));

  printf("%s", mem);
}
