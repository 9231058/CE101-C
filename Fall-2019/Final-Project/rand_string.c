/*
 * In The Name of God
 * =======================================
 * [] File Name : rand_string.c
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
#include <stdlib.h>
#include <time.h>

char *rand_string(size_t size) {
  // reserves a place for NULL
  char *str = malloc((size + 1) * sizeof(char));

  const char charset[] = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";

  for (size_t i = 0; i < size; i++) {
    int key = rand() % (sizeof(charset) / sizeof(char) - 1);
    str[i] = charset[key];
  }
  str[size] = '\0';

  return str;
}

int main() {
  srand(time(NULL));

  char *str = rand_string(5);
  printf("rand string[5]: %s\n", str);
  free(str);
}
