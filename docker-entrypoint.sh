#!/bin/bash

# Start bluetooth services
service dbus start
service bluetooth start

sleep 15

# Run app
./flowercare_collector
