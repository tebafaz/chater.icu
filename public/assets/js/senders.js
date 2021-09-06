deleteMessage = async (id) => {
    let data = {
        id: id
    }
    headers = getHeaders(true)
    return await dataSend('/api/v1/user/message', data, 'DELETE', headers)
}

login = async () => {
    var data = {
        username: getById('login-username').value,
        password: getById('login-password').value
    }
    headers = getHeaders(false)
    return await dataSend('/api/v1/login', data, 'POST',  headers)
}

register = async () => {
    var data = {
        username: getById('register-username').value,
        password: getById('register-password').value
    }
    headers = getHeaders(false)
    return await dataSend('/api/v1/register', data, 'POST', headers)
}

logout = async () => {
    if(getCookie('username') == null || getCookie('Authorization') == null) {
        state_registered = false
        deleteAuthUserCookie()
        return
    }
    headers = getHeaders(true)
    deleteAuthUserCookie()
    return await dataSend('/api/v1/user/logout', null, 'POST', headers)
}

sendMessage = async () => {
    headers = getHeaders(state_registered)
    if (state_registered) {
        var data = {
            message: getById('message-area').value
        }
        return await dataSend('/api/v1/user/message', data, 'POST', headers)
    } else {
        var data = {
            username: getById('username-field').value,
            message: getById('message-area').value
        }
        return await dataSend('/api/v1/guest/message', data, 'POST', headers)
    }
}