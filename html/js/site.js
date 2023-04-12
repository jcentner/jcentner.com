// darken navbar when at top
const nav = document.querySelector('.navbar');

window.onscroll = function() {
	const top = window.scrollY;
	
	if (top >= 100) {
		nav.classList.add('navbarDarken');
	}
	else {
		nav.classList.remove('navbarDarken');
	}
}

// collapse navbar after clicked on mobile
const navLinks = document.querySelectorAll('.nav-item')
const menuToggle = document.getElementById('navbarSupportedContent')

navLinks.forEach((link) => {
	link.addEventListener('click', () => { new boostrap.Collapse(menuToggle).toggle() })
})
