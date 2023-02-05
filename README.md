# Todo (CLI)

> Command-line tool to manage the Todo lists

# Create a Todo App

## Requirement Commands

* Add a todo

    Ex:  
    `add` Go to market  
    `add` Buy a Food  

* List all todo

    Ex:  
    `ls`  
    ```
    (1) Go to market  
    (2) Buy a Food
    ```

* Remove a todo

    Ex:  
    `rm` 1  
    `ls`
    ```
    (1) Buy a Food
    ```

* Check a todo

    Ex:  
    `check` 1  
    `ls`
    ```
    (1) Go to market ✅  
    (2) Buy a Food
    ```

* Uncheck a todo

    Ex:  
    `uncheck` 1  
    `ls`
    ```
    (1) Go to market  
    (2) Buy a Food
    ```

### Coding style
*Make sure the code is readable, either by refactoring into `Controller` class or function.

## Additional Commands
* Toggle a todo
* Search todo

## References
[Swift readLine](https://www.digitalocean.com/community/tutorials/swift-readline-swift-print)