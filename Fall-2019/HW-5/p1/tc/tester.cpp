/*
 * tester.cpp
 */

#include <iostream>
#include <fstream>
#include <cmath>

using namespace std;

int main(int argc, char const *argv[])
{

    // opens required files from given files
    // ifstream test_in(argv[1]);    /* This stream reads from test's input file   */
    ifstream test_out(argv[2]);   /* This stream reads from test's output file  */
    ifstream user_out(argv[3]);   /* This stream reads from user's output file  */

    int a, b, c;
    int x, y, z;
    test_out >> a >> b >> c;
    user_out >> x >> y >> z;

    if (x > y) swap(x, y);
    if (y > z) swap(y, z);
    if (x > y) swap(x, y);

    if (a > b) swap(a, b);
    if (b > c) swap(b, c);
    if (a > b) swap(a, b);

    if (a == x && b == y && c == z) {
        return 0;
    }
    return 1;
}
