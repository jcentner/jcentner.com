// get parameters
let page = window.location.href;
let referrer = document.referrer;

console.log(`page visited: ${page}`)
console.log(`referrer address: ${referrer}`)

const data = {
	page: window.location.href,
	referrer: document.referrer
}

// make API call
// ,****replace with actual API call next
fetch("api/v1/status")
	.then(response => response.json())
	.then(data => console.log(data));

/*
fetch("api/v1/visit", {
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
	.catch(error => console.error("Error:", error);
});
*/
