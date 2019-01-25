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

# Create input files
echo -e "This is my 1 file\nfrom London to L.A." > "fi.txt"
echo -e "THIS IS MY one FILE\nFROM LONDON TO L.A." > "fa.txt"

# Prepare results file
echo "" > results.txt

# Compile & Run
for student in `ls .`; do
	if [ ! -d $student ]; then
		continue
	fi
	echo "- $student"
	echo
	for file in `ls $student/*.c`; do
		# Create stdin arguments
		echo -e "fi.txt\n${student}_fo.txt" > "in.txt"

		# Compile
		gcc $file &> "${student}_log.txt"
		if [ ! $? -eq 0 ]; then
			echo "$student: Compilation failure" >> results.txt
			continue
		else
			# Remove old outputs if exists
			rm -f "${stduent}_fo.txt"

			./a.out < "in.txt" &> /dev/null
			if [ ! $? -eq 0 ]; then
				echo "$student: Run failure" >> results.txt
				continue
			else
				cat "${student}_fo.txt"
				diff -q "${student}_fo.txt" "fa.txt"
				if [ ! $? -eq 0 ]; then
					echo "$student: Wrong Answer" >> results.txt
					echo "$student: Two column diff:" >> results.txt
					diff -y "${student}_fo.txt" "fa.txt" >> results.txt
					echo "" >> results.txt
				else
					echo "$student: Accepted" >> results.txt
				fi
			fi
		fi
	done
	rm -f a.out

	echo
	echo "- $student"
done
