<p align="center">
<img src=https://via.placeholder.com/150/>
</p>

***Replace any content text in bold.***

# Test Introduction

At Reach Industries, we are building innovative products that leverage high-volume data streaming and cloud based AI. Our LabEye camera systems are specifically designed for lab environments where precision and detail are paramount. Our customers seamlessly integrate LabEyes into their setups to capture high-quality visual data. This raw data is then securely transmitted to Lumi, our advanced cloud-based platform.

Lumi harnesses the power of state-of-the-art algorithms to process this data, converting the visual inputs into actionable quantitative values. One of the standout features of our system is its ability to overlay the derived data directly onto the original video in real-time. This not only allows researchers and lab technicians to see the data but also to interactively visualize the insights and patterns that Lumi has discerned, enhancing their understanding and decision-making processes.

This assessment focuses on developing a Golang-based backend service, primarily handling and responding to high-volume data streams through Kafka. Your task is to architect and implement a solution that not only meets the immediate requirements but is also structured for future enhancements and team collaboration. Although your code won't be integrated into a live system, it's essential to approach this task with an eye towards long-term maintainability and scalability. Consider your design and coding choices as foundational elements that could be built upon and refined in a collaborative, evolving development environment.

# Task

## Task Description

Your objective is to create a Golang service that implements a multifunctional counter, demonstrating proficiency in data streaming and message handling using Kafka. This service involves several key components and modes of operation, each responding dynamically to Kafka messages.

### Core Features
1. **Counter Functionality**: The service should include a continuously incrementing counter, increasing by one every second.

2. **Kafka Message Handling**: It must listen to a Kafka message topic named "switch". The messages received on this topic will dictate the service's operation mode. The possible messages are:
   - `Counter`
   - `Time`
   - `Text`
   - `Dump`

3. **Dynamic Response Modes**:
   - **Counter Mode**: Initially, the service operates in this mode, outputting the current value of the counter to another Kafka topic every 3 seconds.
   - **Time Mode**: Upon receiving a "Time" message, the service switches to outputting the current system time every 3 seconds.
   - **Text Mode**: Receiving a "Text" message triggers the service to generate and output a random text string (16 characters) every 3 seconds.
   - **Dump Mode**: The "Dump" message prompts the service to compile all previously generated outputs into a file, timestamped and organized.

4. **Endpoint for Data Retrieval**: The service should also offer a RESTful endpoint (`/getDump`) that allows for retrieval of the saved dump file. The data should be returned as a JSON response.

### Message Format
- The Kafka messages follow this structure: `{"org": "2", "msg": "Counter"}`. The `org` field will be used in the bonus task.

### Bonus Tasks (only if you have time)
- **Enhanced Organizational Output**: Modify the service to handle separate output threads for each organization ID (`org`) received in the Kafka messages. This requires the service to initiate an output thread for a specific `org` upon receiving its first message.
- **Secured Data Access**: Implement JWT authentication for the `/getDump` endpoint. Validate the incoming requests using a JWT secret, and modify the endpoint to return the appropriate dump file based on the `org` claim within the JWT. (Note: The JWT secret should be communicated separately for testing purposes.)

This task is designed to challenge your skills in backend development, focusing on real-time data processing and adaptive system behavior. Your solution should not only meet the functional requirements but also exhibit clean, efficient, and maintainable code practices.


# Setup & Environment

***Guidance Text: Detail any setup steps, required software, or specific environment configurations necessary to undertake the test. If there are any dependencies to be installed or initial steps to be taken before the task, mention them here.***

# Deliverables & How to submit

***Guidance Text: Define clearly what the candidate should submit at the end of the test. For instance, do they need to submit code, a report, a presentation, or all of the above? Additionally, provide a clear procedure on how they should submit their work. This could be via a pull request, email, a specific platform, etc.***

# Evaluation Criteria

***Guidance Text: Explain how you will be evaluating their submissions. What are the key criteria or metrics you will be looking at? This helps the candidate focus on what's essential for the test.***

# FAQ
## 1.

## 2. 

## 3.

# Troubleshooting & Support
Test@Reach.Industries?

# Terms & Conditions for Interview Tests

## 1. Purpose of the Test
This test is designed to evaluate the technical prowess, problem-solving skills, and overall knowledge of the applicant. The test tasks and questions reflect challenges and scenarios the company has already addressed internally. Our goal is to gauge your approach and compare it with solutions we have already devised.

## 2. Use of Test Results
- We emphasize that the tasks presented in the tests have already been achieved by our internal team to a high standard.
- The solutions provided by candidates will purely be for evaluation purposes and will not be used in any of our products or services.
- Any resemblance between our products and a test submission is purely coincidental. We have no intention or motive to replicate your solution in our offerings.
- We will not store or retain your test results post-evaluation.

## 3. Feedback and Next Steps
Feedback may or may not be provided post-test based on company discretion. Candidates are encouraged to engage in the review meeting for a detailed understanding of their performance.

## 4. Modifications & Updates
Reach Industries reserves the right to modify these terms and conditions without prior notice. It's advised to review these T&Cs before proceeding with the test.

## 5. Consent
By participating in this test, candidates agree to these terms and conditions.
