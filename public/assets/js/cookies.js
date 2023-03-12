function createCookie (name, value, minutes) {
  const now = new Date()
  now.setTime(now.getTime() + (minutes * 60 * 1000))
  document.cookie = name + '=' + value + '; expires=' + now.toUTCString() + '; path=/; SameSite=strict;'
}

function deleteCookie (name) {
  document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/; SameSite=strict'
}

function getCookie (name) {
  const value = `; ${document.cookie}`
  const parts = value.split(`; ${name}=`)
  if (parts.length === 2) return parts.pop().split(';').shift()
  else return null
}

function setAuthUserCookie (token, username) {
  createCookie('Authorization', token, 30)
  createCookie('username', username, 30)
}

function deleteAuthUserCookie () {
  deleteCookie('Authorization')
  deleteCookie('username')
}

function updateAuthUserCookie () {
  createCookie('Authorization', getCookie('Authorization'), 30)
  createCookie('username', getCookie('username'), 30)
}

export { createCookie, deleteCookie, getCookie, setAuthUserCookie, deleteAuthUserCookie, updateAuthUserCookie }
