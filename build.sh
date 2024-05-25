#!/bin/bash
LOG_FILE="http/build.log"
touch $LOG_FILE 2>&1 | tee /dev/stderr;

echo "Script started at: $(date)" | tee -a $LOG_FILE;

if [ -f http/go.wasm ]; then
  echo "http/go.wasm exists, removing it..." | tee -a $LOG_FILE;
  if ! rm -f http/go.wasm; then
    echo "Failed to remove http/go.wasm" | tee -a $LOG_FILE >&2;
    exit 1;
  fi
else
  echo "http/go.wasm doesn't exist." | tee -a $LOG_FILE;
fi

echo "Building a new file..." | tee -a $LOG_FILE;
# Check for error in building go.wasm and redirect stderr to log file
GOOS=js GOARCH=wasm go build -o=http/go.wasm -buildvcs=false 2>&1 | tee -a $LOG_FILE >&2;
echo "http/go.wasm was built." | tee -a $LOG_FILE;

if [ -f http/wasm_exec.js ]; then
  echo "http/wasm_exec.js exists, removing it..." | tee -a $LOG_FILE;
  if ! rm -f http/wasm_exec.js; then
    echo "Failed to remove http/wasm_exec.js" | tee -a $LOG_FILE >&2;
    exit 1;
  fi
else
  echo "http/wasm_exec.js doesn't exist." | tee -a $LOG_FILE;
fi

echo "Fetching a new file..." | tee -a $LOG_FILE;
# Check for error in copying wasm_exec.js and redirect stderr to log file
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" http/ 2>&1 | tee -a $LOG_FILE >&2; 
echo "http/wasm_exec.js was fetched from \$GOROOT/misc/wasm/" | tee -a $LOG_FILE;
 
echo "Script ended at: $(date)" | tee -a $LOG_FILE;

# TinyGo (wasm_exec.js:303 syscall/js.finalizeRef not implemented)
  # tinygo build -o=http/go.wasm -target=wasm wasm.go;
  # cp "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" http/;