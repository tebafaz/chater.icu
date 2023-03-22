async function fetcher (fetchFrom, func) {
  try {
    const res = await fetch(fetchFrom)
    const data = await res.json()
    if (res.status >= 400 && res.status <= 600) {
      return data
    }
    func(data)
  } catch (error) {
    console.log(error)
  }
  return null
}

async function fetcherForSub (fetchFrom, func) {
  try {
    const res = await fetch(fetchFrom)
    if (res.status >= 400 && res.status <= 600) {
      return res
    }
    const data = await res.json()
    if (data.timer == null) {
      func(data)
    }
  } catch (error) {
    console.log(error)
  }
  return null
}

async function dataSend (path, data, method, headers) {
  try {
    const res = await fetch(path, {
      method,
      body: JSON.stringify(data),
      headers
    })
    data = await res.json()
    if (res.status < 400 && stateRegistered) {
      updateAuthUserCookie()
    }
    return data
  } catch (error) {
    console.log(error)
  }
}

function getHeaders (isRegistered) {
  return (isRegistered
    ? {
        'Content-Type': 'application/json',
        Authorization: getCookie('Authorization')
      }
    : {
        'Content-Type': 'application/json'
      })
}

function getById (id) {
  return document.getElementById(id)
}
function createElement (element) {
  return document.createElement(element)
}
