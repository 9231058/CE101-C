#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct student {
  int id;
  char first_name[255];
  char last_name[255];
};

struct node {
  struct student *st;
  struct node *next;
};

struct node *new_node() {
  struct node *nn = malloc(sizeof(struct node));
  nn->next = NULL;
  nn->st = NULL;
  return nn;
}

void list_insert(struct node **head, struct node *new) {
  if (*head == NULL) {
    *head = new;
    return;
  }

  struct node *itr = *head;
  while (itr->next != NULL) {
    itr = itr->next;
  }
  itr->next = new;
}

void student_print(const struct student *std) {
  printf("First Name: %s\n", std->first_name);
  printf("Last Name: %s\n", std->last_name);
  printf("ID: %d\n", std->id);
}

void student_list_print(struct node *head) {
  struct node *itr = head;
  int i = 0;

  while (itr != NULL) {
    i++;
    printf("-- %d --\n", i);
    student_print(itr->st);
    printf("-- - --\n");
    itr = itr->next;
  }
}

void student_list_insert(struct node **head, int id, const char *first_name, const char *last_name) {
  struct node *nn = new_node();
  struct student *st = malloc(sizeof(struct student));
  st->id = id;
  if (strlen(first_name) <= 255) {
    strcpy(st->first_name, first_name);
  } else {
    strcpy(st->first_name, "-");
  }
  if (strlen(last_name) <= 255) {
    strcpy(st->last_name, last_name);
  } else {
    strcpy(st->last_name, "-");
  }
  nn->st = st;

  list_insert(head, nn);
}

int main() {
  struct node *student_list = NULL;

  while (1) {
    int select;
    int id;
    char first_name[255];
    char last_name[255];

    printf("1) Insert\n");
    printf("2) List\n");
    printf("3) Quit\n");
    printf("> ");

    scanf("%d", &select);

    switch (select) {
      case 1:
        printf("ID: ");
        scanf("%d", &id);
        printf("First Name: ");
        scanf("%s", first_name);
        printf("Last Name: ");
        scanf("%s", last_name);

        student_list_insert(&student_list, id, first_name, last_name);
        break;
      case 2:
        student_list_print(student_list);
        break;
      case 3:
        return 0;
        break;
      default:
        printf("Please select a correct option\n");
    }
  }
}
