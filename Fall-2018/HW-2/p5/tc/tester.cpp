/*
 * tester.cpp
 */

#include <iostream>
#include <fstream>
#include <cmath>

using namespace std;

const double err = 0.1;

int main(int argc, char const *argv[])
{

    // opens required files from given files
    // ifstream test_in(argv[1]);    /* This stream reads from test's input file   */
    ifstream test_out(argv[2]);   /* This stream reads from test's output file  */
    ifstream user_out(argv[3]);   /* This stream reads from user's output file  */

    // perimeter validation
    string perimeter;
    double test_p, user_p;
    test_out >> perimeter >> test_p;
    user_out >> perimeter >> user_p;
    if (perimeter != "Perimeter:") {
        return 1;
    }
    if (abs(test_p - user_p) > err) {
        return 1;
    }

    // area validation
    string area;
    double test_a, user_a;
    test_out >> area >> test_a;
    user_out >> area >> user_a;
    if (area != "Area:") {
        return 1;
    }
    if (abs(test_a - user_a) > err) {
        return 1;
    }

    return 0;
}
