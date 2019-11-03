/*
 * tester.cpp
 */

#include <iostream>
#include <fstream>

using namespace std;

// validate checks existence of r in [0, n] and r parity (r must be even).
// if r is valid this function returns 0 otherwise 1.
int validate(int n, int r) {
    if (r > n || r < 0) {
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
    int n;

    test_in >> n;
    user_out >> r1 >> r2 >> r3 >> r4;

    // 0 means user passes the test and 1 otherwise
    return validate(n, r1) || validate(n, r2) || validate(n, r3) || validate(n, r4);
}
