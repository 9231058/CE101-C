/*
 * In The Name Of God
 * ========================================
 * [] File Name : p3.c
 *
 * [] Creation Date : 29-11-2016
 *
 * [] Created By : Parham Alvani (parham.alvani@gmail.com)
 * =======================================
 */
/*
 * Copyright (c) 2016 Parham Alvani.
 */
#include <stdio.h>
#include <stdlib.h>

int start = 0, mid = 20, end = 40;

void print_abc(int ring)
{
	int i;

	printf("A");
	for (i = 1; i < ring + 3; i++)
		printf(" ");
	printf("B");
	for (i = 1; i < ring + 3; i++)
		printf(" ");
	printf("C");
	printf("\n\n");
}

void print_star(int n, int a[60][20])
{
	int i, j;

	for (j = 0; j < n; j++) {
		for (i = 0; i < n; i++) {
			if (a[i][j] == 0)
				printf(" ");
			else
				printf("*");

		}
		for (i = 0; i < 3; i++)
			printf(" ");
		for (i = n; i < 2 * n; i++) {
			if (a[i][j] == 0)
				printf(" ");
			else
				printf("*");

		}
		for (i = 0; i < 3; i++)
			printf(" ");
		for (i = 2 * n; i < 3 * n; i++) {
			if (a[i][j] == 0)
				printf(" ");
			else
				printf("*");

		}
		printf("\n");
	}
	print_abc(n);

}

void move(int n, int ring, int start, int end, int a[60][20])
{

	int i, m, j;

	i = (start * ring) / 20;
	j = 0;
	while (a[i][j] == 0)
		j++;
	for (m = 0; m < n; m++)
		a[i + m][j] = 0;
	i = (end * ring) / 20;
	j = ring - 1;
	while (a[i][j] == 1)
		j--;
	for (m = 0; m < n; m++)
		a[i + m][j] = 1;
	print_star(ring, a);
}

void hanoi(int n, int ring, int start, int mid, int end, int a[60][20])
{
	if (n == 1) {
		move(n, ring, start, end, a);
	} else {
		hanoi(n - 1, ring, start, end, mid, a);
		move(n, ring, start, end, a);
		hanoi(n - 1, ring, mid, start, end, a);

	}

}

int main(int argc, char *argv[])
{
	int n, i, j;
	int a[60][20];

	scanf("%d", &n);

	for (i = 0; i < 3 * n; i++) {
		for (j = 0; j < n; j++)
			a[i][j] = 0;
	}
	for (i = 0; i < n; i++) {
		for (j = 0; j <= i; j++) {
			printf("*");
			a[j][i] = 1;
		}
		printf("\n");
	}
	int ring = n;

	print_abc(ring);
	hanoi(n, ring, start, mid, end, a);
	return 0;
}

