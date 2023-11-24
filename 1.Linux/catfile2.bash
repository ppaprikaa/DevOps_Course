#!/bin/bash
printf "Введите имя файла: "
read filename
if [ ! -f $filename ]
then
   printf "Простите, но файла %s не существует\n" $filename
   exit 1
fi

printf "Содержимое файла %s: \n" $filename
cat $filename
