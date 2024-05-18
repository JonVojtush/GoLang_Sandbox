'use strict';

let wasm;
const go = new Go();

function init(wasmObj) {
  wasm = wasmObj.instance;
  go.run(wasm);

  console.log(goLog());
  console.log(goLog("Arg1"));

  //console.log(varFromGoToJS); // TASK: Set in main() or wasm.go

  //Event Listeners
  // Update result when one of the 2 numbers are updated
  document.querySelector('#a').oninput = wasm.exports.updateUI(); // BUG: Not finding function
  document.querySelector('#b').oninput = wasm.exports.updateUI();
}

if ('instantiateStreaming' in WebAssembly) {
  WebAssembly.instantiateStreaming(fetch("go.wasm"), go.importObject).then(wasmObj => {
    console.log("instantiateStreaming = True");
    init(wasmObj);
  })
} else {
  fetch("go.wasm").then(resp =>
    resp.arrayBuffer()
  ).then(bytes =>
    WebAssembly.instantiate(bytes, go.importObject).then(wasmObj => {
      console.log("instantiateStreaming = False");
      init(wasmObj);
    })
  )
}