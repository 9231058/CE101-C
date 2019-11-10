#!/bin/bash
# In The Name Of God
# ========================================
# [] File Name : p.sh
#
# [] Creation Date : 27-10-2016
#
# [] Updated By : Parham Alvani (parham.alvani@gmail.com)
#
# [] Created By : Elahe Jalalpour (el.jalalpour@gmail.com)
# =======================================
current_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

die() {
        echo $@
        exit
}

usage() {
        echo "this script create the zip folder for the given problem that can be uploaded to Quera website"
        echo "$0 [year] [semester] [hw-?] [p?]"
}

if [ $# -ne 4 ]; then
        usage
        exit
fi

echo $1 | grep -E -q '^[0-9]{4}$' || die "4-digit year required, $1 provided"
echo $2 | grep -E -q 'Fall|Spring' || die "semester must be 'Fall' or 'Spring', $2 provided"

year=$1
semester=$2
hw="HW-$3"
p="p$4"
tc="$current_dir/../$semester-$year/$hw/$p/tc"    # problem test cases directory
go="$current_dir/../$semester-$year/$hw/$p/$p.go" # problem solution in Go

if ! [ -d $tc ]; then
        echo "there is no test case folder in $tc"
        exit
fi

# if there is a tester.cpp in test cases folder there is no need to generate
# output files
if ! [ -f $tc/tester.cpp ]; then
        # if there is a generator.go in test cases folder first generates inputs based on it.
        if [ -f $tc/generator.go ]; then
                OIFS=$IFS
                IFS=';' # split output based on ';' instead of newline
                i=0
                echo "[beta] generates test case based on generator.go"

                start=$(date +'%s')

                for input in $(go run $tc/generator.go); do
                        i=$(( i + 1 ))
                        echo $input > "$tc/in/input$i.txt"

                        echo "case $i: $tc/in/input$i.txt"
                        echo
                done
                IFS=$OIFS

                took=$(( $(date +'%s') - $start ))
                printf "$tc/generator.go took %ds.\n" $took
        fi

        if ! [ -f $go ]; then
                echo "there is no go source in $go"
                echo "try to generate zip based on current status of tc folder"
                if ! [ -d $tc/in ] || ! [ -d $tc/out ]; then
                        exit
                fi
        else
                for input in "$tc/in/"*.txt; do
                        i=$(basename $input)
                        i=${i#input}
                        i=${i%.txt}

                        start=$(date +'%s')

                        go run $go < $input > "$tc/out/output$i.txt"

                        took=$(( $(date +'%s') - $start ))
                        printf "$go took %ds.\n" $took

                        echo "test $i: $input --> $tc/out/output$i.txt"
                        echo
                done
        fi
else
        echo "you are using Quera tester.cpp, please make sure you know what you are doing"
fi

cd $tc && zip -r "$current_dir/$hw-$p.zip" *; cd -
echo "quera TC zip is ready"
