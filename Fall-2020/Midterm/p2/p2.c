#include <stdio.h>

int main() {
  int x, y;
  scanf("%d%d", &x, &y);
  int last = x > y ? x : y;
  for(int i = 2; i < last; i++)
    if ((x % i) == (y % i))
      printf("%d\n", i);
  return 0;
}
