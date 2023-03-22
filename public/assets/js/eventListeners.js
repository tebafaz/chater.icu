const sendButtonClickListener = async () => {
  getById('message-error').textContent = ''
  const res = await sendMessage()
  if (res.error !== undefined) {
    getById('message-error').textContent = res.error
  }
  getById('message-area').value = ''
}

const sendButtonKeydownListener = async (event) => {
  if (event.key === 'Enter' && getById('message-area').value !== '') {
    getById('message-error').textContent = ''
    const res = await sendMessage()
    if (res.error !== undefined) {
      getById('message-error').textContent = res.error
    }
    getById('message-area').value = ''
  }
}

const sendButtonKeyupListener = (event) => {
  if (event.key === 'Enter') {
    getById('message-area').value = ''
  }
}

const deleteMessageListener = async (event) => {
  getById('message-error').textContent = ''
  const res = await deleteMessage(parseInt(event.target.parentElement.id))
  if (res.error !== undefined) {
    getById('message-error').textContent = res.error
  }
}

const scrollListener = async () => {
  if (getById('chat').scrollTop === 0 && getFirstID() !== 1) {
    getById('message-error').textContent = ''
    const height = getById('chat').scrollHeight
    const res = await fetcher(`/api/v1/messages?last_id=${getFirstID()}&limit=50`, fetchFromPriorMessages)
    console.log(`asa ${res}`)
    if (res != null && res.error !== undefined) {
      getById('message-error').textContent = res.error
      return
    }
    getById('chat').scrollTo(0, getById('chat').scrollHeight - height)
  }
}

const loginClickListener = async () => {
  await loginHandler()
}

const loginKeydownListener = async (event) => {
  if (event.key === 'Enter' && getById('login-username').value !== '' && getById('login-password').value !== '') {
    await loginHandler()
  }
}

const loginHandler = async () => {
  getById('login-error').innerHTML = ''
  const res = await login()
  if (res.error !== undefined) {
    getById('login-error').textContent = res.error
    return
  }
  setAuthUserCookie(res.token, getById('login-username').value)
  checkSession()
  getById('login-modal').style.display = 'none'
  getById('login-username').value = ''
  getById('login-password').value = ''
}

const registerClickListener = async () => {
  await registerHandler()
}

const registerKeydownListener = async (event) => {
  if (event.key === 'Enter' && getById('register-username').value !== '' && getById('register-password').value !== '') {
    await registerHandler()
  }
}

const registerHandler = async () => {
  getById('register-error').innerHTML = ''
  if (getById('register-password').value !== getById('register-password-retype').value && getById('register-password').value !== '') {
    getById('register-error').innerHTML = "passwords don't match"
    return
  }
  const res = await register()
  if (res.error !== undefined) {
    getById('register-error').textContent = res.error
    return
  }
  setAuthUserCookie(res.token, getById('register-username').value)
  checkSession()
  getById('register-modal').style.display = 'none'
  getById('register-username').value = ''
  getById('register-password').value = ''
}

const logoutListener = async () => {
  getById('message-error').textContent = ''
  const res = await logout()
  if (res.error !== undefined) {
    getById('message-error').textContent = res.error
    return
  }
  deleteAuthUserCookie()
  checkSession()
}

const loginModal = getById('login-modal')
const registerModal = getById('register-modal')

const closeLogin = getById('close-login')
const closeRegister = getById('close-register')

getById('login-button').onclick = function () {
  loginModal.style.display = 'block'
  getById('login-username').focus()
}

getById('register-button').onclick = function () {
  registerModal.style.display = 'block'
  getById('register-username').focus()
}

closeLogin.onclick = function () {
  loginModal.style.display = 'none'
}

closeRegister.onclick = function () {
  registerModal.style.display = 'none'
}

window.onmousedown = function (event) {
  if (event.target === loginModal) {
    loginModal.style.display = 'none'
  }
  if (event.target === registerModal) {
    registerModal.style.display = 'none'
  }
}
