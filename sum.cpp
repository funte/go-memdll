#include <Windows.h>

extern "C" __declspec(dllexport) int sum(int a, int b)
{
	return a + b;
}

BOOL APIENTRY DllMain(
	_In_ HINSTANCE hinstDLL,
	_In_ DWORD fdwReason,
	_In_ LPVOID lpvReserved
) {
	return TRUE;
}
