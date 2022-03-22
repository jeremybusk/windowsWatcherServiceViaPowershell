# windowsWatcherServiceViaPowershell
Powershell service from powershell script that restarts service when fails

# Create Event Log Application and test event
```
New-EventLog –LogName Application –Source "watchAuth0LDAPConn"
Write-EventLog –LogName Application –Source "watchAuth0LDAPConn" –EntryType Information –EventID 1 –Message "This is a test message."
```

# Create Script
C:\app\watchAuth0LDAPConn.ps1
```
$url = "127.0.0.1:49948"
$sleep = 3
$msg = "Starting $url connection watcher."
Write-Output $msg
Write-EventLog -LogName "Application" -Source "watchAuth0LDAPConn" -EventID 1 -EntryType Information -Message $msg
while($true){
  try {
    $r = Invoke-WebRequest $url
  }
  catch {
    $_.Exception.Message
    $msg = "URL $url is unresponsive. Restarting Auth0LDAP service."
    Write-Output $msg
    Write-EventLog -LogName "Application" -Source "watchAuth0LDAPConn" -EventID 2 -EntryType Warning -Message $msg
  }
  Start-Sleep -s $sleep
}
```

# Create Service

Download and use nssm.exe to create service - https://nssm.cc/download
```
$Binary = (Get-Command Powershell).Source
$Arguments = '-ExecutionPolicy Bypass -NoProfile -File "C:\app\watchAuth0LDAPConn.ps1"'
.\nssm.exe install watchAuth0LDAPConn $Binary $Arguments
```

# Start Service
```
start-service watchAuth0LDAPConn
get-service watchAuth0LDAPConn
```

Remove Service
```
nssm.exe remove watchAuth0LDAPConn confirm
```




you may be able to use sc.exe but I had issues with New-Service command limits because it was ps1 and not exe

