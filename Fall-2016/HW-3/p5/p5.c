#include <stdio.h>

int main(int argc, char *argv[])
{
	int x[3];
	int y[3];
	int answer[6];
	int i = 0;
	int j = 0;
	int k = 0;

	scanf("%d %d %d", x, (x + 1), (x + 2));
	scanf("%d %d %d", y, (y + 1), (y + 2));

	while (i < 3 && j < 3) {
		if (x[i] < y[j]) {
			answer[k] = x[i];
			i++;
			k++;
		} else {
			answer[k] = y[j];
			j++;
			k++;
		}
	}
	for (; i < 3; i++) {
		answer[k] = x[i];
		k++;
	}
	for (; j < 3; j++) {
		answer[k] = y[j];
		k++;
	}

	for (i = 5; i > 0; i--)
		printf("%d ", answer[i]);
	printf("%d\n", answer[0]);

}
