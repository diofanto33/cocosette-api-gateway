# Register
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"email":"example@example.com","password":"mypassword"}' \
  http://localhost:8080/register

# Login
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"email":"example@example.com","password":"mypassword"}' \
  http://localhost:8080/login

# Validate
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"token":"your_token_here"}' \
  http://localhost:8080/validate
