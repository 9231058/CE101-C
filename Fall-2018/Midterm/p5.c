/*
 * In The Name of God
 * =======================================
 * [] File Name : p5.c
 *
 * [] Creation Date : 26-11-2019
 *
 * [] Created By : Parham Alvani <parham.alvani@gmail.com>
 * =======================================
*/
/*
* Copyright (c)  2019 Parham Alvani.
*/

#include <stdio.h>

int T(int, int);

int main() {
  int x1, x2, n;

  scanf("%d %d %d", &x1, &x2, &n);

  int z = 0;
  for (int i = x1; i <= x2; i++) {
    z += T(n, i);
  }
  printf("%d\n", z);
}

int T(int n, int x) {
  if (n == 0) {
    return 1;
  }

  if (n == 1) {
    return x;
  }

  return x * T(n - 1, x) + T(n - 2, x);
}
