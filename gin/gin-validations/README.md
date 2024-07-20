# Gin Validations
This Uses ShouldBindJSON to bind the incoming JSON request to the User struct and checks for validation errors.

If there are errors, it returns a 400 Bad Request status with the error message.

If the binding and validation are successful, it responds with a 200 OK status and a success message along with the user data.

## Successful POST
![Screeshot for Successful POST](https://github.com/pillaiharish/golang-codes/blob/main/gin/gin-validations/successful_user_validation.png)

## Invalid country code
![Screeshot for invalid country code](https://github.com/pillaiharish/golang-codes/blob/main/gin/gin-validations/invalid_country_code.png)

## Invalid email id
![Screeshot for email id](https://github.com/pillaiharish/golang-codes/blob/main/gin/gin-validations/invalid_email_id.png)

## Invalid phone number
![Screeshot for invalid phone number](https://github.com/pillaiharish/golang-codes/blob/main/gin/gin-validations/invalid_phone_number.png)