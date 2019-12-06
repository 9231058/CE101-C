#include <stdio.h>

int main(int argc, const char *argv[]) {
    int n;
    scanf("%d", &n);

    int D[n]; // or D = malloc(n * sizeof(int));

    D[0] = 1; // f(0) = 1, fibonacci sequence number is 1
    D[1] = 1; // f(1) = 1, fibonnaci sequence number is 2
    for (int i = 2; i < n; i++) {
        D[i] = D[i - 1] + D[i - 2];
    }
    printf("%d\n", D[n - 1]);
}
