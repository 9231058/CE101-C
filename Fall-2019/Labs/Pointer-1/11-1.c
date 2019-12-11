#include <stdio.h>

void fun(int *p) {
  int q = 10;
  // change the pointer itself doesn't reflect outside of function
  p = &q;
}

int main() {
  int r = 20;
  int *p = &r;
  fun(p);
  printf("%d\n", *p);
  return 0;
}
