#!/bin/bash

printf "Введите имя каталога: "
read dirname
printf "Содержимое каталога %s: \n" $dirname
ls $dirname
