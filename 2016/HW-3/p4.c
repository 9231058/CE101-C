#include <stdio.h>
#include <math.h>

int main(int argc, char *argv[])
{
	double edge1, edge2, angle, radian, area;
	
	scanf("%lf", &edge1);
	scanf("%lf", &edge2);
	scanf("%lf", &angle);

	if (angle == 90)
		if (edge1 == edge2)
			printf("This is square\n");
		else
			printf("This is rectangle\n");
	else
		if (edge1 == edge2)
			printf("This is diamond\n");
		else
			printf("This is parallelogram\n");
	radian = angle * M_PI / 180;
	area = edge1 * edge2 * sin(radian);
	printf("%.2lf\n", area);
}
