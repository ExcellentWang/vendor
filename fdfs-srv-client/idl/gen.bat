@echo off
echo --------------------------------------------
echo Gen golang stub from IDL

cd /d %~dp0/

cd ..
mkdir tf

cd /d %~dp0/

call thrift -r -out ../tf --gen go:package_prefix=fdfs-srv-client/tf/ proto.thrift

echo ---------------------------------------------
echo.Gen success.

pause