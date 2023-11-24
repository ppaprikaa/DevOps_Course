#!/bin/bash
logs_w_errs=($(grep --dereference-recursive --files-with-matches --no-messages error /var/log))
len=${#logs_w_errs[@]}
if [ $len -lt 1 ]
then
   printf "Простите, но ошибочек то нет, оказывается...\n"
   exit 1
fi

printf "Файлики в которых встретились логи уровня error:\n"
for (( i=0; i<${len}; i++ ));
do
   echo "${logs_w_errs[$i]}"
done
