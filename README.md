# preetyjson
 preetyjson is a tool for processing JSON inputs, applying the
 given filter to its JSON text inputs and producing the
 filter's results as JSON on standard output.

##usage
```
    # format json input from a file source
    preetyjson -f inputFile.json 
 
    # format json input and use tabs for indentation
    preetyjson -f inputFile.json -i

    # format json input in bold format
    preetyjson -f inputFile.json -F bold
   
    # format json input in bold and italic format
    preetyjson -f inputFile.json -F bold,italic		

    # format based on the JSON passed into stdin.
    cat inputFile.json | preetyjson
    curl -XGET http://ip:port/url/path | preetyjson`
```

##install
```
wget https://github.com/oshankkumar/preetyjson/releases/download/v0.1.0/preetyjson -O /usr/local/bin/preetyjson

chmod +x /usr/local/bin/preetyjson 
```