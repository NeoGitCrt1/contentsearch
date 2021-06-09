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
 you may get 
 ```shell
Usage of C:\work\tmp\findMailId.exe:
  -exclude string
        exclude file extension. eg. .zip,.exe (default ".zip,.exe")
  -out string
        search result output file fullname (default "C:/work/tmp/r.txt")
  -p int
        index of the capture group, -1 for whole line (default 1)
  -reg string
        regex for search line (default "MailMapper\\.getMailBody - ==> Parameters:\\s(?P<name>.*)\\(String\\)")
  -root string
        search root (default "C:/work/f")
 ```
 
 - next version feature
1. add support for bloom filter
