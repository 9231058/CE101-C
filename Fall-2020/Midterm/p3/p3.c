#include <stdio.h>

int check_right_left_digits(int num) {
  if(num < 10)
    return 0;
  int right_digit = num % 10, chandgan = 1, left_digit;
  int temp = num / 10;
  while(temp > 0) {
    left_digit = temp % 10;
    chandgan *= 10;
    temp /= 10;
  }
  if(right_digit != left_digit)
    return -1;
  return (num - left_digit * chandgan) / 10;
}

int digit_count(int num) {
  int cnt = 0;
  while(num > 0) {
    cnt++;
    num /= 10;
  }
  return cnt;
}

int check_num(int num) {
  int res = num;
  int cnt1, cnt2;
  while(res > 0) {
    cnt1 = digit_count(res);
    res = check_right_left_digits(res);
    cnt2 = digit_count(res);
    for(int i = 0; i < cnt1 - cnt2 - 2; i++) {
      if(res % 10 != 0)
        return 0;
      res /= 10;
    }
  }
  if(res == 0)
    return 1;
  else
    return 0;
}

int main() {
  int num;
  scanf("%d", &num);
  for(int i = 1; i < num; i++) {
    if (check_num(i))
      printf("%d\n", i);
  }
  return 0;
}
