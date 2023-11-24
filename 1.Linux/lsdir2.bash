#!/bin/bash
printf "Введите имя каталога: "
read dirname
if [ ! -d $dirname ]
then
   printf "Простите, но каталога %s не существует\n" $dirname
   exit 1
fi

printf "Содержимое каталога %s: \n" $dirname
ls $dirname
