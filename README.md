# Hack Assembler
Hack Programming Language Assembler

In order to run project you must have golang installed.
then run:
1. go build -o assembler (in order to build project)
2. ./assembler [--source-file=filename] [--source-folder==folderName] [--dest-folder=destinationFolder]  
(runs program with specified folder and file
by default 
--source-file = Asm.asm
--source-folder == resources/
--dest-folder=results/ (creates it if not exists)
) 

or you can run processing_test.go which takes files from resources folder and checks assembly code with .hack files