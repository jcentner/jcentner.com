// status API call
/*
fetch("api/v1/status")
	.then(response => response.json())
	.then(data => console.log(data));
*/



// visit API call (on site load)

// get parameters

const visit_data = {
	page: window.location.href,
	referrer: ((document.referrer) ? document.referrer : 'Direct')
}

fetch("/api/v1/visit", {
	method: "POST",
	header: {
		"Content-Type": "application/json",
	},
	body: JSON.stringify(visit_data),
})
	.then(response => response.json())
	.then(visit_data => {
		console.log("Success:", visit_data);
	})
	.catch(error => console.error("Visit error:", error));



// socialclick API call (on social icon click)

// get parameters

var socialclick_data = {
	button: '',
	page: window.location.href
}

const socialIcons = document.querySelectorAll('.socialclick-icon')

socialIcons.forEach((icon) => {
	icon.addEventListener('mousedown', (event) => {
		if (event.button === 0 || event.button === 2) {
			
			const clickedElement = event.target;

			// what icon was clicked?

			switch(clickedElement.attr('href')) {
				case "https://twitter.com/JacobCentner": 
					socialclick_data.button = 'twitter';
					break;
				case "https://linedin.com/in/jacob-centner/": 
					socialclick_data.button = 'linkedin';
					break;
				case "https://github.com/jcentner/": 
					socialclick_data.button = 'github';
					break;
			}

			fetch("/api/v1/socialclick", {
				method: "POST",
				header: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify(socialclick_data),
			})
			.then((socialclick_data) => {
				console.log(socialclick_data);
				//then navigate away? 
			})
			.catch(error => console.error("Social icon click error: ", error));
		}
	})
})
