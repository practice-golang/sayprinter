#if defined(_WIN32) && !defined(__MINGW32__)
#define MY_API __declspec(dllexport)
#else
#define MY_API
#endif

#ifdef __cplusplus
extern "C" {
#endif

MY_API int say_goodbye();

#ifdef __cplusplus
}
#endif
