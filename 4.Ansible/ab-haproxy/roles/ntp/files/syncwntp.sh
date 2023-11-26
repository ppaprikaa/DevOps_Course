#!/bin/sh
sudo timedatectl set-ntp false
sudo service ntp restart
sudo timedatectl set-ntp true
echo "NTP Sync OK"
