#!/bin/bash

echo "Compiling..."
gcc -std=c99 -o ${1}.out -Wall ${1}.c
echo "Compilation was successful"

counter=1
for i in "${1}/in/"*.txt; do
	output=$(printf "output%d.txt" ${counter})
	./${1}.out < "${i}" > "${1}/out/${output}"
	echo "Problem ${1} - Test ${counter}: ${i} --> ${output}"
	let counter=counter+1
done

zip -r ${1} "${1}/"*
echo "Zip file created for quera..."
