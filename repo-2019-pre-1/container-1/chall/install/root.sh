#!/bin/bash
set -ue
apt update && \
DEBIAN_FRONTEND=noninteractive apt upgrade -y && \
DEBIAN_FRONTEND=noninteractive apt install -y --fix-missing \
build-essential \
git \
file \
grep \
zip \
ssh \
python3-dev \
python3-pip \
tcpdump \
wget \
curl
