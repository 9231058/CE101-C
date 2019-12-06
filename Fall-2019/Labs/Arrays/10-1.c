#include <stdio.h>

int fibonacci(int);

int main(int argc, const char *argv[]) {
  int n;
  
  scanf("%d", &n);

  printf("%d\n", fibonacci(n));
}

int fibonacci(int n) {
    if (n == 1 || n == 2) {
        return 1;
    }
    return fibonacci(n - 1) + fibonacci(n - 2);
}
