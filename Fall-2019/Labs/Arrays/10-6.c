/*
 * In The Name of God
 * =======================================
 * [] File Name : 10-7.c
 *
 * [] Creation Date : 06-12-2019
 *
 * [] Created By : Parham Alvani <parham.alvani@gmail.com>
 * =======================================
*/
/*
* Copyright (c)  2019 Parham Alvani.
*/
#include <stdio.h>
#include <string.h>
#include <ctype.h>

int main(int argc, char *argv[]) {
  char input[200];
  char str[200];

  fgets(input, 200, stdin);

  for (int i = 0; i < strlen(input); i++) {
    input[i] = tolower(input[i]);
  }

  int idx = 0;
  for (int i = 0; i < strlen(input); i++) {
    if (!isspace(input[i])) {
      str[idx++] = input[i];
    }
  }
  str[idx] = 0;

  for (int i = 0; i < strlen(str); i++) {
    if (str[i] != str[strlen(str) - 1 - i]) {
      printf("False\n");
      return 0;
    }
  }
  printf("True\n");
}
