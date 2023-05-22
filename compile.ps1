cd goodbye

cl.exe /nologo /LD goodbye.cpp /Fe:./

gendef ./goodbye.dll
dlltool -k -d ./goodbye.def -l ./libgoodbye.a

mv ./goodbye.dll ../goodbye.dll
mv ./libgoodbye.a ../libgoodbye.a

rm -f goodbye.obj
rm -f goodbye.exp
rm -f goodbye.lib
rm -f goodbye.def
