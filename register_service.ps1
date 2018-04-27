go build -o app\server.exe app\server.go

$CommandStr =
@"
New-Service -Name SapiServer -BinaryPathName 'C:\Program Files (x86)\Windows Resource Kits\Tools\srvany.exe'
New-Item -Path 'HKLM:\SYSTEM\CurrentControlSet\Services\SapiServer\Parameters'
Set-ItemProperty -Path 'HKLM:\SYSTEM\CurrentControlSet\Services\SapiServer\Parameters' -Name Application "$($PSScriptRoot)\app\server.exe"
Write-Host '`Start-Service SapiServer` to start service'
Write-Host '`sc.exe delete SapiServer` to delete service'
"@
$Command = [Scriptblock]::Create($CommandStr)

Start-Process PowerShell -ArgumentList "-NoExit -Command & { $Command }" -Verb RunAs
