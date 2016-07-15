# XML Parser

Command line tool that:

* Reads a spreadsheet and generates a series of XML elements per record, grouped into a parent root element. Each XML record element is structured as column names
* Parses an XML document that contains a series of consistent XML elements into a spreadsheet. First element drives the columns

## Run

### Windows

Place file `xmlparser.exe` from `binaries\windows` folder in any folder
Open Windows command line or PowerShell and go to the folder 
Run `xmlparser -h` for help with command line args 
By default it expects a file `toparse.xml` in the same folder. Take sample file from `data` folder
The results will be in `parsed.xml`

### OS X

Same as for windows. Find binary in `binaries\osx`
Run `./xmlparser`