/*
 * In The Name Of God
 * ========================================
 * [] File Name : p3.c
 *
 * [] Creation Date : 15-11-2016
 *
 * [] Created By : Parham Alvani (parham.alvani@gmail.com)
 * =======================================
*/
/*
 * Copyright (c) 2016 Parham Alvani.
*/
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int main(int argc, char *args[])
{
	int n;
	int i;
	int m = 0;

	scanf("%d", &n);

	srand(time(NULL));
	
	for (i = 0; i < n; i++) {
		long long x, y;
		x = rand();
		y = rand();
		if (x * x + y * y <= (long long) RAND_MAX * RAND_MAX)
			m++;
	}
	
	double p = (double) m / n;
	printf("%.4lf\n", 4 * p);
}
