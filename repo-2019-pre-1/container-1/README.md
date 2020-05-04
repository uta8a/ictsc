# ICTSC 2019 一次予選 コンテナが作れない
- [公式解説](https://blog.icttoracon.net/2019/08/27/ictsc2019-%E4%B8%80%E6%AC%A1%E4%BA%88%E9%81%B8-%E5%95%8F%E9%A1%8C%E8%A7%A3%E8%AA%AC-%E3%82%B3%E3%83%B3%E3%83%86%E3%83%8A%E3%81%8C%E4%BD%9C%E3%82%8C%E3%81%AA%E3%81%84/)


```
~ $
vagrant up
vagrant ssh

vagrant@vagrant $
chmod +x ~/src/install/*.sh
cd ~/src/install/
sudo ./root.sh # update & install 
sudo ./install.docker.sh # docker install
sudo ./install.docker-compose.sh # docker-compose install
source ~/.bashrc
# docker-compose -v # show version
```


```
sudo docker-compose build
sudo docker-compose up -d
sudo docker-compose exec backend /bin/sh
## down
sudo docker-compose down
```

```
docker #
go run main.go
vagrant $
curl 0.0.0.0:8080/hc
```

```
host PC Access to:
http://192.168.33.10:8080/
```
