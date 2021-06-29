function useChatWebsocket() {
    const ws = new WebSocket(`ws://${window.location.host}/ws`);

    ws.onopen = (event) => {
        ws.send(JSON.stringify({
            message: "hello world!"
        }))
    }

    ws.onmessage = (event) => {
        console.log(event.data)
    }
}

document.addEventListener("DOMContentLoaded", () => {
    console.log("👨‍💻 + ✉️ +👩‍💻 = Chat")

    useChatWebsocket()
})