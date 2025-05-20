FROM node:22

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json files
COPY package*.json ./

# Install all dependencies. Dependencies will be cached if the package.json and package-lock.json files are not changed
RUN npm install

# Copy the source code into the container
COPY . .

# Build the React app
RUN npm run build

# Set environment port
ENV PORT=3000

# Expose port
EXPOSE 3000

# Command to run the app
CMD ["npm", "start"]