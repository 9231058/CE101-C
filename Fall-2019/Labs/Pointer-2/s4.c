/*
 * In The Name of God
 * =======================================
 * [] File Name : s4.c
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

int menu(void);
char *add(void);

int main() {
  int size = 10;
  int last_memory = 0;
  char *mem;
  char **nmemories = NULL;
  char **memories = malloc(size * sizeof(char *));

  while (1) {
    int choice = menu();

    switch(choice) {
      case 1:
        printf("----\n");
        mem = add();
        memories[last_memory++] = mem;
        if (last_memory == size) {
          nmemories = malloc(size * 2 * sizeof(char *));
          memcpy(nmemories, memories, size * sizeof(char *));
          free(memories);
          memories = nmemories;
          size *= 2;
        }
        printf("----\n");
        break;
      case 2:
        printf("----\n");
        for (int i = 0; i < last_memory; i++) {
          printf("[%d] > %s", i, memories[i]);
        }
        printf("----\n");
        break;
      default:
        return 0;
    }
  }
}

int menu(void) {
  printf("1) Add\n");
  printf("2) Print\n");

  int choice;
  scanf("%d", &choice);
  getchar();
  return choice;
}

char *add(void) {
  char *mem = NULL;

  char input[200];
  do {
    if (fgets(input, 200, stdin) == NULL) {
      return NULL;
    }

    char *nmem;
    if (mem == NULL) {
      nmem = malloc(sizeof(char) * strlen(input));
      strcpy(nmem, input);
    } else {
      nmem = malloc(sizeof(char) * strlen(input) + strlen(mem));
      strcpy(nmem, mem);
      free(mem);
      strcpy(nmem + strlen(mem), input);
    }
    mem = nmem;
  } while(strcmp(input, "\n"));

  printf("> %s", mem);
  return mem;
}
