const toggle = document.querySelector(".toggle");
const menu = document.querySelector(".menu");
const trim = document.querySelector(".trim");
const overlay = document.querySelector(".overlay");
const result = document.querySelector(".result");
const close = document.querySelector(".close");
const btn = document.querySelector(".btn");

var payload = {};
var form = document.getElementById("urlForm");

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


// Remove Overlay
function removeOverlay(){
    overlay.classList.remove("show");
    result.classList.remove("show");
}


btn.addEventListener("click", (e) => {
    e.preventDefault()

    const tooltip = document.querySelector(".tooltip");
    tooltip.classList.add("show");

})

form.addEventListener("submit", (e) => {
    e.preventDefault()
        // Get user input
    var input = document.getElementById("urlInput").value
    payload["url"] = input
        // Make the request
    axios.post("/trim", payload)
        .then((response) => {
            let status = response.status
            let data = response.data
            console.log(data.msg)
                //Update Dom        
            let p = document.getElementById("trim_url")
            p.innerHTML = data.msg

            Trim()
        })
        .catch((error) => {
            console.log(error)
        })
})


// Copy To clip Board function

new ClipboardJS('.btn');


/* Event Listener */
toggle.addEventListener("click", toggleMenu, false);

close.addEventListener("click", removeOverlay, false);