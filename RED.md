/Users/astep/GolandProjects/filebrowser/filebrowser config set -r /Users/astep/FilesDEMO
/Users/astep/GolandProjects/filebrowser/filebrowser config set -p 8085
filebrowser -d=/Users/astep/GolandProjects/filebrowser/filebrowser.db config cat
filebrowser -d=/Users/astep/GolandProjects/filebrowser/filebrowser.db users ls

filebrowser -d=/Users/astep/GolandProjects/filebrowser/filebrowser.db -r /Users/astep/FilesDEMO -p 8081
/Users/astep/GolandProjects/filebrowser/filebrowser -d=/Users/astep/GolandProjects/filebrowser/filebrowser.db -r /Users/astep/FilesDEMO -p 8081
filebrowser -r /Users/astep/FilesDEMO -p 8081

filebrowser -r /Users/astep/FilesDEMO -p 8081

# docker
docker run \
-v /Users/astep/FilesDEMO:/srv \
-v /Users/astep/GolandProjects/filebrowser2/filebrowser.db:/database/filebrowser.db \
-v /Users/astep/GolandProjects/filebrowser2/docker/root/defaults/settings.json:/config/settings.json \
-e PUID=$(id -u) \
-e PGID=$(id -g) \
-p 8081:80 \
filebrowser/filebrowser:s6