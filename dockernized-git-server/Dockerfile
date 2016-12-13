FROM ubuntu

RUN apt-get -y update \
	&& apt-get -y install git \
	&& apt-get -y install openssh-server
RUN apt-get -y install git-sh

RUN adduser git
RUN mkdir /home/git/.ssh
RUN chown git /home/git/.ssh
RUN chgrp git /home/git/.ssh
RUN mkdir /opt/git
RUN chown git /opt/git
RUN chgrp git /opt/git
# mk sample project.
RUN mkdir /opt/git/project.git
RUN chown git /opt/git/project.git
RUN chgrp git /opt/git/project.git
RUN git init --bare --shared /opt/git/project.git

WORKDIR /home/git

COPY docker-entrypoint.sh /
RUN chmod +x /docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]

EXPOSE 22
