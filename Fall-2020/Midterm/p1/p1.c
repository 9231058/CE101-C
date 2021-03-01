#include <stdio.h>

int check_sequence(int x) {
  int last_num = 1, n = 1;
  while (last_num < x) {
    n++;
    last_num += 2 * n;
  }
  if (last_num == x)
    return n;
  else
    return -1;
}

int main() {
  int num, res;
  scanf("%d", &num);
  res = check_sequence(num);
  printf("%d", res);
}
