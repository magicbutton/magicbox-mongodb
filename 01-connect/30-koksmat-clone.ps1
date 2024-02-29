<#---
title: Login to Azure
---


#>
Push-Location
set-location (join-path $PSScriptRoot ".." ".." )
gh repo clone https://github.com/koksmat-com/koksmat
cd koksmat
go install

Pop-Location