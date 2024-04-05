import multiprocessing
import threading
from urllib.request import urlopen
import time


urls_to_load = [
    "http://google.com",
    "http://python.org",
    "http://ruby-lang.org",
    "http://golang.org",
]


limit = .5  # seconds


def read_url(url, q):
    now = time.time()
    data = urlopen(url).read()
    dt = time.time() - now
    q.put((url, len(data), dt))


def fetch_parallel():
    start = time.time()
    queue = multiprocessing.Queue()
    threads = [threading.Thread(target=read_url, args=(url, queue)) for url in urls_to_load]
    for t in threads:
        t.start()
    for t in threads:
        t.join(timeout=limit - (time.time() - start))
    while not queue.empty():
        print("%s %d [%.4fs]" % queue.get())
    print("Total time: %.4fs" % (time.time() - start))


if __name__ == "__main__":
    fetch_parallel()
