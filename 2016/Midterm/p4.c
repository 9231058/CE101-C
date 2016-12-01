/*
 * In The Name Of God
 * ========================================
 * [] File Name : p4.c
 *
 * [] Creation Date : 01-12-2016
 *
 * [] Created By : Parham Alvani (parham.alvani@gmail.com)
 * =======================================
*/
/*
 * Copyright (c) 2016 Parham Alvani.
*/
#include <stdio.h>

int fibonacci(int n)
{
	if (n == 1 || n == 2)
		return 1;
	else
		return fibonacci(n - 1) + fibonacci(n - 2);
}

int fibonacci_check_and_generate(int n)
{
	int i = 1;

	while (fibonacci(i) < n)
		i++;

	if (fibonacci(i) == n)
		return fibonacci(i + 1);
	else
		return -1;
}

int main(int argc, char *argv[])
{
	int n;
	int retval;

	scanf("%d", &n);

	retval = fibonacci_check_and_generate(n);
	if (retval == -1)
		printf("Error !\n");
	else
		printf("%d %d\n", retval - n, retval);
}
