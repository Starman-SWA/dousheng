@echo off

set /p wsl_ip=Enter WSL IP:

set /p wlan_ip=Enter WLAN IP:

@echo on

netsh interface portproxy reset

netsh interface portproxy add v4tov4 listenaddress=%wlan_ip% listenport=8080 connectaddress=%wsl_ip%  connectport=8080

netsh advfirewall firewall add rule name="Open Port 8080 for WSL2" dir=in action=allow protocol=TCP localport=8080

netsh interface portproxy show all

pause