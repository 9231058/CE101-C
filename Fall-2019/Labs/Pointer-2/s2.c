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

  char input[200];
  do {
    if (fgets(input, 200, stdin) == NULL) {
      return 0;
    }

    char *nmem;
    if (mem == NULL) {
      nmem = malloc(sizeof(char) * strlen(input));
      strcpy(nmem, input);
    } else {
      nmem = malloc(sizeof(char) * strlen(input) + strlen(mem));
      strcpy(nmem, mem);
      strcpy(nmem + strlen(mem), input);
    }
    mem = nmem;
  } while(strcmp(input, "\n"));

  printf("%s", mem);
}
