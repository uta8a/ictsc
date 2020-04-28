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
# down
sudo docker-compose down
```

```
docker #
go run main.go
vagrant $
curl 0.0.0.0:8080/hc
```
