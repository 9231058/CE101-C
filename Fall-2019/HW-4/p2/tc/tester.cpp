/*
 * tester.cpp
 */

#include <iostream>
#include <fstream>
#include <limits>
#include <cmath>

using namespace std;

int main(int argc, char const *argv[])
{

    // opens required files from given files
    ifstream test_in(argv[1]);    /* This stream reads from test's input file   */
    // ifstream test_out(argv[2]);   /* This stream reads from test's output file  */
    ifstream user_out(argv[3]);   /* This stream reads from user's output file  */

    int count;
    test_in >> count;

    int zero;
    user_out >> zero;
    if (zero != 0) {
        return 1;
    }

    int last = 0;
    for (int i = 0; i < count; i++) {
        int current;
        user_out >> current;

        int diff = current - last;
        if (diff > 1 || diff < -1) {
            return 1;
        }
        last = current;
    }

    return 0;
}
