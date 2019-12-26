/*
 * In The Name of God
 * =======================================
 * [] File Name : 13-1.c
 *
 * [] Creation Date : 26-12-2019
 *
 * [] Created By : Parham Alvani <parham.alvani@gmail.com>
 * =======================================
*/
/*
* Copyright (c)  2019 Parham Alvani.
*/
#include <stdio.h>

struct s1 {
  int f1;
  int f2;
  char f3[12];
};

struct s2 {
  int f1;
  char f2[3];
  char f3[2];
};

struct s3 {
  int f1;
  char f2[3];
  char f3[2];
} __attribute__((packed));

int main() {
  printf("%lu\n", sizeof(struct s1));
  printf("%lu\n", sizeof(struct s2));
  printf("%lu\n", sizeof(struct s3));
}
