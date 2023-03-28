// get parameters
let page = window.location.href;
let referrer = document.referrer;

console.log(`page visited: ${page}`)
console.log(`referrer address: ${referrer}`)

$.getJSON("https://api.myip.com", function(data) {
	console.log(`visitor ip: ${data.ip}`)
	console.log(`visitor country: ${data.country}`)

	// make API call
	// ,****replace with actual API call next
	fetch("api/v1/status")
		.then(response => response.json())
		.then(data => console.log(data));
}
