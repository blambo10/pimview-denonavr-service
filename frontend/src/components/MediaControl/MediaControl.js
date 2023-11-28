import { useState } from 'react';
import * as React from 'react';
import { useTheme, styled } from '@mui/material/styles';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import VolumeDownIcon from '@mui/icons-material/VolumeDown';
import VolumeUpIcon from '@mui/icons-material/VolumeUp';
import VolumeOffIcon from '@mui/icons-material/VolumeOff';
import { Slider } from '@mui/material';

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

function HandleClick(e, client) {
    console.log("clicked");
  
    switch(e) {
      case 'volumeup':
        client.publish("denonavr/volume", "up");
      break;
      case 'volumedown':
        client.publish("denonavr/volume", "down");
      break;
      case 'mute':
        client.publish("denonavr/volume", "mute");
      break;
      case 'remoteup':
        client.publish("denonavr/remote", "up");
      break;
      case 'remotedown':
        client.publish("denonavr/remote", "down");
      break;
      case 'remoteok':
        client.publish("denonavr/remote", "ok");
      break;
      default:
    }
}

export default function DenonAVRDevice(props) {
  const theme = useTheme();
  const topic = "denonavr/currentvolume";

  const [updatingSlider, setUpdatingSlider] = useState(false);

  // const handleSliderChange = (event, newValue) => {
  //     props.updateVolume(newValue);
  // };

  const handleSliderCommit = (event, newValue) => {
    setUpdatingSlider(true);
    props.handleSliderChange(newValue)
    setUpdatingSlider(false);

    console.log("Event: " + event)

    console.log(newValue);
  };

  return (
    <Widget>
        <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
          {props.Title}
        </Typography>
        <Typography sx={{ display: 'flex', alignItems: 'right', pl: 1, pb: 1 }} variant="body2">
            
        <IconButton aria-label="volumedown" id="volumedown" onClick={(e) => HandleClick("volumedown", props.mqttclient)}>
          <VolumeDownIcon id="volumedown" />
        </IconButton>
        <Slider 
            aria-label="Small" 
            value={props.volume}
            onChange={props.handleSliderChange}
            onChangeCommitted={props.handleSliderCommit} 
            valueLabelDisplay="auto" 
            min={0}
            max={75}
            display='flex' 
            step={0.5}
        />
        <IconButton aria-label="volumeup" id="volumeup" onClick={(e) => HandleClick("volumeup", props.mqttclient)}>
          <VolumeUpIcon id="volumeup"/>
        </IconButton>
        <IconButton aria-label="mute" id="mute" onClick={((e) => HandleClick("mute", props.mqttclient))}>
            <VolumeOffIcon sx={{}} />
        </IconButton>
        </Typography>
    </Widget>
  );
}
