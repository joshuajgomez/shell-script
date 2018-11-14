#!/bin/bash  
## Go directly to destinations.    

option=$1

if [ $# -eq 0 ]
then
	echo "Going to ~/meaa/repo/vendor/quest/Phone/DeviceManagerService/"
	cd ~/meaa/repo/vendor/quest/Phone/DeviceManagerService/
elif [ $option == "ps" ]
then
	echo "Going to ~/meaa/repo/vendor/quest/Phone/PhoneService/"
	cd ~/meaa/repo/vendor/quest/Phone/PhoneService/
elif [ $option == "ph" ]
then
	echo "Going to ~/meaa/repo/vendor/quest/Phone/PhoneHMI/"
	cd ~/meaa/repo/vendor/quest/Phone/PhoneHMI/
elif [ $option == "dms" ]
then
	echo "Going to ~/meaa/repo/vendor/quest/Phone/DeviceManagerService/"
	cd ~/meaa/repo/vendor/quest/Phone/DeviceManagerService/
else
	echo "Invalid option"
fi

