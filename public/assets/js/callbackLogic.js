let firstID
let lastID
let stateRegistered = false

const checkSession = () => {
  if (getCookie('username') != null && getCookie('Authorization') != null) {
    stateRegistered = true
    getById('rule-1').style.display = 'none'
    getById('rule-3').style.display = 'none'
    getById('username-field').style.display = 'none'
    getById('login-register-span').style.display = 'none'
    getById('logout').style.display = 'unset'
    getById('username-label').textContent = `username: ${getCookie('username')}`
    updateChat()
  } else {
    stateRegistered = false
    getById('rule-1').style.display = 'unset'
    getById('rule-3').style.display = 'unset'
    getById('username-field').style.display = 'unset'
    getById('login-register-span').style.display = 'unset'
    getById('logout').style.display = 'none'
    getById('username-label').textContent = 'username: '
    updateChat()
  }
}

const updateChat = () => {
  if (stateRegistered) {
    document.querySelectorAll(`[username="${getCookie('username')}"]`).forEach((element) => {
      element.firstChild.setAttribute('class', 'my-username')
      element.querySelector('.message-delete').style.display = 'unset'
    })
  } else {
    document.querySelectorAll('[registered="true"]').forEach((element) => {
      if (element.getAttribute('username') === 'tebafaz') {
        element.firstChild.setAttribute('class', 'priveleged-username')
        return
      }
      element.firstChild.setAttribute('class', 'registered-username')
      element.querySelector('.message-delete').style.display = 'none'
    })
  }
}

async function main () {
  checkSession()
  let err = await fetcher('/api/v1/last-messages', fetchFromConnect)
  if (err !== undefined) {
    console.log('cannot connect to server')
  }
  let fewRequests = true
  const tooManyRequest = 429
  try {
    while (fewRequests) {
      err = await fetcherForSub(`/api/v1/subscribe?id=${lastID + 1}`, fetchFromSubscribe)
      if (err !== undefined && err.status === tooManyRequest) {
        console.log('cannot subscribe to server')
        fewRequests = false
      }
    }
  } catch (error) {
    console.log(err)
    return
  }

  setInterval(async () => {
    await fetcherForSub(`/api/v1/subscribe?id=${lastID + 1}`, fetchFromSubscribe)
  }, 1000)
}

window.addEventListener('load', () => {
  main()
})
