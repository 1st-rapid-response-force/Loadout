@echo off
go build -buildmode=c-archive

gcc -shared -pthread -o "C:\Program Files (x86)\Steam\steamapps\common\Arma 3\x\rrf\rrf_loadouts_x64.dll" "./native/RRFInterchange.cpp" loadout.a -lWinMM -lntdll -lWS2_32