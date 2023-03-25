@echo off 
cls

set "token=%1"

if [%token%]==[] (goto:undefined) else (goto:define)

:undefined
echo Token is not defined
set /p "token=provide a token : " 
:define
echo Token is %token%
go run . %token%

pause