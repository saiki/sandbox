FROM node

WORKDIR /opt/crowi
RUN git clone https://github.com/crowi/crowi.git /opt/crowi
RUN npm install -g webpack && npm install -g gulp && npm install && gulp && webpack

ENV MONGO_URI mongodb://db/crowi

COPY docker-entry.sh /opt/crowi
RUN chmod +x /opt/crowi/docker-entry.sh

ENTRYPOINT ["/opt/crowi/docker-entry.sh"]

EXPOSE 3000
