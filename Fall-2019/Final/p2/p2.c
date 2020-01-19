#include <stdio.h>
#include <stdlib.h>

int *set_union(int *a, int size_a, int *b, int size_b, int *result_size) {
  int *result = malloc((size_a + size_b) * sizeof(int));
  int size = 0;

  for (int i = 0; i < size_a; i++) {
    result[size++] = a[i];
  }

  for (int i = 0; i < size_b; i++) {
    int f = 0;
    for (int j = 0; j < size_a; j++) {
      if (b[i] == a[j]) {
        f = 1;
        break;
      }
    }
    if (f == 0) {
      result[size++] = b[i];
    }
  }

  *result_size = size;

  return result;
}

int *set_union3(int *a, int size_a, int *b, int size_b, int *c, int size_c, int *result_size) {
  int rs;
  int *r1 = set_union(a, size_a, b, size_b, &rs);
  int *r2 = set_union(r1, rs, c, size_c, result_size);
  free(r1);
  return r2;
}

int main() {
  int a[] = {1, 2, 3};
  int b[] = {1, 2, 3, 4};
  int c[] = {1, 5, 6, 4, 10};

  int rs;
  int *r = set_union3(a, sizeof(a) / sizeof(int), b, sizeof(b) / sizeof(int), c, sizeof(c) / sizeof(int), &rs);

  for (int i = 0; i < rs; i++) {
    printf("%d ", r[i]);
  }
  printf("\n");
}
