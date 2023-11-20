# week9
This assignment is to explore the code generator and ai assist tool in recreation of assignment AQ

Under cmd, the folder ai_generate contain main.go file that generated with ChatGPT 3.5 and in ai_assist folder, the main.go file is modified from week2 assignment using GitHub copilot and unit test is created with go generator tool. 

## Code Generator and AI assisted tool
The tools used in programming are GitHub copilot which is the extension in VS code and go generator which is standard library that also in vs code.

With orginal assignment 2 code, I started the conversation with Copilot by asking suggestion to revise the code to be more idiomatic. The feedback from GitHub Copilot are straightforward as: Use time.Since instead of manually calculating the time difference. 2 Use range instead of traditional for loop if you don't need the index.3 Use := for variable declaration and assignment if the variable is not declared before. 3 Group the data into a slice and iterate over it, instead of calling the function separately for each data set.

There were some mistakes that generated from Copilot when it changes the data type as it does not also change the function input which lead to invalid functions. 

As for unit testing, go generator that in standard library has really come in handy which is big surprise for me because of the testing code is very standard written that able to apply in multiple cases for testing. Go generator is more complicated to use as first has throughout understanding of the target programming task and object and then understand the go generator function which I think is better to code yourself than benefit from go generation. 

I do see the possibility of using code generation tool and AI assistant programming as it will help to create functions or data type as individuel task.

## AI-generated programming

In the folder ai_generate, the main.go is created by ChatGPT which is greatly reducing workload in creating code that generally help with the assignment. In this assignment, the issue with ChatGPT is that it does not able to fully validate the program due to the usage of third party package, which need to give it hint in the prompt to revise the code. But in general, ChatGPT 3.5 is very powerful tool in completion of small programming task like this. What it cannot do is to provide the result of the code which cannot help in comparing Python and R. The full conversation with ChatGPT is store in md file as well for reference.

The training material is the original code from assignment 2 as well as Python and R code. There are no extra material being used in this excercise.

## Recommendation

For startup that utilize automation tool, I would recommend using less AI programming assisting for the core business and writing auxiliary program using LLM service or ai assisted tool. The core business schema is private and important information, and also require developer to be throughly understand the work where ai asssisted programming could help to write the piece code that increase the productivity and time saving. Another aspect is that creating the rule or displine to use AI assisting and generator as those tool may create random variable name or function name that not being specified by user. Especailly is important when several people work on the same project.


