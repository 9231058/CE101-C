/*
 * In The Name Of God
 * ========================================
 * [] File Name : p2.c
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
	int a, b;
	int c, d;

	scanf("%d %d %d %d", &a, &b, &c, &d);

	if (a > b) {
		b = a + b;
		a = b - a;
		b = b - a;
	}
	if (c > d) {
		d = c + d;
		c = d - c;
		d = d - c;
	}

	printf("%d %d\n", c - b, d - a);
}
