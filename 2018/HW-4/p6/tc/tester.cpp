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
    int digit;
    test_in >> number >> digit;

new_game:
    for (int run = 0; run < 5; run++) {
        int corrects = 0;
        bool state[number];

        // reads user generated numbers and compare them with test case numbers
        for (int i = 0; i < number; i++) {
            int user_number;
            int test_number;

            user_out >> user_number;
            test_in >> test_number;

            if (user_number == test_number) {
                state[i] = true;
                corrects++;
            } else {
                state[i] = false;
            }
        }

        // reads user generated status
        // correct: Correct :) :D
        // incorrect: Incorrect :( :P
        user_out.ignore(numeric_limits<streamsize>::max(), '\n'); // consumes all new-line characters
        for (int i = 0; i < number; i++) {
            string user_status;
            getline(user_out, user_status);
            if (state[i] && user_status != "Correct :) :D") {
                return 1;
            }
            if (!state[i] && user_status != "Incorrect :( :P") {
                return 1;
            }
        }

        double user_ratio;
        user_out >> user_ratio;

        double test_ratio = (double) corrects / number;

        if (abs(user_ratio - test_ratio) > err) {
            return 1;
        }
    }

    user_out.ignore(numeric_limits<streamsize>::max(), '\n'); // consumes all new-line characters
    // reads user's menu
    // each menu on its line
    string menu_1;
    string menu_2;
    string menu_3;
    string menu_4;

    getline(user_out, menu_1);
    getline(user_out, menu_2);
    getline(user_out, menu_3);
    getline(user_out, menu_4);

    /*
     * 1) Continue
     * 2) Increase numbers
     * 3) Increase digits
     * 4) End
     */
    if (menu_1 != "1)Continue") {
        return 1;
    }
    if (menu_2 != "2)Increase numbers") {
        return 1;
    }
    if (menu_3 != "3)Increase digits") {
        return 1;
    }
    if (menu_4 != "4)End") {
        return 1;
    }

    int choice;
    test_in >> choice;

    switch (choice) {
        case 1:
            goto new_game;
            break;
        case 2:
            number++;
            goto new_game;
            break;
        case 3:
            digit++;
            goto new_game;
            break;
        case 4:
            break;
        default:
            return 1;
    }

    return 0;
}
