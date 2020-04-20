# !/bin/bash
kill -9 $(pgrep devops)
cd ~/devops
git pull https://github.com/laihezhao/devops.git
chmod +x devops
./devops &