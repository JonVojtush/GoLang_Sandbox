'use strict';

let wasm;
const go = new Go();
// Create an empty object to hold all of the WebAssembly exports
const wasmExports = {};

function eventListeners() {
  // Update result when one of the 2 numbers are updated
  document.querySelector('#a').oninput = wasmExports.updateUI;
  document.querySelector('#b').oninput = wasmExports.updateUI;
}

function main(wasmObj) {
  // Store all exports into an object
  Object.assign(wasmExports, wasmObj.instance.exports);

  wasm = wasmObj.instance;
  go.run(wasm);

  eventListeners();
  console.log(varFromGoToJS);
}

// Go --> WASM
if ('instantiateStreaming' in WebAssembly) {
  WebAssembly.instantiateStreaming(fetch("go.wasm"), go.importObject).then(wasmObj => main(wasmObj));
} else {
  fetch("go.wasm").then(resp => resp.arrayBuffer()).then(bytes => WebAssembly.instantiate(bytes, go.importObject)).then(wasmObj => main(wasmObj));
}





