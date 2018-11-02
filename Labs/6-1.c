#include <stdio.h>

int main() {
    double x, y;
    scanf("%lf %lf", &x, &y);

    double area = x * y;

    printf("%lf\n", area);
    printf("%.02lf\n", area);
    printf("%.2lf\n", area);
    printf("%02.02lf\n", area);
    printf("%g\n", area);
}