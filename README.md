# Test Introduction

At Reach Industries, we are building innovative products that leverage high-volume data streaming and cloud based AI. Our LabEye camera systems are specifically designed for lab environments where precision and detail are paramount. Our customers seamlessly integrate LabEyes into their setups to capture high-quality visual data. This raw data is then securely transmitted to Lumi, our advanced cloud-based platform.

Lumi harnesses the power of state-of-the-art algorithms to process this data, converting the visual inputs into actionable quantitative values. One of the standout features of our system is its ability to overlay the derived data directly onto the original video in real-time. This not only allows researchers and lab technicians to see the data but also to interactively visualize the insights and patterns that Lumi has discerned, enhancing their understanding and decision-making processes.

This assessment challenges you to develop a Golang-based backend service focused on efficiently running and managing parallel data streams, orchestrated through a Kafka message queue. It is essential to design the service with scalability in mind, accommodating an increasing number of organizations and escalating data generation, ensuring seamless future expansion. Your task is to architect and implement a solution that not only meets the immediate requirements but is also structured for future enhancements and team collaboration. Although your code won't be integrated into a live system, it's essential to approach this task with an eye towards long-term maintainability and scalability. Consider your design and coding choices as foundational elements that could be built upon and refined in a collaborative, evolving development environment.

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

- **Languages**: Golang
- **Container** https://github.com/Reach-Industries/Backend-Golang-Developer-Test/tree/main/Test-Source
- **Technologies**: Docker, Kafka

# Deliverables & How to submit

- Code uploaded to GitHub (either as a public repo or private repo).
- The repo should be easily runnable or buildable without additional intervention from the evaluators.
- Detailed notes in the repository's README.md. Explain the libraries you've used, any resources you've referenced, your design approach, and any constraints you faced. Your insights and the clarity of your explanation in this section are critical.

# Evaluation Criteria

You will be marked on:
- Your code quality.
- Your project organisation.
- The efficiency of your approach to providing a solution.
- The rationale you provide for the approach you have taken (in comments or in supporting documents).

# FAQ
## 1. Do I need to set up any specific environment or version of Golang and Docker for this assessment?
- No, as long as you include version numbers and it builds. (Golang is already installed in the provided container

## 2. Does the JWT need to be a real secret
- No this can be any string, but please provide it in the documentation for testing. 

## 3. Can I use libraries or packages to complete technical parts of the test?
- Absolutely. Please approach this test in the same way you would approach a standard piece of work. As stated previously, the only stipulation is you must use Golang and the provided docker container.

## 4. What should I do if I'm taking longer than 4-5 hours to complete the test?

We understand that the scope of this test might require more time than initially anticipated. If you find yourself exceeding the 4-5 hour mark, you have two options:

#### Option 1: Continue Until the Deadline
If you feel that you can complete the test within the allotted deadline, you are encouraged to continue. The comprehensive completion of tasks is appreciated and can be indicative of your dedication and thorough approach.

#### Option 2: Summarize Remaining Work
If completing the test is not feasible within the deadline, we value a summary of how you would have proceeded. Please include

- **Detailed Plan**: Outline the steps you would take to complete the remaining tasks.
- **Methods and Tools**: Describe the methods and tools you would employ to tackle the remaining challenges.
- **Time Estimation**: Provide an estimated timeline for these tasks.
- **Potential Challenges**: Anticipate any potential challenges or roadblocks you might encounter and how you would address them.
- **Final Objectives**: Clarify the objectives you would aim to achieve with the completed work.

This approach allows us to understand your problem-solving process, planning skills, and how you manage time and priorities. It also gives us insight into your technical knowledge and approach to unforeseen challenges.

Remember, the quality of your work and the thought process behind it are more critical than the quantity. We are looking for insights into your approach and decision-making process, which can be effectively demonstrated through either a completed test or a well-articulated plan for unfinished parts.

# Troubleshooting & Support
Test@Reach.Industries

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
