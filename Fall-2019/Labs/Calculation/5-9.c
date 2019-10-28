/*
 * In The Name of God
 * =======================================
 * [] File Name : 6-2.c
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
#include <math.h>
#include <stdlib.h>

int main() {
    int pass;
    int n;

    scanf("%d", &n);

    pass = rand() % (int) pow(10, n);

    printf("%0*d\n", n, pass);
}
