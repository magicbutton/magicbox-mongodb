<#---
title: App deploy to production
tag: appdeployproduction
api: post
---
#>

$appname = "magicbox-mongodb"
$imagename = "magicbox-mongodb"

$inputFile = join-path  $env:KITCHENROOT $appname ".koksmat", "koksmat.json"

if (!(Test-Path -Path $inputFile) ) {
  Throw "Cannot find file at expected path: $inputFile"
} 
$json = Get-Content -Path $inputFile | ConvertFrom-Json
$version = "v$($json.version.major).$($json.version.minor).$($json.version.patch).$($json.version.build)"

<#
Envs
#>

$envs = @()
function env($name, $value ) {
  if ($null -eq $value) {
    throw "Environment value for $name is not set"
  }
  return @{name = $name; value = $value }
}


$envs += env "EXCHCERTIFICATEPASSWORD" $env:EXCHCERTIFICATEPASSWORD
$envs += env "EXCHAPPID" $env:EXCHAPPID
$envs += env "EXCHORGANIZATION" $env:EXCHORGANIZATION
$envs += env "EXCHCERTIFICATE" $env:EXCHCERTIFICATE
$envs += env "PNPAPPID" $env:PNPAPPID
$envs += env "PNPTENANTID" $env:PNPTENANTID
$envs += env "PNPCERTIFICATE" $env:PNPCERTIFICATE
$envs += env "PNPSITE" $env:PNPSITE
$envs += env "NATS" "nats://nats:4222"
$envs += env "KITCHENROOT" "/kitchens"
$configEnv = ""
foreach ($item in $envs) {

  $configEnv += @"
          - name: $($item.name)
            value: $($item.value)

"@
}

<#
Then we build the deployment file
#>

$image = "ghcr.io/365admin/$($imagename)-app:$($version)"

$config = @"
apiVersion: apps/v1
kind: Deployment
metadata:
  name: $appname-app
spec:
  selector:
    matchLabels:
      app: $appname-app
  replicas: 1
  template:
    metadata:
      labels:
        app: $appname-app
    spec: 
      containers:
      - name: $appname-app
        image: $image
        command: [$appname]
        args: ["service"]               
        env:
$configEnv                           

"@

write-host "Applying config" -ForegroundColor Green

write-host $config -ForegroundColor Gray

$config |  kubectl apply -f -