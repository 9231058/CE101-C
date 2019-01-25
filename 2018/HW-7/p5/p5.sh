#!/bin/bash
# In The Name Of God
# ========================================
# [] File Name : p5.sh
#
# [] Creation Date : 18-01-2018
#
# [] Created By : Parham Alvani <parham.alvani@gmail.com>
# =======================================

# This script is created based on Quera's
# per question user submissions format so it assumes
# following format:
# /
# -- 9231058/hello.c
# -- 9131036/main.c
# -- 9331032/iman.c

# Prepares the results file
echo "" > results.txt

# Prepares the input generator
if [ ! -f p5.c ]; then
        echo "Please add p5.c to generate the input"
        exit
fi
gcc p5.c -o p5.out

# Compiles and Runs
for student in `ls .`; do
	if [ ! -d $student ]; then
		continue
	fi
	echo "- $student"
	echo
	for file in `ls $student/*.c`; do
		# Create stdin arguments
                # please note that `fi.bin` is created by p5.c
                ./p5.out
		echo -e "fi.bin" > "in.txt"

		# Compile
		gcc $file &> "${student}_log.txt"
		if [ ! $? -eq 0 ]; then
			echo "$student: Compilation failure" >> results.txt
			continue
		else
			# Remove old outputs if exists
			rm -f "${stduent}_fo.txt"

			./a.out < "in.txt" > "${student}_fo.txt" 2> /dev/null
			if [ ! $? -eq 0 ]; then
				echo "$student: Run failure" >> results.txt
				continue
			else
				echo "" >> results.txt
				echo "--------" >> results.txt
				echo "" >> results.txt
				echo "$student: Answer is:" >> results.txt
				cat "${student}_fo.txt" >> results.txt
				echo "" >> results.txt
				echo "--------" >> results.txt
				echo "" >> results.txt
			fi
		fi
	done
	rm -f a.out

	echo
	echo "- $student"
done
