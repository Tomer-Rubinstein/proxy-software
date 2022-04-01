import requests

proxies = {'http': 'http://127.0.0.1:8000/'}
r = requests.get('http://info.cern.ch/', proxies=proxies)

print(r.text)
