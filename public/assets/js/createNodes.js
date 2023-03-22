const createChatCenterElement = message => {
  const element = createElement('p')
  element.append(message)
  element.setAttribute('class', 'chat-center')
  return element
}

const createMessage = element => {
  const usernameNode = createUsernameNode(element)
  const dateNode = createDateNode(element)
  const textMenuNode = createTextMenuNode(element)
  const messageNode = createMessageNode(element)

  const wrapper = createElement('p')
  if (element.is_registered) {
    wrapper.setAttribute('registered', 'true')
  } else {
    wrapper.setAttribute('registered', 'false')
  }
  wrapper.setAttribute('username', element.username)
  wrapper.setAttribute('id', element.id)
  wrapper.append(usernameNode, ' - ', dateNode, ': ', textMenuNode, createElement('br'), messageNode)
  return wrapper
}

const createUsernameNode = ({ username, isRegistered }) => {
  const usernameNode = createElement('span')
  if (getCookie('username') === username) {
    usernameNode.setAttribute('class', 'my-username')
  } else if (username === 'tebafaz') {
    usernameNode.setAttribute('class', 'priveleged-username')
  } else if (isRegistered) {
    usernameNode.setAttribute('class', 'registered-username')
  } else {
    usernameNode.setAttribute('class', 'not-registered-username')
  }
  usernameNode.append(username)
  return usernameNode
}

const createDateNode = ({ sentAt }) => {
  const dateNode = createElement('span')
  dateNode.setAttribute('class', 'date')
  dateNode.append(sentAt)
  return dateNode
}

const createTextMenuNode = ({ username }) => {
  const span = createElement('span')
  span.setAttribute('class', 'message-delete')
  span.setAttribute('title', 'delete this message')
  span.setAttribute('onclick', 'deleteMessageListener(event);')
  span.innerHTML = '&times;'
  if (getCookie('username') !== username) {
    span.style.display = 'none'
  }
  return span
}

const createMessageNode = ({ message }) => {
  const messageNode = createElement('span')
  messageNode.setAttribute('class', 'message')
  messageNode.append(message)
  return messageNode
}

