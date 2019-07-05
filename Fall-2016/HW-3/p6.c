/*
 * In The Name Of God
 * ========================================
 * [] File Name : p6.c
 *
 * [] Creation Date : 08-11-2016
 *
 * [] Created By : Parham Alvani (parham.alvani@gmail.com)
 * =======================================
*/
/*
 * Copyright (c) 2016 Parham Alvani.
*/
#include <stdio.h>

int main(int argc, char *argv[])
{
	int n;
	int a, b, c, d, e;

	scanf("%d", &n);

	if (n < 10000 || n > 99999) {
		printf("Error\n");
		return 0;
	}

	a = n % 10;
	n /= 10;
	b = n % 10;
	n /= 10;
	c = n % 10;
	n /= 10;
	d = n % 10;
	n /= 10;
	e = n % 10;

	if (a == e && b == d)
		printf("True\n");
	else
		printf("False\n");
}
