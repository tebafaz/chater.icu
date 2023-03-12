const deleteMessage = async (id) => {
  const data = {
    id: id
  }
  headers = getHeaders(true)
  return await dataSend('/api/v1/user/message', data, 'DELETE', headers)
}

const login = async () => {
  const data = {
    username: getById('login-username').value,
    password: getById('login-password').value
  }
  headers = getHeaders(false)
  return await dataSend('/api/v1/login', data, 'POST',  headers)
}

const register = async () => {
  const data = {
    username: getById('register-username').value,
    password: getById('register-password').value
  }
  headers = getHeaders(false)
  return await dataSend('/api/v1/register', data, 'POST', headers)
}

const logout = async () => {
  if(getCookie('username') == null || getCookie('Authorization') == null) {
    state_registered = false
    deleteAuthUserCookie()
    return
  }
  headers = getHeaders(true)
  deleteAuthUserCookie()
  return await dataSend('/api/v1/user/logout', null, 'POST', headers)
}

const sendMessage = async () => {
  headers = getHeaders(state_registered)
  if (state_registered) {
    const data = {
      message: getById('message-area').value
    }
  return await dataSend('/api/v1/user/message', data, 'POST', headers)
  } else {
    const data = {
      username: getById('username-field').value,
      message: getById('message-area').value
    }
    return await dataSend('/api/v1/guest/message', data, 'POST', headers)
  }
}