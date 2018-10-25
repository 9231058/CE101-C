/*
 * tester.cpp
 */

#include <iostream>
#include <fstream>

using namespace std;

// validate checks existence of r in [n, m] and r parity (r must be even).
// if r is valid this function returns 0 otherwise 1.
int validate(int n, int m, int r) {
    if (r < n || r > m) {
        return 1;
    }
    return r % 2;
}

int main(int argc, char const *argv[])
{

    // opens required files from given files
    ifstream test_in(argv[1]);    /* This stream reads from test's input file   */
    // ifstream test_out(argv[2]);   /* This stream reads from test's output file  */
    ifstream user_out(argv[3]);   /* This stream reads from user's output file  */

    int r1, r2, r3, r4;
    int m, n;

    test_in >> n >> m;
    user_out >> r1 >> r2 >> r3 >> r4;

    // 0 means user passes the test and 1 otherwise
    return validate(n, m, r1) || validate(n, m, r2) || validate(n, m, r3) || validate(n, m, r4);
}
