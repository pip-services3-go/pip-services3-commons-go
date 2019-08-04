#!/usr/bin/env pwsh

Set-StrictMode -Version latest
$ErrorActionPreference = "Stop"

# Get component data and set necessary variables
$component = Get-Content -Path "component.json" | ConvertFrom-Json
$buildImage="$($component.registry)/$($component.name):$($component.version)-$($component.build)-build"
$container=$component.name

# Get buildnumber from teamcity agent
$component.build = $env:BUILD_NUMBER
Set-Content -Path "component.json" -Value $($component | ConvertTo-Json)

# Remove build files
if (Test-Path "dist") {
    Remove-Item -Recurse -Force -Path "dist"
}

# Build docker image
docker build -f docker/Dockerfile.build -t $buildImage .

# Create and copy compiled files, then destroy
docker create --name $container $buildImage
docker cp "$($container):/go/bin" ./dist
docker rm $container

if (!(Test-Path ./dist) -and $env:RETRY -eq $true) {
    # if build failed and retries enabled run build again
    Write-Host "Build failed, but retries enabled, so restarting build script again..."
    ./build.ps1
} elseif (!(Test-Path ./dist)) {
    Write-Host "dist folder doesn't exist in root dir. Build failed. Watch logs above."
    exit 1
}