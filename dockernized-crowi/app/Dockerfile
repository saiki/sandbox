FROM node

WORKDIR /opt/crowi
RUN git clone https://github.com/crowi/crowi.git /opt/crowi
RUN npm install -g webpack && npm install -g gulp && npm install && gulp && webpack

EXPOSE 3000

ENV MONGO_URI mongodb://db/crowi

#ENTRYPOINT ["npm start"]
CMD cd /opt/crowi && npm start
