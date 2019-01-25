/*
 * In The Name of God
 * =======================================
 * [] File Name : p5.c
 *
 * [] Creation Date : 25-01-2019
 *
 * [] Created By : Parham Alvani <parham.alvani@gmail.com>
 * =======================================
 */
/*
 * Copyright (c)  2019 Parham Alvani.
 */
/*
 * This file is uesed to write time with the following
 * format into binary files.
 */

#include <stdio.h>

struct time {
  int hour;
  int min;
  int sec;
  int mili_sec;
};

int main(int argc, const char *argv[]) {
  struct time t1 = { 10, 10, 10, 100 };
  struct time t2 = { 20, 20, 20, 200 };
  struct time t3 = { 12, 12, 12, 300 };
  struct time t4 = { 13, 13, 13, 400 };

  FILE *output = fopen("fi.bin", "wb");
  fwrite(&t1, sizeof(struct time), 1, output);
  fwrite(&t2, sizeof(struct time), 1, output);
  fwrite(&t3, sizeof(struct time), 1, output);
  fwrite(&t4, sizeof(struct time), 1, output);
}
