/*
 * In The Name of God
 * =======================================
 * [] File Name : 13-3.c
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
#include <stdlib.h>

struct menu_item {
  char name[255];
  int price;
};

int less_price(const void *a, const void *b) {
  return ((struct menu_item *) a)->price - ((struct menu_item *) b)->price;
}

struct dish {
  char name[255];
  int price;
};

struct dish dishes[] = {
  {"Ghost Dog", 38000},
  {"Boca Juniors", 45000},
  {"Jogo Bonito", 45000},
  {"Hitchcock", 49000},
  {"Hamlet", 40000},
  {"Pulp Fiction", 45000}, 
  {"Hidden", 40000}, 
};

struct herbal {
  char name[255];
  int price;
};

struct herbal herbals[] = {
  {"The Fish Fall in Love", 18000},
  {"Honey with Lemon", 16000},
  {"Green", 15000},
  {"Mint", 13000},
  {"Orange Blossom", 15000},
};

struct mocktail {
  char name[255];
  int price;
};

struct mocktail mocktails[] = {
  {"Mojito", 20000},
  {"Limonade", 20000},
  {"Dark Knight", 25000},
  {"Heath Ledger", 25000},
};

struct brewed_coffee {
  char name[255];
  int price;
  int is_double;
};

struct brewed_coffee brewed_coffees[] = {
  { "V60", 33000, 1 },
  { "French Press", 17000, 0 },
  { "V60", 25000, 0 },
};

int main() {
  printf("Dishes:\n");
  int number_of_dishes = sizeof(dishes) / sizeof(struct dish);

  qsort(dishes, number_of_dishes, sizeof(struct dish), less_price);
  for (int i = 0; i < number_of_dishes; i++) {
    printf("name: %s\n", dishes[i].name);
    printf("price: %d\n", dishes[i].price);
  }

  printf("\n\n");
  printf("Herbals:\n");
  int number_of_herbals = sizeof(herbals) / sizeof(struct herbal);

  qsort(herbals, number_of_herbals, sizeof(struct herbal), less_price);
  for (int i = 0; i < number_of_herbals; i++) {
    printf("name: %s\n", herbals[i].name);
    printf("price: %d\n", herbals[i].price);
  }

  printf("\n\n");
  printf("Mocktails:\n");
  int number_of_mocktails = sizeof(mocktails) / sizeof(struct mocktail);

  qsort(mocktails, number_of_mocktails, sizeof(struct mocktail), less_price);
  for (int i = 0; i < number_of_mocktails; i++) {
    printf("name: %s\n", mocktails[i].name);
    printf("price: %d\n", mocktails[i].price);
  }

  printf("\n\n");
  printf("Brewed Coffees:\n");
  int number_of_brewed_coffees = sizeof(brewed_coffees) / sizeof(struct brewed_coffee);

  qsort(brewed_coffees, number_of_brewed_coffees, sizeof(struct brewed_coffee), less_price);
  for (int i = 0; i < number_of_brewed_coffees; i++) {
    printf("name: %s\n", brewed_coffees[i].name);
    printf("price: %d\n", brewed_coffees[i].price);
    printf("is_double: %d\n", brewed_coffees[i].is_double);
  }
}
