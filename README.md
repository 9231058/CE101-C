# C101
## Introduction
Feel free to visit and use some sample codes and homeworks problem source codes in C and GoLang for the
Introduction to Programming course at the Amirkabir University of Tech.
under supervision of [Prof. Bakhshi](http://ceit.aut.ac.ir/~bakhshis/).

For homework problems and more materials about the course see [here](http://ceit.aut.ac.ir/~bakhshis/c).

**Thanks to [Quera](https://quera.ir/) for supporting our course!**

## Trello boards
- [Fall-2017](https://trello.com/b/2HlMa6yF)
- [Fall-2018](https://trello.com/b/HGB2XpUD)

## Requirements
Source codes of Homework are written in [Go](https://golang.org/) and C.
Documentations are in Word and needs [Sahel](https://github.com/rastikerdar/sahel-font) and [Vazir](https://github.com/rastikerdar/vazir-font) fonts.

## Problem Structure
Each problem have the following structure:

```
HW-1\
    |
    |- p1\
    |    |
    |    |- p1.md
    |    |- p1.go
    |    |- tc\
    |    |     |- in
    |    |     |- out
    |
    |- p2\
```

`p.sh` creates the zip file that can be uplaoded to Quera website based on the above structure. It feeds the inputs
from `in` folder into `p.go` then puts the results into `out` folder.

If there is a file with `generator.go` name in `tc` folder the `p.sh` use it to generate the test cases into
`in` folder.

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("%d %d;", 1, 10)
	fmt.Printf("%d %d;", 5, 10)

	for i := 0; i < 20; i++ {
		m := rand.Intn(100)
		n := rand.Intn(m)

		fmt.Printf("%d %d;", n, m)
	}
}
```

## Topics
- Lecture 3: Basic C Programming
- Lecture 4: Calculations
- Lecture 5: Interaction
- Lecture 6: Making Decisions
- Lecture 7: Repeating Statements
- Lecture 8: Functions
- Lecture 9: Arrays
- Lecture 10: Pointers & Dynamic Memory
- Lecture 11: Structures
- Lecture 12: Files
- Lecture 13: Miscellaneous

## Assignments
- HW1: Algorithm Design

- HW2: Caculation (with an introduction to `math.h`) and Interaction (`scanf()` and `printf()`)

- HW3: Decision Making

- HW4: Loops

- HW5: Functions

- HW6: Mid-Term Project

- HW7: Pointer, Structs and Files

  

**Exercises are categorized by topic and semester order in this** [link](https://docs.google.com/spreadsheets/d/1SAPncbl116FTfDpmxd_Duw4hvQ280_bsEuRlWDDcNFU/edit#gid=1228963258)

The difficulty of each question from 10 and terms that contain the question are included in the sheet along with their link.

## Workshops
- Introduction-1
- Introduction-2
- Algorithm
- Algorithm
- I/O, Calculation
- Decision
- Loop
- Midterm Review
- Git
- Arrays & DP
- Pointers
- Pointers and Arguments
- Linked Lists & Structs

## Teaching Assistants over Semesters
### Fall-2014
* [Shirin Shirazi](https://ir.linkedin.com/in/shirin-ha-shirazi)
* [Bita Azari](http://ceit.aut.ac.ir/~azari/)
* [Ahmad Asadi](https://github.com/ahmad-asadi)
* [Parham Alvani](https://github.com/1995parham)

### Fall-2015
* [Elahe Jalalpour](https://github.com/elahejalalpour)
* [Shiva Zamani](https://github.com/shiva-z)
* [Iman Tabrizian](https://github.com/Tabrizian)
* [Parham Alvani](https://github.com/1995parham)

### Fall-2016
* [Mahtab Farrokh](https://github.com/mahtabfarrokh)
* [Ramtin Shakeri](https://github.com/RamtinSh7596)
* [Esmail Naderi]()
* [Amirmohammad Haghollahi](https://github.com/AMIRmh)
* [Parham Alvani](https://github.com/1995parham)

### Fall-2017
* [Paya Faghani](https://github.com/pfaghani)
* [Ehsan Souri](https://github.com/ehsansouri23)
* [Parsa Eskandarnejad](https://github.com/parsaaes)
* [Soroush Barmakie](https://github.com/sbarmak1377)
* [Parham Alvani](https://github.com/1995parham)

### Fall-2018
* [Parsa Eskandarnejad](https://github.com/parsaaes)
* [Sepehr Akhoundi](https://github.com/Sepehr1812)
* [Amirhosein Khoshbin](https://github.com/AOptimist)
* [Muhammad Azhdari](https://github.com/mmdaz)
* [Mohsen Mohammadi](https://github.com/MrMiM77)
* [Mohammad Sami](https://github.com/MohammadMDSA)
* [AliAkbar Badri](https://github.com/aabadri) (Lab)
* [Amirhossein Bavand](https://github.com/ahbavand) (Lab)
* [Saeid Dadkhah](https://github.com/SaeidDadkhah) (Lab)
* [Mahdi Aghajani](https://github.com/mmaghajani) (Lab)
* [Parham Alvani](https://github.com/1995parham)

### Spring-2019
* [MohamadHasan Taghadosi](https://github.com/taghad) (Lab)
* [Paya Faghani](https://github.com/pfaghani)
* [Mohsen Mohammadi](https://github.com/MrMiM77)
* [Roya Taheri](https://github.com/RoyaTaheri) (Lab)
* [Parham Alvani](https://github.com/1995parham)

### Fall-2019

* [Amin Rashidbeigi](https://github.com/aminrashidbeigi) (Lab)
* [Mohammad Mashayekh]() (Lab)
* [MohamadHasan Taghadosi](https://github.com/taghad) (Lab)
* [Mahvash Siavashpour](https://github.com/mahvash-siavashpour) (Lab)
* [Mohsen Mohammadi](https://github.com/MrMiM77)
* [Mohamad Fatemi](https://github.com/smf8)
* [Saman Hoseini](https://github.com/saman2000hoseini)
* [Mahin Mirshams](https://github.com/mahinmirshams)
* [Mohammad Javad Ardestani](https://github.com/mohammadjavadArdestani) (Lab)
* [Niki Pourazin](https://github.com/npourazin)
* [Parham Alvani](https://github.com/1995parham) (Lab)

### Spring-2020 

* [Kian Kashfipour](https://github.com/kian79)
* [Mohammad Fatemi](https://github.com/smf8)
* [MohamadHasan Taghadosi](https://github.com/taghad) (Lab)
* [Roya Taheri](https://github.com/RoyaTaheri) (Lab)
* [Saman Hoseini](https://github.com/saman2000hoseini)
* [Maryam Alikarami](https://github.com/malikarami)
* [Negin Haji Sobhani](https://github.com/neginhsobhani)
* [Darya Zare](https://github.com/DaryaZareM)
* [Saeed Maroof](https://gitlab.com/saeed.maroof)
* [Bahar Kaviani](https://github.com/Baharkaviani)
* [Parham Alvani](https://github.com/1995parham) (Course instructor)

### Fall-2020

- [Mohammad Fatemi](https://github.com/smf8)
- [Saman Hoseini](https://github.com/saman2000hoseini)
- [Parham Alvani](https://github.com/1995parham)
- [MohamadHasan Taghadosi](https://github.com/taghad)
- [Amirparsa Salmankhah](http://Github.com/Amirparsa-Sal)
- [Kasra Zarei](https://github.com/kasrazarei39)
- [Mina Beiki](https://github.com/mina-beiki)
- [Ali Ansari](http://Github.com/aliaa80)
- [Mahla Sharifi](http://github.com/mahlashrifi)
- [Amirreza Hashemi](https://github.com/AmirSpurs)
- [Mahdi Rahmani](https://github.com/Mahdi-Rahmani)
- [Mohammad Salehi](https://github.com/m-salehi-v)
- [Bardia Ardakanian](https://github.com/bardia-ardakanian)
- [Amir Naziri](http://github.com/Amir79naziri)
- [Korosh Roohi](https://github.com/KoroshRH)
- [Mohammad Javad Rajabi](https://github.com/rajabi2001)
- [Parham Moonesi](https://github.com/Pmoonesi)
- [Farshid Nooshi](https://github.com/FarshidNooshi)
- [Ghazale Noroozi](https://github.com/GhazaleNoroozi)
- [Faraz Farangizade](https://github.com/farazff)
- [Amin Habibollah](https://github.com/aminhbl)

