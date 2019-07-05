/*
 * In The Name of God
 * =======================================
 * [] File Name : 12-4.c
 *
 * [] Creation Date : 14-12-2018
 *
 * [] Created By : Parham Alvani <parham.alvani@gmail.com>
 * =======================================
*/
/*
* Copyright (c)  2018 Parham Alvani.
*/

#include <stdio.h>

int f(int* p) {
  printf("a = %d\n", p); // a = 10?
}

int main() {
  int a = 10;
  f((int *) a);
}
