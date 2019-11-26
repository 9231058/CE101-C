#include <stdio.h>

int check1(int, int, int, int, int, int, int, int);
int check2(int, int, int, int, int, int, int, int);
int distance2(int, int, int, int);

int main() {
}

int distance2(int x1, int y1, int x2, int y2) {
  return (x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2);
}

int check1(int x1, int y1, int x2, int y2, int x3, int y3, int x4, int y4) {
  int c = 0;

  if (distance2(x1, y1, x2, y2) == distance2(x3, y3, x4, y4)) {
    c++;
  }
  if (distance2(x1, y1, x4, y4) == distance2(x2, y2, x3, y3)) {
    c++;
  }
  if (distance2(x1, y1, x2, y2) == distance2(x2, y2, x3, y3)) {
    c++;
  }

  if (c == 3) {
    return 1;
  }
  return 0;
}

int check2(int x1, int y1, int x2, int y2, int x3, int y3, int x4, int y4) {
  int mx1 = (x1 + x2) / 2;
  int my1 = (y1 + y2) / 2;

  int mx2 = (x2 + x3) / 2;
  int my2 = (y2 + y3) / 2;

  int mx3 = (x3 + x4) / 2;
  int my3 = (x3 + x4) / 2;

  int mx4 = (x4 + x1) / 2;
  int my4 = (y4 + y1) / 2;

  return check1(mx1, my1, mx2, my2, mx3, my3, mx4, my4);
}
