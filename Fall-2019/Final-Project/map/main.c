/*
 * In The Name of God
 * =======================================
 * [] File Name : main.c
 *
 * [] Creation Date : 17-01-2020
 *
 * [] Created By : Parham Alvani <parham.alvani@gmail.com>
 * =======================================
*/
/*
* Copyright (c)  2020 Parham Alvani.
*/

#include <stdio.h>
#include <stdint.h>

enum blocks {
  ENERGY = 1,
  MITOSIS,
  FORBIDDEN,
  NORMAL
};

int main() {
  uint32_t n;

  scanf("%d", &n);

  FILE *f = fopen("map.bin", "w");
  if (!f) {
    perror("fopen:");
  }

  fwrite(&n, sizeof(uint32_t), 1, f);

  printf("ENERGY = %d\n", ENERGY);
  printf("MITOSIS = %d\n", MITOSIS);
  printf("FORBIDDEN = %d\n", FORBIDDEN);
  printf("NORMAL = %d\n", NORMAL);

  // reads blocks mode from user and stores them with
  // row-major format.
  for (int i = 0; i < n; i++) {
    for (int j = 0; j < n; j++) {
      uint8_t mode;
      printf("[%d][%d] = ", i, j);
      getchar();
      scanf("%c", &mode);
      fwrite(&mode, sizeof(uint8_t), 1, f);
    }
  }

  fclose(f);
}
