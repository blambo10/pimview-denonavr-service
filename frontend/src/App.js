
import mqtt from "mqtt";
import { useState, useEffect } from 'react';
import * as React from 'react';
import { useTheme, styled } from '@mui/material/styles';
import MediaControl from './components/MediaControl/MediaControl'


const mqttAddress = process.env.REACT_APP_PIMVIEW_RABBITMQ_MQTT_ADDRESS;
const mqttUsername = process.env.REACT_APP_PIMVIEW_RABBITMQ_MQTT_USER;
const mqttPassword = process.env.REACT_APP_PIMVIEW_RABBITMQ_MQTT_PASSWORD;
const mqttPort = process.env.REACT_APP_PIMVIEW_RABBITMQ_MQTT_PORT;

var options = {
  
  protocol: "ws",
  username: mqttUsername,
  password: mqttPassword,
  keepalive: 20,
  // clientId uniquely identifies client
  // choose any string you wish
  clientId: "mqttjs_" + Math.random().toString(16).substr(2, 8),
};

var client = mqtt.connect("mqtt://" + mqttAddress + ":" + mqttPort + "/ws", options);

const Widget = styled('div')(({ theme }) => ({
  padding: 16,
  borderRadius: 16,
  width: 343,
  maxWidth: '100%',
  margin: 'auto',
  position: 'relative',
  zIndex: 1,
  backgroundColor:
    theme.palette.mode === 'dark' ? 'rgba(0,0,0,0.6)' : 'rgba(255,255,255,0.4)',
  backdropFilter: 'blur(40px)',
}));

export default function DenonAVRDevice(props) {
  const theme = useTheme();
  const topic = "denonavr/currentvolume";

  const [currentVolume, setCurrentVolume] = useState(5);

  const handleSliderChange = (event, newValue) => {
    setCurrentVolume(newValue);
  };

  const handleSliderCommit = (event, newValue) => {
    client.publish("denonavr/volume", newValue + '');
  };

  // Similar to componentDidMount and componentDidUpdate:
  useEffect(() => {
    console.log('checking volume')

    client.subscribe(topic, (err) => {
      if (err) {
        console.log(err);
      }
    });

    client.on("message", (topic, message) => {
      // message is Buffer
      console.log(topic);
      console.log(message.toString());
      if (currentVolume != message.toString()) {
        console.log("updating volume slider")
        setCurrentVolume(message.toString())
      }
    });
  }, []);

  return (
    <MediaControl 
      Title={"Home Theatre"} 
      handleSliderCommit={handleSliderCommit} 
      handleSliderChange={handleSliderChange}
      volume={currentVolume}>
      

    </MediaControl>
  );
}
