#!/bin/bash  
# Go directly to destinations.
#
# -Installation-
# 1. Checkout git repo in home.
# 2. Add below lines to .bashrc file
##### PATH=$PATH:~/shell-script
##### alias go='source go'   

option=$1

if [ $# -eq 0 ]
then
	echo "Going to ~/source/vendor/quest/Phone/DeviceManagerService/"
	cd ~/source/vendor/quest/Phone/DeviceManagerService/
elif [ $option == "ps" ]
then
	echo "Going to ~/source/vendor/quest/Phone/PhoneService/"
	cd ~/source/vendor/quest/Phone/PhoneService/
elif [ $option == "ph" ]
then
	echo "Going to ~/source/vendor/quest/Phone/PhoneHMI/"
	cd ~/source/vendor/quest/Phone/PhoneHMI/
elif [ $option == "dms" ]
then
	echo "Going to ~/source/vendor/quest/Phone/DeviceManagerService/"
	cd ~/source/vendor/quest/Phone/DeviceManagerService/
else
	echo "Invalid option"
fi

