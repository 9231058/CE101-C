/* tester.cpp
 */
#include<bits/stdc++.h>
#include <fstream>
#include<stdlib.h>
#include<math.h>
using namespace std;
int main(int argc, char const *argv[])
{
        // opens required files from given files
    ifstream test_in(argv[1]);    /* This stream reads from test's input file   */
    // ifstream test_out(argv[2]);   /* This stream reads from test's output file  */
    ifstream user_out(argv[3]);   /* This stream reads from user's output file  */


    int n;
    test_in>>n;



    int start = pow(10,n-1);
    int stop = pow(10,n);
    bool c1,c2,c3,c4;
    c1=c2=c3=false;
    
    if (n==1){ char ans[10]; user_out>>ans;  if(strcmp(ans,"no corona"))return 1; else return 0;}

    int ans;
    user_out>>ans;
    int chek=ans;
    if (ans<stop,ans>start)c1=true;
    if (ans>0)c2=true;
    c3=true;
        while(ans>0){
            int dig=ans%10;
            if ( dig ==0){c3=false;break;}
            if (chek%dig == 0 ){c3=false;break;}
            ans/=10;
        }

    return (!c1 || !c2 || !c3) ;


    }
