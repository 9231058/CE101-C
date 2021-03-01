#include <stdio.h>
#include <stdlib.h>

int main(){
    int n;
    scanf("%d", &n);
    int populations[n];
    for (int i = 0; i < n; ++i) {
        scanf("%d", populations + i);
    }
    long distances[n];
    for (int i = 0; i < n; ++i) {
        long distance = 0;
        for (int j = 0; j < n; ++j) {
            distance += populations[j] * abs(i - j);
        }
        distances[i] = distance;
    }
    int idx = 0;
    for (int i = 1; i < n; ++i) {
        if (distances[i] < distances[idx])
            idx = i;
    }
    printf("%d", idx + 1);
}
