async function fetcher(fetchFrom, func) {
    try {
        let res = await fetch(fetchFrom)
        let data = await res.json()
        if (res.status >= 400 && res.status <= 600) {
            return data
        }
        func(data)
    } catch (error) {
        console.log(error)
    }
    return null
}

async function fetcherForSub(fetchFrom, func) {
    try {
        let res = await fetch(fetchFrom)
        if (res.status >= 400 && res.status <= 600) {
            return res
        }
        let data = await res.json()
        if (data.timer == null) {
            func(data)
        }
    } catch (error) {
        throw error
    }
    return null
}

async function dataSend(path, data, method, headers) {
    try {
        let res = await fetch(path, {
            method: method,
            body: JSON.stringify(data),
            headers: headers
        })
        data = await res.json()
        if (res.status < 400 && state_registered) {
            updateAuthUserCookie()
        }
        return data    
    } catch (error) {
        console.log(error)
    }
}

function getHeaders(is_registered) {
    return (is_registered ? {
        'Content-Type': 'application/json',
        'Authorization': getCookie('Authorization')
    } : {
        'Content-Type': 'application/json',
    })
}


function getById(id) {
    return document.getElementById(id)
}
function createElement(element) {
    return document.createElement(element)
}