#!/bin/bash

# Make the palworld-admin executable
chmod u+x /root/palworld-admin

# Create the systemd service file
echo "[Unit]
Description=Palworld Admin Service
After=network.target

[Service]
ExecStart=/root/palworld-admin
WorkingDirectory=/root
User=root
Group=root
Restart=always

[Install]
WantedBy=multi-user.target" | sudo tee /etc/systemd/system/palworld-admin.service

# Start and enable the service
sudo systemctl start palworld-admin
sudo systemctl enable palworld-admin

# Check the service status
sudo systemctl status palworld-admin
