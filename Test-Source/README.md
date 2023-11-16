# Setting Up the Test Environment

## Clone the Repository
- First, clone the repository from [Reach Industries Backend Test](https://github.com/Reach-Industries/backendtest.git) to your local machine.
- Use the command: `git clone https://github.com/Reach-Industries/backendtest.git`

## Navigate to the Project Directory
- After cloning, navigate to the project directory: `cd backendtest`

## Build the Switcher Service
- Inside the repository, you'll find a `switcher` folder. This contains the Dockerfile necessary to build the switcher image.
- Navigate to the switcher directory: `cd switcher`
- Build the Docker image for the switcher service using: `docker build -t switcher-service .`
  - Ensure that Docker is running on your machine before executing this command.

## Run the Docker Compose
- Return to the root directory of the project: `cd ..`
- Use Docker Compose to set up the entire environment. This will start all the necessary services, including Kafka, Zookeeper, and the switcher service you just built.
- Run the command: `docker-compose up`
- This step will take some time as it needs to download and set up the required Docker images.

## Verify the Setup
- Once the Docker Compose process is complete, and all services are up and running, you can verify the setup.
- Open your web browser and navigate to `localhost:2181`. This is the default port for Zookeeper, which is part of your Kafka setup.
- Using a tool like Offset Explorer (previously known as Kafka Tool), you can connect to the Kafka instance and monitor the topics and messages. Make sure you have Offset Explorer installed on your machine.
- In Offset Explorer, connect to the Kafka instance using the provided Zookeeper URL (localhost:2181) to observe the messages being produced and consumed in the Kafka topic.

## Writing and Testing Your Service
- Now that your environment is set up, you can start writing your service as per the test requirements.
- After writing your service, integrate it within the Docker Compose environment and ensure it interacts correctly with the Kafka setup.

## Troubleshooting
- If you encounter any issues during setup, double-check the steps and ensure all commands are executed in the correct directories.
- Ensure Docker and Docker Compose are correctly installed and updated on your system.
- If Kafka or Zookeeper services are not running as expected, check the docker-compose logs for any error messages.

## Submitting Your Work
- After completing your service, follow the submission guidelines provided in the test instructions.
