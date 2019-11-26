#include <stdio.h>

int power(int, int);

int main() {
  int n;

  scanf("%d", &n);

  for (int i = 0; i  * i <= n; i++) {
    for (int j = 0; j * j <= n; j++) {
      for (int k = 0; k * k <= n; k++) {
        int p = power(j, k);
        if (power(i, p) == n) {
          printf("a = %d b = %d c = %d\n", i, j, k);
        }
      }
    }
  }
  printf("There is no a^b^c\n");
}

int power(int a, int b) {
  if (a == 0) {
    return 0;
  }

  int r = 1;

  for (int i = 0; i < b; i++) {
    r = r * a;
  }

  return r;
}
