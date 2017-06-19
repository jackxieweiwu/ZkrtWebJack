# coding=utf-8
import requests
import time
import os, sys, re
import subprocess

filePath = "d:\\TestServerApi\\"
ServerIP = "d:\\TestServerApi\\"

def testApi():
    num = 0
    while True:
        num += 1;
        gps = requests.get("http://localhost:8080/v1/drone_gps/",
                           params={'drone_lat': '22.5865488', 'drone_lng': "132.356844"}).text;

        drone_message = requests.get("http://localhost:8080/v1/drone_message/",
                                     params={'drone_alt': '150', 'drone_y': '3.3', 'drone_r': '0.0',
                                             'drone_p': '1.2'}).text;
        if (len(gps) > 35):
            print "---------*****GPS*******------------"
            print gps
            print "------------------"
        else:
            print 'no(gps)'
            #logWrite(gps,"/////GPS////////------num")
            #NetWrodApi()

        if (len(drone_message) > 75):
            print "---------******Message******------------"
            print drone_message
            print "------------------"
        else:
            print 'no(message)'
            #logWrite(drone_message, "/////Message////////------num")
            #NetWrodApi()

def mkdir():
    import os
    path = filePath.strip()
    path = path.rstrip("\\")
    isExists = os.path.exists(path)
    if not isExists:
        print path + ' 创建成功'
        os.makedirs(path)
        return True
    else:
        print path + ' 目录已存在'
        return False

    open(path+'log.txt', 'w')

def NetWrodApi():
    try:
        p = subprocess.Popen(["ping -c 1 -w 1 " + ServerIP], stdout=subprocess.PIPE, stderr=subprocess.PIPE, shell=True)
        out = p.stdout.read()
        regex = re.compile('100% packet loss')
        if len(regex.findall(out)) == 0:
            print ServerIP + ': host up'
            return 'UP'
        else:
            print ServerIP + ': host down'
            return 'DOWN'
    except:
        logWrite("NetCheck work error!","/////netWork")
        return 'ERR'

def logWrite(error,type):
    f = open(filePath + 'log.txt', 'a')
    f.write(str(time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(time.time())) + error+type))
    f.write('\n')
    f.close()

if __name__ == '__main__':
    #mkdir()
   # NetWrodApi()
    testApi()