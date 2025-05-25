FROM node:18-alpine

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the source code
COPY . .

# Install Vue CLI globally
RUN npm install -g @vue/cli

# Build the application
RUN npm run build

# Install serve to run the application
RUN npm install -g serve

# Expose port
EXPOSE 3000

# Command to run the application
CMD ["serve", "-s", "dist", "-l", "3000"] 