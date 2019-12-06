#include <stdio.h>


int main(int argc, const char *argv[]) {
  int n;

  scanf("%d", &n);

  int a = 1;
  int b = 1;

  for (int i = 3; i <= n; i++) {
    if (i % 2 == 1) {
      a = a + b;
    } else {
      b = a + b;
    }
  }

  if (n % 2 == 1) {
    printf("%d\n", a);
  } else {
    printf("%d\n", b);
  }

}
