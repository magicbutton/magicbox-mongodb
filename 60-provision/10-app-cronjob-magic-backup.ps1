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


$envs += env "AZURE_STORAGE_ACCOUNT" $env:AZURE_STORAGE_ACCOUNT
$envs += env "AZURE_STORAGE_KEY" $env:AZURE_STORAGE_KEY
$envs += env "AZURE_STORAGE_CONTAINER" $env:AZURE_STORAGE_CONTAINER

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

$image = "ghcr.io/magicbutton/$($imagename)-app:$($version)"

$config = @"
apiVersion: batch/v1
kind: CronJob
metadata:
  name: $appname-jobs
spec:
  schedule: "* 3 * * *"
  jobTemplate:
    spec:
      parallelism: 1
      completions: 1
      backoffLimit: 3

      template:
        metadata:
          labels:
            app: $appname-app
        spec: 
          restartPolicy: Never
          volumes:
          - name: shared-data
            emptyDir: {}          
          containers:
        
          - name: $appname-app
            image: $image
            command: [$appname]
            args: ["magic","backup"]
                    
            volumeMounts:
            - name: shared-data
              mountPath: /niels            
            env:
$configEnv 

"@


write-host "Applying config" -ForegroundColor Green

write-host $config -ForegroundColor Gray

$config |  kubectl apply -f -