#! /bin/sh
touch .env

if [ $# -ne 0 ]; then
  echo "===> Overriding env params with args ..."
  for var in "$@"
  do
    export "$var"
  done
fi

cd /usr/share/nginx/html/

sed -i "s/{{MQTT_ADDRESS}}/$REACT_APP_PIMVIEW_RABBITMQ_MQTT_ADDRESS/g" *
sed -i "s/{{MQTT_USERNAME}}/$REACT_APP_PIMVIEW_RABBITMQ_MQTT_USER/g" *
sed -i "s/{{MQTT_PASSWORD}}/$REACT_APP_PIMVIEW_RABBITMQ_MQTT_PASSWORD/g" *
sed -i "s/{{MQTT_PORT}}/$REACT_APP_PIMVIEW_RABBITMQ_MQTT_PORT/g" *