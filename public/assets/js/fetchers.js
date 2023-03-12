import { createChatCenterElement, createMessage } from './createNodes'
import { getById } from './apiCallback'
import { setFirstID, setLastID } from './callbackLogic'

const fetchFromConnect = data => {
  const elements = data.messages.map(element => {
    if (element.id === 1) {
      const startText = createChatCenterElement('Start of the chat')
      getById('chat').prepend(startText)
    }
    const wrapper = createMessage(element)
    const chatDiv = getById('chat')
    chatDiv.appendChild(wrapper)
    chatDiv.scrollTop = chatDiv.scrollHeight - chatDiv.clientHeight
    return element
  })
  console.log(elements)
  setFirstID(data.messages[0].id)
  setLastID(data.last_id)
}

const fetchFromSubscribe = (data) => {
  if (data.deleted_id != null) {
    getById(`${data.deleted_id}`).remove()
    return
  }
  setLastID(data.last_id)
  const elements = data.messages.map(element => {
    const wrapper = createMessage(element)
    const chatDiv = getById('chat')
    if (chatDiv.scrollTop > chatDiv.scrollHeight - chatDiv.clientHeight - 100) {
      chatDiv.appendChild(wrapper)
      chatDiv.scrollTop = chatDiv.scrollHeight - chatDiv.clientHeight
    } else {
      chatDiv.appendChild(wrapper)
    }
    return element
  })
  console.log(elements)
}

const fetchFromPriorMessages = (data) => {
  setFirstID(data.messages[0].id)
  const fragment = document.createDocumentFragment()
  const elements = data.messages.map(element => {
    if (element.id === 1) {
      const startText = createChatCenterElement('Start of the chat')
      fragment.prepend(startText)
    }
    const wrapper = createMessage(element)
    fragment.appendChild(wrapper)
    return element
  })
  const chatDiv = getById('chat')
  chatDiv.prepend(fragment)
  console.log(elements)
}

export { fetchFromConnect, fetchFromSubscribe, fetchFromPriorMessages }
