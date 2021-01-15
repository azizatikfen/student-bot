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
    console.log("Id = ", id, "Passwd = ", passwd)
}