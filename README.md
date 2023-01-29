# Pimview Backend

This project is the back end to "pimview".

The core is a broker of whichever flavour is desired,
then the front end should have components built to suit and the sub layer should process the queue 
and translate to any iot device that isnt natively compatible with mqtt.

### `go run main.go run sub`

Ensure the broker variables are configured appropriately
