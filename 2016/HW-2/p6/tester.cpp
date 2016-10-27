/*
 * tester.cpp
 *
*/

#include <iostream>
#include <fstream>
#include <string>

using namespace std;

int main(int argc, char const *argv[])
{

	ifstream test_in(argv[1]);    /* This stream reads from test's input file   */
	ifstream test_out(argv[2]);   /* This stream reads from test's output file  */
	ifstream user_out(argv[3]);   /* This stream reads from user's output file  */

	/* Your code here */
	/* If user's output is correct, return 0, otherwise return 1 */

	int n, m;

	test_in >> m >> n;

	for (int i = 0; i < 4; i++) {
		int v;

		user_out >> v;
		if (v > n || v < m)
			return 1;
	}

	return 0;
}
