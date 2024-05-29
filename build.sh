#!/bin/bash
LOG_FILE="web/build.log"
touch $LOG_FILE 2>&1 | tee /dev/stderr;
echo "Script started at: $(date)" | tee -a $LOG_FILE;


if [ -f web/go.wasm ]; then
  echo "web/go.wasm exists, removing it..." | tee -a $LOG_FILE;
  if ! rm -f web/go.wasm; then
    echo "Failed to remove web/go.wasm" | tee -a $LOG_FILE >&2;
    exit 1;
  fi
else
  echo "web/go.wasm doesn't exist." | tee -a $LOG_FILE;
fi


echo "Building a new file..." | tee -a $LOG_FILE;
GOOS=js GOARCH=wasm go build -o=web/go.wasm -buildvcs=false 2>&1 | tee -a $LOG_FILE >&2; # Check for error in building go.wasm and redirect stderr to log file
### TinyGo (wasm_exec.js:303 syscall/js.finalizeRef not implemented)
  # tinygo build -o=web/go.wasm -target=wasm main.go;
  # cp "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" web/;
echo "web/go.wasm was built." | tee -a $LOG_FILE;


if [ -f web/wasm_exec.js ]; then
  echo "web/wasm_exec.js exists, removing it..." | tee -a $LOG_FILE;
  if ! rm -f web/wasm_exec.js; then
    echo "Failed to remove web/wasm_exec.js" | tee -a $LOG_FILE >&2;
    exit 1;
  fi
else
  echo "web/wasm_exec.js doesn't exist." | tee -a $LOG_FILE;
fi


echo "Fetching a new file..." | tee -a $LOG_FILE;
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" web/ 2>&1 | tee -a $LOG_FILE >&2; # Check for error in copying wasm_exec.js and redirect stderr to log file 
echo "web/wasm_exec.js was fetched from \$GOROOT/misc/wasm/" | tee -a $LOG_FILE;


echo "Script ended at: $(date)" | tee -a $LOG_FILE;
echo "----------------------------------------------------------------------------------------------------" | tee -a $LOG_FILE;