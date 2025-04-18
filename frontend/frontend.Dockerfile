# Using an existing node alpine build as a base
FROM node:alpine3.21

# Creating an application directory
RUN mkdir -p /app

# Setting /app to be the working directory
WORKDIR /app

# Copy the config files containing the list of dependencies
COPY frontend/package*.json .

# Install all the node packages onto the container dir
RUN npm run install

# Copy the rest of the frontend dir into the current container dir
COPY backend/ .

# Build the frontend
RUN npm run build

# Expose the port the frontend in running on
# TODO: Update the static :5173 port with this env variable
EXPOSE $PORT

# Run the frontend server
CMD ['npm', 'start']