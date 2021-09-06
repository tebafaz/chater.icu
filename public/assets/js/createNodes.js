createChatCenterElement = message => {
    var element = createElement('p')
    element.append(message)
    element.setAttribute('class', 'chat-center')
    return element
}

createMessage = element => {
    var usernameNode = createUsernameNode(element)
    var dateNode = createDateNode(element)
    var textMenuNode = createTextMenuNode(element)
    var messageNode = createMessageNode(element)

    var wrapper = createElement('p')
    if (element.is_registered) {
        wrapper.setAttribute('registered', 'true')
    } else {
        wrapper.setAttribute('registered', 'false')
    }
    wrapper.setAttribute('username', element.username)
    wrapper.setAttribute('id', element.id)
    wrapper.append(usernameNode, ' - ', dateNode, `: `, textMenuNode, createElement('br'), messageNode)
    return wrapper
}

createUsernameNode = ({username, is_registered}) => {
    var usernameNode = createElement('span')
    if (getCookie("username") == username) {
        usernameNode.setAttribute('class', 'my-username')
    } else if (username == 'tebafaz') {
        usernameNode.setAttribute('class', 'priveleged-username')
    } else if (is_registered){
        usernameNode.setAttribute('class', 'registered-username')
    } else {
        usernameNode.setAttribute('class', 'not-registered-username')
    }
    usernameNode.append(username)
    return usernameNode
}

createDateNode = ({sent_at}) => {
    var dateNode = createElement('span')
    dateNode.setAttribute('class', 'date')
    dateNode.append(sent_at)
    return dateNode
}

createTextMenuNode = ({username}) => {
    var span = createElement('span')
    span.setAttribute('class', 'message-delete')
    span.setAttribute('title', 'delete this message')
    span.setAttribute('onclick', 'deleteMessageListener(event);')
    span.innerHTML = '&times;'
    if (getCookie("username") != username) {
        span.style.display = "none"
    }
    return span
}

createMessageNode = ({message}) => {
    var messageNode = createElement('span')
    messageNode.setAttribute('class', 'message')
    messageNode.append(message)
    return messageNode
}