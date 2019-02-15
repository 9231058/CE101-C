#include <stdio.h>

// stores fibonacci sequences in the array.
// global variables are initialized to zero in c
int memory[1000];

int fibonacci(int n) {
    if (n == 1 || n == 2) {
        return 1;
    }
    if (memory[n] == 0) {
        memory[n] = fibonacci(n - 1) + fibonacci(n - 2); // memorizes what it calculated now.
    }
    return memory[n];
}