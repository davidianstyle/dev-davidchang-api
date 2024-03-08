# webhook-app/Dockerfile

# Use a Node.js base image
FROM node:14

# Create and set the working directory
WORKDIR /usr/src/app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Expose the port that your app is running on
EXPOSE 8000

# Start the application
CMD ["node", "index.js"]
