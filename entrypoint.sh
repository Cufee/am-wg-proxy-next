#!/bin/sh
set -x

pid=0

# SIGUSR1-handler
cleanup_handler() {
  # This will free up a dns name
  ./tailscale logout
}

# SIGTERM-handler
term_handler() {
  if [ $pid -ne 0 ]; then
    kill -SIGTERM "$pid"
    wait "$pid"
  fi
  exit 143; # 128 + 15 -- SIGTERM
}

# setup handlers
# on callback, kill the last background process, which is `tail -f /dev/null` and execute the specified handler
trap 'kill ${!}; cleanup_handler' SIGUSR1
trap 'kill ${!}; term_handler' SIGTERM

# Run binary and tailscale
./tailscaled --tun=userspace-networking &
./tailscale up --authkey=${TAILSCALE_AUTHKEY} --advertise-tags=tag:services --hostname=${TAILSCALE_APP_NAME} &
./binary &
pid="$!"

# wait forever
while true
do
  tail -f /dev/null & wait ${!}
done