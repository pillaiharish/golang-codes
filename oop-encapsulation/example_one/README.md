# ENCAPSULATION


Encapsulation is the practice of bundling the data (attributes) and the methods (functions) that operate on the data into a single unit or class and restricting access to some of the object's components. This helps in hiding the internal state of the object and only exposing a controlled interface to the outside world.


## Bundling Data and Methods:

The Account struct contains data fields (holder_name and balance).
The methods Deposit and Balance operate on these fields.

## Controlled Access:
The balance field is not directly accessible from outside the Account struct. It can only be modified through the Deposit method.
The Balance method provides a controlled way to access the balance field's value.



## Benefits of Encapsulation
- Data Hiding: The internal state (balance) is protected from unauthorized access and modification.
- Modularity: The Account struct and its methods form a self-contained module, making the code easier to maintain and understand.
- Flexibility and Maintainability: Changes to the internal implementation of the Account struct can be made without affecting the code that uses it, as long as the method interfaces remain the same.