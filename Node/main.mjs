// Run with Node 21

const urls = [
    "http://google.com",
    "http://python.org",
    "http://ruby-lang.org",
    "http://golang.org",
]

const startTime = Date.now()
process.on("exit", function() {
    console.log("Total time:", (Date.now() - startTime) / 1000)
})

setTimeout(() => {
    process.exit(0)
}, 500)

const promises = urls.map(url => {
    return new Promise(async (resolve, reject) => {
        const start = Date.now()
        const response = await fetch(url)
        const body = await response.text()
        const time = (Date.now() - start) / 1000
        console.log(url, body.length, time)
    })
})

await Promise.all(promises)
