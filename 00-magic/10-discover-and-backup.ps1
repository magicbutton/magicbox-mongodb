<#---
title: Backup All
tag: backup
icon: magic.png
api: post
---

#






#>



if ($env:WORKDIR -eq $null) {
    $env:WORKDIR = join-path $psscriptroot ".." ".koksmat" "workdir"
}
$workdir = $env:WORKDIR

if (-not (Test-Path $workdir)) {
    $x = New-Item -Path $workdir -ItemType Directory 
}

$workdir = Resolve-Path $workdir

write-host "Workdir: $workdir"


write-host "Getting Databases"
magicbox-mongodb discover databases

write-host "Backing Up"
magicbox-mongodb backup all (join-path $workdir "databaseservices.json") NO

