FROM node:latest
WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY package.json ./
COPY package-lock.json ./
COPY bootstrap.sh ./bootstrap.sh
RUN chmod +x ./bootstrap.sh

#TODO issue with react try running commented with explicit react dnd install and see if error can be reproduced. 

RUN npm install
RUN npm install react-scripts url
RUN npm install mqtt
RUN npm install url
RUN npm i buffer process
RUN npm install @mui/material
RUN npm install @emotion/styled

COPY . ./

CMD ["./bootstrap.sh"]