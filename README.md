# prettyjson
 prettyjson is a tool for processing JSON inputs, applying the
 given filter to its JSON text inputs and producing the
 filter's results as JSON on standard output.

## usage
```
    # format json input from a file source
    prettyjson -f inputFile.json 
 
    # format json input and use tabs for indentation
    prettyjson -f inputFile.json -i

    # format json input in bold format
    prettyjson -f inputFile.json -F bold
   
    # format json input in bold and italic format
    prettyjson -f inputFile.json -F bold,italic		

    # format based on the JSON passed into stdin.
    cat inputFile.json | prettyjson
    curl -XGET http://ip:port/url/path | prettyjson`
```

## install
```
sudo wget https://github.com/oshankkumar/prettyjson/releases/download/v0.1.0/prettyjson -O /usr/local/bin/prettyjson

sudo chmod +x /usr/local/bin/prettyjson 
```
