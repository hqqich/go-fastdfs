import requests

url = 'http://172.22.90.1:8080/group1/upload'
files = {'file': open('fileserver.exe', 'rb')}
options={'output':'json','path':'','scene':''} #参阅浏览器上传的选项
r = requests.post(url,data=options, files=files)
print(r.text)