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
#include <math.h>

#define awesome pow(pow(3.0, 0.05), 20.0);

int main() {
  double d = awesome;

  printf("Method 1 -- based on equality\n");
  if (lround(d) == d) {
    printf("%lf is an integer\n", d);
  } else {
    printf("%lf isn't an integer (%d)\n", d, (int)d);
  }

  printf("Method 2 -- based on error rage\n");
  if (fabs(lround(d) - d) <= 0.000001) { // d is an integer
    printf("%lf is an integer\n", d);
  } else {
    printf("%lf isn't an integer (%d)\n", d, (int)d);
  }
}
