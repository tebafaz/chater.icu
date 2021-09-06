sendButtonClickListener = async () => {
    getById('message-error').textContent = ""
    var res = await sendMessage()
    if(res.error !== undefined){
        getById('message-error').textContent = res.error
    }
    getById('message-area').value = ""
}

sendButtonKeydownListener =  async (event) => {
    if (event.key == "Enter" && getById('message-area').value != "") {
        getById('message-error').textContent = ""
        var res = await sendMessage()
        if(res.error !== undefined){
            getById('message-error').textContent = res.error
        }
        getById('message-area').value = ""
    }
}

sendButtonKeyupListener = (event) => {
    if (event.key == "Enter") {
        getById('message-area').value = ""
    }
}

deleteMessageListener = async (event) => {
    getById('message-error').textContent = ""
    var res = await deleteMessage(parseInt(event.target.parentElement.id))
    if(res.error !== undefined){
        getById('message-error').textContent = res.error
    }
    
}

scrollListener = async () => {
    
    if (getById('chat').scrollTop == 0 && firstID != 1) {
        getById('message-error').textContent = ""
        var height = getById('chat').scrollHeight
        var res = await fetcher(`/api/v1/messages?last_id=${firstID}&limit=50`, fetchFromPriorMessages)
        console.log(`asa ${res}`)
        if (res != null && res.error !== undefined) {
            getById('message-error').textContent = res.error
            return
        }
        getById('chat').scrollTo(0, getById('chat').scrollHeight - height)
    }
}

loginClickListener = async () => {
    getById('login-error').innerHTML = ""
    var res = await login()
    if (res.error !== undefined) {
        getById('login-error').textContent = res.error
        return
    }
    setAuthUserCookie(res.token, getById('login-username').value)
    checkSession()
    getById('login-modal').style.display = "none"
    getById('login-username').value = ""
    getById('login-password').value = ""
}

loginKeydownListener = async (event) => {
    if (event.key == "Enter" && getById('login-username').value != "" && getById('login-password').value != "") {
        getById('login-error').innerHTML = ""
        var res = await login()
        if (res.error !== undefined) {
            getById('login-error').textContent = res.error
            return
        }
        setAuthUserCookie(res.token, getById('login-username').value)
        checkSession()
        getById('login-modal').style.display = "none"
        getById('login-username').value = ""
        getById('login-password').value = ""    
    }
}

registerClickListener = async () => {
    getById('register-error').innerHTML = ""
    if (getById('register-password').value != getById('register-password-retype').value && getById('register-password').value != "") {
        getById('register-error').innerHTML = "passwords don't match"
        return
    }
    var res = await register()
    if (res.error !== undefined) {
        getById('register-error').textContent = res.error
        return
    }
    setAuthUserCookie(res.token, getById('register-username').value)
    checkSession()
    getById('register-modal').style.display = "none"
    getById('register-username').value = ""
    getById('register-password').value = ""

}

registerKeydownListener = async (event) => {
    if (event.key == "Enter" && getById('register-username').value != "" && getById('register-password').value != "") {
        getById('register-error').innerHTML = ""
    if (getById('register-password').value != getById('register-password-retype').value && getById('register-password').value != "") {
        getById('register-error').innerHTML = "passwords don't match"
        return
    }
    var res = await register()
    if (res.error !== undefined) {
        getById('register-error').textContent = res.error
        return
    }
    setAuthUserCookie(res.token, getById('register-username').value)
    checkSession()
    getById('register-modal').style.display = "none"
    getById('register-username').value = ""
    getById('register-password').value = ""

    }
}

logoutListener = async () => {
    getById('message-error').textContent = ""
    var res = await logout()
    if (res.error !== undefined) {
        getById('message-error').textContent = res.error
        return
    }
    deleteAuthUserCookie()
    checkSession()
}

var loginModal = getById('login-modal')
var registerModal = getById('register-modal')

var closeLogin = getById('close-login')
var closeRegister = getById('close-register')

getById('login-button').onclick = function() {
    loginModal.style.display = "block"
    getById('login-username').focus()
}

getById('register-button').onclick = function() {
    registerModal.style.display = "block"
    getById('register-username').focus()
}

closeLogin.onclick = function() {
    loginModal.style.display = "none"
}

closeRegister.onclick = function() {
    registerModal.style.display = "none"
}

window.onmousedown = function(event) {
    if (event.target == loginModal) {
        loginModal.style.display = "none"
    }
    if (event.target == registerModal) {
        registerModal.style.display = "none"
    } 
}