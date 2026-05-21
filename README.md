<h1 align="center">pre-requisites: make sure you have go installed:</h2>


<h2 align="center">
    <code> go -v </code>
</h2>
<h2 align="center">or:</h2>

<h2 align="center">
    <code> go --v </code>
</h2>

<br>
<h3 align="center">(only supports post/get requests for now, but will extend later on).</h3>
<br>
<h1 align="center">get started:</h2>

<h2 align="center">setting headers while testing:</h2>

<h3 align="center"> 
    <code> go run . [-H key:value] [URL] </code>
</h>


<br>
<h2 align="center">testing post/get requests:</h2>
<h3 align="center">
    <code> go run . [-x POST/GET] [URL] </code>
</h3>

<br>

<h2 align="center">with streaming enabled: (streaming gives live response data back.)</h2>

<h3 align="center">
    <code> go run . [-H key:value] [-stream true/false] [URL] </code>
</h3>

<h3 align="center">(important to note that you can include 1 of many flags while testing aswell, but make sure to always include the URL at the end).</h3>

<br>

<h2 align="center">include response body logging:</h2>

<h3 align="center">
    <code> go run . [-b true/false -s 1204] [URL] </code>
</h3>

<h4 align="center">(important to note you can set a limit on the response body size to log (positive integers only), as with the -s flag, if you want it to be the default (1024), simply do not include the flag).</h4>

<br>

<h2 align="center">session mode:</h2>

<h3 align="center">Instead of manually pasting, you can enter session mode to dynamically store urls, headers, etc in a variable and then use that variable to test API's (in which the variable will hold the header/url etc).</h3>

<br>

<h2 align="center">to enter session mode, run the following:</h2>

<h3 align="center">
    <code> go run . [-session true/false] </code>
</h3>

<h3 align="center">when entering session mode, the url does not need to be present at the end.</h3>
