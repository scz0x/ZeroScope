@echo off
cd /d %~dp0
setlocal ENABLEEXTENSIONS

set "STEP=0"
set "TOTAL=7"

:: Helper for progress bar
call :step "Cleaning old files..."
del /f /q ZeroScope.exe >nul 2>&1
del /f /q cmd-icon.syso >nul 2>&1
del /f /q ZeroScope_patched.exe >nul 2>&1

:: Ensure folders
call :step "Checking folders..."
if not exist assets mkdir assets
if not exist tools mkdir tools
if not exist tools\ResourceHacker mkdir tools\ResourceHacker

:: Generate syso
if exist assets\ZeroScope.ico (
    call :step "Generating syso from icon..."
    rsrc -ico assets\ZeroScope.ico -o cmd-icon.syso
) else (
    echo [!] Icon file not found. Skipping syso.
)

:: Build binary
call :step "Building ZeroScope.exe..."
go build -o ZeroScope.exe ./cmd
if not exist ZeroScope.exe (
    echo [✗] Build failed.
    pause
    exit /b
)

:: Download Resource Hacker (PowerShell or certutil)
if not exist tools\ResourceHacker\ResourceHacker.exe (
    call :step "Downloading Resource Hacker..."

    set "ZIP=%TEMP%\rh.zip"
    set "RH_URL=https://www.angusj.com/resourcehacker/resource_hacker.zip"

    where powershell >nul 2>&1
    if %errorlevel%==0 (
        powershell -Command "Invoke-WebRequest -Uri '%RH_URL%' -OutFile '%ZIP%'"
    ) else (
        certutil -urlcache -split -f "%RH_URL%" "%ZIP%" >nul
    )

    if exist "%ZIP%" (
        powershell -Command "Expand-Archive -Path '%ZIP%' -DestinationPath tools\ResourceHacker -Force" 2>nul
        del "%ZIP%" >nul 2>&1
    ) else (
        echo [✗] Could not download Resource Hacker. Please download manually.
        pause
        exit /b
    )
)

:: Inject icon
call :step "Injecting icon..."
tools\ResourceHacker\ResourceHacker.exe -open ZeroScope.exe -save ZeroScope_patched.exe -action addoverwrite -res assets\ZeroScope.ico -mask ICONGROUP,MAINICON,
if exist ZeroScope_patched.exe (
    move /Y ZeroScope_patched.exe ZeroScope.exe >nul
    echo [✓] Icon injection completed.
) else (
    echo [✗] Icon injection failed.
)

call :step "Done! ZeroScope.exe is ready."
echo.
pause
exit /b

:step
set /a STEP+=1
echo [▌%STEP%/%TOTAL%] %~1
goto :eof