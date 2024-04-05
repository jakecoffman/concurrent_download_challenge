// Run with Node 21

const urls = [
    "http://google.com",
    "http://python.org",
    "http://ruby-lang.org",
    "http://golang.org",
]

setTimeout(() => {
    process.exit(0)
}, 500)

const promises = urls.map(url => {
    return new Promise(async (resolve, reject) => {
        const start = Date.now()
        const response = await fetch(url)
        const body = await response.text()
        const time = (Date.now() - start) / 1000
        resolve({ url, body, time })
    })
})

const results = await Promise.all(promises)
results.forEach(result => {
    console.log(result.url, result.body.length, result.time)
})
