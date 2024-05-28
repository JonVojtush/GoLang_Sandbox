'use strict';

let wasm;
const go = new Go();

function eventListeners() {
  // Update result when one of the 2 numbers are updated
  document.querySelector('#a').oninput = updateUI;
  document.querySelector('#b').oninput = updateUI;
}

function init(wasmObj) {
  go.run(wasmObj.instance);
  eventListeners();
  updateUI(); // Set first result
  console.log(varFromGoToJS);
}

if ('instantiateStreaming' in WebAssembly) { 
  WebAssembly.instantiateStreaming(fetch("go.wasm"), go.importObject).then(wasmObj => {
    init(wasmObj);
  })
} else {
  fetch("go.wasm").then(resp =>
    resp.arrayBuffer()
  ).then(bytes =>
    WebAssembly.instantiate(bytes, go.importObject).then(wasmObj => {
      init(wasmObj);
     })
   )
}