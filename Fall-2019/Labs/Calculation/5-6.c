/*
 * In The Name of God
 * =======================================
 * [] File Name : 5-6.c
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

int main() {
  double l, w;

  scanf("%lf %lf", &l, &w);

  double area = l * w;

  printf("%lf\n", area);
  printf("%.2lf\n", area);
  printf("%.02lf\n", area);
  printf("%04.2lf\n", area);
  printf("%05.2lf\n", area);
  printf("%g\n", area);
}
