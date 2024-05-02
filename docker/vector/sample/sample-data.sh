#!/bin/bash

names=("Yamada" "Suzuki" "Inoue" "Otani")
ages=(5 10 15 18 20 21 22 24 26 27 29 30 31 34 37 40)
addresses=("Tokyo" "Osaka" "Aich" "Hokkaido" "Fukuoka")

for name in ${names[@]}; do
    for age in ${ages[@]}; do
        for address in ${addresses[@]}; do
            echo "{name:"$name", age:"$age", address:"$address"}" >> test.log
        done
    done
done
