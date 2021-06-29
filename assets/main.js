function useChatWebsocket() {
    const ws = new WebSocket(`ws://${window.location.host}/ws`);

    ws.onopen = (event) => {
        setInterval(() => {
            const msg = {
                message: "Hello World!"
            }
            console.log(msg)
            ws.send(JSON.stringify(msg))
        }, 5000)
        
    }

    ws.onmessage = (event) => {
        console.log(event.data)
    }
}

document.addEventListener("DOMContentLoaded", () => {
    console.log("👨‍💻 + ✉️ +👩‍💻 = Chat")

    useChatWebsocket()
})