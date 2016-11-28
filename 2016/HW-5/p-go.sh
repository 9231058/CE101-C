#!/bin/bash
# In The Name Of God
# ========================================
# [] File Name : p-go.sh
#
# [] Creation Date : 04-11-2015
#
# [] Created By : Elahe Jalalpour (el.jalalpour@gmail.com)
# =======================================


counter=1
for i in "${1}/in/"*.txt; do
	output=$(printf "output%d.txt" ${counter})
	go run ${1}.go < "${i}" > "${1}/out/${output}"
	echo "Problem ${1} - Test ${counter}: ${i} --> ${output}"
	let counter=counter+1
done

cd ${1}; zip -r "../${1}" *; cd ..
echo "Zip file created for quera..."
