const go = new Go()
const wasm_url = "./wasm/dist/output.wasm";
let wasm

const add_el = document.getElementById("add") as HTMLElement
if("instantiateStreaming" in WebAssembly) {
    WebAssembly.instantiateStreaming(fetch(wasm_url), go.importObject).then((result) => {
        wasm = result.instance
        go.run(wasm)
    })
} else {
    fetch(wasm_url).then(resp =>
        resp.arrayBuffer()
    ).then(bytes =>
        WebAssembly.instantiate(bytes, go.importObject).then(function (obj) {
            wasm = obj.instance;
            go.run(wasm);
        })
    )
}