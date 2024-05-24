'use strict';

let wasm;
const go = new Go();

function setEventListeners() {
  // Update result when one of the 2 numbers are updated
  document.querySelector('#a').oninput = wasm.exports.updateUI();
  document.querySelector('#b').oninput = wasm.exports.updateUI();
}

function testLogs() {
  console.log(goLog());
  console.log(goLog("Arg1"));
  console.log(varFromGoToJS);
}

function init(wasmObj) {
  wasm = wasmObj.instance;
  go.run(wasm);

  testLogs();
  setEventListeners();
}

if ('instantiateStreaming' in WebAssembly) {
  WebAssembly.instantiateStreaming(fetch("go.wasm"), go.importObject).then(wasmObj => {
    //console.log("instantiateStreaming originally = True");
    init(wasmObj);
  })
} else {
  fetch("go.wasm").then(resp =>
    resp.arrayBuffer()
  ).then(bytes =>
    WebAssembly.instantiate(bytes, go.importObject).then(wasmObj => {
      //console.log("instantiateStreaming originally = False");
      init(wasmObj);
    })
  )
}