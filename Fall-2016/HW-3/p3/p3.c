#include <stdio.h>

int main(int argc, char *argv[])
{
	int hw[4];
	int i;
	const double hw_p[4] = {0.25, 0.15, 0.35, 0.25};
	double total = 0;

	scanf("%d", hw);
	scanf("%d", hw + 1);
	scanf("%d", hw + 2);
	scanf("%d", hw + 3);

	for (i = 0; i < 4; i++)
		total += (hw[i] * hw_p[i]);

	printf("%.2lf\n", total);

	for (i = 0; i < 4; i++)
		printf("%.2lf\n", hw[i] * hw_p[i]);
}
