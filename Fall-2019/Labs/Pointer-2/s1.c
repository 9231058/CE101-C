/*
 * In The Name of God
 * =======================================
 * [] File Name : s1.c
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

int main() {
  char input[200];

  if (fgets(input, 200, stdin) == NULL) {
    return 0;
  }

  printf("%s", input);
}
