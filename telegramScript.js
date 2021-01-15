d = document.getElementsByClassName("im_message_text");
for (let i = 1; i < d.length + 1; i++) {
    let msgText = d[d.length - i].innerText;
    let linkIndex = msgText.indexOf("https://us04web.zoom.us/j/");
    let passwdIndex = msgText.indexOf("?pwd=");
    let id, passwd;
    if (linkIndex !== -1) {
        id = msgText.substring(linkIndex + 26, passwdIndex);
        if (linkIndex !== 0) {
            passwd = msgText.substring(passwdIndex + 5, msgText.indexOf("\n", passwdIndex))
        } else {
            passwd = msgText.substring(passwdIndex + 5, msgText.lenght)
        }
    };
    if (id !== undefined && passwd !== undefined){
        fetch("http:/localhost:8080/meeting",{method:"POST",body:`{"ID":"`+id+`","Password":"`+passwd+`"}`}).then((res)=>{console.log(res);if (res === "Success"){}})
    }
    console.log("Id = ", id, "Passwd = ", passwd)
}