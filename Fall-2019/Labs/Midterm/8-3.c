/*
 * In The Name of God
 * =======================================
 * [] File Name : 8-3.c
 *
 * [] Creation Date : 21-11-2019
 *
 * [] Created By : Parham Alvani <parham.alvani@gmail.com>
 * =======================================
*/
/*
* Copyright (c)  2019 Parham Alvani.
*/

#include <stdio.h>

int f();

int main() {
  for (int i = 0; i < 10; i++) {
    printf("%d: %d\n", i + 1, f());
  }
}

int f() {
  static int k = 0;
  k++;
  if (k % 2 == 0) {
    return 0;
  } else {
    return 1;
  }
}
