#!/bin/bash

if [ -f http/go.wasm ]; then
  rm -f http/go.wasm;
  echo "http/go.wasm was removed; building a new file...";
else
  echo "http/go.wasm doesn't exist; Building one...";
fi

echo -en '\n'

#tinygo build -o=http/go.wasm -target=wasm wasm.go;
GOOS=js GOARCH=wasm go build -o=http/go.wasm -buildvcs=false
echo "http/go.wasm was built.";

echo -en '\n'

if [ -f http/wasm_exec.js ]; then
  rm -f http/wasm_exec.js;
  echo "http/wasm_exec.js was removed; fetching a new file...";
else
  echo "http/wasm_exec.js doesn't exist; Fetching one...";
fi

echo -en '\n'

#cp "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" http/; #wasm_exec.js:303 syscall/js.finalizeRef not implemented
#echo "http/wasm_exec.js was fetched from tinygo.";
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" http/; 
echo "http/wasm_exec.js was fetched from wasm.";

echo -en '\n'

#if error
  #remain open until closed manually
#else
  echo "Finished, this will close automatically."
  sleep 5
  exit
#fi