var payload = {}

var form = document.getElementById("urlForm")
form.addEventListener("submit", (e) => {
    e.preventDefault()
    // Get user input
    var input = document.getElementById("urlInput").value
    payload["url"] = input
    // Make the request
    axios.post("/trim", payload)
        .then((response) => {
            var status = response.status
            var data = response.data
            console.log(data.msg)
            //Update Dom        
            var p = document.getElementById("shortenedURL")
            p.innerHTML = data.msg
        })
        .catch((error) => {
            console.log(error)
        })
})