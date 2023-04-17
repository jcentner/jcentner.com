// darken navbar when at top or when mobile menu shows
const nav = document.querySelector('.navbar');
const navLinks = document.querySelectorAll('.nav-item')
const menuToggle = document.getElementById('navbarSupportedContent')

window.onscroll = function() {

	toggleNavbarDarken();
}

function toggleNavbarDarken() {
	const top = window.scrollY;
	
	if (top >= 100 || menuToggle.classList.contains('show')) {
		nav.classList.add('navbarDarken');
	}
	else {
		nav.classList.remove('navbarDarken');
	}
}

menuToggle.addEventListener('show.bs.collapse', () => {
	nav.classList.add('navbarDarken');
});

menuToggle.addEventListener('hidden.bs.collapse', () => {
	toggleNavbarDarken();
});

// collapse navbar after clicked on mobile
navLinks.forEach((link) => {
	link.addEventListener('click', () => { new bootstrap.Collapse(menuToggle).toggle() })
})
