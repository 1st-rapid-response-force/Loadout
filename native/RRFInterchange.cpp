#define WIN32_LEAN_AND_MEAN

#include "windows.h"
#include "interface.h"

BOOL APIENTRY DllMain( HMODULE hModule, DWORD  ul_reason_for_call, LPVOID lpReserved)
{
	switch (ul_reason_for_call)
	{
	case DLL_PROCESS_ATTACH:
	case DLL_THREAD_ATTACH:
	case DLL_THREAD_DETACH:
	case DLL_PROCESS_DETACH:
		break;
	}
	return TRUE;
}

extern "C"
{
    __declspec(dllexport) void __stdcall RVExtension(char *output, int outputSize, const char *function);
};

void __stdcall RVExtension(char *output, int outputSize, const char *function)
{	
	char* funcCopy = strdup(function);
    GoString response = HandleRVInput(funcCopy);

	outputSize = response.n + 1;
	strncpy_s(output, outputSize, response.p, _TRUNCATE);
}