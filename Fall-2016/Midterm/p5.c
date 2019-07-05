/*
 * In The Name Of God
 * ========================================
 * [] File Name : p5.c
 *
 * [] Creation Date : 01-12-2016
 *
 * [] Created By : Parham Alvani (parham.alvani@gmail.com)
 * =======================================
*/
/*
 * Copyright (c) 2016 Parham Alvani.
*/
#define TRUE 1
#define FALSE 0

int overlap(int a1, int b1, int a2, int b2);

int cover(int a1, int b1, int a2, int b2, int a3, int b3)
{
	/* [a1, b1] -> [a2. b2] <- [a3, b3] */
	if (overlap(a1, b1, a2, b2) && overlap(a3, b3, a2, b2)) {
		return TRUE;
	}
	/* [a2, b2] -> [a1. b1] <- [a3, b3] */
	if (overlap(a2, b2, a1, b1) && overlap(a3, b3, a1, b1)) {
		return TRUE;
	}
	/* [a1, b1] -> [a3. b3] <- [a2, b2] */
	if (overlap(a1, b1, a3, b3) && overlap(a2, b2, a3, b3)) {
		return TRUE;
	}

	return FALSE;
}
