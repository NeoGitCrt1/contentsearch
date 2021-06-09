# contentsearch
search content by regex match and output to a file with dupliacation emitted

- usage example
```cmd
 .\findMailId.exe -root "C:\bro\html" -reg ">(?P<name>[A-Z]{3}-[A-Z]{3}-.*)<\/td>" -out "C:/work/tmp/r2.txt"
 ```
 
 - full help 
 ```cmd
 .\findMailId.exe -?
 ```
 
 - next version feature
1. add support for bloom filter
