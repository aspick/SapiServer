go build app\server.go

$Command = {
New-Service -Name SapiServer -BinaryPathName 'C:\Program Files (x86)\Windows Resource Kits\Tools\srvany.exe'
New-Item -Path 'HKLM:\SYSTEM\CurrentControlSet\Services\SapiServer\Parameters'
Set-ItemProperty -Path 'HKLM:\SYSTEM\CurrentControlSet\Services\SapiServer\Parameters' -Name Application '$PSScriptRoot\app\server.exe'
Write-Host '`Start-Service SapiServer` to start service'
}
Start-Process PowerShell -ArgumentList "-NoExit -Command & { $Command }" -Verb RunAs
