# reverse

Converting the graphic representation into a text

## Author

* [@smslash](https://github.com/smslash)

## Objectives

You must follow the same [instructions](https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art/README.md) as in the first subject but this time the process will be reversed. desrever fo dnik siht toN.

Ascii-reverse consists on reversing the process, converting the graphic representation into a text. You will have to create a text file containing a graphic representation of a random `string` given as an argument.

The argument will be a **flag**, `--reverse=<fileName>`, in which `--reverse` is the flag and `<fileName>` is the file name. The program must then print this `string` in **normal text**.

- The flag must have exactly the same format as above, any other formats must return the following usage message:

```bash
Usage: go run . [OPTION]

EX: go run . --reverse=<fileName>
```

If there are other `ascii-art` optional projects implemented, the program should accept other correctly formatted `[OPTION]` and/or `[BANNER]`.

Additionally, the program must still be able to run with a single `[STRING]` argument.

## Instructions

- Your project must be written in **Go**.

- The code must respect the [good practices](https://01.alem.school/git/root/public/src/branch/master/subjects/good-practices/README.md).

- It is recommended to have **test files** for [unit testing](https://go.dev/doc/tutorial/add-a-test).

## Usage

```
$ cat file.txt
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
$
$ go run . --reverse=file.txt
hello
$
```

## Allowed packages

- Only the [standard Go](https://pkg.go.dev/std) packages are allowed
This project will help you learn about :

- The Go file system(**fs**) API

- Data manipulation