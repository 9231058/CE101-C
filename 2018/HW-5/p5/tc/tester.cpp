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

    string user_prompt;

    int vector1_size;
    int vector2_size;
    test_in >> vector1_size >> vector2_size;

    // read vectors size
    getline(user_out, user_prompt);
    if (user_prompt != "Enter size1 and size2")
        return 1;

    int vector1[vector1_size];
    int vector2[vector2_size];

    // read vector 1
    getline(user_out, user_prompt);
    if (user_prompt != "Enter vector 1")
        return 1;

    for (int i = 0; i < vector1_size; i++) {
        test_in >> vector1[i];
    }

    // read vector 2
    getline(user_out, user_prompt);
    if (user_prompt != "Enter vector 2")
        return 1;


    for (int i = 0; i < vector2_size; i++) {
        test_in >> vector2[i];
    }

menu:
    // reads user's menu
    // each menu on its line
    string menu_1;
    string menu_2;
    string menu_3;
    string menu_4;
    string menu_5;
    string menu_6;
    string menu_7;
    string menu_8;

    getline(user_out, menu_1);
    getline(user_out, menu_2);
    getline(user_out, menu_3);
    getline(user_out, menu_4);
    getline(user_out, menu_5);
    getline(user_out, menu_6);
    getline(user_out, menu_7);
    getline(user_out, menu_8);

    /*
     * 1) Add
     * 2) Sub
     * 3) Multiply
     * 4) Max
     * 5) Min
     * 6) Change
     * 7) Print
     * 8) Exit
     */
    if (menu_1 != "1) Add") {
        return 1;
    }
    if (menu_2 != "2) Sub") {
        return 1;
    }
    if (menu_3 != "3) Multiply") {
        return 1;
    }
    if (menu_4 != "4) Max") {
        return 1;
    }
    if (menu_5 != "5) Min") {
        return 1;
    }
    if (menu_6 != "6) Change") {
        return 1;
    }
    if (menu_7 != "7) Print") {
        return 1;
    }
    if (menu_8 != "8) Exit") {
        return 1;
    }

    int choice;
    test_in >> choice;
    int chosen_vector;

    switch (choice) {
        case 1: // add
            if (vector1_size != vector2_size) {
                getline(user_out, user_prompt);
                if (user_prompt != "Invalid Operation")
                    return 1;
            } else {
                for (int i = 0; i < vector1_size; i++) {
                    int result;
                    user_out >> result;
                    if (result != vector1[i] + vector2[i]) {
                        return 1;
                    }
                }
                user_out.ignore(numeric_limits<streamsize>::max(), '\n'); // consumes all new-line characters
            }
            goto menu;
            break;
        case 2: // subtract
            if (vector1_size != vector2_size) {
                getline(user_out, user_prompt);
                if (user_prompt != "Invalid Operation")
                    return 1;
            } else {
                for (int i = 0; i < vector1_size; i++) {
                    int result;
                    user_out >> result;
                    if (result != vector1[i] - vector2[i]) {
                        return 1;
                    }
                }
                user_out.ignore(numeric_limits<streamsize>::max(), '\n'); // consumes all new-line characters
            }
            goto menu;
            break;
        case 3: // inner product
            if (vector1_size != vector2_size) {
                getline(user_out, user_prompt);
                if (user_prompt != "Invalid Operation")
                    return 1;
            } else {
                int result = 0;
                for (int i = 0; i < vector1_size; i++) {
                    result += vector1[i] * vector2[i];
                }
                int user_result;
                user_out >> user_result;
                if (result != user_result) {
                    return 1;
                }
                user_out.ignore(numeric_limits<streamsize>::max(), '\n'); // consumes all new-line characters
            }
            goto menu;
            break;
        case 4: // min
            getline(user_out, user_prompt);
            if (user_prompt != "Choose vector")
                return 1;
            test_in >> chosen_vector;
            int min_index;
            user_out >> min_index;
            switch (chosen_vector) {
                case 1:
                    for (int i = 0; i < vector1_size; i++) {
                        if (vector1[min_index] > vector1[i]) {
                            return 1;
                        }
                    }
                    break;
                case 2:
                    for (int i = 0; i < vector2_size; i++) {
                        if (vector2[min_index] > vector2[i]) {
                            return 1;
                        }
                    }
                    break;
                default:
                    return 1;
            }
            user_out.ignore(numeric_limits<streamsize>::max(), '\n'); // consumes all new-line characters
            goto menu;
            break;
        case 5: // max
            getline(user_out, user_prompt);
            if (user_prompt != "Choose vector")
                return 1;
            test_in >> chosen_vector;
            int max_index;
            user_out >> max_index;
            switch (chosen_vector) {
                case 1:
                    for (int i = 0; i < vector1_size; i++) {
                        if (vector1[min_index] > vector1[i]) {
                            return 1;
                        }
                    }
                    break;
                case 2:
                    for (int i = 0; i < vector2_size; i++) {
                        if (vector2[min_index] > vector2[i]) {
                            return 1;
                        }
                    }
                    break;
                default:
                    return 1;
            }
            user_out.ignore(numeric_limits<streamsize>::max(), '\n'); // consumes all new-line characters
            goto menu;
            break;
        case 6: // change
            getline(user_out, user_prompt);
            if (user_prompt != "Choose vector")
                return 1;
            test_in >> chosen_vector;
            getline(user_out, user_prompt);
            switch (chosen_vector) {
                case 1:
                    if (user_prompt != "Enter new vector 1")
                        return 1;
                    for (int i = 0; i < vector1_size; i++) {
                        int element;
                        test_in >> element;
                        vector1[i] = element;
                    }
                case 2:
                    if (user_prompt != "Enter new vector 2")
                        return 1;
                    for (int i = 0; i < vector2_size; i++) {
                        int element;
                        test_in >> element;
                        vector2[i] = element;
                    }
                default:
                    return 1;
            }
            goto menu;
            break;
        case 7: // print
            getline(user_out, user_prompt);
            if (user_prompt != "Choose vector")
                return 1;
            test_in >> chosen_vector;
            switch (chosen_vector) {
                case 1:
                    for (int i = 0; i < vector1_size; i++) {
                        int element;
                        user_out >> element;
                        if (element != vector1[i]) {
                            return 1;
                        }
                    }
                case 2:
                    for (int i = 0; i < vector2_size; i++) {
                        int element;
                        user_out >> element;
                        if (element != vector2[i]) {
                            return 1;
                        }
                    }
                default:
                    return 1;
            }
            user_out.ignore(numeric_limits<streamsize>::max(), '\n'); // consumes all new-line characters
            goto menu;
            break;
        case 8: // exit
            break;
        default:
            return 1;
    }

    return 0;
}
