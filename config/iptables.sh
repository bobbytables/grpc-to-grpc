#!/usr/bin/env sh

# Create a chain that goes into envoy that we can jump to for multiple listeners
# Chains can redirect to envoy by using -j ENVOY
iptables -t nat -N ENVOY
iptables -t nat -A ENVOY -p tcp -j REDIRECT --to-port 1337

# Create an output chain that eventually redirects to Envoy for all TCP traffic
iptables -t nat -N SEAPLANE_OUTPUT
iptables -t nat -A OUTPUT -p tcp -j SEAPLANE_OUTPUT

# When envoy is the one trying to do outbound traffic, let it do its thing
iptables -t nat -A SEAPLANE_OUTPUT -m owner --gid-owner 1337 -j RETURN

# If the output is attempting to go to a local container / host, just return
# so we dont go to an envoy sidecar
iptables -t nat -A SEAPLANE_OUTPUT -d 127.0.0.1/32 -j RETURN

# Otherwise, send all outbound traffic to the envoy process
iptables -t nat -A SEAPLANE_OUTPUT -j ENVOY
