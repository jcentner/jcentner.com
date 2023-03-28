// get parameters
const data = {
	page: window.location.href,
	referrer: document.referrer
}

console.log(`page visited: ${data.page}`)
console.log(`referrer address: ${data.referrer}`)

// test API call
/*
fetch("api/v1/status")
	.then(response => response.json())
	.then(data => console.log(data));
*/

fetch("/api/v1/visit", {
	method: "POST",
	header: {
		"Content-Type": "application/json",
	},
	body: JSON.stringify(data),
})
	.then(response => response.json())
	.then(data => {
		console.log("Success:", data);
	})
	.catch(error => console.error("Error:", error));

