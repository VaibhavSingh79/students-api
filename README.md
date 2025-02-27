# 📚 Database Managing API

🚀 A simple Go-based API that manages  data, providing endpoints for adding, retrieving, updating, and deleting  information. 
This project leverages Postman for documentation and testing🎯. It is designed for beginners in Go and API development, focusing on core CRUD operations.

<h2> 📖 Table of Contents</h2>
<ul>
<li>✨Features</li>
<li>⚙️Getting Started</li>
<li>🌐Prerequisites</li>
<li>🛠️Installation</li>
<li>🛜API Endpoints</li>
<li>🔧Postman Integration</li>
<li>💡Usage Examples</li>
<li>📜License</li>
</ul>

<h2>✨Features</h2>
<ul>
<li>✔️Add new user records.
<li>✔️Retrieve all  details or a specific user's data.
<li>✔️Update an existing user’s information.
<li>✔️Delete a user record.
<li>✔️Includes Postman examples for easy testing and integration.
</ul>




<h2>🚀Getting Started</h2>
<h3>🔧Prerequisites</h3>
Ensure you have the following installed on your system:
<ul><li>Go (Golang): Minimum version 1.20.0. or above🐹
<li>Postman for API testing and automation.🚀
<li>Git for cloning the repository.📂 </ul>


<h3>⚙️Installation</h3>
<ol>
<li>Clone the repository:

```bash

git clone https://github.com/VaibhavSingh79/students-api.git
cd students-api
```

<li>Build and run: </li>

```bash

go build
student./students-api
```

<li>The API service will start at</li>

```bash
http://localhost:8080
```

<h3>📌API Endpoints</h3>
Below is a summary of available endpoints:


<h3>1️⃣Create a User</h3>
<ul> 
<li>Endpoint: <b><i>POST /users</i></b></li>

<li>Request Body (JSON):</li>

  ```json

{
  "name": "John Doe",
  "age": 22,
  "class": "10th Grade"
}
```
<li>Description: Adds a new user to the database.
</ul>
  
<h3>2️⃣Retrieve All Students</h3>
<ul>
<li>Endpoint: <i>GET /users</i>
<li>Description: Fetches a list of all users.
</ul>

<h3>3️⃣Retrieve a  by ID</h3>
<ul>
<li>Endpoint: <b><i>GET /users/{id}</i></b>
<li>Description: Fetches a single user's details with the given ID.
</ul>


<h3>4️⃣Update a User</h3>
<ul>
<li>Endpoint: PUT /users/{id}
<li>Request Body (JSON):

  ```json

{
  "name": "Jane Doe",
  "age": 23,
  "class": "11th Grade"
}
```
<li>Description: Modifies existing user's details.
</ul>

  
<h3>Delete a User</h3>
<ul>
<li>Endpoint: <b><i>DELETE /s/{id}</i></b>
<li>Description: Removes a user record from the system.
</li></ul>


<h3>📮Postman Integration</h3>
Postman has been used to document and test the API. Use the provided Postman collection to simplify testing:
<ol>  
<li>Import the Postman collection file included in the repository <b><i>(users-api.postman_collection.json)</i></b>. It contains pre-configured requests for all available endpoints.

<li>Utilize Postman’s built-in features to view examples, authentication settings, and environment variables associated with the API collection 1 2 5 10 .
</ol>
  
<h3>Example: Testing with Postman</h3>
<ul>
<li>Open Postman and import the collection.
<li>Run a <b><i>POST</i></b> request to add a user.
<li>Execute a sequence of API calls to form a workflow, such as adding a user, retrieving their details, then updating and deleting them.
</ul>
  
<h3>💡Usage Examples</h3>

<h3>Adding a New user with Postman</h3>
<ul>
<li>Select the <b><i>POST /users</i></b> request.
<li>In the body, input:
  
  ```json

{
  "name": "Alice Smith",
  "age": 21,
  "class": "Final Year"
}
```
<li>Send the request and verify the response.
</ul>

<h3>Retrieving Users</h3>
Run the <b><i>GET /s</i></b> request to view all users, or use <b><i>GET /users/{id}</i></b> to fetch a specific .

