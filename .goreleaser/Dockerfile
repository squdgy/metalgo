FROM debian:11-slim AS execution

# Maintain compatibility with previous images
RUN mkdir -p /metalgo/build
WORKDIR /metalgo/build

# Copy the executables into the container
COPY build/plugins/evm plugins/evm
COPY metalgo .

CMD [ "./metalgo" ]
