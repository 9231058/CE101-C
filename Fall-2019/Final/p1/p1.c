#include <stdio.h>
#include <stdlib.h>

void change(int **p) {
  int i, j;

  for (i = 0; i < 2; i++) {
    for (j = 0; j < 2; j++) {
      int tmp = *((*p + i) + j);
      *((*p + i) + j) = *(*(p + i + 1) + j);
      *(*(p + i + 1) + j) = tmp;
    }
  }
}

void print(int *p, int size) {
  int i;
  for (i = 0; i < size; i++)
    printf("[%d] = %d ", i, p[i]);
  printf("\n");
}

int main() {
  int i, j, *arr, *parr[3];

  arr = (int *)malloc(9 * sizeof(int));
  for (i = 0; i < 3; i++) {
    parr[i] = arr + i * 3;
    for (j = 0; j < 3; j++)
      arr[i * 3 + j] = (i + 2) * (j + 3);
  }

  print(arr, 9);
  change(&(parr[0]));
  print(arr, 9);

  return 0;
}
