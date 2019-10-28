/*
 * In The Name of God
 * =======================================
 * [] File Name : 5-2.c
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
#include <math.h>

int main(int argc, char *argv[]) {
  int x1, y1;
  int x2, y2;
  int x3, y3;
  int x4, y4;

  scanf("%d %d", &x1, &y1);
  scanf("%d %d", &x2, &y2);
  scanf("%d %d", &x3, &y3);
  scanf("%d %d", &x4, &y4);

  int s1 = (x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2);
  int s2 = (x2 - x3) * (x2 - x3) + (y2 - y3) * (y2 - y3);
  int s3 = (x3 - x4) * (x3 - x4) + (y3 - y4) * (y3 - y4);
  int s4 = (x4 - x1) * (x4 - x1) + (y4 - y1) * (y4 - y1);

  printf("%lf\n", sqrt(s1));
  printf("%lf\n", sqrt(s2));
  printf("%lf\n", sqrt(s3));
  printf("%lf\n", sqrt(s4));
}
