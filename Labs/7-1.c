/*
 * In The Name of God
 * =======================================
 * [] File Name : 7-1.c
 *
 * [] Creation Date : 09-11-2018
 *
 * [] Created By : Parham Alvani <parham.alvani@gmail.com>
 * =======================================
*/
/*
* Copyright (c)  2018 Parham Alvani.
*/
#include <stdio.h>

int main() {
  double d;
  scanf("%lf", &d);

  if ((int) d == d) { // d is an integer
    printf("%lf is an integer\n", d);
  } else {
    printf("%lf isn't an integer (%d)\n", d, (int)d);
  }
}
