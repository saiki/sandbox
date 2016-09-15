FROM node

RUN git clone https://github.com/crowi/crowi.git /opt/crowi
RUN cd /opt/crowi && npm install

EXPOSE 3000

ENV MONGO_URI mongodb://db/crowi

WORKDIR /opt/crowi

#ENTRYPOINT ["npm start"]
CMD cd /opt/crowi && npm start
