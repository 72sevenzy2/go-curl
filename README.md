 <h2 align="center">pre-requisites: make sure you have go installed:</h2>
```
<h3 align="center">go -v</h3>

<h3 align="center">or</h3>

<h3 align="center">go --v</h3>
```
<br>
<h3 align="center">(this project only suppots post/get requests for API-testing for now, but will extend later on.)</h3>
<br>

<h2 align="center"> get started: </h2>
<h3 align="center">setting headers while testing:</h3>
```
<h4 align="center">go run . [-H key:value] <URL></h4>
```
<br>
<h3 align="center">testing post/get requests: </h3>
```
<h4 align="center">go run . [-x <POST/GET>] <URL></h4>
```
<br>
<h3 align="center">with streaming enabled: (streaming gives live response data back.)</h3>
```
<h4 align="center">go run . [-H key:value] -stream <true/false> <URL></h4>
```
<h3 align="center">(^ important to note that you can include 1 of many flags while testing aswell, but make sure to always include the URL at the end.)</h3>
<br>
<h3 align="center">include response body logging:</h3>
```
<h4 align="center">go run . [-b <true/false> -s <1024>] <URL></h4>
```
<h4 align="center">also important to note you can set a limit on the response body size to log (positive integers only), as with the -s flag, if you want it to be the default (1024), simply do not include the flag.</h3>

<br>

<h3 align="center">session mode:</h3>
<h5 align="center">Instead of manually pasting urls over and over, you can enter session mode to dynamically store urls, headers, etc in a variable and then use that variable to test API's (in which the variable will hold the header/url etc).</h5>
<br>
<h4 align="center">to enter session mode, run the following:</h4>
```
<h4 align="center">go run . [-session <true/false>]</h4>
```
<h5 align="center">when entering session mode, the url does not need to be present at the end.</h5>