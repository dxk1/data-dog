# 拉取基础镜像
FROM csighub.tencentyun.com/scf/component:scf_baseimage_v9

# set
LABEL maintainer="kangding <kangding@tencent.com>"

WORKDIR /usr/local/services/data_dog

COPY entrypoint.sh /etc/kickStart.d/entrypoint.sh
COPY gracestop.sh /usr/local/services/data_dog/gracestop.sh
COPY DataDog /usr/local/services/data_dog/DataDog
