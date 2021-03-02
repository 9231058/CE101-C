#include <stdio.h>
#include <stdlib.h>

int search(const int *arr, int size, int num) {
    for (int i = 0; i < size; ++i) {
        if (arr[i] == num)
            return 1;
    }
    return 0;
}

int * intersection(const int *a, int size_a, const int *b, int size_b, int *size_result) {
    int min_size = size_a < size_b? size_a: size_b;
    int *res = (int *) calloc(sizeof(int), min_size);
    int idx = 0;
    for (int i = 0; i < size_a; ++i) {
        if (search(b, size_b, a[i]))
            res[idx++] = a[i];
    }
    *size_result = idx;
    return res;
}

int main(){
    int m, n, size, idx;
    scanf("%d %d", &m, &n);
    int a[m], b[n];
    for (int i = 0; i < m; ++i) {
        scanf("%d", a + i);
    }
    for (int i = 0; i < n; ++i) {
        scanf("%d", b + i);
    }
    int *my_intersection = intersection(a, m, b, n, &size);
    int out_size = m + n - size;
    int *res = (int *) calloc(sizeof(int), out_size);
    for (idx = 0; idx < m; ++idx) {
        res[idx] = a[idx];
    }
    for (int i = 0; i < n; ++i) {
        if (!search(my_intersection, size, b[i]))
            res[idx++] = b[i];
    }
    printf("%d", res[0]);
    for (int i = 1; i < out_size; ++i) {
        printf(" %d", res[i]);
    }
    free(my_intersection);
    free(res);
}
