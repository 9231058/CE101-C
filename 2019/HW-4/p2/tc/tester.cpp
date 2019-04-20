/*
 * tester.cpp
 */

#include <iostream>
#include <fstream>
#include <limits>
#include <cmath>

using namespace std;

const double err = 0.1;

int main(int argc, char const *argv[])
{

    // opens required files from given files
    ifstream test_in(argv[1]);    /* This stream reads from test's input file   */
    // ifstream test_out(argv[2]);   /* This stream reads from test's output file  */
    ifstream user_out(argv[3]);   /* This stream reads from user's output file  */

    int number;
    int count;
    test_in >> number >> count;

    for (int i = 0; i < count; i++) {
        bool *check = new bool[number];
        for (int j = 0; j < number; j++) {
            check[j] = false;
        }

        for (int j = 0; j < number; j++) {
            if (check[j] == false) {
                check[j] = true;
            } else {
                return 1;
            }
        }

        delete[] check;
    }

    return 0;
}
