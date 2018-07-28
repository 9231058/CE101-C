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

usage() {
        echo "$0 [hw-?] [p?]"
}

if [ $# -ne 2 ]; then
        usage
        exit
fi

hw="HW-$1"
p="p$2"
tc="$current_dir/$hw/$p/tc"
go="$current_dir/$hw/$p/$p.go"

if ! [ -d $tc ]; then
        echo "There is no test case in $tc"
        exit
fi

if ! [ -f $go ]; then
        echo "There is no go source in $go"
        exit
fi

for input in "$tc/in/"*.txt; do
        i=$(basename $input)
        i=${i#input}
        i=${i%.txt}

        start=$(date +'%s')

        go run $go < $input > "$tc/out/output$i.txt"

        took=$(( $(date +'%s') - $start ))
        printf "$go Took %ds.\n" $took

        echo "Test $i: $input --> $tc/out/output$i.txt"
        echo
done

cd $tc && zip -r "$current_dir/$hw-$p.zip" *; cd -
echo "Quera TC zip is ready"
