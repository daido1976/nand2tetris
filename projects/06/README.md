# projects 06

```sh
# Install Java
# $ chmod +x ./tools/Assembler.sh
$ ./tools/Assembler.sh

# Compile and execute Go program
$ go run ./projects/06/assembler/parser.go ./projects/06/assembler/main.go ./projects/06/assembler/code.go ./projects/06/assembler/symbol_table.go projects/06/rect/Rect.asm

# Compile Go program
$ go build -o ./projects/06/assembler/main ./projects/06/assembler
# Execute Go program
$ ./projects/06/assembler/main projects/06/rect/Rect.asm
```
