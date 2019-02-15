/*
 * In The Name of God
 * =======================================
 * [] File Name : 12-5.c
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

int g = 5;

int main() {
  int a = 10;
  int b = 100;

  printf("g: %p\n", &g);
  printf("a: %p\n", &a);
  printf("b: %p\n", &b);
}
