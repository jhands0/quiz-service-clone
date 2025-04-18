# Using an existing golang alpine build as a base
FROM golang:1.23.8-alpine3.21

# Creating an application directory
RUN mkdir -p /app

# Setting /app to be the working directory
WORKDIR /app

# Copy the backend dir into the current container dir
COPY backend/ .

# Build the backend
RUN go build

# Expose the port the backend in running on
# TODO: Update the static :3000 port with this env variable
EXPOSE $PORT

# Run the backend server
CMD ['go', 'run', '.']