#!/bin/bash
if [ $UID -eq 0 ]; then
   printf "You are root\n"
   exit 0
fi;

printf "You are not root. You are $(whoami)\n"
exit 1
