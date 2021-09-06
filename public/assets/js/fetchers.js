fetchFromConnect = data => {
    data.messages.map(element => {
        if(element.id == 1){
            var startText = createChatCenterElement('Start of the chat')
            getById('chat').prepend(startText)
        }
        var wrapper = createMessage(element)
        var chatDiv = getById('chat')
        chatDiv.appendChild(wrapper)
        chatDiv.scrollTop = chatDiv.scrollHeight - chatDiv.clientHeight
    })
    firstID = data.messages[0].id
    lastID = data.last_id
}

fetchFromSubscribe = (data) => {
    if (data.deleted_id != null){
        getById(`${data.deleted_id}`).remove()
        return
    }
    lastID = data.last_id
    data.messages.map(element => {
        var wrapper = createMessage(element)
        var chatDiv = getById('chat')
        if (chatDiv.scrollTop > chatDiv.scrollHeight - chatDiv.clientHeight - 100) {
            chatDiv.appendChild(wrapper)
            chatDiv.scrollTop = chatDiv.scrollHeight - chatDiv.clientHeight
        } else {
            chatDiv.appendChild(wrapper)
        }
    })
}

fetchFromPriorMessages = (data) => {
    firstID = data.messages[0].id
     var fragment = document.createDocumentFragment()
     data.messages.map(element => {
         if(element.id == 1){
             var startText = createChatCenterElement('Start of the chat')
             fragment.prepend(startText)
         }
         var wrapper = createMessage(element)
         fragment.appendChild(wrapper)
     })
     var chatDiv = getById('chat')
     chatDiv.prepend(fragment)
 }