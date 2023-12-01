# Windows Install

param ($token)

$ServiceName= "mmesh-node"
$InstallationFolder= "C:\Program Files\mmesh"
$MmeshNodeBinary= "mmesh-node.exe"
$MmeshNodeDownloaded= $false
$MmeshNodeBinaryChecksum = "mmesh-node.exe_checksum.sha256"
$WintunBinary= "wintun.dll"
$WintunVersion="wintun-0.14.1"
$ConfigFile="mmesh-node.yml"
$UriMmesh="https://dl.mmesh.io/binaries/stable/latest/windows/amd64/$MmeshNodeBinary"
$UriMmeshChecksum= "https://dl.mmesh.io/binaries/stable/latest/windows/amd64/$MmeshNodeBinaryChecksum"
$UriWintun="https://www.wintun.net/builds/$WintunVersion.zip"

## Functions

## Get-TimeStamp Function
function Get-TimeStamp {
    
    return "{0:yyyy/MM/dd} {0:HH:mm:ss.fff}" -f (Get-Date)
    
}

## Write-Log Fuction

function Write-Log {
    param (
        [Parameter(Mandatory=$true, Position=0)]
        [string] $LogLevel,
        [Parameter(Mandatory=$true, Position=1)]
        [string] $Message
    )

    if ($LogLevel -eq "info"){
        Write-Host "[ info] " -f Blue -NoNewLine; 
    
    } elseif ($LogLevel -eq "warn") {
        Write-Host "[ warn] " -f Yellow -NoNewLine;
    }

    Write-Host "$(Get-TimeStamp) " -f DarkGray -NoNewLine;
    Write-Host "$Message"
    
}

## Download Function
function Download {
    param (
        [Parameter(Mandatory=$true, Position=0)]
        [string] $Uri,
        [Parameter(Mandatory=$true, Position=1)]
        [string] $OutFile
    )

    try {
        $ProgressPreference = 'SilentlyContinue'
        Invoke-WebRequest -Uri $Uri -OutFile "$OutFile"
        $ProgressPreference = 'Continue'
    } catch {
        # Get the status code...
        $StatusCode = [int]$_.Exception.Response.StatusCode

        if  ($statusCode -eq 0) {
            throw " Error: The file $InstallationFolder\$MmeshNodeBinary cannot be replaced because it is being used by another process."
        }
        elseif ($statusCode -ne 200) {
            throw "Download Error: $StatusCode"
        }
    } 
}

## Test-Hash Function
function Test-Hash {
    param(
    [Parameter(Mandatory=$true, Position=0)]
    [string] $file,
    [Parameter(Mandatory=$true, Position=1)]
    [string] $hash
    )
   
    if ((Get-FileHash $file -Algorithm SHA256).Hash.ToUpper() -eq $hash.ToUpper()) {

        Write-Log -LogLevel "info" -Message "Check binary checksum: OK"

    } else {
        throw ("$(Get-TimeStamp) Check binary checksum: FAIL")
    }

}

## Test-Administrator Function
function Test-Administrator  
{  
    [OutputType([bool])]
    param()
    process {
        [Security.Principal.WindowsPrincipal]$user = [Security.Principal.WindowsIdentity]::GetCurrent();
        return $user.IsInRole([Security.Principal.WindowsBuiltinRole]::Administrator);
    }
}

## Main 

## Check if script is execute as Administrator
if(-not (Test-Administrator))
{ 
    Write-Log -LogLevel "warn" -Message "This script must be executed as Administrator!"
    exit 1;
}

## Create installation folder
if (!(Test-Path $InstallationFolder -PathType Container)) 
{
    Write-Log -LogLevel "info" -Message "Create installation folder."
    New-Item -ItemType Directory -Force -Path $InstallationFolder | Out-Null
}

## Create Configuration file
if (!(Test-Path "$InstallationFolder\$ConfigFile"))
{
   New-Item -path "$InstallationFolder" -name "$ConfigFile" -type "file" -value "Token: $token" | Out-Null
   Write-Log -LogLevel "info" -Message "Configuration file created."
} else {
    Move-Item -Path "$InstallationFolder\$ConfigFile" -Destination "$InstallationFolder\old-$ConfigFile" -Force
    New-Item -path "$InstallationFolder" -name "$ConfigFile" -type "file" -value "Token: $token" | Out-Null
    Write-Log -LogLevel "info" -Message "Configuration file replaced."
}

## Download mmmesh-node binary

if (Test-Path "$InstallationFolder\$MmeshNodeBinary") {

    $MmeshNodeDownloaded= $true
    Write-Log -LogLevel "info" -Message "mmesh-node binary has already been downloaded!"

    do { $answer = Read-Host -Prompt "Do you want to replace it?(Y/N)"

    } while("yes","no", "y", "n", "Y", "N" -notcontains $answer)
}  

if (($answer -eq "yes") -or ($answer -eq "y") -or ($answer -eq "Y") -or ($MmeshNodeDownloaded -eq $false)) {

    # Stop service if mmesh-node is running
    $arrService = Get-Service -Name $ServiceName -ErrorAction SilentlyContinue

    while ($arrService.Status -eq 'Running')
    {

        Stop-Service $ServiceName
        Write-Log -LogLevel "info" -Message "$ServiceName service is Running! Stopping..."
        Start-Sleep -seconds 10
        $arrService.Refresh()
        if ($arrService.Status -eq 'Stopped')
        {
            Write-Log -LogLevel "info" -Message "$ServiceName is now Stopped!"
        }
    }

    ## Uninstall mmesh-node
    if ($arrService -ne $null)  {

        Write-Log -LogLevel "info" -Message "Uninstall $ServiceName service!"
        & "$InstallationFolder\$MmeshNodeBinary" "service-uninstall"
        Start-Sleep -seconds 5
    }

    Write-Log -LogLevel "info" -Message "Downloading mmmesh-node binary..."
    Download -Uri $UriMmesh -OutFile "$InstallationFolder\$MmeshNodeBinary"
    Write-Log -LogLevel "info" -Message "Done!"

    Write-Log -LogLevel "info" -Message "Downloading mmmesh-node binary checksum..."
    Download -Uri $UriMmeshChecksum -OutFile "$InstallationFolder\$MmeshNodeBinaryChecksum"

    ## Compare checksum
    $downloadedHash= $(Get-Content $InstallationFolder\$MmeshNodeBinaryChecksum).split(" ")[0]

    Test-Hash -file $InstallationFolder\$MmeshNodeBinary -hash $downloadedHash
}  

## Download wintun DLL

if (Test-Path "$InstallationFolder\$WintunBinary") {
    Write-Log -LogLevel "info" -Message "Wintun DLL has already been downloaded! Nothing to do!"

} else {

    Write-Log -LogLevel "info" -Message "Downloading wintun DLL..."
    Download -Uri $UriWintun -OutFile "$InstallationFolder\$WintunVersion.zip"
    Write-Log -LogLevel "info" -Message "Done!"

    ## Install wintun DLL
    $wintunTmpPath = "$InstallationFolder\$WintunVersion\wintun\bin\amd64\$WintunBinary"

    Write-Log -LogLevel "info" -Message "Installing wintun DLL..."
    $zipFiles = Get-ChildItem $InstallationFolder -Filter *.zip

    foreach ($zipFile in $zipFiles) {

        $zipOutPutFolderExtended = $InstallationFolder + "\" + $zipFile.BaseName
        Expand-Archive -Path $zipFile.FullName -DestinationPath $zipOutPutFolderExtended

    }

    Move-Item -Path "$wintunTmpPath" -Destination "$InstallationFolder" -Force

    if (Test-Path "$InstallationFolder\$WintunVersion.zip") {
        Remove-Item "$InstallationFolder\$WintunVersion.zip"
    }

    if (Test-Path "$InstallationFolder\$WintunVersion") {
        Remove-Item "$InstallationFolder\$WintunVersion" -Recurse -Force
    }

    Write-Log -LogLevel "info" -Message "Done!"
}

    
## Install mmesh-node 
$serviceInstalled = Get-Service -Name $ServiceName -ErrorAction SilentlyContinue

if ($serviceInstalled -eq $null)  {

    Write-Log -LogLevel "info" -Message "Install $ServiceName service!"
    & "$InstallationFolder\$MmeshNodeBinary" "service-install"
    Start-Sleep -seconds 5
} else {
    Write-Log -LogLevel "info" -Message "$ServiceName service has already been installed!"
}

## Start mmesh-node Service
$arrService = Get-Service -Name $ServiceName

while ($arrService.Status -ne 'Running')
{

    Start-Service $ServiceName
    Write-Log -LogLevel "info" -Message "Service starting..."
    Start-Sleep -seconds 10
    $arrService.Refresh()
    if ($arrService.Status -eq 'Running')
    {
        Write-Log -LogLevel "info" -Message "Service is now Running!"
    }

}

# Print Service Output
Get-Service -Name $ServiceName
