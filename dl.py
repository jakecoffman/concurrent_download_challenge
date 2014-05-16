import multiprocessing
import urllib2
import Queue
import time


urls_to_load = [
    "http://google.com",
    "http://python.org",
    "http://ruby-lang.org",
    "http://golang.org",
]


limit = .5  # seconds


def read_url(url, queue):
    now = time.time()
    data = urllib2.urlopen(url).read()
    dt = time.time() - now
    queue.put((url, len(data), dt))


def fetch_parallel():
    start = time.time()
    queue = multiprocessing.Queue()
    processes = [multiprocessing.Process(target=read_url, args=(url, queue)) for url in urls_to_load]
    for p in processes:
        p.start()
    for p in processes:
        p.join(timeout=limit - (time.time() - start))
        if p.is_alive():
            p.terminate()
        if not queue.empty():
            print "%s %d [%.4fs]" % queue.get()
    print "Total time: %.4fs" % (time.time() - start)


if __name__ == "__main__":
    fetch_parallel()
