const toggle = document.querySelector(".toggle");
const menu = document.querySelector(".menu");
const trim = document.querySelector(".trim");
const overlay = document.querySelector(".overlay");
const result = document.querySelector(".result");
const close = document.querySelector(".close");

 
/* Toggle mobile menu */
function toggleMenu() {
    if (menu.classList.contains("active")) {
        menu.classList.remove("active");
        
        toggle.querySelector("a").innerHTML = "<i class='fas fa-bars'></i>";
    } else {
        menu.classList.add("active");
        
        toggle.querySelector("a").innerHTML = "<i class='fas fa-times'></i>";
    }
}
 

/* Trim Functions */
function Trim() {
	
	overlay.classList.add("show");
	result.classList.add("show");
}

function removeOverlay(){
	overlay.classList.remove("show");
	result.classList.remove("show");
}


/* Copy To ClipBoard */
function copyToBoard() {
    trim_url = document.getElementById('trim_url');

    trim_url.select();
    trim_url.setSelectionRange(0, 99999);

    document.execCommand("copy");

    /* Alert the copied text */
    alert("Copied the text: " + trim_url.value);
}


/* Event Listener */
toggle.addEventListener("click", toggleMenu, false);

trim.addEventListener("click", Trim, false);

close.addEventListener("click", removeOverlay, false);