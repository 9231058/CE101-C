/*
 * In The Name of God
 * =======================================
 * [] File Name : 6-1.c
 *
 * [] Creation Date : 09-11-2018
 *
 * [] Created By : Parham Alvani <parham.alvani@gmail.com>
 * =======================================
*/
/*
* Copyright (c)  2018 Parham Alvani.
*/

#include <stdio.h>

int main() {
    double x, y;
    scanf("%lf %lf", &x, &y);

    double area = x * y;

    printf("%lf\n", area);
    printf("%.02lf\n", area);
    printf("%.2lf\n", area);
    // it considers fractional part and decimal point in the range
    // that is expressed in the printf before the decimal point.
    printf("%05.02lf\n", area);
    printf("%g\n", area);
}
