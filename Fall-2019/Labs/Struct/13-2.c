/*
 * In The Name of God
 * =======================================
 * [] File Name : 13-2.c
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
#include <string.h>
#include <stdlib.h>

struct film {
  char name[255];
  int year;
};

int main() {
  int n;
  scanf("%d", &n);
  getchar();

  struct film *films = malloc(sizeof(struct film) * n);

  for (int i = 0; i < n; i++) {
    printf("name: ");
    char name[255];
    fgets(name, 255, stdin);

    printf("year: ");
    int year;
    scanf("%d", &year);
    getchar();

    films[i].year = year;
    strcpy(films[i].name, name);
  }

  for (int i = 0; i < n; i++) {
    printf("name: %s", films[i].name);
    printf("year: %d\n", films[i].year);
    printf("--\n");
  }
}
