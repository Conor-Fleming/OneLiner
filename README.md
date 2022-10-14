# OneLiner

## About
One liner is a program to make more efficient the process of writing custom M extensions to run in Epic EMR software. When given a text file OneLiner will return a new txt file with the contents formatted into One line for use as an Extension in the Epic application. In the case that you give a text file that is already in "OneLiner" format it will return a text file with the contents in multiple lines, tabbed and pretty.

### Project status: 
Build passing
    
## Installation and Use
Clone Github Repo: https://github.com/Conor-Fleming/OneLiner
Build and run
```
git clone https://github.com/Conor-Fleming/OneLiner
```
Use build command to generate executable file and run with the file to be converted as a command line argument
```
go build
./Oneliner test.txt
```
## Example
### Converting .txt file from multiple lines to one
```
$ cat test.txt
s check=$$evalRule^elibHULIB22(803864,patID,patDAT) 
s hemOnc=$$provName^LNEWSCH4(ID("SER"),"2")
s infusion=$$gt44181^S2PAFPP4(1)
s output=$S(check=1:hemOnc,1:infusion) 
s cells(col,1)=output
```

Run app
```
$ ./Oneliner test.txt
cat result.txt
s check=$$evalRule^elibHULIB22(803864,patID,patDAT) s hemOnc=$$provName^LNEWSCH4(ID("SER"),"2") s infusion=$$gt44181^S2PAFPP4(1) s output=$S(check=1:hemOnc,1:infusion) s cells(col,1)=output
```
### Converting .txt file from one line to many
```
n return,count d smarty2^JVITALS4("","","","",1,"","","",1,I) f n1=1:1:V(0) {s V(n1)=$REPLACE(V(n1),":",",") s V(n1)=$REPLACE(V(n1),",",",    --")_$CHAR(13,10) s V(n1)=V(n1)_"," s numPieces=$L(V(n1),",") f n2=1:1:numPieces {i ($P(V(n1),",",n2)'["- Other")&($P(V(n1),",",n2)'="") {s count=count+1 s return(count)=$P(V(n1),",",n2) } } } f n3=1:1:count {s V(n3)=return(n3) } s V(n3)=$CHAR(10) s V(0)=count
```
Run app
```
$ ./Oneliner test.txt
$ cat result.txt
n return,count 
d smarty2^JVITALS4("","","","",1,"","","",1,I) 
f n1=1:1:V(0) {
	s V(n1)=$REPLACE(V(n1),":",",")
	s V(n1)=$REPLACE(V(n1),",",",    --")_$CHAR(13,10)
	s V(n1)=V(n1)_"," 
	s numPieces=$L(V(n1),",") 
	f n2=1:1:numPieces {
		i ($P(V(n1),",",n2)'["- Other")&($P(V(n1),",",n2)'="") {
			s count=count+1 s return(count)=$P(V(n1),",",n2)
		}
	} 	
} 
f n3=1:1:count {
	s V(n3)=return(n3)
} 
s V(n3)=$CHAR(10)
s V(0)=count
```


