#!/bin/bash

printf "Введите имя файла: "
read filename
printf "Содержимое файла %s: \n" $filename
cat $filename
