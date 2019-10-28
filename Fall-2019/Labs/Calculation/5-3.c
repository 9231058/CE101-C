/*
 * In The Name of God
 * =======================================
 * [] File Name : 5-3.c
 *
 * [] Creation Date : 28-10-2019
 *
 * [] Created By : Parham Alvani <parham.alvani@gmail.com>
 * =======================================
*/
/*
* Copyright (c)  2019 Parham Alvani.
*/
#include <stdio.h>

const double PI = 3.14159265359;

int main() {
  double r;

  scanf("%lf", &r);

  printf("P = %lf\n", 2 * r * PI);
  printf("S = %lf\n", r * r * PI);
}
