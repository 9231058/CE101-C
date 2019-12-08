/*
 * In The Name of God
 * =======================================
 * [] File Name : 10-5-2.c
 *
 * [] Creation Date : 06-12-2019
 *
 * [] Created By : Parham Alvani <parham.alvani@gmail.com>
 * =======================================
 */
/*
 * Copyright (c)  2019 Parham Alvani.
 */
#include <stdio.h>
#include <stdlib.h>

int C[100][100];

int main(int argc, char *argv[]) {
  // this program will calculate C(n, m)
  // for example C(5, 2) = 10
  int n;
  int m;

  scanf("%d %d", &n, &m);

  for (int i = 0; i <= n; i++) {
    C[i][0] = 1;
  }

  for (int i = 1; i <= n; i++) {
    for (int j = 1; j <= i; j++) {
      C[i][j] = C[i - 1][j - 1] + C[i - 1][j];
    }
  }

  printf("%d\n", C[n][m]);
}
