// +build !wasm

package main

const (
	template_string = `
<!doctype html>
<html>
<head>
{{if .Title}}
<title>{{.Title}}</title>
{{else}}
<title>dev42 - {{.Request.URL.Path}}</title>
{{end}}
<meta charset="utf-8"/>
<link rel="apple-touch-icon" sizes="180x180" href="static/icons/root/apple-touch-icon.png">
<link rel="icon" type="image/png" sizes="32x32" href="static/icons/root/favicon-32x32.png">
<link rel="icon" type="image/png" sizes="16x16" href="static/icons/root/favicon-16x16.png">
<link rel="manifest" href="static/icons/root/site.webmanifest">
<link rel="mask-icon" href="static/icons/root/safari-pinned-tab.svg" color="#054306">
{{if .MetaTags}}{{range $k, $v := .MetaTags}}
<meta name="{{$k}}" content="{{$v}}"/>
{{end}}{{end}}
<meta name="msapplication-TileColor" content="#00aba9">
<meta name="theme-color" content="#ffffff">
{{if .CSSFiles}}{{range $f := .CSSFiles}}
<link rel="stylesheet" href="{{$f}}" />
{{end}}{{end}}
<script src="https://cdn.jsdelivr.net/npm/text-encoding@0.7.0/lib/encoding.min.js"></script> <!-- MS Edge polyfill -->
<script src="/wasm_exec.js"></script>
</head>
<body>
<div id="root_mount_parent">
{{if .ServerRenderedOutput}}{{.ServerRenderedOutput}}{{else}}
<img style="position: absolute; top: 50%; left: 50%;" src="https://cdnjs.cloudflare.com/ajax/libs/galleriffic/2.0.1/css/loader.gif">
{{end}}
</div>
<script>
var wasmSupported = (typeof WebAssembly === "object");
if (wasmSupported) {
	if (!WebAssembly.instantiateStreaming) { // polyfill
		WebAssembly.instantiateStreaming = async (resp, importObject) => {
			const source = await (await resp).arrayBuffer();
			return await WebAssembly.instantiate(source, importObject);
		};
	}
	const go = new Go();
	WebAssembly.instantiateStreaming(fetch("/main.wasm"), go.importObject).then((result) => {
		go.run(result.instance);
	});
} else {
	document.getElementById("root_mount_parent").innerHTML = 'This application requires WebAssembly support.  Please upgrade your browser.';
}
</script>
</body>
</html>
`
	title    = "Dice Fairness Analyzer - dev42"
	icons    = "static/icons/root/*"
	cssFonts = "static/css/fonts.css"
	cssIndex = "static/css/index.css"
)
